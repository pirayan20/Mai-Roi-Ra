package routers

import (
	"log"
	"net/http"

	controllers "github.com/2110366-2566-2/Mai-Roi-Ra/backend/controllers"
	st "github.com/2110366-2566-2/Mai-Roi-Ra/backend/pkg/struct"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
)

func SetupRouter(c *dig.Container) *gin.Engine {
	r := gin.Default()

	// Swagger setup
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Example of setting up a route using the container to resolve dependencies
	err := c.Invoke(func(eventController *controllers.EventController) {
		r.POST("/api/v1/events", func(ctx *gin.Context) {
			var req st.CreateEventRequest
			if err := ctx.ShouldBindJSON(&req); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			eventController.CreateEvent(ctx, &req)
		})
		r.GET("/api/v1/events", func(ctx *gin.Context) {
			req := &st.GetEventListsRequest{}
			eventController.GetEventLists(ctx, req)
		})
		r.GET("/api/v1/events/:id", func(ctx *gin.Context) {
			req := st.GetEventDataByIdRequest{
				EventId: ctx.Param("id"),
			}
			eventController.GetEventDataById(ctx, req)
		})
	})

	if err != nil {
		log.Println(err)
		return nil
	}

	err = c.Invoke(func(locationController *controllers.LocationController) {
		r.GET("/api/v1/locations/:id", func(ctx *gin.Context) {
			req := st.GetLocationByIdRequest{
				LocationId: ctx.Param("id"),
			}
			locationController.GetLocationById(ctx, req)
		})
	})

	if err != nil {
		log.Println(err)
		return nil
	}

	err = c.Invoke(func(testController *controllers.TestController) {
		r.GET("/api/v1/test", func(ctx *gin.Context) {
			// var req st.CreateEventRequest
			// if err := ctx.ShouldBindJSON(&req); err != nil {
			// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			// 	return
			// }
			testController.GetTest(ctx)
		})
	})

	if err != nil {
		log.Println(err)
		return nil
	}

	return r
}
