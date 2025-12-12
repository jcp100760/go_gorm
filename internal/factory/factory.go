package factory

import (
	"github.com/gorilla/mux"
	"github.com/jcp100760/go_gorm/internal/handlers"
	"github.com/jcp100760/go_gorm/internal/repository"
	"github.com/jcp100760/go_gorm/internal/service"
)

// Factory es el ensamblador principal de la aplicación.
type Factory struct {
	repo repository.Repository
	svc  service.Service
	hnd  *handlers.Handlers
	rtr  *mux.Router
}

// NewFactory crea y arma todas las piezas (repository -> service -> handlers -> router)
func NewFactory() *Factory {
	// 1) Crear repositorio concreto (in-memory)
	repo := repository.NewMemoryRepository()

	// 2) Crear servicio inyectando la interface del repositorio
	svc := service.NewClientService(repo)

	// 3) Crear handlers con el servicio
	hnd := handlers.NewHandlers(svc)

	// 4) Configurar router y registrar rutas
	router := mux.NewRouter()
	hnd.RegisterRoutes(router)

	return &Factory{
		repo: repo,
		svc:  svc,
		hnd:  hnd,
		rtr:  router,
	}
}

// Router expone el router configurado para que main lo use.
func (f *Factory) Router() *mux.Router {
	return f.rtr
}
