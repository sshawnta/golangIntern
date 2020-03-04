package service

import (
	`context`
	`fmt`

	`github.com/sshawnta/TestTransport/models`
)

type Service interface {
	GetClientId(ctx context.Context, request *models.GetClientId) (response models.Response, err error)
	PostOrder(ctx context.Context, request *models.PostClientId) (response models.Response, err error)
	GetCount(ctx context.Context, request *models.GetClientId) (response models.Response, err error)
	GetOrder(ctx context.Context) (response models.Response, err error)
}

type service struct {
}

func (s *service) GetClientId(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	if request.Id <= 0 {
		response.Error = true
		response.ErrorText = "err"
		err = fmt.Errorf("err")
		return
	}
	response.Data = &models.Data{Res:true}
	return
}

func (s *service) PostOrder(ctx context.Context, request *models.PostClientId) (response models.Response, err error) {
	if request.Id <= 0 {
		response.Error = true
		response.ErrorText = "err"
		err = fmt.Errorf("err")
		return
	}
	response.Data = &models.Data{Res:true}
	return
}

func (s *service) GetCount(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	if request.Id <= 0 {
		response.Error = true
		response.ErrorText = "err"
		err = fmt.Errorf("err")
		return
	}
	response.Data = &models.Data{Res:true}
	return
}

func (s *service) GetOrder(ctx context.Context) (response models.Response, err error) {
	response.Data = &models.Data{Res:true}
	return
}

func NewService() Service {
	return &service{}
}
