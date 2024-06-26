package repository

import (
	"errors"
	"log"
	"time"

	"github.com/2110366-2566-2/Mai-Roi-Ra/backend/models"
	st "github.com/2110366-2566-2/Mai-Roi-Ra/backend/pkg/struct"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

type IEventRepository interface {
	CreateEvent(req *models.Event) (*st.CreateEventResponse, error)
	GetEventLists(req *st.GetEventListsRequest) ([]*models.Event, *int, *int, error)
	GetEventListsByStartDate(endDate string) ([]*models.Event, error)
	GetEventDataById(string) (*models.Event, error)
	UpdateEvent(req *models.Event) (*st.MessageResponse, error)
	DeleteEventById(req *st.EventIdRequest) (*st.MessageResponse, error)
	GetAdminAndOrganizerEventById(eventId string) (*string, *string, error)
	VerifyEvent(req *st.VerifyEventRequest) (*st.MessageResponse, error)
}

func NewEventRepository(
	db *gorm.DB,
) IEventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) CreateEvent(req *models.Event) (*st.CreateEventResponse, error) {
	log.Println("[Repo: CreateEvent]: Called")
	trans := r.db.Begin().Debug()
	if err := trans.Create(&req).Error; err != nil {
		trans.Rollback()
		log.Println("[Repo: CreateEvent]: Insert data in Events table error:", err)
		return nil, err
	}
	if err := trans.Commit().Error; err != nil {
		trans.Rollback()
		log.Println("[Repo: CreateEvent]: Call orm DB Commit error:", err)
		return nil, err
	}
	return &st.CreateEventResponse{
		EventId: req.EventId,
	}, nil
}

func (r *EventRepository) UpdateEvent(req *models.Event) (*st.MessageResponse, error) {
	log.Println("[Repo: UpdateEvent]: Called")

	// Find the event by event_id
	var modelEvent models.Event
	if err := r.db.Where("event_id = ?", req.EventId).First(&modelEvent).Error; err != nil {
		log.Println("[Repo: UpdateEvent] event_id not found")
		return nil, err
	}

	// Update the fields
	if !req.StartDate.IsZero() {
		modelEvent.StartDate = req.StartDate
	}

	if !req.EndDate.IsZero() {
		modelEvent.EndDate = req.EndDate
	}

	if req.Status != "" {
		modelEvent.Status = req.Status
	}

	if req.ParticipantFee != 0.0 {
		modelEvent.ParticipantFee = req.ParticipantFee
	}

	if req.Description != "" {
		modelEvent.Description = req.Description
	}

	if req.EventName != "" {
		modelEvent.EventName = req.EventName
	}

	if !req.Deadline.IsZero() {
		modelEvent.Deadline = req.Deadline
	}

	if req.Activities != "" {
		modelEvent.Activities = req.Activities
	}

	if req.EventImage != nil {
		modelEvent.EventImage = req.EventImage
	}

	modelEvent.UpdatedAt = time.Now()

	// Save the updated version
	if err := r.db.Save(&modelEvent).Error; err != nil {
		log.Println("[Repo: UpdateEventInformation] Error updating event in the database:", err)
		return nil, err
	}

	return &st.MessageResponse{
		Response: req.EventId,
	}, nil
}

func (r *EventRepository) DeleteEventById(req *st.EventIdRequest) (*st.MessageResponse, error) {
	log.Println("[Repo: DeleteEventById]: Called")
	eventModel := models.Event{}

	// Delete the event from the database
	if result := r.db.Where("event_id = ?", req.EventId).First(&eventModel); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("[Repo: DeleteEventById] no records found")
			return nil, result.Error
		} else {
			log.Println("[Repo: DeleteEventById] something wrong when deleting from database:", result.Error)
			return nil, result.Error
		}
	} else {
		if err := r.db.Delete(&eventModel).Error; err != nil {
			log.Println("[Repo: DeleteEventById] errors when delete from database")
			return nil, err
		}
	}

	// Return a success message
	return &st.MessageResponse{
		Response: "success",
	}, nil
}

func (r *EventRepository) GetEventLists(req *st.GetEventListsRequest) ([]*models.Event, *int, *int, error) {
	log.Println("[Repo: GetEventLists] Called")
	var eventLists []*models.Event

	query := r.db.Model(&models.Event{})

	if req.OrganizerId != "" {
		query = query.Where(`organizer_id = ?`, req.OrganizerId)
	}

	if req.Filter != "" {
		query = query.Where(`status=?`, req.Filter)
	}

	if req.Search != "" {
		search := "%" + req.Search + "%"
		query = query.Where(`event_name ILIKE ? OR description ILIKE ?`, search, search)
	}

	if req.Sort != "" {
		query = query.Order(req.Sort)
	} else {
		query = query.Order("start_date ASC")
	}

	// Pagination logic
	offset := req.Offset
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	totalEventsQuery := query
	var totalEventsInt64 int64
	if err := totalEventsQuery.Count(&totalEventsInt64).Error; err != nil {
		log.Println("[Repo: GetEventLists]: cannot count the events:", err)
		return nil, nil, nil, err
	}

	query = query.Offset(offset).Limit(limit)

	if err := query.Find(&eventLists).Error; err != nil {
		log.Println("[Repo: GetEventLists]: cannot query the events:", err)
		return nil, nil, nil, err
	}

	totalEvents := int(totalEventsInt64)
	totalPages := int(totalEvents) / limit

	if int(totalEvents)%limit > 0 {
		totalPages++
	}

	return eventLists, &totalEvents, &totalPages, nil
}

func (r *EventRepository) GetEventListsByStartDate(startDate string) ([]*models.Event, error) {
	log.Println("[Repo: GetEventListsByStartDate] Called")

	var eventLists []*models.Event

	// Find events where start_date is equal to the input startDate
	if err := r.db.Where("start_date = ?", startDate).Find(&eventLists).Error; err != nil {
		log.Println("[Repo: GetEventListsByStartDate] Error querying the events:", err)
		return nil, err
	}
	return eventLists, nil
}

func (r *EventRepository) GetEventDataById(eventId string) (*models.Event, error) {
	log.Println("[Repo: GetEventDataById]: Called")
	var event models.Event
	if err := r.db.Where(`event_id=?`, eventId).Find(&event).Error; err != nil {
		log.Println("[Repo: GetEventDataById]: cannot find event_id:", err)
		return nil, err
	}
	return &event, nil
}

func (r *EventRepository) GetAdminAndOrganizerEventById(eventId string) (*string, *string, error) {
	log.Println("[Repo: GetAdminAndOrganizerEventById]: Called")
	var eventModel models.Event
	if err := r.db.Where(`event_id = ?`, eventId).Find(&eventModel).Error; err != nil {
		log.Println("[Repo: GetAdminAndOrganizerEventById]: cannot find event_id:", err)
		return nil, nil, err
	}
	return &eventModel.AdminId, &eventModel.OrganizerId, nil
}

func (r *EventRepository) VerifyEvent(req *st.VerifyEventRequest) (*st.MessageResponse, error) {
	log.Println("[Repo: VerifyEvent]: Called")
	if err := r.db.Model(&models.Event{}).Where("event_id = ?", req.EventId).
		Updates(map[string]interface{}{
			"AdminId":   req.AdminId,
			"Status":    req.Status,
			"UpdatedAt": time.Now(),
		}).Error; err != nil {
		log.Println("[Repo: UpdateEvent] Error updating event in Events table:", err)
		return nil, err
	}
	res := &st.MessageResponse{
		Response: "Verify Event Successful",
	}
	return res, nil
}
