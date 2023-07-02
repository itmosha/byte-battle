package apiserver

import (
	"byte-battle_backend/internal/app/model"
	"byte-battle_backend/pkg/loggers"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (s *APIserver) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "POST" {

			// Read the request's body

			body, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Invalid data provided"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "register/", http.StatusBadRequest)
				return
			}

			// Decode the request's body into an object

			decodedBody := UserBody{}
			json.Unmarshal(body, &decodedBody)

			// Check if necessary data was provided

			if decodedBody.Username == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Username was not provided"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "register/", http.StatusBadRequest)
				return
			}
			if decodedBody.Email == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Email was not provided"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "register/", http.StatusBadRequest)
				return
			}
			if decodedBody.Password == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Password was not provided"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "register/", http.StatusBadRequest)
				return
			}

			// Check uniqueness of username and email

			if uniqueUsername, err := s.store.User().CheckUniqueValue("username", decodedBody.Username); !uniqueUsername {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "User with the same name already exists"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "register/", http.StatusBadRequest)
				return
			} else if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not query the database"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "register/", http.StatusInternalServerError)
				return
			}

			if uniqueEmail, err := s.store.User().CheckUniqueValue("email", decodedBody.Email); !uniqueEmail {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "User with the same email already exists"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "register/", http.StatusBadRequest)
				return
			} else if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not query the database"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "register/", http.StatusInternalServerError)
				return
			}

			// Encrypt the password

			h := sha256.New()
			h.Write([]byte(decodedBody.Password))
			encrypted_pwd := hex.EncodeToString(h.Sum(nil))

			// Create object instance

			user := &model.User{
				Username:     decodedBody.Username,
				Email:        decodedBody.Email,
				Role:         1,
				EncryptedPwd: encrypted_pwd,
			}

			// Insert object into the database

			user, err = s.store.User().CreateInstance(user)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not query the database"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "register/", http.StatusInternalServerError)
				return
			}

			// Set Created status and return the created user's id

			w.WriteHeader(http.StatusCreated)
			resp, _ := json.Marshal(map[string]int{"id": user.ID})
			w.Write(resp)
		}
	}
}

func (s *APIserver) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else if r.Method == "POST" {

			// Read the request's body

			body, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Invalid data provided"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "login/", http.StatusBadRequest)
				return
			}

			// Decode the request's body into an object

			decodedBody := UserBody{}
			json.Unmarshal(body, &decodedBody)

			// Handle cases where the necessary data was not provided

			if decodedBody.Username == "" && decodedBody.Email == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "No username or password provided"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "login/", http.StatusBadRequest)
				return
			}
			if decodedBody.Password == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Password was not provided"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "login/", http.StatusBadRequest)
				return
			}

			// Encrypt the password

			h := sha256.New()
			h.Write([]byte(decodedBody.Password))
			encryptedPwd := hex.EncodeToString(h.Sum(nil))

			// Create User object instance

			user := &model.User{
				ID:           0,
				Username:     decodedBody.Username,
				Email:        decodedBody.Email,
				Role:         0,
				EncryptedPwd: encryptedPwd,
			}

			// Query the database

			token, err := s.store.User().Login(user)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Could not log in with provided credentials"})
				w.Write(resp)

				loggers.ApiRequestFailure("POST", "login/", http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			resp, _ := json.Marshal(map[string]string{"token": token})
			w.Write(resp)

			loggers.ApiRequestSuccess("POST", "login/", http.StatusOK)
		}
	}
}
