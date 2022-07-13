package controllers

import (
	"encoding/json"
	"net/http"
	"task-peer-api/models"
	u "task-peer-api/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // İstek gövdesi decode edilir, hatalı ise hata döndürülür
	if err != nil {
		u.Respond(w, u.Message(false, "Geçersiz istek. Lütfen kontrol ediniz!"))
		return
	}

	resp := account.Create() // Hesap yaratılır
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // İstek gövdesi decode edilir, hatalı ise hata döndürülür
	if err != nil {
		u.Respond(w, u.Message(false, "Geçersiz istek. Lütfen kontrol ediniz!"))
		return
	}

	resp := models.Login(account.Email, account.Password) // Giriş yapılır
	u.Respond(w, resp)
}
