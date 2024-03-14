package product

import (
	"github.com/drizzleent/inHouseAd/internal/service"
	"github.com/drizzleent/inHouseAd/internal/service/product/control"
)

type Service struct {
	service.ControlService
}

func NewProductService() service.ProductService {
	return &Service{
		ControlService: control.NewControlService(),
	}
}
