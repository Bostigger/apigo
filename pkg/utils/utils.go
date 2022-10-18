package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DigestBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		err := json.Unmarshal([]byte(body), x)
		if err != nil {
			return
		}
	}
}
