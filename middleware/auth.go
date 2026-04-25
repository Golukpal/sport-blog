package middleware

import(
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"sport-blog/utils"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		tokenstring := strings.Split(header, " ")[1]

		token,_ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error){
			return utils.SecretKey, nil
		
	})
	 if !token.Valid {
		http.Error(w, "Invalid auth token", http.StatusUnauthorized)
		return
	 }
	    next.ServeHTTP(w, r)
	})
	 
}