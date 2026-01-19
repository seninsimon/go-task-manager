package services

import (
	"errors"
	"task-manager/internal/models"
	"task-manager/internal/repositories"
)

type TaskService struct {
	repo *repositories.TaskRepository
}

func NewTaskService() *TaskService {
	return &TaskService{
		repo : &repositories.TaskRepository{},
	}
}

func (s *TaskService) Create(title string, userId uint) error {
	if title == "" {
		return errors.New("title is required")
	}

	task := &models.Task{

		Title:  title,
		UserID: userId,
	}

	return  s.repo.Create(task)

}

func (s *TaskService) GetAll(userid uint) ([]models.Task,  error) {
	return s.repo.FindByUser(userid)
}


func (s *TaskService) Update(id uint , userID uint , title string , status string) error {
	task , err := s.repo.FindById(id , userID) 
	if err != nil {
		return errors.New("tasks not found");
	}

	if title != "" {
		task.Title = title
	}

	if status != "" {
		task.Status = status
	}	

	return s.repo.Update(task)
}


func (s *TaskService) Delete(id uint , userId uint) error {
	return s.repo.Delete(id ,userId)
}
