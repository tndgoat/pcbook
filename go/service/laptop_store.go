package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/tndgoat/pcbook/pb"
)

// ErrAlreadyExists is returned when a record already exists in the store.
var ErrAlreadyExists = errors.New("record already exists")

// LaptopStore defines the interface for storing laptops.
type LaptopStore interface {
	// Save saves a laptop to the in-memory store.
	Save(laptop *pb.Laptop) error
	// Find retrieves a laptop from the store by its ID.
	Find(id string) (*pb.Laptop, error)
}

// InMemoryLaptopStore is an in-memory implementation of the LaptopStore interface.
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// NewInMemoryLaptopStore creates a new instance of InMemoryLaptopStore.
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if _, exists := store.data[laptop.Id]; exists {
		return ErrAlreadyExists
	}

	// Store a deep copy of the laptop to avoid external modifications.
	other := &pb.Laptop{} // Save saves a laptop to the in-memory store.
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[other.Id] = other

	return nil
}

// Find retrieves a laptop from the store by its ID.
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop, exists := store.data[id]
	if !exists {
		return nil, nil
	}

	// Return a deep copy of the laptop to avoid external modifications.
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}
