package member_repository_interface

import "coba-firebase-golang/app/database/model"

type Member_Repository_Interface interface {
	Create(member *model.Member) error
	GetAll() ([]model.Member, error)
	GetById(ID string) (*model.Member, error)
	Update(ID string, member *model.Member) error
	Delete(ID string) error
}
