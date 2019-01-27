package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  _ "github.com/go-sql-driver/mysql"
)

func register(w http.ResponseWriter, r *http.Request) {
  var buffer [255]byte
  var body map[string]interface{}
  len, _ := r.Body.Read(buffer[:])
  if err:= json.Unmarshal(buffer[:len], &body); err != nil {
    panic(err)
  }
  uid := int(body["uid"].(float64))
  fmt.Println("uid <-", uid)
  password, ok := body["password"].(string)
  if !ok {
    w.Write([]byte("No password"))
    return
  }
  stmt, err := db.PrepareContext(ctx, "INSERT INTO users(uid, password) VALUES(?, ?)")
  if err != nil {
    panic(err)
  }
  defer stmt.Close()
  if _, err := stmt.Exec(uid, password); err != nil {
    msg := map[string]interface{} {
      "success": false,
      "error": err.Error(),
    }
    ob, _ := json.Marshal(msg)
    w.Write(ob)
    return
  }
  w.Write([]byte("Success!"))
}