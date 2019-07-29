package usecase

import (
	"github.com/diegoahg/journey/app/domain"
	"github.com/stretchr/testify/mock"
)

type MockJourneyRepository struct {
	mock.Mock
}

func (m *MockJourneyRepository) Save(car *domain.Journey) error {
	args := m.Called(car)
	return args.Error(0)
}

func (m *MockJourneyRepository) RemoveAll() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockJourneyRepository) GetQueueing() ([]*domain.Journey, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Journey), args.Error(1)
}

func (m *MockJourneyRepository) Update(car *domain.Journey) error {
	args := m.Called(car)
	return args.Error(0)
}

func (m *MockJourneyRepository) FindByID(id int) (*domain.Journey, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Journey), args.Error(1)
}

func (m *MockJourneyRepository) RemoveByID(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
