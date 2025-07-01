package main

import (
    "encoding/json"
    "log"
    "net/http"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

type Model struct {
    gorm.Model
    Name    string
    Version string
    Path    string
    Metrics string
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    dsn := "host=localhost user=postgres password=yourpassword dbname=goserve sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        http.Error(w, "DB connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()
    db.AutoMigrate(&Model{})
    model := Model{Name: "mnist", Version: "v1.0", Path: "/models/mnist.onnx", Metrics: "{\"accuracy\": 0.97}"}
    db.Create(&model)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(model)
}

func main() {
    http.HandleFunc("/register", registerHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}