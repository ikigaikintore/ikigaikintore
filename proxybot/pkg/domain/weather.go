package domain

import (
	"github.com/ikigaikintore/ikigaikintore/proxybot/pkg/service"
	"strconv"
)

type Today struct {
	Temperature float64
	Humidity    int32
	WindSpeed   float64
	Icon        string
}

func Into(resp *service.WeatherCurrent) Today {
	return Today{
		Temperature: resp.GetTemperature(),
		Humidity:    resp.GetHumidity(),
		WindSpeed:   resp.GetWindSpeed(),
		Icon:        resp.GetIcon(),
	}
}

func (t Today) String() string {
	return strconv.FormatFloat(t.Temperature, 'f', 1, 64) + "C " +
		strconv.FormatInt(int64(t.Humidity), 10) + "% " +
		strconv.FormatFloat(t.WindSpeed, 'f', 1, 64) + " m/s"
}
