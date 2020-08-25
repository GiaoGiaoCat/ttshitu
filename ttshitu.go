package ttshitu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	// Read response data
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Println(err)

	// 	return
	// }
	// Debug: fmt.Printf("%s\n", data)
	// fmt.Printf("%s\n", data)

	var resBody ResponseBody
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		log.Println(err)

		return
	}

	fmt.Printf("%+v", resBody)

	result = resBody.Data.Result

	// fmt.Printf("%+v", resBody)
	// fmt.Printf("%+v", resBody)

	// return resBody.Data.Result, nil
	// result = "111"
	// result, err = decoder(res.Body)
	// if err != nil {
	// 	log.Println(err)

	// 	return
	// }

	return result
}

// func decoder(body io.Reader) (result string, err error) {
// 	var resBody ResponseBody
// 	// Try to decode the request body into the struct. If there is an error,
// 	// respond to the client with the error message and a 400 status code.
// 	err = json.NewDecoder(body).Decode(&resBody)
// 	if err != nil {
// 		log.Println(err)

// 		return "", err
// 	}

// 	return resBody.Data.Result, nil
// }

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
