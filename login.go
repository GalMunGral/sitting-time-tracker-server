package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  jwt "github.com/dgrijalva/jwt-go"
)

// Expects `uid` and `password` from request body
// Sends back generated JWT token
func login(w http.ResponseWriter, r *http.Request) {
  var buffer [255]byte
  var body map[string]interface{}
  len, _ := r.Body.Read(buffer[:])
  if err:= json.Unmarshal(buffer[:len], &body); err != nil {
    panic(err)
  }
  uid := int(body["uid"].(float64))
  password, ok := body["password"].(string)
  if !ok {
    w.Write([]byte("No password"))
    return
  }
  stmt, err := db.PrepareContext(ctx, "SELECT * FROM users WHERE uid = ? AND password = ?")
  if err != nil {
    panic(err)
  }
  rows, err := stmt.Query(uid, password)
  if err != nil {
    msg := map[string]interface{} {
      "success": false,
      "error": err.Error(),
    }
    ob, _ := json.Marshal(msg)
    w.Write(ob)
    return
  }
  if !rows.Next() {
    msg := map[string]interface{} {
      "success": false,
      "error": "Invalid Credentials",
    }
    ob, _ := json.Marshal(msg)
    w.Write(ob)
    return
  }
  // Create JWT token
  fmt.Println("uid >> ", uid)
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "uid": uid,
    "test": "something",
  })
  str, err := token.SignedString([]byte("test-test"))
  w.Write([]byte(str))
}