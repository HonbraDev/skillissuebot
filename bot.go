package skillissue

import (
	"log"
	"os"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
)

type Bot struct {
	Config
	Client *bot.Client
	Player *basic.Player
	Logger *log.Logger
}

func NewBot(c Config) *Bot {
	client := bot.NewClient()
	player := basic.NewPlayer(client, basic.Settings{
		Locale:              "en_US",
		ViewDistance:        0,
		ChatMode:            0,
		DisplayedSkinParts:  basic.Jacket | basic.LeftSleeve | basic.RightSleeve | basic.LeftPantsLeg | basic.RightPantsLeg | basic.Hat,
		MainHand:            1,
		ChatColors:          true,
		EnableTextFiltering: false,
		AllowListing:        true,
		Brand:               "Samsung Smart Fridge",
	})
	logger := log.New(os.Stdout, "", log.LstdFlags)
	b := &Bot{
		Config: c,
		Client: client,
		Player: player,
		Logger: logger,
	}
	basic.EventsListener{
		Disconnect: b.handleDisconnect,
		ChatMsg:    b.handleChat,
		Death:      b.handleDeath,
		GameStart:  b.handleGameStart,
	}.Attach(client)
	return b
}

func (b *Bot) Run(host string) error {
	err := b.Client.JoinServer(host)
	if err != nil {
		return err
	}
	b.Logger.Println("Joined server")
	return b.Client.HandleGame()
}
