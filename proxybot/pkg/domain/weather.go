package domain

import (
	"github.com/ikigaikintore/ikigaikintore/proxybot/pkg/service"
	"strconv"
	"time"
)

type Today struct {
	Temperature float64
	Humidity    int32
	WindSpeed   float64
	Icon        string
}

func IntoToday(resp *service.WeatherCurrent) Today {
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

type TemperatureRange struct {
	Max, Min float64
}
type Point struct {
	Timestamp        time.Time
	Temperature      float64
	Humidity         int32
	TemperatureRange TemperatureRange
}

type Points []Point

func IntoPoint(resp *service.WeatherDailyPoint) Point {
	return Point{
		Timestamp:   time.Unix(int64(resp.GetTimestamp()), 0),
		Temperature: resp.GetTemperature(),
		Humidity:    resp.GetHumidity(),
		TemperatureRange: TemperatureRange{
			Max: resp.GetTemperatureRange().GetMax(),
			Min: resp.GetTemperatureRange().GetMin(),
		},
	}
}

func (p Point) String() string {
	return p.Timestamp.Format(time.DateTime) + " -> " +
		strconv.FormatFloat(p.Temperature, 'f', 1, 64) + "C " +
		strconv.FormatInt(int64(p.Humidity), 10) + "% " +
		strconv.FormatFloat(p.TemperatureRange.Min, 'f', 1, 64) + "C~" +
		strconv.FormatFloat(p.TemperatureRange.Max, 'f', 1, 64) + "C"
}
