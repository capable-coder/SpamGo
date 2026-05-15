package modules

import (
	"fmt"
	"strings"
	"time"

	"github.com/amarnathcjd/gogram/telegram"
)

var IsBotMode bool

// Reply sends an HTML reply with flood wait handling (max 3 retries)
func Reply(m *telegram.NewMessage, text string) (*telegram.NewMessage, error) {
	retries := 0
	for retries < 3 {
		msg, err := m.Reply(text, &telegram.SendOptions{
			ParseMode: telegram.HTML,
		})
		if err != nil {
			errText := strings.ToUpper(err.Error())
			if strings.Contains(errText, "FLOOD_WAIT") {
				wait := extractFloodWait(errText)
				if wait <= 0 {
					wait = 5
				}
				time.Sleep(time.Duration(wait+1) * time.Second)
				retries++
				continue
			}
			// Reply failed — fallback to SendMessage
			_, _ = m.Client.SendMessage(m.ChatID(), text, &telegram.SendOptions{
				ParseMode: telegram.HTML,
			})
			return nil, err
		}
		return msg, nil
	}
	return nil, fmt.Errorf("reply failed after retries")
}

// GetArgs returns all text after the first word (the command itself)
func GetArgs(m *telegram.NewMessage) string {
	text := m.Text()
	if idx := findSpace(text); idx != -1 {
		return text[idx+1:]
	}
	return ""
}

func findSpace(s string) int {
	for i, c := range s {
		if c == ' ' {
			return i
		}
	}
	return -1
}

// extractFloodWait parses wait seconds from a FLOOD_WAIT error string
func extractFloodWait(errText string) int {
	var wait int
	_, _ = fmt.Sscanf(errText, "FLOOD_WAIT_%d", &wait)
	return wait
}

// getReplyMsgID returns the replied-to message ID directly from the message
// object without any network call.
// GetReplyMessage() does a network fetch which can return nil — this is the
// root cause of reply not working. We read ReplyTo from the raw message instead.
func getReplyMsgID(m *telegram.NewMessage) int32 {
	if !m.IsReply() {
		return 0
	}
	// m.Message is the raw *tg.Message object — ReplyTo holds the header
	// which contains the replied-to message ID without a network call
	if msg := m.Message; msg != nil {
		if replyTo := msg.ReplyTo; replyTo != nil {
			switch r := replyTo.(type) {
			case *telegram.MessageReplyHeader:
				return r.ReplyToMsgID
			}
		}
	}
	// Fallback: try network fetch if struct access didn't work
	replyMsg, err := m.GetReplyMessage()
	if err == nil && replyMsg != nil {
		return int32(replyMsg.ID)
	}
	return 0
}
