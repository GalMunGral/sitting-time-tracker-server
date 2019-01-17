package main

import (
	"net/http"
  "strings"
  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

func test(w http.ResponseWriter, r *http.Request) {
  var (
    id int
    start string
    end string
  )
  rows, err := db.Query("select * from records")
  if err != nil {
  	panic(err)
  }
  defer rows.Close()
  for rows.Next() {
  	err := rows.Scan(&id, &start, &end)
  	if err != nil {
  		panic(err)
    }
    fmt.Printf("%d, %s, %s", id, start, end);
  }
  err = rows.Err()
  if err != nil {
  	panic(err)
  }
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
  fmt.Printf("yo\n")
}

func main() {
  var err error
  db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/sitting_time_tracker")
  if err != nil { panic(err) }
  if err = db.Ping(); err != nil { fmt.Printf("YOYO") }
  defer db.Close()

  http.HandleFunc("/test", test)
	if err = http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}