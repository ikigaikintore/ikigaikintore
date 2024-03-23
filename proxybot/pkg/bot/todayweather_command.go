package bot

import (
	"log"

	"github.com/ikigaikintore/ikigaikintore/backend/pkg/proto"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func NewHandlerTodayWeather(cli proto.WeatherClient) Command {
	return &cmdHandler{
		fn: func(bot *telego.Bot, update telego.Update) {
			resp, err := cli.GetWeather(update.Context(), &proto.WeatherRequest{WeatherFilter: &proto.WeatherFilter{Location: "Tokyo"}})
			if err != nil {
				log.Println(err)
				return
			}
			_, _ = bot.SendMessage(
				tu.Messagef(
					tu.ID(update.Message.Chat.ID),
					"Got somethin' that might interest ya'! %v",
					resp.GetWeatherCurrent().GetTemperature(),
				),
			)
		},
		cmds: []th.Predicate{th.CommandEqual("todayWeather")},
	}
}
