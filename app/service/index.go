package service

import (
	"coba-firebase-golang/app/database/model"
	"coba-firebase-golang/app/helper"
	"coba-firebase-golang/app/interface/repository/member_repository_interface"
	"coba-firebase-golang/app/request"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Member_Service struct {
	repository member_repository_interface.Member_Repository_Interface
}

func NewMemberServiceRegistry(Member_Repository member_repository_interface.Member_Repository_Interface) *Member_Service {
	return &Member_Service{
		repository: Member_Repository,
	}
}

func (s *Member_Service) Create(request *request.Member_Request) (*model.Member, error) {

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name Cannot Be Empty!")
	}

	if request.Email == nil || *request.Email == "" {
		return nil, helper.NewBadRequest("Email Cannot Be Empty")
	}

	if request.Phone == nil || *request.Phone == "" {
		return nil, helper.NewBadRequest("Phone Cannot Be Empty")
	}

	NewUUID := uuid.NewString()
	JoinDate := time.Now()
	MemberCode := helper.GenerateRandomNumber()

	member := &model.Member{
		ID:          NewUUID,
		Name:        request.Name,
		Email:       request.Email,
		Phone:       request.Phone,
		Join_Date:   JoinDate,
		Member_Code: MemberCode,
	}

	err := s.repository.Create(member)

	if err != nil {
		return nil, helper.NewInternalServerError("An error occurred while creating new member data. : " + err.Error())
	}

	return member, nil
}

func (s *Member_Service) GetAll() ([]model.Member, error) {

	member, errGet := s.repository.GetAll()

	if errGet != nil {
		return nil, helper.NewInternalServerError("An error occurred while retrieving all member data. : " + errGet.Error())
	}

	return member, nil

}

func (s *Member_Service) GetById(ID string) (*model.Member, error) {

	member, errGet := s.repository.GetById(ID)

	if errGet != nil {

		if status.Code(errGet) == codes.NotFound {
			return nil, helper.NewNotFound("Member with that ID was not found")
		}

		return nil, helper.NewInternalServerError("An error occurred while retrieving member data by id. : " + errGet.Error())
	}

	return member, nil

}

func (s *Member_Service) Update(ID string, request *request.Member_Request) (*model.Member, error) {

	if request.Name == nil || *request.Name == "" {
		return nil, helper.NewBadRequest("Name Cannot Be Empty!")
	}

	if request.Email == nil || *request.Email == "" {
		return nil, helper.NewBadRequest("Email Cannot Be Empty")
	}

	if request.Phone == nil || *request.Phone == "" {
		return nil, helper.NewBadRequest("Phone Cannot Be Empty")
	}

	member, errGet := s.repository.GetById(ID)

	if errGet != nil {

		if status.Code(errGet) == codes.NotFound {
			return nil, helper.NewNotFound("Member with that ID was not found")
		}

		return nil, helper.NewInternalServerError("An error occurred while retrieving member data by id. : " + errGet.Error())
	}

	member.Name = request.Name
	member.Email = request.Email
	member.Phone = request.Phone

	errUpdate := s.repository.Update(ID, member)

	if errUpdate != nil {
		return nil, helper.NewInternalServerError("An error occurred while update member data. : " + errUpdate.Error())
	}

	return member, nil
}

func (s *Member_Service) Delete(ID string) error {

	_, errGet := s.repository.GetById(ID)

	if errGet != nil {

		if status.Code(errGet) == codes.NotFound {
			return helper.NewNotFound("Member with that ID was not found")
		}

		return helper.NewInternalServerError("An error occurred while retrieving member data by id. : " + errGet.Error())
	}

	errDelete := s.repository.Delete(ID)

	if errDelete != nil {
		return helper.NewInternalServerError("An error occurred while delete member data. : " + errDelete.Error())
	}

	return nil

}
