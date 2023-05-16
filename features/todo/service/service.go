package service

import (
	"todolistapi/features/todo"
	"todolistapi/helper"

	"github.com/go-playground/validator"
)

type todoService struct {
	qry todo.TodoData
	vld *validator.Validate
}

func New(td todo.TodoData) todo.TodoService {
	return &todoService{
		qry: td,
		vld: validator.New(),
	}
}

// Create implements todo.TodoService
func (ts *todoService) Create(newTodo todo.Core) (todo.Core, error) {
	err := helper.Validation(newTodo)
	if err != nil {
		return newTodo, err
	}

	res, err := ts.qry.Create(newTodo)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Delete implements todo.TodoService
func (ts *todoService) Delete(id uint) error {
	err := ts.qry.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements todo.TodoService
func (ts *todoService) GetAll(actID uint) ([]todo.Core, error) {
	res, err := ts.qry.GetAll(actID)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetOne implements todo.TodoService
func (ts *todoService) GetOne(id uint) (todo.Core, error) {
	res, err := ts.qry.GetOne(id)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Update implements todo.TodoService
func (ts *todoService) Update(id uint, updatedTodo todo.Core) (todo.Core, error) {
	res, err := ts.qry.Update(id, updatedTodo)
	if err != nil {
		return updatedTodo, err
	}

	return res, nil
}
