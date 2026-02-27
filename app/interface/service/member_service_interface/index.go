package member_service_interface

import (
	"coba-firebase-golang/app/database/model"
	"coba-firebase-golang/app/request"
	"context"
)

type Member_Service_Interface interface {
	Create(request *request.Member_Request, ctx context.Context) (*model.Member, error)
	GetAll(ctx context.Context) ([]model.Member, error)
	GetById(ID string, ctx context.Context) (*model.Member, error)
	Update(ID string, request *request.Member_Request, ctx context.Context) (*model.Member, error)
	Delete(ID string, ctx context.Context) error
}
