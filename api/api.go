package api

import (
	"awesomeProject/db"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var con = db.Open("postgresql://postgres:postgres@localhost/testTaskAvito")

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

func RunApi() {
	router := gin.Default()
	router.POST("/Segment", addSegment)
	router.DELETE("/Segment", delSegment)

	router.Run("localhost:8080")
}
