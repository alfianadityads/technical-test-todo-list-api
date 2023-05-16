package data

import (
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
	panic("unimplemented")
}

// GetAll implements todo.TodoData
func (tq *todoQuery) GetAll(actID uint) ([]todo.Core, error) {
	panic("unimplemented")
}

// GetOne implements todo.TodoData
func (tq *todoQuery) GetOne(id uint) (todo.Core, error) {
	panic("unimplemented")
}

// Update implements todo.TodoData
func (tq *todoQuery) Update(id uint, updatedTodo todo.Core) (todo.Core, error) {
	panic("unimplemented")
}
