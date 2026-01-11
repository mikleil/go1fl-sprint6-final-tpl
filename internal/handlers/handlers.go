package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type service interface {
	ConvertString(input string) (string, error)
}

type Handler struct {
	logger  *log.Logger
	service service
}

func New(logger *log.Logger, service service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./index.html")
}

/*
	func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
		filepath := filepath.Join("../index.html")
		http.ServeFile(w, r, filepath)
	}
*/
func (h *Handler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(100000000000); err != nil {
		h.logger.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		h.logger.Printf("Error returning file: %v", err)
		http.Error(w, "Error returning file", http.StatusBadRequest)
		return
	}

	input, err := io.ReadAll(file)
	if err != nil {
		h.logger.Printf("Error reading file: %v", err)
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	result, err := h.service.ConvertString(string(input))
	if err != nil {
		h.logger.Printf("Error converting string: %v", err)
		http.Error(w, "Error converting string", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}

func WriteToFile(str string) {
	os.WriteFile(time.Now().String()+".txt", []byte(str), 0644)
}
