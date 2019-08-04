package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/caioeverest/ingresso-watcher/config"
)

func GetEventById(conf *config.Config, id string) ([]interface{}, error) {

	url, err := selectApi(id)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", conf.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err := successStatusCode(res.StatusCode); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(body, &decoded); err != nil {
		return nil, err
	}

	presentations := decoded["data"].(map[string]interface{})["presentations"]
	items := presentations.(map[string]interface{})["items"].([]interface{})

	return items, nil
}

func successStatusCode(code int) error {
	if code >= 200 && code <= 299 {
		return nil
	}
	return errors.New(fmt.Sprintf("Request receive HTTP status %d", code))
}

func selectApi(code string) (string, error) {
	spliter := strings.Split(code, "-")
	if len(spliter) < 2 {
		return "", errors.New(fmt.Sprintf("Unable to find api ref"))
	}

	apiRef := spliter[1]
	id := spliter[0]

	switch apiRef {
	case "1":
		return fmt.Sprintf("https://bff-sales-api-cdn.ingressorapido.com.br/api/v1/events/%s", id), nil
	case "2":
		return fmt.Sprintf("https://bff-sales-api-cdn.bileto.sympla.com.br/api/v1/events/%s", id), nil
	default:
		return "", errors.New(fmt.Sprintf("Api ref not acceptable %s", apiRef))
	}
}
