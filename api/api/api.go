package api

import (
	"api/db"
	"api/docs"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
)

var con = db.Open("postgresql://postgres:postgres@postgres/testTaskAvito")

// @Summary Add segment
// @Description Add segment
// @ID add-segment
// @Accept  json
// @Produce  json
// @Param name query string true "segment name"
// @Success 201 "ok"
// @Failure 400 "bad request"
// @Failure 409 "segment with this name alredy exist"
// @Failure 500 "internal server error"
// @Router /Segment [post]
func addSegment(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := db.CreateSegment(con, name)
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

// @Summary Delete segment
// @Description Delete segment
// @ID del-segment
// @Accept  json
// @Produce  json
// @Param name query string true "segment name"
// @Success 202 "ok"
// @Failure 400 "bad request"
// @Failure 404 "segment with this name not found"
// @Failure 500 "internal server error"
// @Router /Segment [delete]
func delSegment(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err := db.DeleteSegment(con, name)
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

// @Summary Update user
// @Description Update user
// @ID put-user
// @Accept  json
// @Produce  json
// @Param id query int true "user id"
// @Param add body db.QueueUpdateUser true "Segments to add and del"
// @Success 202 "ok"
// @Failure 400 "bad request"
// @Failure 404 "user with this id not found"
// @Failure 500 "internal server error"
// @Router /User [put]
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

// @Summary Get user
// @Description Get user
// @ID get-user
// @Accept  json
// @Produce  json
// @Param id query int true "user id"
// @Success 200 {array} string "segments"
// @Failure 400 "bad request"
// @Failure 404 "user with this id not found"
// @Failure 500 "internal server error"
// @Router /User [get]
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

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.POST("/Segment", addSegment)
	router.DELETE("/Segment", delSegment)

	router.PUT("/User", putUser)
	router.GET("/User", getUser)

	router.Run("api:8080")
}
