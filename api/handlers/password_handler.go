package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/syd2/psg/db/db"
	"github.com/syd2/psg/utils"
)

func main() {
	fmt.Println("")
}

func CreatePasswordHandler(q db.Queries, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"error": "something went wrong!"}
		params := db.CreatePasswordParams{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Fatal(err)
			utils.Json(w, resp, http.StatusInternalServerError)
		}
	}
}
