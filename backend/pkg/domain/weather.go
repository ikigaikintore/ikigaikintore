package domain

import (
	"github.com/ikigaikintore/ikigaikintore/backend/internal/output/weather"
)

type WeatherType int

const (
	UNKNOWN WeatherType = iota
	THUNDERSTORM
	DRIZZLE
	RAIN
	SNOW
	MIST
	SMOKE
	HAZE
	DUST
	FOG
	SAND
	ASH
	SQUALL
	TORNADO
	CLEAR
	CLOUDS
)

type WeatherRequest struct {
	Cnt int

	Latitude   float64
	Longitude  float64
	APIKey     string
	ModeToday  weather.GetCurrentParamsMode
	Mode3H     weather.GetForecast3HourParamsMode
	Lang       string
	UnitsToday weather.GetCurrentParamsUnits
	Units3H    weather.GetForecast3HourParamsUnits
}

type WeatherPoint struct {
	Timestamp        uint64
	Temperature      float64
	Humidity         int32
	Icon             string
	TemperatureRange struct {
		Min, Max float64
	}
	Weather WeatherType
}

type WeatherCurrent struct {
	Temperature float64
	Icon        string
	WindSpeed   float64
	Timestamp   uint64
	Humidity    int32
	Weather     WeatherType
}

type WeatherResponse struct {
	WeatherPoints  []WeatherPoint
	WeatherCurrent WeatherCurrent
}

func WeatherDefaultParameters() WeatherRequest {
	return WeatherRequest{
		ModeToday:  weather.GetCurrentParamsModeJson,
		UnitsToday: weather.GetCurrentParamsUnitsMetric,
		Units3H:    weather.GetForecast3HourParamsUnitsMetric,
		Mode3H:     weather.GetForecast3HourParamsModeJson,
	}
}
