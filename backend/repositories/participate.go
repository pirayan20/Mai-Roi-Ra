package repository

import (
	"errors"
	"log"
	"time"

	"github.com/2110366-2566-2/Mai-Roi-Ra/backend/models"
	st "github.com/2110366-2566-2/Mai-Roi-Ra/backend/pkg/struct"
	"gorm.io/gorm"
)

// UserRepository represents the repository for the User model.
type ParticipateRepository struct {
	DB *gorm.DB
}

type IParticipateRepository interface {
	RegisterEvent(req *st.RegisterEventRequest) (*st.RegisterEventResponse, error)
	CancelRegisterEvent(req *st.RegisterEventRequest) (*st.RegisterEventResponse, error)
}

// NewUserRepository creates a new instance of the UserRepository.
// oldone-func NewUserRepository(c *gin.Context, db *gorm.DB) *UserRepository {
func NewParticipateRepository(
	DB *gorm.DB,
) IParticipateRepository {
	return &ParticipateRepository{
		DB: DB,
	}
}

func (r *ParticipateRepository) RegisterEvent(req *st.RegisterEventRequest) (*st.RegisterEventResponse, error) {
	participateModel := models.Participate{
		UserId:    req.UserId,
		EventId:   req.EventId,
		CreatedAt: time.Time{},
	}

	trans := r.DB.Begin().Debug()
	if err := trans.Create(&participateModel).Error; err != nil {
		trans.Rollback()
		log.Println("[Repo: RegisterEvent]: Insert data in Participate table error:", err)
		return nil, err
	}

	if err := trans.Commit().Error; err != nil {
		trans.Rollback()
		log.Println("[Repo: CreateUser]: Call orm DB Commit error:", err)
		return nil, err
	}

	message := st.RegisterEventResponse{
		Message: "Registered Successful",
	}

	return &message, nil
}

func (r *ParticipateRepository) CancelRegisterEvent(req *st.RegisterEventRequest) (*st.RegisterEventResponse, error) {
	log.Println("[Repo: CancelRegisterEvent]: Called")
	participateModel := models.Participate{}

	if result := r.DB.Where("event_id = ? AND user_id = ?", req.EventId, req.UserId).First(&participateModel); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("[Repo: CancelRegisterEvent] no records found")
			return nil, result.Error
		} else {
			log.Println("[Repo: CancelRegisterEvent] something wrong when query in database")
		}
	} else {
		if err := r.DB.Delete(&participateModel).Error; err != nil {
			log.Println("[Repo: CancelRegisterEvent] errors when delete from database")
			return nil, err
		}
	}
	// Return a success message
	return &st.RegisterEventResponse{
		Message: "Cancel succesful",
	}, nil
}
