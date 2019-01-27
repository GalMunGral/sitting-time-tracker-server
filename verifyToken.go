package main

import (
  "net/http"
  "fmt"
  "strings"
  _ "github.com/go-sql-driver/mysql"
  jwt "github.com/dgrijalva/jwt-go"
)

func verifyToken(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  tokenString := strings.Join(query["token"], "")
  token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte("test-test"), nil
  })
  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    uid := int(claims["uid"].(float64))
    fmt.Println("VERIFIED: ", uid)
  }
}
