package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"api-court-rental/api/auth"
	"api-court-rental/api/models"
	"api-court-rental/api/responses"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	AccessToken string
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	owner := models.Owner{}
	err = json.Unmarshal(body, &owner)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	owner.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(owner.Email, owner.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	credentials := Credentials{AccessToken: token}
	responses.JSON(w, http.StatusOK, credentials)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	owner := models.Owner{}

	err = server.DB.Debug().Model(models.Owner{}).Where("email = ?", email).Take(&owner).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(owner.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(owner.ID)
}
