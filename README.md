<div align="center">

# 🍹 ᴘʙxsᴘᴀᴍɢᴏ

<p align="center">
<a href="https://t.me/PBXCHATS">
<img src="https://files.tgvibes.online/5JreGgKB.jpg" width="600">
</a>
</p>

**Multi-Bot Spam & Raid Tool — Built with Go + gogram**

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev)
[![Telegram](https://img.shields.io/badge/Telegram-Bot-blue?style=flat&logo=telegram)](https://core.telegram.org/bots)
[![Developer](https://img.shields.io/badge/Dev-@BadmundaXd-purple?style=flat)](https://t.me/BadmundaXd)

</div>

---

## 📋 Commands

### 🚀 Spam Commands

| Command | Usage | Description |
|---------|-------|-------------|
| `.spam` | `.spam 10 hello` | Normal text spam |
| `.ds` | `.ds 1.5 10 hello` | Delay spam (delay in seconds) |
| `.delayspam` | `.delayspam 1 10 hello` | Alias for `.ds` |
| `.sspam` | `.sspam 10` | Sticker/media spam (reply to media) |
| `.hang` | `.hang 10` | Hang spam |
| `.pspam` | `.pspam 10` | Porn text spam |
| `.stopspam` | `.stopspam` | Stop all active spam |

---

### ⚔️ Raid Commands

| Command | Usage | Description |
|---------|-------|-------------|
| `.raid` | `.raid 20` | Hindi raid |
| `.hraid` | `.hraid 20` | Haryanvi raid |
| `.eraid` | `.eraid 20` | English raid |
| `.punraid` | `.punraid 20` | Punjabi raid |
| `.praid` | `.praid 20` | Porn text + video raid |
| `.shayari` | `.shayari 10` | Shayari raid |
| `.stopraid` | `.stopraid` | Stop all active raids |

> **Reply Mode:** Reply to any user's message before using raid commands — all messages will go as reply to that specific message.

---

### 👁 Reply Raid Commands (Watcher Mode)

| Command | Usage | Description |
|---------|-------|-------------|
| `.replyraid` | Reply to user + `.replyraid` | Watch target — raid when they next msg |
| `.hreplyraid` | Reply to user + `.hreplyraid` | Haryanvi reply raid |
| `.ereplyraid` | Reply to user + `.ereplyraid` | English reply raid |
| `.preplyraid` | Reply to user + `.preplyraid` | Punjabi reply raid |
| `.preplyraid` | Reply to user + `.preplyraid` | Porn reply raid |

> Reply raid sends **one random message** when target user sends any message, then auto-stops.

---

### ⚙️ Extra Commands

| Command | Usage | Description |
|---------|-------|-------------|
| `.ping` | `.ping` | Check bot ping & status |
| `.restart` | `.restart` | Restart the bot |
| `.logs` | `.logs` | Fetch bot logs (pm2) |
| `.addsudo` | `.addsudo [id]` or reply | Add sudo user |
| `.rmsudo` | `.rmsudo [id]` or reply | Remove sudo user |
| `.sudolist` | `.sudolist` | List all sudo users |
| `/start` | `/start` | Start message |
| `/help` | `/help` | Help menu |

---

## ⚙️ Environment Variables

Create a `.env` file with the following:

```env
# Required
APP_ID=12345678
APP_HASH=0123456789abcdef0123456789abcdef
OWNER_ID=your_telegram_id
BOT_TOKEN1=123456789:AAxxxxxxxxxx

# Multiple bots (optional)
BOT_TOKEN2=123456789:AAxxxxxxxxxx
BOT_TOKEN3=123456789:AAxxxxxxxxxx

# Optional
MONGO_URL=mongodb+srv://user:pass@cluster.mongodb.net/
START_PIC=https://files.tgvibes.online/5JreGgKB.jpg
HELP_PIC=https://files.tgvibes.online/5JreGgKB.jpg
```

> Get `APP_ID` and `APP_HASH` from [my.telegram.org](https://my.telegram.org)

---

## 🚀 Hosting Guide

### 🖥️ VPS / Linux

```bash
# Install Go
sudo apt update
sudo apt install golang git -y

# Clone repo
git clone https://github.com/badmunda05/SpamGo.git
cd SpamGo

# Create .env
cp sample.env .env
nano .env   # fill in your values

# Build & run
go mod tidy
go build -o pbxspamgo .
./pbxspamgo

# Run with pm2 (recommended)
npm install -g pm2
pm2 start pbxspamgo --name spamgo
pm2 save
pm2 startup
```

---

### 📱 Termux (Android)

```bash
# One-line setup
bash termux-install.sh
```

Or manually:

```bash
pkg update -y && pkg install golang git -y
git clone https://github.com/badmunda05/SpamGo.git
cd SpamGo
cp sample.env .env
nano .env
go mod tidy
go build -o pbxspamgo .
./pbxspamgo
```

Keep running after closing Termux:
```bash
nohup ./pbxspamgo &
```

---

### 🐳 Docker

```bash
# Build image
docker build -t pbxspamgo .

# Run with env file
docker run --env-file .env pbxspamgo

# Or with individual env vars
docker run \
  -e APP_ID=12345678 \
  -e APP_HASH=abcdef \
  -e OWNER_ID=123456 \
  -e BOT_TOKEN1=xxx \
  pbxspamgo
```
### ✈️ Fly.io

```bash
# Install flyctl
curl -L https://fly.io/install.sh | sh

# Login
flyctl auth login

# Deploy
flyctl launch
flyctl secrets set APP_ID=xxx APP_HASH=xxx OWNER_ID=xxx BOT_TOKEN1=xxx
flyctl deploy
```
---
## 🚀 Deploy Now

<div align="center">

### ☁️ Heroku

[![Deploy on Heroku](https://img.shields.io/badge/Deploy%20On%20Heroku-430098?style=for-the-badge&logo=heroku&logoColor=white)](https://dashboard.heroku.com/new?template=https://github.com/Badmunda05/SpamGo)

---

### 🟢 Koyeb

[![Deploy on Koyeb](https://img.shields.io/badge/Deploy%20On%20Koyeb-121212?style=for-the-badge&logo=koyeb&logoColor=white)](https://app.koyeb.com/deploy?type=git&repository=github.com/Badmunda05/SpamGo&branch=main&name=SpamGo)

---

### 🟣 Render

[![Deploy on Render](https://img.shields.io/badge/Deploy%20On%20Render-46E3B7?style=for-the-badge&logo=render&logoColor=black)](https://render.com/deploy?repo=https://github.com/Badmunda05/SpamGo)

---

### 🔵 Railway

[![Deploy on Railway](https://img.shields.io/badge/Deploy%20On%20Railway-0B0D0E?style=for-the-badge&logo=railway&logoColor=white)](https://railway.app/new/template?template=https://github.com/Badmunda05/SpamGo)

</div>

---
## 📁 Project Structure

```
SpamGo/
├── main.go              # Entry point
├── .env                 # Your config (not committed)
├── sample.env           # Example env file
├── app.json             # Heroku config
├── Dockerfile           # Docker build
├── Procfile             # Heroku/pm2 process
├── railway.json         # Railway config
├── render.yaml          # Render config
├── fly.toml             # Fly.io config
├── termux-install.sh    # Termux auto-setup
├── client/
│   └── client.go        # Bot client manager
├── config/
│   ├── config.go        # Env loader
│   └── data.go          # Raid/spam text data
├── database/
│   ├── db.go            # MongoDB connection
│   ├── raid.go          # Active raid persistence
│   └── sudo.go          # Sudo user management
└── modules/
    ├── basic.go         # Start, help, ping, sudo
    ├── module.go        # Handler registration
    ├── raid.go          # All raid commands
    ├── spam.go          # All spam commands
    ├── praid.go         # Porn raid + spam
    ├── shayari.go       # Shayari raid
    └── utils.go         # Helpers
```

---

## 👨‍💻 Developer

**BadMunda** — [@BadmundaXd](https://t.me/BadmundaXd)

Channel: [@PBX_UPDATE](https://t.me/PBX_UPDATE) | Support: [@PBXCHATS](https://t.me/PBXCHATS)

---

<div align="center">
Made with ❤️ by BadMunda
</div>
