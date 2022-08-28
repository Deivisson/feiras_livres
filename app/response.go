package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func getResourceParam(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}

func getBodyParams(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}
