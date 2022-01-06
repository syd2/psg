package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/syd2/psg/utils"
)

func AuthWrapper(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp := map[string]interface{}{"error": "unauthorized", "message": "missing token"}
			var h = r.Header.Get("Authorization")
			h = strings.TrimSpace(h)
			if h == "" {
				utils.Json(w, resp, 401)
				return
			}

			//need to know what this does
			token, err := jwt.Parse(h, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				if err.Error() == "Token is expired" {
					resp["error"] = err.Error()
					utils.Json(w, resp, 401)
					return
				}
				resp["error"] = "Couldnt verify that you are authorized. make sure to check if your credentials are correct"
				utils.Json(w, resp, 401)
				log.Println(err.Error())
				return

			}
			//need to check this too out
			claims, _ := token.Claims.(jwt.MapClaims)
			userid, err := strconv.Atoi(claims["userid"].(string))
			if err != nil {
				resp["error"] = "something went wrong!!"
				utils.Json(w, resp, 500)
				log.Println(err.Error())
				return
			}
			ctx := context.WithValue(r.Context(), "userid", userid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
