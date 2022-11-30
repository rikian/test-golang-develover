package http

import (
	"bytes"
	"encoding/json"
	"go/service1/config"
	"log"
	"net/http"
	"net/url"
)

var client = &http.Client{}

type HttpResponse struct {
	Message string
}

type HttpImpl struct {
	r *HttpResponse
}

func NewHttpRequest() HttpRequest {
	return &HttpImpl{
		r: &HttpResponse{},
	}
}

func (h *HttpImpl) CreateDirectoryImage(id string) (*HttpResponse, error) {
	formUserId := url.Values{}
	formUserId.Set("id", id)
	formUserId.Set("token", "12345")
	payload := bytes.NewBuffer([]byte(formUserId.Encode()))
	request, err := http.NewRequest("POST", config.HttpAddress+"/user/create-directory-image/", payload)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(h.r)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return h.r, nil
}

func (h *HttpImpl) DeleteImage(userId, productId string) (*HttpResponse, error) {
	formDeleteImageProduct := url.Values{}
	formDeleteImageProduct.Set("userId", userId)
	formDeleteImageProduct.Set("productId", productId)
	formDeleteImageProduct.Set("token", "12345")
	payload := bytes.NewBuffer([]byte(formDeleteImageProduct.Encode()))
	request, err := http.NewRequest("POST", config.HttpAddress+"/user/delete-directory-image/", payload)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(h.r)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return h.r, nil
}
