package controllers

import (
	"URLShortener/DAO"
	models "URLShortener/Models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	err = UserRegister(requestData.FirstName, requestData.LastName, requestData.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario registrado con éxito"))
}

func UserRegister(FirstName, LastName, Email string) error {
	userID := uuid.New().String()
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "TuApp",
		AccountName: Email,
	})
	if err != nil {
		return err
	}

	newUser := models.User{
		ID:        userID,
		FirstName: FirstName,
		LastName:  LastName,
		Email:     Email,
		CreatedAt: time.Now(),
		Secret:    secret.Secret(), // Guardar el secreto para la generación futura de TOTP.
	}

	dao, _ := DAO.NewURLDao()

	err = dao.CreateUser(&newUser)
	if err != nil {

		return err
	}

	// TODO: Enviar correo electrónico al usuario con instrucciones para configurar TOTP.

	return nil
}
