package handlers

import (
	// "context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/syd2/psg/db/db"
)

func HelloWorldHandler(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		ctx := r.Context()
		userid := int(ctx.Value("userid").(int))
		//w.Write([]byte())
		// w.Write([]byte("Hello World"))
		user, err := q.GetUserById(ctx, int64(userid))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(strconv.Itoa(int(userid)) + ":")
		fmt.Printf(fmt.Sprintf("hello, %s \n", string(user.Username)))
		// w.Write([]byte(fmt.Sprintf("hello, %s ", user.Username)))

	}

}
