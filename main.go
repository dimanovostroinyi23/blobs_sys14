package main

import (
  "database/sql"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
  var err error
  // Замініть налаштування з'єднання з вашою реальною інформацією доступу
  connStr := "host=localhost port=5432 user=postgres dbname=postgres password=z14042004 sslmode=disable"
  db, err = sql.Open("postgres", connStr)
  if err != nil {
    log.Fatal(err)
  }
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  // Виправлення тут: обрамляємо SQL-запит у лапки, перетворюючи його на рядок
  createTableSQL := `CREATE TABLE IF NOT EXISTS blobs (
        id SERIAL PRIMARY KEY,
        data JSONB NOT NULL
    );`
  _, err = db.Exec(createTableSQL)
  if err != nil {
    log.Fatal("Failed to create table: ", err)
  }

  fmt.Println("Successfully connected to the database and ensured the 'blobs' table exists.")
}

type Blob struct {
  ID   int             `json:"id,omitempty"`
  Data json.RawMessage `json:"data"`
}

func createBlob(w http.ResponseWriter, r *http.Request) {
  var b Blob
  if err := json.NewDecoder(r.Body).Decode(&b.Data); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  err := db.QueryRow("INSERT INTO blobs(data) VALUES($1) RETURNING id", b.Data).Scan(&b.ID)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(b)
}

func getBlob(w http.ResponseWriter, r *http.Request) {
  idStr := mux.Vars(r)["id"]
  id, err := strconv.Atoi(idStr)
  if err != nil {
    http.Error(w, "Invalid blob ID", http.StatusBadRequest)
    return
  }

  var b Blob
  err = db.QueryRow("SELECT id, data FROM blobs WHERE id = $1", id).Scan(&b.ID, &b.Data)
  if err != nil {
    http.Error(w, "Blob not found", http.StatusNotFound)
    return
  }

  json.NewEncoder(w).Encode(b)
}

func deleteBlob(w http.ResponseWriter, r *http.Request) {
  idStr := mux.Vars(r)["id"]
  id, err := strconv.Atoi(idStr)
  if err != nil {
    http.Error(w, "Invalid blob ID", http.StatusBadRequest)
    return
  }

  _, err = db.Exec("DELETE FROM blobs WHERE id = $1", id)
  if err != nil {
    http.Error(w, "Failed to delete blob", http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusNoContent)
}

func listBlobs(w http.ResponseWriter, r *http.Request) {
  rows, err := db.Query("SELECT id, data FROM blobs")
  if err != nil {
    http.Error(w, "Failed to retrieve blobs", http.StatusInternalServerError)
    return
  }
  defer rows.Close()

  blobs := []Blob{}
  for rows.Next() {
    var b Blob
    if err := rows.Scan(&b.ID, &b.Data); err != nil {
      http.Error(w, "Failed to scan blob", http.StatusInternalServerError)
      return
    }
    blobs = append(blobs, b)
  }

  json.NewEncoder(w).Encode(blobs)
}

func main() {
  initDB()

  r := mux.NewRouter()
  r.HandleFunc("/blobs", createBlob).Methods("POST")
  r.HandleFunc("/blobs/{id}", getBlob).Methods("GET")
  r.HandleFunc("/blobs/{id}", deleteBlob).Methods("DELETE")
  r.HandleFunc("/blobs", listBlobs).Methods("GET")

  fmt.Println("Server is running on port 8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}