package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type service interface {
	ConvertString(input string) (string, error)
}

type Handler struct {
	logger  *log.Logger
	service service
}

func getProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
			return cwd, nil
		}

		parent := filepath.Dir(cwd)
		if parent == cwd {
			break
		}
		cwd = parent
	}

	return "", fmt.Errorf("go.mod not found")
}

func New(logger *log.Logger, service service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}
func (h *Handler) HandlerRoot(w http.ResponseWriter, r *http.Request) {
	rootDir, err := getProjectRoot()
	if err != nil {
		h.logger.Printf("Internal error: %v", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	indexPath := filepath.Join(rootDir, "index.html")
	http.ServeFile(w, r, indexPath)
}

func (h *Handler) HandlerUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(100000000); err != nil {
		h.logger.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		h.logger.Printf("Error reading file: %v", err)
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		h.logger.Printf("Error reading file: %v", err)
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	result, err := h.service.ConvertString(string(input))
	if err != nil {
		h.logger.Printf("Error converting string: %v", err)
		http.Error(w, "Error converting string", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
