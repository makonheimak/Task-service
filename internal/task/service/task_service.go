package service

import (
	"github.com/makonheimak/task-service/internal/task/orm"
	"github.com/makonheimak/task-service/internal/task/repository"
)

type TaskService interface {
	CreateTask(req orm.Task) (orm.Task, error)
	GetAllTasks() ([]orm.Task, error)
	GetTaskByID(id int64) (orm.Task, error)
	GetTasksByUserID(userID int64) ([]orm.Task, error)
	UpdateTask(id int64, task string) (orm.Task, error)
	DeleteTask(id int64) error
}

type Service struct {
	repo repository.TaskRepository
}

func NewService(r repository.TaskRepository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateTask(req orm.Task) (orm.Task, error) {
	err := s.repo.CreateTask(&req)
	if err != nil {
		return orm.Task{}, err
	}
	return req, nil
}

func (s *Service) GetAllTasks() ([]orm.Task, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *Service) GetTaskByID(id int64) (orm.Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return orm.Task{}, err
	}
	return task, nil
}

func (s *Service) GetTasksByUserID(userID int64) ([]orm.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

func (s *Service) UpdateTask(id int64, newText string) (orm.Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return orm.Task{}, err
	}

	task.Task = newText

	if err := s.repo.UpdateTask(task); err != nil {
		return orm.Task{}, err
	}

	return task, nil
}

func (s *Service) DeleteTask(id int64) error {
	err := s.repo.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
