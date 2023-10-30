package route

import "github.com/yarikTri/web-transport-cards/internal/models"

type Usecase interface {
	GetByID(routeID uint32) (models.Route, error)
	List() ([]models.Route, error)
	Search(subString string) ([]models.Route, error)
}

type Repository interface {
	GetByID(routeID uint32) (models.Route, error)
	List() ([]models.Route, error)
	Search(subString string) ([]models.Route, error)
	DeleteByID(routeID uint32) error
}

type Tables interface {
	Routes() string
	RoutesStations() string
	Stations() string
}
