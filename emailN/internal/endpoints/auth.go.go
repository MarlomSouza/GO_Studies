package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"

	jwt_go "github.com/dgrijalva/jwt-go/v4"
)

func Auth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "Request does not contain authorization header"})

			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/providerGO")

		if err != nil {
			fmt.Println(err)
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": "error to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})

		_, err = verifier.Verify(r.Context(), tokenString)
		if err != nil {
			fmt.Println(err)
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}

		token, _ := jwt_go.Parse(tokenString, nil)
		claims := token.Claims.(jwt_go.MapClaims)
		email := claims["email"]

		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
