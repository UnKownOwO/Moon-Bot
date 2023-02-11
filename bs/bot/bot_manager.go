package bot

var (
	BOT_MANAGER *BotManager
)

type BotManager struct {
}

func NewBotManager() *BotManager {
	botManager := new(BotManager)

	// 初始化管理器
	BOT_MANAGER = botManager

	return botManager
}
