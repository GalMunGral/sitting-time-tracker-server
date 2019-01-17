package main

import (
  "net/http"
  "fmt"
  "strings"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func test(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  id := strings.Join(query["id"], "")

  var (
    start string
    end string
  )
  rows, err := db.Query("select start_time, end_time from records where user_id = ?", id)
  if err != nil { panic(err) }
  defer rows.Close()
  for rows.Next() {
  	err := rows.Scan(&start, &end)
  	if err != nil { panic(err) }
    fmt.Printf("%s, %s\n", start, end);
  }
  if rows.Err() != nil { panic(err) }
  
  message := start + " => " + end + "\n"
  w.Write([]byte(message))
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

  if err = http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}