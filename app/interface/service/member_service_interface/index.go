package member_service_interface

import (
	"coba-firebase-golang/app/database/model"
	"coba-firebase-golang/app/request"
)

type Member_Service_Interface interface {
	Create(request *request.Member_Request) (*model.Member, error)
	GetAll() ([]model.Member, error)
	GetById(ID string) (*model.Member, error)
	Update(ID string, request request.Member_Request) (*model.Member, error)
	Delete(ID string) error
}
