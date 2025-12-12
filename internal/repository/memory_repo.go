package repository

import (
	"errors"
	"sync"
)

// Client representa la entidad cliente.
type Client struct {
	ID      int    `json:"id"`
	Nombre  string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad    int    `json:"edad"`
}

// Repository define la abstracción para el acceso a datos.
type Repository interface {
	FindAll() ([]Client, error)
	Save(c Client) (Client, error)
}

// memoryRepo implementa Repository con un map en memoria.
type memoryRepo struct {
	mu    sync.RWMutex
	data  map[int]Client
	lastID int
}

// NewMemoryRepository crea una instancia del repositorio en memoria.
func NewMemoryRepository() Repository {
	return &memoryRepo{
		data:   make(map[int]Client),
		lastID: 0,
	}
}

var ErrInvalidClient = errors.New("client: datos inválidos")

// FindAll retorna todos los clientes actuales.
func (r *memoryRepo) FindAll() ([]Client, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]Client, 0, len(r.data))
	for _, v := range r.data {
		list = append(list, v)
	}
	return list, nil
}

// Save añade un cliente al repositorio, asignando un ID incremental.
func (r *memoryRepo) Save(c Client) (Client, error) {
	if c.Nombre == "" || c.Apellido == "" {
		return Client{}, ErrInvalidClient
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++
	c.ID = r.lastID
	r.data[c.ID] = c
	return c, nil
}
