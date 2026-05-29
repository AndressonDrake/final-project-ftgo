package repository

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"gw-heatlcare.com/domain"
	l "gw-heatlcare.com/helper/logger"
)

var (
	client = resty.New()
)

type coreRepository struct {
	API_KEY  string
	BASE_URL string
}

func CoreRepository(API_KEY, BASE_URL string) domain.CoreRepository {
	return &coreRepository{API_KEY: API_KEY, BASE_URL: BASE_URL}
}

func (c *coreRepository) GetCore(pathUrl string) (response interface{}, err error) {
	baseUrl := c.BASE_URL + pathUrl
	l.Log.Info(l.Fields{
		"hit_core": baseUrl,
	}, nil, baseUrl)

	fmt.Printf(`Hit Core : %s`, baseUrl)
	fmt.Println()

	request := client.R()
	request.SetHeader("api-key", c.API_KEY)
	request.SetResult(&response)
	request.SetError(&response)

	resp, err := request.Get(baseUrl)
	if err != nil {
		l.Log.Error(l.Fields{
			"error": err.Error(),
			"url":   baseUrl,
		}, pathUrl, "fail to request core")
		return
	}

	if resp.StatusCode() != 200 {
		err = fmt.Errorf(resp.Status())
		l.Log.Error(l.Fields{
			"error":     err.Error(),
			"resp_code": resp.StatusCode(),
		}, pathUrl, "fail request code")
		return
	}

	l.Log.Info(l.Fields{
		"response": response,
		"url":      baseUrl,
	}, pathUrl, "response from core")

	return
}
