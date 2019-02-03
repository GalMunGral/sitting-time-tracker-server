package main

import (
  "net/http"
  "encoding/json"
  jwt "github.com/dgrijalva/jwt-go"
)

func login(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var user User
  var uid int
  if (r.Method != "POST") {
    w.WriteHeader(405)
    return
  }

  dec := json.NewDecoder(r.Body)
  if err:= dec.Decode(&user); err != nil {
    sendError(w, 409, err.Error())
    return
  }

  stmt, _ := db.PrepareContext(ctx, "SELECT * FROM users WHERE username = ? AND password = ?")
  defer stmt.Close()
  rows, err := stmt.Query(user.Username, user.Password)
  if err != nil {
    sendError(w, 409, err.Error())
    return
  }
  if !rows.Next() {
    sendError(w, 401, "Invalid Credentials")
    return
  }
  rows.Scan(&uid, nil, nil)
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "uid": uid,
  })
  str, err := token.SignedString([]byte(SECRET))
  ob, _ := json.Marshal(map[string]interface{} {
    "token": str,
  }) 
  w.Write(ob)
}