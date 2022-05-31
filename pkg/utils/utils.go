package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(x interface{}, r *http.Request){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return 
		}
	}
}