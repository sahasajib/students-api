package student

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sahasajib/students-api/internal/storage"
	"github.com/sahasajib/students-api/internal/types"
)

func handleCors(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control_Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func HanleOpt(w http.ResponseWriter, r *http.Request){
	if r.Method == "OPTIONS"{
		w.WriteHeader(http.StatusCreated)
	}
}

func sendData(w http.ResponseWriter,  data interface{}, statusCode int){
	
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)

	encoder.Encode(data)
}

var studentList[] types.Student
func New(storage storage.Storage) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request) {
		// Handler logic to get students
		handleCors(w)
		HanleOpt(w, r)
		
		if r.Method != "POST" {
			http.Error(w, "Please use POST method", http.StatusMethodNotAllowed)
			return
		}
		var newStudent types.Student
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newStudent)

		if err != nil{
			fmt.Println("Error decoding JSON:", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		newStudent.ID =  int64(len(studentList) + 1)
		studentList = append(studentList, newStudent)
	
		if err := validator.New().Struct(newStudent)
			err != nil {
				fmt.Println("Validation error: ", err)
				http.Error(w, "Validation error", http.StatusBadRequest)
				return
			}

			slog.Info("User created successfully", slog.String("name", newStudent.Name))

			lastId, err := storage.CreateStudent(
				newStudent.Name,
				newStudent.Email,
				newStudent.Age,
			)
			if err != nil {
				http.Error(w, "Failed to create student", http.StatusInternalServerError)
				return 
			}

		sendData(w, map[string]int64{"id": lastId}, http.StatusOK)

	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleCors(w)
		HanleOpt(w, r)

		if r.Method != "GET" {
			http.Error(w, "Please use GET method", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		slog.Info("Fetching student by ID", slog.String("id", id))

		// Here you would typically fetch the student from the storage
		// For now, we will just return a dummy student	
	}
}