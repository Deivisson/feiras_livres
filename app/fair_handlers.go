package app

import (
	"net/http"

	"github.com/Deivisson/feiras_livres/service"
)

type FairHandlers struct {
	service service.FairService
}

func (fh *FairHandlers) create(w http.ResponseWriter, r *http.Request) {
	bodyParams, err := getBodyParams(r)
	if err != nil {
		return
	}

	if fair, err := fh.service.Create(bodyParams); err != nil {
		writeResponse(w, err.Code, err.ToMessage())
	} else {
		writeResponse(w, http.StatusOK, fair)
	}
}

func (fh *FairHandlers) update(w http.ResponseWriter, r *http.Request) {
	id := getResourceParam(r, "id")
	bodyParams, err := getBodyParams(r)
	if err != nil {
		return
	}

	if fair, err := fh.service.Update(bodyParams, id); err != nil {
		writeResponse(w, err.Code, err.ToMessage())
	} else {
		writeResponse(w, http.StatusOK, fair)
	}
}

func (fh *FairHandlers) search(w http.ResponseWriter, r *http.Request) {
	bodyParams, err := getBodyParams(r)
	if err != nil {
		return
	}

	if fairs, err := fh.service.Search(bodyParams); err != nil {
		writeResponse(w, err.Code, err.ToMessage())
	} else {
		writeResponse(w, http.StatusOK, fairs)
	}
}

func (fh *FairHandlers) getById(w http.ResponseWriter, r *http.Request) {
	id := getResourceParam(r, "id")

	if fair, err := fh.service.GetById(id); err != nil {
		writeResponse(w, err.Code, err.ToMessage())
	} else {
		writeResponse(w, http.StatusOK, fair)
	}
}

func (fh *FairHandlers) delete(w http.ResponseWriter, r *http.Request) {
	id := getResourceParam(r, "id")

	if err := fh.service.Delete(id); err != nil {
		writeResponse(w, err.Code, err.ToMessage())
	} else {
		writeResponse(w, http.StatusOK, map[string]string{
			"id": id,
		})
	}
}
