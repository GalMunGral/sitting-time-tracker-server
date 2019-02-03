package main

import (
  // "fmt"
  "net/http"
  "encoding/json"
  jwt "github.com/dgrijalva/jwt-go"
  _ "github.com/go-sql-driver/mysql"
)

func register(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var user User
  if (r.Method != "POST") {
    w.WriteHeader(405)
    return
  }

  dec := json.NewDecoder(r.Body)
  if err:= dec.Decode(&user); err != nil {
    sendError(w, 409, err.Error())
    return
  }

  stmt, _ := db.PrepareContext(ctx, "INSERT INTO users(username, password) VALUES(?, ?)")
  defer stmt.Close()
  result, err := stmt.Exec(user.Username, user.Password)
  if err != nil {
    sendError(w, 409, err.Error())
    return
  }
  uid, _ := result.LastInsertId()
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "uid": uid,
  })
  str, _ := token.SignedString([]byte(SECRET))
  ob, _ := json.Marshal(map[string]interface{} {
    "token": str,
  }) 
  w.Write(ob)
}
