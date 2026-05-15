package modules

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"pbxgo/config"
	"pbxgo/database"

	"github.com/amarnathcjd/gogram/telegram"
)

/*
	PbxGo Raid Module
	Created By: BadMunda
*/

// ─────────────────────────────────────────────
// Active Raid Tracker
// ─────────────────────────────────────────────

var (
	activeRaidChats   = make(map[int64]bool)
	activeRaidChatsMu sync.Mutex

	replyRaidWatchers   = make(map[int64]func())
	replyRaidWatchersMu sync.Mutex
)

func setRaidActive(chatID int64) {
	activeRaidChatsMu.Lock()
	activeRaidChats[chatID] = true
	activeRaidChatsMu.Unlock()
}

func setRaidStopped(chatID int64) {
	activeRaidChatsMu.Lock()
	delete(activeRaidChats, chatID)
	activeRaidChatsMu.Unlock()
}

func isRaidActive(chatID int64) bool {
	activeRaidChatsMu.Lock()
	defer activeRaidChatsMu.Unlock()
	return activeRaidChats[chatID]
}

// ─────────────────────────────────────────────
// HANDLERS
// ─────────────────────────────────────────────

func raidHandler(m *telegram.NewMessage) error {
	return genericRaid(m, config.RAID, "raid", 700*time.Millisecond)
}

func hraidHandler(m *telegram.NewMessage) error {
	return genericRaid(m, config.HRAID, "hraid", 700*time.Millisecond)
}

func eraidHandler(m *telegram.NewMessage) error {
	return genericRaid(m, config.ERAID, "eraid", 700*time.Millisecond)
}

func punraidHandler(m *telegram.NewMessage) error {
	return genericRaid(m, config.PUNRAID, "punraid", 700*time.Millisecond)
}

// ─────────────────────────────────────────────
// Reply Raid Handlers
// ─────────────────────────────────────────────

func replyRaidHandler(m *telegram.NewMessage) error {
	return genericReplyRaid(m, config.RAID, "raid", 700*time.Millisecond)
}

func hreplyRaidHandler(m *telegram.NewMessage) error {
	return genericReplyRaid(m, config.HRAID, "hraid", 700*time.Millisecond)
}

func ereplyRaidHandler(m *telegram.NewMessage) error {
	return genericReplyRaid(m, config.ERAID, "eraid", 700*time.Millisecond)
}

func preplyRaidHandler(m *telegram.NewMessage) error {
	return genericReplyRaid(m, config.PUNRAID, "punraid", 700*time.Millisecond)
}

// ─────────────────────────────────────────────
// STOP RAID
// ─────────────────────────────────────────────

func stopRaidHandler(m *telegram.NewMessage) error {
	chatID := m.ChatID()

	setRaidStopped(chatID)

	replyRaidWatchersMu.Lock()
	if cancel, ok := replyRaidWatchers[chatID]; ok {
		cancel()
		delete(replyRaidWatchers, chatID)
	}
	replyRaidWatchersMu.Unlock()

	database.DeleteRaid(chatID)

	setShayariStopped(chatID)

	_, _ = m.Delete()

	Reply(m, "🛑 <b>ᴀʟʟ ʀᴀɪᴅs sᴛᴏᴘᴘᴇᴅ!</b>")

	return nil
}

// ─────────────────────────────────────────────
// FLOOD WAIT
// ─────────────────────────────────────────────

func handleFlood(err error) bool {

	if err == nil {
		return false
	}

	errText := strings.ToUpper(err.Error())

	if strings.Contains(errText, "FLOOD_WAIT") {

		var wait int

		fmt.Sscanf(errText, "FLOOD_WAIT_%d", &wait)

		if wait <= 0 {
			wait = 5
		}

		time.Sleep(time.Duration(wait+1) * time.Second)

		return true
	}

	return false
}

// ─────────────────────────────────────────────
// GENERIC RAID
// FIXED REPLY SUPPORT
// ─────────────────────────────────────────────

func genericRaid(
	m *telegram.NewMessage,
	raidList []string,
	raidType string,
	delay time.Duration,
) error {

	args := GetArgs(m)

	var count int

	fmt.Sscanf(args, "%d", &count)

	if count < 1 {
		Reply(m, "❌ ᴄᴏᴜɴᴛ ᴍᴜsᴛ ʙᴇ ᴀᴛ ʟᴇᴀsᴛ 1.")
		return nil
	}

	// ───── FIXED REPLY ID ─────

	var replyToID int32

	if m.IsReply() {

		replyMsg, err := m.GetReplyMessage()

		if err == nil && replyMsg != nil {
			replyToID = int32(replyMsg.ID)
		}
	}

	// ──────────────────────────

	database.SaveRaid(database.RaidSession{
		ChatID:    m.ChatID(),
		ReplyToID: replyToID,
		RaidType:  raidType,
		Count:     count,
	})

	_, _ = m.Delete()

	setRaidActive(m.ChatID())

	defer func() {
		setRaidStopped(m.ChatID())
		database.DeleteRaid(m.ChatID())
	}()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {

		if !isRaidActive(m.ChatID()) {
			break
		}

		text := raidList[rng.Intn(len(raidList))]

		opts := &telegram.SendOptions{
			ParseMode: telegram.HTML,
		}

		// ───── FIXED REPLY ─────

		if replyToID > 0 {
			opts.ReplyID = replyToID
		}

		// ───────────────────────

		_, err := m.Client.SendMessage(
			m.ChatID(),
			text,
			opts,
		)

		if err != nil {

			if handleFlood(err) {
				i--
				continue
			}

			time.Sleep(2 * time.Second)
			continue
		}

		time.Sleep(delay)
	}

	return nil
}

// ─────────────────────────────────────────────
// GENERIC REPLY RAID
// ALWAYS REPLIES TO TARGET MESSAGE
// ─────────────────────────────────────────────

func genericReplyRaid(
	m *telegram.NewMessage,
	raidList []string,
	raidType string,
	delay time.Duration,
) error {

	if !m.IsReply() {

		Reply(
			m,
			"↩️ ʀᴇᴘʟʏ ᴛᴏ ᴀ ᴜsᴇʀ ᴍᴇssᴀɢᴇ ꜰɪʀsᴛ.",
		)

		return nil
	}

	args := GetArgs(m)

	var count int

	fmt.Sscanf(args, "%d", &count)

	if count < 1 {
		count = 1
	}

	replyMsg, err := m.GetReplyMessage()

	if err != nil || replyMsg == nil {

		Reply(
			m,
			"❌ ꜰᴀɪʟᴇᴅ ᴛᴏ ɢᴇᴛ ʀᴇᴘʟɪᴇᴅ ᴍᴇssᴀɢᴇ.",
		)

		return nil
	}

	replyToID := int32(replyMsg.ID)

	database.SaveRaid(database.RaidSession{
		ChatID:    m.ChatID(),
		ReplyToID: replyToID,
		RaidType:  raidType,
		Count:     count,
	})

	_, _ = m.Delete()

	setRaidActive(m.ChatID())

	defer func() {
		setRaidStopped(m.ChatID())
		database.DeleteRaid(m.ChatID())
	}()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {

		if !isRaidActive(m.ChatID()) {
			break
		}

		text := raidList[rng.Intn(len(raidList))]

		_, err := m.Client.SendMessage(
			m.ChatID(),
			text,
			&telegram.SendOptions{
				ReplyID:   replyToID,
				ParseMode: telegram.HTML,
			},
		)

		if err != nil {

			if handleFlood(err) {
				i--
				continue
			}

			time.Sleep(2 * time.Second)

			continue
		}

		time.Sleep(delay)
	}

	return nil
}

// ─────────────────────────────────────────────
// RESUME RAIDS
// ─────────────────────────────────────────────

func ResumeRaids(cl *telegram.Client) {

	sessions := database.LoadActiveRaids()

	for _, s := range sessions {

		go func(sess database.RaidSession) {

			raidList := getRaidList(sess.RaidType)

			if len(raidList) == 0 {
				return
			}

			rng := rand.New(rand.NewSource(time.Now().UnixNano()))

			activeRaidChatsMu.Lock()
			activeRaidChats[sess.ChatID] = true
			activeRaidChatsMu.Unlock()

			defer func() {
				setRaidStopped(sess.ChatID)
				database.DeleteRaid(sess.ChatID)
			}()

			for i := 0; i < sess.Count; i++ {

				if !isRaidActive(sess.ChatID) {
					break
				}

				text := raidList[rng.Intn(len(raidList))]

				opts := &telegram.SendOptions{
					ParseMode: telegram.HTML,
				}

				// FIXED REPLY SUPPORT

				if sess.ReplyToID > 0 {
					opts.ReplyID = sess.ReplyToID
				}

				_, err := cl.SendMessage(
					sess.ChatID,
					text,
					opts,
				)

				if err != nil {

					if handleFlood(err) {
						i--
						continue
					}

					time.Sleep(2 * time.Second)

					continue
				}

				time.Sleep(800 * time.Millisecond)
			}

		}(s)
	}
}

// ─────────────────────────────────────────────
// GET RAID LIST
// ─────────────────────────────────────────────

func getRaidList(raidType string) []string {

	switch raidType {

	case "raid":
		return config.RAID

	case "hraid":
		return config.HRAID

	case "eraid":
		return config.ERAID

	case "punraid":
		return config.PUNRAID

	default:
		return config.RAID
	}
}

// ─────────────────────────────────────────────
// REGISTER
// ─────────────────────────────────────────────

func init() {

	Register(ModuleInfo{
		Name:        "Raid",
		Description: "Multi Language Raid Commands",
		Commands: []CommandInfo{
			{Pattern: "raid", Handler: raidHandler, Sudo: true},
			{Pattern: "hraid", Handler: hraidHandler, Sudo: true},
			{Pattern: "eraid", Handler: eraidHandler, Sudo: true},
			{Pattern: "punraid", Handler: punraidHandler, Sudo: true},

			{Pattern: "replyraid", Handler: replyRaidHandler, Sudo: true},
			{Pattern: "hreplyraid", Handler: hreplyRaidHandler, Sudo: true},
			{Pattern: "ereplyraid", Handler: ereplyRaidHandler, Sudo: true},
			{Pattern: "preplyraid", Handler: preplyRaidHandler, Sudo: true},

			{Pattern: "stopraid", Handler: stopRaidHandler, Sudo: true},
		},
	})
}
