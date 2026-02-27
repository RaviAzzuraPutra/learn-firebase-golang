package member_repository_interface

import (
	"coba-firebase-golang/app/database/model"
	"context"
)

type Member_Repository_Interface interface {
	Create(member *model.Member, ctx context.Context) error
	GetAll(ctx context.Context) ([]model.Member, error)
	GetById(ID string, ctx context.Context) (*model.Member, error)
	Update(ID string, member *model.Member, ctx context.Context) error
	Delete(ID string, ctx context.Context) error
}
