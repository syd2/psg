package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"github.com/syd2/psg/db/db"
	"github.com/syd2/psg/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(q *db.Queries, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg_resp := map[string]string{"error": "something went wrong!!"}
		params := db.CreateUserParams{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Fatal(err)
			utils.Json(w, msg_resp, http.StatusInternalServerError)
			return
		}

		if params.Username == "" || params.Password == "" {
			msg_resp["error"] = "Username or Password cant be empty, enter a username and a password"
			utils.Json(w, msg_resp, http.StatusBadRequest)
			return
		}
		if len(params.Password) < 8 {
			msg_resp["error"] = "Password is too short, must be greater than 8 characters"
			utils.Json(w, msg_resp, http.StatusBadRequest)
			return
		}
		password := strings.TrimSpace(params.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			msg_resp["error"] = fmt.Sprintf("failed to hash password: %v", err)
			utils.Json(w, msg_resp, http.StatusInternalServerError)
			return
		}

		params.Password = string(hashedPassword)
		params.Username = html.EscapeString(params.Username)
		user, err := q.CreateUser(r.Context(), params)
		if err != nil {

			if pqError, k := err.(*pq.Error); k {
				switch pqError.Code.Name() {
				case "unique_violation":
					msg_resp["error"] = fmt.Sprintf("the username '%s' already exists, chose another unique username", params.Username)
					utils.Json(w, msg_resp, http.StatusBadRequest)
				}
			}
			msg_resp["error"] = fmt.Sprintf("Error creating user: %v", err)
			utils.Json(w, msg_resp, http.StatusInternalServerError)
			return
		}
		token, err := utils.GenAuthToken(strconv.Itoa(int(user.ID)), secret)
		if err != nil {
			msg_resp["error"] = fmt.Sprintf("Error generating token: %v", err)
			utils.Json(w, msg_resp, http.StatusInternalServerError)
			return
		}
		msg_resp = map[string]string{"message": "user created " + params.Username, "token": token}
		utils.Json(w, msg_resp, http.StatusCreated)
		return

	}
}
