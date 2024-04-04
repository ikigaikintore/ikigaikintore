package bot

import (
	"context"
	"github.com/ikigaikintore/ikigaikintore/proxybot/pkg/domain"
	"github.com/ikigaikintore/ikigaikintore/proxybot/pkg/service"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
)

func NewHandlerTodayWeather(cli service.WeatherClient) CommandUpdate {
	return &cmdUpdateHandler{
		fn: func(bot *telego.Bot, update telego.Update) {
			respWeather, err := cli.GetWeather(update.Context(), &service.WeatherRequest{WeatherFilter: &service.WeatherFilter{Location: "Tokyo"}})
			if err != nil {
				log.Println(err)
				return
			}
			today := domain.IntoToday(respWeather.GetWeatherCurrent())
			_, _ = bot.SendPhoto(
				tu.Photo(
					tu.ID(update.Message.Chat.ID),
					tu.FileFromURL("https://openweathermap.org/img/wn/"+today.Icon+"@2x.png"),
				).WithCaption(today.String()),
			)
		},
		cmds: []th.Predicate{th.CommandEqualArgv("weather", "now")},
	}
}

func NewHandlerFuture(cli service.WeatherClient) CommandUpdate {
	return &cmdUpdateHandler{
		fn: func(bot *telego.Bot, update telego.Update) {
			respWeather, err := cli.GetWeather(update.Context(), &service.WeatherRequest{WeatherFilter: &service.WeatherFilter{Location: "Tokyo"}})
			if err != nil {
				log.Println(err)
				return
			}
			weatherPoints := respWeather.GetWeatherPoint()
			var msg string
			for _, v := range weatherPoints {
				msg += domain.IntoPoint(v).String() + "\r\n"
			}
			_, _ = bot.SendMessage(
				tu.Message(
					tu.ID(update.Message.Chat.ID),
					msg,
				),
			)
		},
		cmds: []th.Predicate{th.CommandEqualArgv("weather", "3h")},
	}
}

func NewHandlerLocation() CommandUpdate {
	return &cmdUpdateHandler{
		fn: func(bot *telego.Bot, update telego.Update) {
			_, _ = bot.SendMessage(
				tu.Message(
					tu.ID(update.Message.Chat.ID),
					"send me your location",
				).
					WithDisableNotification().
					WithReplyMarkup(
						tu.Keyboard(
							tu.KeyboardRow(
								tu.KeyboardButton("Send location").
									WithRequestLocation(),
							),
						).
							WithOneTimeKeyboard(),
					),
			)
		},
		cmds: []th.Predicate{th.CommandEqual("location")},
	}
}

// add cache here to save temporally the location
// as default, i will return the weather from tokyo

func NewHandlerResponseLocation() CommandMessage {
	return &cmdMessageHandler{
		fn: func(ctx context.Context, bot *telego.Bot, message telego.Message) {
			if message.Location != nil {

			}
		},
		cmds: []th.Predicate{th.AnyEditedMessageWithFrom()},
	}
}
