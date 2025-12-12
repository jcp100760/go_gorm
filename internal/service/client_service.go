package service

import "github.com/jcp100760/go_gorm/internal/repository"

// Service define la abstracción de la capa de negocio para clientes.
type Service interface {
	GetAllClients() ([]repository.Client, error)
	AddClient(c repository.Client) (repository.Client, error)
}

// clientService es la implementación concreta de Service.
type clientService struct {
	repo repository.Repository
}

// NewClientService crea un servicio de clientes con la repo inyectada.
func NewClientService(r repository.Repository) Service {
	return &clientService{repo: r}
}

func (s *clientService) GetAllClients() ([]repository.Client, error) {
	return s.repo.FindAll()
}

func (s *clientService) AddClient(c repository.Client) (repository.Client, error) {
	// Aquí podrían ir reglas de negocio (validaciones, eventos, etc.)
	return s.repo.Save(c)
}
