package service

import (
	`context`

	`github.com/stretchr/testify/mock`

	`github.com/sshawnta/TestTransport/models`
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetClientId(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	args := m.Called(context.Background(),request)
	if a, ok := args.Get(0).(models.Response); ok{
		return a, args.Error(1)
	}
	return response, args.Error(1)
}

func (m *MockService) PostOrder(ctx context.Context, request *models.PostClientId) (response models.Response, err error) {
	args := m.Called(context.Background(),request)
	if a, ok := args.Get(0).(models.Response); ok{
		return a, args.Error(1)
	}
	return response, args.Error(1)
}

func (m *MockService) GetCount(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	args := m.Called(context.Background(),request)
	if a, ok := args.Get(0).(models.Response); ok{
		return a, args.Error(1)
	}
	return response, args.Error(1)
}

func (m *MockService) GetOrder(ctx context.Context) (response models.Response, err error) {
	args := m.Called(context.Background())
	if a, ok := args.Get(0).(models.Response); ok{
		return a, args.Error(1)
	}
	return response, args.Error(1)
}

