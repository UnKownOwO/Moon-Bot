package bot

import "moon-bot/bs/model"

type LoginManager struct {
	botMap map[int64]*model.Bot // 内存机器人数据
}

func NewLoginManager() *LoginManager {
	r := new(LoginManager)

	r.botMap = make(map[int64]*model.Bot)

	return r
}

// AddBot 内存数据中添加机器人
func (l *LoginManager) AddBot(bot *model.Bot) {
	l.botMap[bot.Account] = bot
}

// DeleteBot 内存数据中删除机器人
func (l *LoginManager) DeleteBot(account int64) {
	delete(l.botMap, account)
}
