package jsonutils

import (
	"encoding/json"
	"net/http"
)

func JsonDecode(res *http.Response, item interface{}) error {
	if err := json.NewDecoder(res.Body).Decode(&item); err != nil {
		return err
	}
	return nil
}
