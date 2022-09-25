package controllers

import (
	"io"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"projects/api-court-rental/api/models"
	"projects/api-court-rental/api/responses"
)

func (server *Server) CreateOwner(w http.ResponseWriter, r *http.Request) {
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
			panic(err)
	}

	owner := models.Owner{}
 	err = json.Unmarshal(buffer, &owner)

	owner.Prepare()
	ownerCreated, err := owner.SaveOwner(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}

	responses.JSON(w, http.StatusCreated, ownerCreated)
}

func (server *Server) DeleteOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	owner := models.Owner{}

	id := vars["id"]

	_, err := owner.DeleteOwner(server.DB, id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, id)
}
