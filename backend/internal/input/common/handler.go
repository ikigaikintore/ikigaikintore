package common

import (
	"context"
	"log"
	"os"

	"github.com/ikigaikintore/ikigaikintore/backend/pkg/domain"
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/proto"
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/usecase"
)

type handler struct {
	svc usecase.WeatherService
}

type Handler interface {
	GetWeather(ctx context.Context, request *proto.WeatherRequest) (*proto.WeatherReply, error)
}

func NewHandler() *handler {
	return &handler{svc: usecase.NewWeatherService()}
}

func (w handler) GetWeather(ctx context.Context, request *proto.WeatherRequest) (*proto.WeatherReply, error) {
	filt := request.GetWeatherFilter()
	params := domain.WeatherDefaultParameters()
	resp, err := w.svc.GetWeather(ctx, domain.WeatherRequest{
		Cnt:        15,
		Latitude:   filt.GetLatitude(),
		Longitude:  filt.GetLongitude(),
		APIKey:     os.Getenv("OPENWEATHER_API_KEY"),
		ModeToday:  params.ModeToday,
		Mode3H:     params.Mode3H,
		Lang:       params.Lang,
		UnitsToday: params.UnitsToday,
		Units3H:    params.Units3H,
	})
	if err != nil {
		log.Println("getWeather", err)
		return nil, err
	}
	return &proto.WeatherReply{
		WeatherPoint:   toWeatherPoints(resp.WeatherPoints),
		WeatherCurrent: toWeatherCurrent(resp.WeatherCurrent),
	}, nil
}

func toWeatherPoints(points []domain.WeatherPoint) []*proto.WeatherDailyPoint {
	wp := make([]*proto.WeatherDailyPoint, len(points))
	for i, point := range points {
		wp[i] = &proto.WeatherDailyPoint{
			Timestamp:   point.Timestamp,
			Temperature: point.Temperature,
			Humidity:    point.Humidity,
			TemperatureRange: &proto.TemperatureRange{
				Max: point.TemperatureRange.Max,
				Min: point.TemperatureRange.Min,
			},
			Weather: proto.WeatherType(point.Weather),
			Icon:    point.Icon,
		}
	}
	return wp
}

func toWeatherCurrent(curr domain.WeatherCurrent) *proto.WeatherCurrent {
	return &proto.WeatherCurrent{
		Timestamp:   curr.Timestamp,
		Temperature: curr.Temperature,
		Humidity:    curr.Humidity,
		Icon:        curr.Icon,
		WindSpeed:   curr.WindSpeed,
		Weather:     proto.WeatherType(curr.Weather),
	}
}
