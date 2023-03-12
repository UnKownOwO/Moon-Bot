package bot

type ModBase interface {
	initData()     // 初始化数据
	loadData(*Bot) // 加载数据
	saveData()     // 保存数据
}

const (
	modIdBot uint8 = iota
)

type Bot struct {
	UserId    int64             // 用户Id
	modManage map[uint8]ModBase // 模块
	exitTime  int64             // 离线时间
}

func NewBot(userId int64) *Bot {
	bot := new(Bot)
	bot.UserId = userId
	bot.modManage = map[uint8]ModBase{
		modIdBot: new(ModBot),
	}
	// 初始化模块
	bot.initMod()

	return bot
}

// initMod 初始化模块
func (b *Bot) initMod() {
	for _, mod := range b.modManage {
		mod.loadData(b)
	}
}

// close bot被清理时将执行关闭
func (b *Bot) close() {
	b.GetModBot().closeChan <- true
}

// getMod 获取某一模块
func (b *Bot) getMod(modId uint8) ModBase {
	return b.modManage[modId]
}

// GetModBot 获取基础模块
func (b *Bot) GetModBot() *ModBot {
	return b.getMod(modIdBot).(*ModBot)
}
