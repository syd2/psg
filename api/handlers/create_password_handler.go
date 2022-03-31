package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"

	"github.com/syd2/psg/db/db"
	"github.com/syd2/psg/utils"
)

func main() {
	fmt.Println("")
}

func CreatePasswordHandler(q *db.Queries, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"error": "something went wrong!"}
		params := db.CreatePasswordParams{}
		ctx := r.Context()
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Fatal(err)
			utils.Json(w, resp, http.StatusInternalServerError)
			return
		}

		if params.AppName == "" || params.AppPassword == "" {
			resp["error"] = "the app name and password cant be empty"
			utils.Json(w, resp, http.StatusBadRequest)
			return
		}

		params.AppName = html.EscapeString(params.AppName)
		params.AppPassword = strings.TrimSpace(params.AppPassword)
		params.UserID = int64(ctx.Value("userid").(int))
		password, err := q.CreatePassword(ctx, params)

		if err != nil {
			if pqError, k := err.(*pq.Error); k {
				switch pqError.Code.Name() {
				case "unique_violation":
					resp["error"] = fmt.Sprintf("this app '%s' already has a password", params.AppName)
					utils.Json(w, resp, http.StatusBadRequest)
				}
			}
			resp["error"] = fmt.Sprintf("Error creating user: %v", err)
			utils.Json(w, resp, http.StatusInternalServerError)
			return
		}
		res := map[string]string{"message": fmt.Sprintf("your new password created successfully")}
		res_2 := fmt.Sprintf("app name : %v ,  password : %v ", password.AppName, password.AppPassword)
		utils.Json(w, res, http.StatusCreated)
		utils.Json(w, res_2, http.StatusCreated)
		return

	}
}
