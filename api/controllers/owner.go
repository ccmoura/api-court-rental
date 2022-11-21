package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"api-court-rental/api/models"
	"api-court-rental/api/responses"
)

func (server *Server) CreateOwner(w http.ResponseWriter, r *http.Request) {
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	owner := models.Owner{}
	err = json.Unmarshal(buffer, &owner)

	err = owner.FindDuplicatedEmail(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusConflict, err)
		return
	}

	err = owner.FindDuplicatedCpf(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusConflict, err)
		return
	}

	err = owner.FindDuplicatedPhone(server.DB, "")
	if err != nil {
		responses.ERROR(w, http.StatusConflict, err)
		return
	}

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

	_, err := owner.FindOwnerById(server.DB, id)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	_, err = owner.DeleteOwner(server.DB, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, id)
}

func (server *Server) UpdateOwner(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	owner := models.Owner{}

	_, err := owner.FindOwnerById(server.DB, id)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &owner)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = owner.FindDuplicatedPhone(server.DB, id)
	if err != nil {
		responses.ERROR(w, http.StatusConflict, err)
		return
	}

	owner.Prepare()

	updateOwner, err := owner.UpdateOwner(server.DB, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, updateOwner)
}
