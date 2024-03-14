package control

import "github.com/drizzleent/inHouseAd/internal/service"

type Service struct {
}

func NewControlService() service.ControlService {
	return &Service{}
}
