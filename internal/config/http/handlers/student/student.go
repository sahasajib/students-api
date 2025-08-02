package student

import "net/http"


func New() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request) {
	// Handler logic to get students
	w.Write([]byte("List of students"))
}
}