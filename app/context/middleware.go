package context

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	usersResolver "projectt/app/graphql/resolver/users"

	"github.com/gin-gonic/gin"
)

type UseCaseContext int

const (
	ConfigKey        UseCaseContext = 0
	UsersResolverKey UseCaseContext = 1
)

// ApplyContextMiddleware
func ApplyContextMiddleware(
	next http.Handler,
	UsersResolver usersResolver.UsersResolver,
) http.Handler {
	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			res.Header().Set("Access-Control-Allow-Credentials", "true")
			res.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, token")

			ret := make(map[string]interface{})

			newContext := context.WithValue(req.Context(), "headers", ret)

			newContext = context.WithValue(newContext, UsersResolverKey, UsersResolver)

			log.Printf(
				"%s %s %s \n",
				req.RemoteAddr,
				req.Method,
				req.URL,
			)

			next.ServeHTTP(res, req.WithContext(newContext))
		},
	)
}

// ApplyAuthenticationMiddleware ...
func ApplyAuthenticationMiddleware(next http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		next.ServeHTTP(c.Writer, c.Request)
	}
}

// respondWithJSON
func respondWithJSON(w http.ResponseWriter, code int, message interface{}) {
	response, _ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
