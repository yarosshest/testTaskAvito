package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
	"testTaskAvito/db"
	"testTaskAvito/docs"
)

var con = db.Open("postgresql://postgres:postgres@localhost/testTaskAvito")

// @Summary Add segment
// @Description Add segment
// @ID add-segment
// @Accept  json
// @Produce  json
// @Param segment query string true "segment name"
// @Success 201 "ok"
// @Failure 400 "bad request"
// @Failure 409 "segment with this name alredy exist"
// @Failure 500 "internal server error"
// @Router /Segment [post]
func addSegment(c *gin.Context) {
	var postSegment db.Segment
	if err := c.BindJSON(&postSegment); err != nil {
		segmentMarshalled, _ := json.Marshal(db.Segment{Name: "test"})
		c.NegotiateFormat(string(segmentMarshalled))
		return
	}

	_, err := db.CreateSegment(con, postSegment)
	if err != nil {
		if errors.Is(err, db.NonUniqueFiledErr) {
			c.Writer.WriteHeader(http.StatusConflict)
			return
		}
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
	return
}

func delSegment(c *gin.Context) {
	var postSegment db.Segment
	if err := c.BindJSON(&postSegment); err != nil {
		segmentMarshalled, _ := json.Marshal(db.Segment{Name: "test"})
		c.NegotiateFormat(string(segmentMarshalled))
		return
	}

	err := db.DeleteSegment(con, postSegment)
	if err != nil {
		if errors.Is(err, db.NotFoundErr) {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusAccepted)
	return
}

func putUser(c *gin.Context) {
	p := c.Query("id")
	id, err := strconv.Atoi(p)

	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var queue db.QueueUpdateUser
	if err := c.BindJSON(&queue); err != nil {
		userMarshalled, _ := json.Marshal(db.QueueUpdateUser{Add: []string{"test1"}, Dell: []string{"test2"}})
		c.NegotiateFormat(string(userMarshalled))
		return
	}

	_, err = db.UpdateUser(con, id, queue)

	if err != nil {
		if errors.Is(err, db.NotFoundErr) {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusAccepted)
	return
}

func getUser(c *gin.Context) {

	p := c.Query("id")
	id, err := strconv.Atoi(p)

	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	segments, err := db.GetUserSegments(con, id)

	if err != nil {
		if errors.Is(err, db.NotFoundErr) {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, segments)
	return
}

func RunApi() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.POST("/Segment", addSegment)
	router.DELETE("/Segment", delSegment)

	router.PUT("/User", putUser)
	router.GET("/User", getUser)

	router.Run("localhost:8080")
}
