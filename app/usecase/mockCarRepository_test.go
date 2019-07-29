package usecase

import (
	"github.com/diegoahg/journey/app/domain"
	"github.com/stretchr/testify/mock"
)

type MockCarRepository struct {
	mock.Mock
}

func (m *MockCarRepository) Save(car *domain.Car) error {
	args := m.Called(car)
	return args.Error(0)
}

func (m *MockCarRepository) RemoveAll() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCarRepository) GetEmptys() ([]*domain.Car, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Car), args.Error(1)
}

func (m *MockCarRepository) Update(car *domain.Car) error {
	args := m.Called(car)
	return args.Error(0)
}

func (m *MockCarRepository) FindByID(id int) (*domain.Car, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Car), args.Error(1)
}
