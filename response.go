package fsspgo

import (
	"encoding/json"
	"errors"
	"strconv"
)

type response struct {
	Status    string          `json:"status"`
	Code      uint            `json:"code"`
	Exception string          `json:"exception"`
	Response  json.RawMessage `json:"response"`
}

func (r *response) checkError() error {
	if r.Code != 0 {
		return errors.New(strconv.FormatUint(uint64(r.Code), 10) + ": " + r.Exception)
	}

	return nil
}

func (r *response) task() (string, error) {
	task := new(struct {
		Task string `json:"task"`
	})

	if err := json.Unmarshal(r.Response, task); err != nil {
		return "", err
	}

	return task.Task, nil
}

type status uint8

const (
	StatusDone status = iota
	StatusWork
	StatusWait
	StatusErr
)

func (r *response) status() (status, error) {
	status := new(struct {
		Status   status `json:"status"`
		Progress string `json:"progress"`
	})

	if err := json.Unmarshal(r.Response, status); err != nil {
		return StatusErr, err
	}

	return status.Status, nil
}

type Result struct {
	Name       string `json:"name"`
	Production string `json:"exe_production"`
	Document   string `json:"details"`
	Subject    string `json:"subject"`
	Department string `json:"department"`
	Bailiff    string `json:"bailiff"`
	End        string `json:"ip_end"`
}

type resultData struct {
	Status    int    `json:"status"`
	TaskStart string `json:"task_start"`
	TaskEnd   string `json:"task_end"`
	Result    []struct {
		Status requestResponseType `json:"status"`
		Query  json.RawMessage     `json:"query"`
		Result []Result            `json:"result"`
	} `json:"result"`
}

func (r *response) result() ([]Result, error) {
	data := new(resultData)

	if err := json.Unmarshal(r.Response, data); err != nil {
		return nil, err
	}

	results := make([]Result, 0)

	for _, elem := range data.Result {
		results = append(results, elem.Result...)
	}

	return results, nil
}
