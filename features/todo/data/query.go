package data

import (
	"errors"
	"fmt"
	"log"
	"todolistapi/features/todo"

	"gorm.io/gorm"
)

type todoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) todo.TodoData {
	return &todoQuery{
		db: db,
	}
}

// Create implements todo.TodoData
func (tq *todoQuery) Create(newTodo todo.Core) (todo.Core, error) {
	todo := CoreToModel(newTodo)

	err := tq.db.Create(&todo).Error
	if err != nil {
		log.Println("Query create a new todo error : ", err.Error())
		return newTodo, err
	}

	return ModelToCore(todo), nil
}

// Delete implements todo.TodoData
func (tq *todoQuery) Delete(id uint) error {
	qryDelete := tq.db.Delete(&Todo{}, id)

	affectedRow := qryDelete.RowsAffected

	if affectedRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Todo with ID %d Not Found", id)
		return errors.New(msg)
	}

	return nil
}

// GetAll implements todo.TodoData
func (tq *todoQuery) GetAll(actID uint) ([]todo.Core, error) {
	allTodo := []Todo{}

	if actID <= 0 {
		err := tq.db.Find(&allTodo).Error
		if err != nil {
			log.Println("Query get All activities error : ", err.Error())
			return []todo.Core{}, err
		}
	} else {
		err := tq.db.Where("activity_group_id = ?", actID).Find(&allTodo).Error
		if err != nil {
			log.Println("Query get All todo error : ", err.Error())
			return []todo.Core{}, err
		}
	}

	todoArr := []todo.Core{}
	for _, val := range allTodo {
		todoArr = append(todoArr, ModelToCore(val))
	}

	return todoArr, nil
}

// GetOne implements todo.TodoData
func (tq *todoQuery) GetOne(id uint) (todo.Core, error) {
	act := Todo{}

	err := tq.db.Where("id = ?", id).First(&act).Error
	if err != nil {
		log.Println("Query get todo by ID error : ", err.Error())
		return todo.Core{}, err
	}

	return ModelToCore(act), nil
}

// Update implements todo.TodoData
func (tq *todoQuery) Update(id uint, updatedTodo todo.Core) (todo.Core, error) {
	cnvUpdate := CoreToModel(updatedTodo)

	qry := tq.db.Model(Todo{}).Where("id = ?", id).Updates(&cnvUpdate)
	toggle := tq.db.Model(&cnvUpdate).Where("id = ?", id).Update("is_active", updatedTodo.IsActive)
	err := qry.Error

	affectedRow := toggle.RowsAffected

	if affectedRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Todo with ID %d Not Found", id)
		return todo.Core{}, errors.New(msg)
	}

	affectedRow = qry.RowsAffected

	if affectedRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Todo with ID %d Not Found", id)
		return todo.Core{}, errors.New(msg)
	}

	if err != nil {
		log.Println("Query update todo by ID error : ", err.Error())
		return todo.Core{}, errors.New("Error")
	}

	var updateRow Todo
	tq.db.First(&updateRow, "id = ?", id)

	return ModelToCore(updateRow), nil
}
