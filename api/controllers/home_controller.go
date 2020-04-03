package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zhuangalbert/app-go/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Hello World")
}

func (server *Server) Getcommodity(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	commodity := make([]responses.Commodity, 0)
	json.Unmarshal(body, &commodity)

	responses.JSON(w, http.StatusOK, commodity)
}

func (server *Server) GetArea(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/option_area")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return

	}

	area := make([]responses.Area, 0)
	json.Unmarshal(body, &area)

	responses.JSON(w, http.StatusOK, area)
}

func (server *Server) GetSize(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/option_size")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return

	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	size := make([]responses.Size, 0)
	json.Unmarshal(body, &size)

	responses.JSON(w, http.StatusOK, size)
}
