package usecase

import "github.com/ervitis/ikigaikintore/backend/pkg/proto"

type weatherToProto struct {
}

type WeatherToProto interface {
	ToRequest() *proto.WeatherRequest
}

func (w weatherToProto) ToRequest() *proto.WeatherRequest {
	return &proto.WeatherRequest{}
}
