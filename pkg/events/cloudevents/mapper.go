package cloudevents

import (
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

type CloudEvents struct{}

func NewMapper() *CloudEvents {
	return &CloudEvents{}
}

func (c *CloudEvents) Request(r *http.Request) {}

func (c *CloudEvents) Response(w http.ResponseWriter, statusCode int, data []byte) (int, error) {
	var response cloudevents.Event

	response.SetID(uuid.New().String())
	response.SetSource("knaive-lambda-runtime")
	// response.SetExtension()
	response.SetData(cloudevents.ApplicationJSON, data)

	return w.Write([]byte(response.String()))
}
