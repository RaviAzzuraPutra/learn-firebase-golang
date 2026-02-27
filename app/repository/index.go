package repository

import (
	"coba-firebase-golang/app/database/connect"
	"coba-firebase-golang/app/database/model"
	"context"

	"cloud.google.com/go/firestore"
)

type Member_Repository struct {
	DB *firestore.Client
}

func NewMemberRepositoryRegistry() *Member_Repository {
	return &Member_Repository{
		DB: connect.FirestoreClient,
	}
}

func (repo *Member_Repository) Create(member *model.Member, ctx context.Context) error {

	_, errCreate := repo.DB.Collection("members").Doc(member.ID).Set(ctx, member)

	return errCreate

}

func (repo *Member_Repository) GetAll(ctx context.Context) ([]model.Member, error) {

	iter := repo.DB.Collection("members").Documents(ctx)

	var members []model.Member

	for {
		doc, err := iter.Next()

		if err != nil {
			break
		}

		var member model.Member

		doc.DataTo(&member)

		members = append(members, member)
	}

	return members, nil

}

func (repo *Member_Repository) GetById(ID string, ctx context.Context) (*model.Member, error) {

	doc, err := repo.DB.Collection("members").Doc(ID).Get(ctx)

	if err != nil {
		return nil, err
	}

	var member model.Member

	doc.DataTo(&member)

	return &member, nil

}

func (repo *Member_Repository) Update(ID string, member *model.Member, ctx context.Context) error {

	_, err := repo.DB.Collection("members").Doc(ID).Set(ctx, member)

	return err

}

func (repo *Member_Repository) Delete(ID string, ctx context.Context) error {

	_, err := repo.DB.Collection("members").Doc(ID).Delete(ctx)

	return err

}
