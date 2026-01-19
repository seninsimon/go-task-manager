package repositories

import (
	"task-manager/internal/db"
	"task-manager/internal/models"
)

type TaskRepository struct{}

func (r *TaskRepository) Create(task *models.Task) error{
	return db.DB.Create(task).Error
}


func (r *TaskRepository) FindByUser(userID uint) ([]models.Task , error){
	var tasks []models.Task
	err := db.DB.Where("user_id = ?",userID).Find(&tasks).Error
	return  tasks , err
}

func (r *TaskRepository) FindById(id uint , userId uint) (*models.Task , error){
	var task models.Task
	err := db.DB.Where("id = ? AND user_id = ?",id , userId).Find(&task).Error
	if err != nil {
		return  nil , err
	}
	return  &task , err
}


func (r *TaskRepository) Update(task *models.Task) error {
	return  db.DB.Save(task).Error
}

func (r *TaskRepository) Delete(id uint , user_id uint) error {	
	var task models.Task
	return  db.DB.Where("id = ? AND user_id = ? ",id , user_id).Delete(&task).Error	
}




