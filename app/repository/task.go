package repository

import "github.com/yyh-gl/go-flat-package/app"

type taskRepository struct {
	memoryDB map[string]*app.TaskDTO
}

var _ app.TaskRepository = (*taskRepository)(nil)

func NewTaskRepository() *taskRepository {
	return &taskRepository{
		memoryDB: map[string]*app.TaskDTO{
			"hoge": &app.TaskDTO{
				Title:       "first task",
				Description: "fffff",
			},
		},
	}
}

func (t *taskRepository) Get(id string) *app.TaskDTO {
	return t.memoryDB[id]
}

