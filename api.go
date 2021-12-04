package fsspgo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
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

func (api *API) Single(param SingleParam) (*Response, error) {
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

func (api *API) Group(params ...GroupParam) (*Response, error) {
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

func (api *API) Status(task string) (*Response, error) {
	params := make(url.Values)

	params.Add("token", api.token)
	params.Add("task", task)

	path, err := url.Parse(host + "/status")
	if err != nil {
		panic(err)
	}

	path.RawQuery = params.Encode()

	data := new(Response)

	resp, err := http.Get(path.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

func (api *API) Result(task string) (*Response, error) {
	params := make(url.Values)

	params.Add("token", api.token)
	params.Add("task", task)

	path, err := url.Parse(host + "/result")
	if err != nil {
		panic(err)
	}

	path.RawQuery = params.Encode()

	data := new(Response)

	resp, err := http.Get(path.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}
