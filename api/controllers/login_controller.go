package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zhuangalbert/app-go/api/auth"
	"github.com/zhuangalbert/app-go/api/models"
	"github.com/zhuangalbert/app-go/api/responses"
	"github.com/zhuangalbert/app-go/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, res)
}

func (server *Server) SignIn(username, password string) (responses.Authentication, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return responses.Authentication{}, err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return responses.Authentication{}, err
	}

	token, err := auth.CreateToken(user.ID)

	result_response := responses.Authentication{
		Username:    user.Username,
		Name:        user.Name,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		AccessToken: token,
		ExpiredIn:   3600,
	}

	return result_response, err
}
