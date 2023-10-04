package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/d-mittal-21/GymkhanaCalendar/Backend_Go/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateEventHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)

	client, err := createClient()
	if err != nil {
		log.Printf("error creating client: %v", err)
		return
	}

	router := gin.New()

	router.POST("/api/event", CreateEventHandler(client))

	payload := []byte(`{
		"title": "Example Event",
		"desc": "This is an example event",
		"start": "2023-06-27T10:00:00Z",
		"end": "2023-06-27T12:00:00Z",
		"venue": "Example Venue",
		"link": "https://example.com"
	}`)

	req, _ := http.NewRequest("POST", "/api/event", strings.NewReader(string(payload)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseEvent types.Event
	err2 := json.Unmarshal(w.Body.Bytes(), &responseEvent)
	assert.nil(t, err2)

	assert.Equal(t, "Example Event", responseEvent.Title)
	assert.Equal(t, "This is an example event", responseEvent.Description)
	assert.Equal(t, "Example Venue", responseEvent.Venue)
	assert.Equal(t, "https://example.com", responseEvent.Link)

}
