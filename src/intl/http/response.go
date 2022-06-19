package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response http.Response

func (resp *Response) Unmarshal(data interface{}) error {
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}
	return nil
}
