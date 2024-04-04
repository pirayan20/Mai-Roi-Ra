package services

import (
	"log"

	"github.com/2110366-2566-2/Mai-Roi-Ra/backend/models"
	st "github.com/2110366-2566-2/Mai-Roi-Ra/backend/pkg/struct"
	repository "github.com/2110366-2566-2/Mai-Roi-Ra/backend/repositories"
)

type ResponseService struct {
	// logger            *log.Logger
	RepositoryGateway repository.RepositoryGateway
}

type IResponseService interface {
	CreateResponse(req *st.CreateResponseRequest) (*st.CreateResponseResponse, error)
}

func NewResponseService(
	repoGateway repository.RepositoryGateway,
) IResponseService {
	return &ResponseService{
		RepositoryGateway: repoGateway,
	}
}

func (s *ResponseService) CreateResponse(req *st.CreateResponseRequest) (*st.CreateResponseResponse, error) {
	log.Println("[Service: CreateResponse]: Called")

	ResponseModel := models.Response{
		OrganizerId: req.OrganizerId,
		PostId: req.PostId,
		Detail: req.Detail,
	}

	res, err := s.RepositoryGateway.ResponseRepository.CreateResponse(&ResponseModel)
	if err != nil {
		return nil, err
	}
	return res, nil
}