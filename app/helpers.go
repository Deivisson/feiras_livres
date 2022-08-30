package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Deivisson/feiras_livres/utils/errs"
	"github.com/gorilla/mux"
)

/// TODO: Create a lib/package to these functions
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

func permittedParams(params []byte, values string) ([]byte, *errs.AppError) {
	var unallowed []string
	inputParams := make(map[string]interface{})
	outputParams := make(map[string]interface{})

	json.Unmarshal(params, &inputParams)
	for k, v := range inputParams {
		if pos := strings.Index(values, k); pos == -1 {
			unallowed = append(unallowed, k)
		} else {
			outputParams[k] = v
		}
	}

	if len(unallowed) > 0 {
		log.Println(fmt.Errorf("params not allowed: %s", strings.Join(unallowed, ", ")))
	}

	params, err := json.Marshal(outputParams)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error on get permitted params")
	}
	return params, nil
}
