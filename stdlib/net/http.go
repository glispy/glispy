package net

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/glispy/glispy/types"
)

// HTTPGetRequest makes an HTTP get request
func HTTPGetRequest(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var url types.String
	if url, err = args.GetString(0); err != nil {
		err = fmt.Errorf("invalid key: %v", err)
		return
	}

	var resp *http.Response
	if resp, err = http.Get(string(url)); err != nil {
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
