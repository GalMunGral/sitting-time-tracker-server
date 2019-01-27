package main

import (
  "net/http"
  "fmt"
  "context"
  // "reflect"
  "strings"
  "encoding/json"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var ctx = context.Background()

func register(w http.ResponseWriter, r *http.Request) {
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
  http.HandleFunc("/register", register)

  if err = http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}