package repository

import (
	"github.com/makonheimak/task-service/internal/task/orm"

	"gorm.io/gorm"
)

// CRUD
type TaskRepository interface {
	CreateTask(req *orm.Task) error
	GetAllTasks() ([]orm.Task, error)
	GetTaskByID(id int64) (orm.Task, error)
	GetTasksByUserID(userID int64) ([]orm.Task, error)
	UpdateTask(task orm.Task) error
	DeleteTask(id int64) error
}

type Repository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateTask(req *orm.Task) error {
	return r.db.Create(&req).Error
}

func (r *Repository) GetAllTasks() ([]orm.Task, error) {
	var tasks = []orm.Task{}
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *Repository) GetTaskByID(id int64) (orm.Task, error) {
	var task orm.Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *Repository) GetTasksByUserID(userID int64) ([]orm.Task, error) {
	var tasks []orm.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *Repository) UpdateTask(task orm.Task) error {
	return r.db.Save(&task).Error
}

func (r *Repository) DeleteTask(id int64) error {
	return r.db.Delete(&orm.Task{}, "id = ?", id).Error
}
