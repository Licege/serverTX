package auth
//Набросок
import (
	"encoding/base64"
	"net/http"
	"strings"
)

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
