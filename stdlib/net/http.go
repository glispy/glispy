package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/glispy/glispy/types"
)

// HTTPGetRequest makes an HTTP GET request
func HTTPGetRequest(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var (
		hcVal types.Atom
		url   types.String
		hc    *http.Client
	)

	if err = args.GetValues(&hcVal, &url); err != nil {
		return
	}

	if hc, err = getHTTPClient(hcVal); err != nil {
		return
	}

	return httpRequest(hc, "GET", string(url), nil)
}

// HTTPPostRequest makes an HTTP POST request
func HTTPPostRequest(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var (
		hcVal types.Atom
		url   types.String
		value types.Atom

		hc   *http.Client
		body io.Reader
	)

	if err = args.GetValues(&hcVal, &url, &value); err != nil {
		return
	}

	if hc, err = getHTTPClient(hcVal); err != nil {
		return
	}

	if body, err = getBody(value); err != nil {
		err = fmt.Errorf("error marshaling value (%+v) as JSON: %v", value, err)
		return
	}

	return httpRequest(hc, "POST", string(url), body)
}

// HTTPPutRequest makes an HTTP PUT request
func HTTPPutRequest(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var (
		hcVal types.Atom
		url   types.String
		value types.Atom

		hc   *http.Client
		body io.Reader
	)

	if err = args.GetValues(&hcVal, &url, &value); err != nil {
		return
	}

	if hc, err = getHTTPClient(hcVal); err != nil {
		return
	}

	if body, err = getBody(value); err != nil {
		err = fmt.Errorf("error marshaling value (%+v) as JSON: %v", value, err)
		return
	}

	return httpRequest(hc, "PUT", string(url), body)
}

func getHTTPClient(val types.Atom) (hc *http.Client, err error) {
	switch n := val.(type) {
	case *http.Client:
		hc = n
	case types.Nil:
		hc = &http.Client{}

	default:
		err = fmt.Errorf("expected *http.Client or nil, received %T", val)
		return
	}

	return
}

func getBody(value interface{}) (body io.Reader, err error) {
	if value == nil {
		return
	}

	var bs []byte
	if bs, err = json.Marshal(value); err != nil {
		return
	}

	body = bytes.NewReader(bs)
	return
}

func httpRequest(hc *http.Client, method, url string, body io.Reader) (exp types.Expression, err error) {
	var req *http.Request
	if req, err = http.NewRequest(method, url, body); err != nil {
		return
	}

	var resp *http.Response
	if resp, err = hc.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	var m map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return
	}

	exp = m
	return
}
