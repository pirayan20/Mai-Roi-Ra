package controllers

import (
	"log"
	"net/http"

	st "github.com/2110366-2566-2/Mai-Roi-Ra/backend/pkg/struct"
	"github.com/2110366-2566-2/Mai-Roi-Ra/backend/services"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	ServiceGateway services.ServiceGateway
}

func NewEventController(
	sg services.ServiceGateway,
) *EventController {
	return &EventController{
		ServiceGateway: sg,
	}
}

// CreateEvent creates a new event.
// @Summary Create new event
// @Description Create a new event with the provided details.
// @Tags events
// @Accept json
// @Produce json
// @Param request body structure.CreateEventRequest true "Create Event Request"
// @Success 200 {object} structure.CreateEventResponse
// @Failure 400 {object} object "Bad Request"
// @Failure 500 {object} object "Internal Server Error"
// @Router /events [post]
func (c *EventController) CreateEvent(ctx *gin.Context, req *st.CreateEventRequest) {
	log.Println("[CTRL: CreateEvent] Input:", req)

	res, err := c.ServiceGateway.EventService.CreateEvent(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("[CTRL: CreateEvent] Output:", res)
	ctx.JSON(http.StatusOK, res)
}

// UpdateEvent updates an existing event.
// @Summary Update existing event
// @Description Update an existing event with the provided details.
// @Tags events
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID" example:"event123"
// @Param request body structure.UpdateEventRequest true "Update Event Request"
// @Success 200 {object} structure.UpdateEventResponse
// @Failure 400 {object} object "Bad Request"
// @Failure 500 {object} object "Internal Server Error"
// @Router /events/{event_id} [put]
func (c *EventController) UpdateEvent(ctx *gin.Context, req *st.UpdateEventRequest) {
	res, err := c.ServiceGateway.EventService.UpdateEvent(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("[CTRL: UpdateEvent] Output:", res)
	ctx.JSON(http.StatusOK, res)
}

// DeleteEventById deletes an event by its ID.
// @Summary Delete event by ID
// @Description Delete an event with the specified ID.
// @Tags events
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID" example:"event123"
// @Success 200 {object} object "OK"
// @Failure 400 {object} object "Bad Request"
// @Failure 500 {object} object "Internal Server Error"
// @Router /events/{event_id} [delete]
func (c *EventController) DeleteEventById(ctx *gin.Context, req *st.DeleteEventRequest) {
	deleteMessage, err := c.ServiceGateway.EventService.DeleteEventById(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("[CTRL: DeleteEvent] Output:", deleteMessage)
	ctx.JSON(http.StatusOK, gin.H{"message": deleteMessage})
}

// @Summary GetEventLists
// @Description Get list of events
// @Tags events
// @Accept json
// @Produce json
// @Param organizer_id query string false "For My events"
// @Param filter query string false "Status query. i.e. Approved"
// @Param sort query string false "Sort order. i.e. start_date ASC"
// @Param offset query int false "offset i.e. 0"
// @Param limit query int false "Items per page i.e. 10"
// @Success 200 {object} structure.GetEventListsResponse
// @Failure 400 {object} object "Bad Request"
// @Failure 500 {object} object "Internal Server Error"
// @Router /events [get]
func (c *EventController) GetEventLists(ctx *gin.Context, req *st.GetEventListsRequest) {
	log.Println("[CTRL: GetEventLists] Input:", req)
	res, err := c.ServiceGateway.EventService.GetEventLists(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("[CTRL: GetEventLists] Output:", res)
	ctx.JSON(http.StatusOK, res)
}

// @Summary GetEventDataById
// @Description Get a test message
// @Tags events
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Success 200 {object} structure.GetEventDataByIdResponse
// @Failure 400 {object} object "Bad Request"
// @Failure 500 {object} object "Internal Server Error"
// @Router /events/{event_id} [get]
func (c *EventController) GetEventDataById(ctx *gin.Context, req st.GetEventDataByIdRequest) {
	log.Println("[CTRL: GetEventDataById] Input:", req)
	res, err := c.ServiceGateway.EventService.GetEventDataById(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("[CTRL: GetEventDataById] Output:", res)
	ctx.JSON(http.StatusOK, res)
}

// @Summary GetParticipantLists
// @Description Get list of all participant in the events
// @Tags events
// @Accept json
// @Produce json
// @Param event_id query string true "event_id"
// @Param offset query int false "offset i.e. 0"
// @Param limit query int false "Items per page i.e. 10"
// @Success 200 {object} structure.GetParticipantListsResponse
// @Failure 400 {object} object "Bad Request"
// @Failure 500 {object} object "Internal Server Error"
// @Router /events/participant [get]
func (c *EventController) GetParticipantLists(ctx *gin.Context, req *st.GetParticipantListsRequest) {
	log.Println("[CTRL: GetParticipantLists] Input:", req)
	res, err := c.ServiceGateway.EventService.GetParticipantLists(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("[CTRL: GetParticipantLists] Output:", res)
	ctx.JSON(http.StatusOK, res)
}
