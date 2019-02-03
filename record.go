package main

import (
  "fmt"
  "net/http"
  "strings"
  "encoding/json"
  jwt "github.com/dgrijalva/jwt-go"
  _ "github.com/go-sql-driver/mysql"
)

func record(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  
  query := r.URL.Query()
  tokenString := strings.Join(query["token"], "")

  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte(SECRET), nil // Return the secret used for signing
  })
  if err != nil {
    sendError(w, 401, err.Error())
    return
  }
  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok || !token.Valid {
    sendError(w, 401, err.Error())
    return
  }
  uid := int(claims["uid"].(float64)) // Extract uid from the token
  fmt.Println("Authorized:", uid)

  switch r.Method {
    case "GET":
      ob, _ := json.Marshal(map[string]interface{} {
        "results": getRecords(uid),
      }) 
      w.Write(ob)
    case "POST":
      var record Record
      dec := json.NewDecoder(r.Body)
      if err := dec.Decode(&record); err != nil || record.Start == "" || record.End == "" {
        if err != nil {
          sendError(w, 400, err.Error())
        } else {
          sendError(w, 400, "start/end field missing.")
        }
        return
      }
      stmt, _ := db.PrepareContext(ctx, "INSERT INTO records(uid, start_time, end_time) VALUES(?, ?, ?)")
      defer stmt.Close()
      _, err := stmt.Exec(uid, record.Start, record.End)
      if err != nil {
        sendError(w, 400, err.Error())
        return
      } else {
        sendConfirmation(w)
      }
  }
}

func getRecords(uid int) []Record {
  var records []Record
  var start, end string

  rows, _ := db.Query("SELECT start_time, end_time FROM records WHERE uid = ?", uid)
  defer rows.Close()

  for rows.Next() {
  	if err := rows.Scan(&start, &end); err != nil {
      panic(err)
    }
    records = append(records, Record{ Uid: uid, Start: start, End: end })
  }
  if rows.Err() != nil {
    panic(rows.Err())
  }
  return records
}