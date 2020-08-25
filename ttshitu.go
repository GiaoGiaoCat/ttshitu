package ttshitu

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type RequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Image    string `json:"image"`
}

type ResponseBody struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Result string `json:"result"`
		ID     string `json:"id"`
	} `json:"data"`
}

var ErrUsernameNotFound = errors.New("用户名或密码错误")

func GetCode(req *RequestBody) (result string) {
	requestBody, err := json.Marshal(req)
	// Debug: fmt.Println(string(requestBody))
	if err != nil {
		log.Println(err)
		return
	}

	res, err := send(bytes.NewBuffer(requestBody))
	// Check for response error
	if err != nil {
		log.Println(err)
		return
	}
	// Close response body
	defer res.Body.Close()
	result, err = decoder(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	return result
}

func decoder(body io.Reader) (result string, err error) {
	var resBody ResponseBody
	err = json.NewDecoder(body).Decode(&resBody)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if !resBody.Success {
		return "", ErrUsernameNotFound
	}

	return resBody.Data.Result, nil
}

func send(body io.Reader) (*http.Response, error) {
	// Pass context.Background() to SendWithContext
	return sendWithContext(context.Background(), body)
}

// Sending an HTTP request and accepting context.
func sendWithContext(ctx context.Context, body io.Reader) (*http.Response, error) {
	// Change NewRequest to NewRequestWithContext and pass context it
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://api.ttshitu.com/base64", body)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
