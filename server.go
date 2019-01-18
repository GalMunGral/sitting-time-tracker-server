package main

import (
  "net/http"
  "fmt"
  "strings"
  "encoding/json"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func test(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  id := strings.Join(query["id"], "")
  type Record struct {
    Start string `json: "start"`
    End string `json: "end"`
  }

  rows, err := db.Query("select start_time, end_time from records where user_id = ?", id)
  if err != nil { panic(err) }
  defer rows.Close()

  var records []Record
  var start string
  var end string
  for rows.Next() {
  	if err := rows.Scan(&start, &end); err != nil { panic(err) }
    records = append(records, Record{Start: start, End: end})
  }
  if rows.Err() != nil { panic(err) }

  ob, err := json.Marshal(records)
  if err != nil { panic(err) }
  w.Write(ob)
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