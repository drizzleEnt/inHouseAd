package service

type AuthService interface {
	Auth()
}

type ProductService interface {
	GetCategoryList() error
	GetGoodInCategory()
	ControlService
}

type ControlService interface {
	AddCategory()
	UpdateCategory()
	DeleteCategory()
	AddGoods()
	UpdateGoods()
	DeleteGoods()
}

type CollectingService interface {
}
