package auth
//Набросок
import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodES512)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "User User"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString("password")

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		hashed, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(hashed), ":", 2)

		if len(pair) != 2 || !authCheck(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func authCheck(username, password string) bool  {
	if username == "test" && password == "test" {
		return true
	}

	return false
}
