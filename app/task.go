package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	repository TaskRepository
}

type TaskRepository interface {
	Get(string) *TaskDTO
}

func NewTaskHandler(repo TaskRepository) *TaskHandler {
	return &TaskHandler{
		repository: repo,
	}
}

type TaskDTO struct {
	Title string
	Description string
}

func (t *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Title string `json:"title"`
		Description string `json:"description"`
	}
	
	vars := mux.Vars(r)
	dto := t.repository.Get(vars["id"])
	res := response{
		Title:       dto.Title,
		Description: dto.Description,
	}
	resp, _ := json.Marshal(res)
	_, _ = w.Write(resp)
}
