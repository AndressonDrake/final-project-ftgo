package helper

import (
	"fmt"
	// "strings"

	"github.com/go-resty/resty/v2"
)

var (
	RClient    = resty.New()
	XAPIKey    string
	URLMAILING string
)

func SendMail(request map[string]interface{}) (err error) {

	req := RClient.R()

	message := request["message"]

	to := request["email"]
	subject := "Invoice"

	requestBody := map[string]interface{}{
		"message": message,
		"to":      to,
		"subject": subject,
	}

	var response map[string]interface{}

	req.SetHeader("Contetn-Type", "application/json")
	req.SetHeader("Accept", "application/json")
	req.SetHeader("x-api-key", XAPIKey)

	req.SetBody(requestBody)
	req.SetResult(&response)
	req.SetError(&response)

	resp, err := req.Post(URLMAILING)

	if err != nil {
		return
	}

	if resp.StatusCode() != 200 {
		err = fmt.Errorf(resp.Status())
		return
	}

	return
}
