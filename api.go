package fsspgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const (
	host = `https://api-ip.fssp.gov.ru/api/v1.0`
)

type API struct {
	token string
}

func NewAPI(token string) *API {
	return &API{token: token}
}

type Response struct {
	Status    string          `json:"status"`
	Code      uint            `json:"code"`
	Exception string          `json:"exception"`
	Response  json.RawMessage `json:"response"`
}

func (r *Response) CheckError() error {
	if r.Code != 0 {
		return errors.New(strconv.FormatUint(uint64(r.Code), 10) + ": " + r.Exception)
	}

	return nil
}

func (r *Response) Task() (string, error) {
	task := new(
		struct {
			Task string `json:"task"`
		},
	)

	if err := json.Unmarshal(r.Response, task); err != nil {
		return "", err
	}

	return task.Task, nil
}

func (api *API) Single(param singleParam) (*Response, error) {
	data := new(Response)

	resp, err := http.Get(param.formatSingleParams(api.token))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

type groupRequest struct {
	Token   string         `json:"token"`
	Request []innerRequest `json:"request"`
}

func (api *API) Group(params ...groupParam) (*Response, error) {
	data := new(Response)
	req := groupRequest{
		Token:   api.token,
		Request: []innerRequest{},
	}

	for _, param := range params {
		req.Request = append(req.Request, param.formatGroupParams())
	}

	buf := bytes.NewBuffer(nil)

	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return nil, err
	}

	resp, err := http.Post(host+"/search/group", "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}
