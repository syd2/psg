package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/syd2/psg/db/db"
	"github.com/syd2/psg/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserHandler(q *db.Queries, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"error": "something went wrong!!"}
		params := db.CreateUserParams{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			utils.Json(w, response, http.StatusInternalServerError)
		}
		if params.Username == "" && params.Password == "" {
			response["error"] = "Username or Password cant be empty, enter a username and a password"
			utils.Json(w, response, http.StatusBadRequest)
			return
		}
		params.Username = html.EscapeString(params.Username)
		user, err := q.GetUser(r.Context(), params.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				response["error"] = fmt.Sprintf("user '%s' not found, there is no user with this username", params.Username)
				utils.Json(w, response, http.StatusBadRequest)
				return
			}
			response["error"] = fmt.Sprintf("somthing went wrong  ,  %s", err.Error())
			utils.Json(w, response, http.StatusInternalServerError)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
		if err != nil {
			response["error"] = fmt.Sprintf("Password incorrect!!!!! %s", err.Error())
			utils.Json(w, response, http.StatusBadRequest)
			return
		}
		token, err := utils.GenAuthToken(strconv.Itoa(int(user.ID)), secret)
		if err != nil {
			log.Println(err)
			utils.Json(w, response, http.StatusInternalServerError)
			return
		}

		response = map[string]string{"token": token}
		utils.Json(w, response, http.StatusOK)
		return
	}
}
