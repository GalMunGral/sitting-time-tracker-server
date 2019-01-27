package main

import (
  "net/http"
  "fmt"
  "context"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var ctx = context.Background()

type Record struct {
  Start string `json: "start"`
  End string `json: "end"`
}

func main() {
  var err error
  db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/sitting_time_tracker")
  if err != nil { panic(err) }
  if err = db.Ping(); err == nil {
    fmt.Printf("Database Connected!\n")
  }
  defer db.Close()

  http.HandleFunc("/test", test)
  http.HandleFunc("/register", register)
  http.HandleFunc("/login", login)
  http.HandleFunc("/verify-token", verifyToken)

  if err = http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}