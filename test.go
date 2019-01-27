package main

import (
  "net/http"
  "strings"
  "encoding/json"
  _ "github.com/go-sql-driver/mysql"
)

func test(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  id := strings.Join(query["id"], "")

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
