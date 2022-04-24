package util

import (
	"io/ioutil"
	"net/http"
)

func GetQueryParameterFromRequest(r *http.Request) map[string]string {
	queryParams := make(map[string]string)
	r.ParseForm()

	for key, value := range r.Form {
		queryParams[key] = value[0]
	}
	return queryParams
}

func WriteResponseMessage(w http.ResponseWriter, status int, responseText []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if len(responseText) > 0 {
		w.Write(responseText)
	}
}

func ReadReqBodyAsStruct(r *http.Request, inter interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	ByteToStruct(data, &inter)
	return nil
}
