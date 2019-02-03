package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "context"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var ctx = context.Background()
var SECRET = "you-are-retarded"

type User struct {
  Username, Password string
}

type Record struct {
  Uid int `json:"uid"`
  Start string  `json:"start"`
  End string `json:"end"`
}

func main() {
  var err error
  db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/sitting_time_tracker")
  if err != nil { panic(err) }
  if err = db.Ping(); err == nil {
    fmt.Printf("Database Connected!\n")
  }
  defer db.Close()

  http.HandleFunc("/register", register)
  http.HandleFunc("/login", login)
  http.HandleFunc("/verify-token", verifyToken)
  http.HandleFunc("/record", record)

  if err = http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}

func sendConfirmation(w http.ResponseWriter) {
    msg, _ := json.Marshal(map[string]interface{} {
      "success": true,
    })
    w.Write(msg)
}

func sendError(w http.ResponseWriter, status int, errString string) {
    msg, _ := json.Marshal(map[string]interface{} {
      "success": false,
      "error": errString,
    })
    w.WriteHeader(status)
    w.Write(msg)
}