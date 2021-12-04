package fsspgo

import (
	"net/url"
	"strconv"
)

type Person struct {
	FirstName  string `json:"firstname"`
	SecondName string `json:"secondname,omitempty"`
	LastName   string `json:"lastname"`
	Birthdate  string `json:"birthdate,omitempty"`
	Region     uint8  `json:"region"`
}

type Legal struct {
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
	Region  uint8  `json:"region"`
}

type Number struct {
	Number string `json:"number"`
}

type SingleParam interface {
	formatSingleParams(token string) string
}

func (p *Person) formatSingleParams(token string) string {
	params := make(url.Values)

	params.Add("token", token)
	params.Add("lastname", p.LastName)
	params.Add("firstname", p.FirstName)
	params.Add("region", strconv.FormatUint(uint64(p.Region), 10))

	if p.SecondName != "" {
		params.Add("secondname", p.SecondName)
	}

	if p.Birthdate != "" {
		params.Add("birthdate", p.Birthdate)
	}

	path, err := url.Parse(host + "/search/physical")
	if err != nil {
		panic(err)
	}

	path.RawQuery = params.Encode()

	return path.String()
}

func (l *Legal) formatSingleParams(token string) string {
	params := make(url.Values)

	params.Add("token", token)
	params.Add("name", l.Name)
	params.Add("region", strconv.FormatUint(uint64(l.Region), 10))

	if l.Address != "" {
		params.Add("address", l.Address)
	}

	path, err := url.Parse(host + "/search/legal")
	if err != nil {
		panic(err)
	}

	path.RawQuery = params.Encode()

	return path.String()
}

func (n *Number) formatSingleParams(token string) string {
	params := make(url.Values)

	params.Add("token", token)
	params.Add("number", n.Number)

	path, err := url.Parse(host + "/search/ip")
	if err != nil {
		panic(err)
	}

	path.RawQuery = params.Encode()

	return path.String()
}

type GroupParam interface {
	formatGroupParams() innerRequest
}

func (p *Person) formatGroupParams() innerRequest {
	return innerRequest{
		Type:   personRequestResponseType,
		Params: p,
	}
}

func (l *Legal) formatGroupParams() innerRequest {
	return innerRequest{
		Type:   legalRequestRequestResponseType,
		Params: l,
	}
}

func (n *Number) formatGroupParams() innerRequest {
	return innerRequest{
		Type:   numberRequestRequestResponseType,
		Params: n,
	}
}
