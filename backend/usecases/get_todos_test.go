package usecases_test

import (
	"fmt"
	"testing"

	"github.com/alifurkangokce/GoWithReactExample/backend/entities"
	"github.com/alifurkangokce/GoWithReactExample/backend/usecases"
	"github.com/gomagedon/expectate"
)

var dummyTodos = []entities.Todo{
	{
		Title:       "todo 1",
		Description: "description of todo 1",
		IsCompleted: true,
	},
	{
		Title:       "todo 2",
		Description: "description of todo 2",
		IsCompleted: false,
	},
	{
		Title:       "todo 3",
		Description: "description of todo 3",
		IsCompleted: false,
	},
}

type BadTodosRepo struct{}

func (BadTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("Something Went Wrong")
}

type MockTodosRepo struct{}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return dummyTodos, nil
}

func TestGetTodos(t *testing.T) {
	t.Run("Returns ErrInternal when TodosRepository returns eror", func(t *testing.T) {
		expect := expectate.Expect(t)
		repo := new(BadTodosRepo)
		todos, err := usecases.GetTodos(repo)
		expect(err).ToBe(usecases.ErrInternal)

		if todos != nil {
			t.Fatalf("Expected todos to be nil; Got:%v", todos)
		}
	})
	t.Run("Returns todos from TodosRepository", func(t *testing.T) {
		expect := expectate.Expect(t)
		repo := new(MockTodosRepo)
		todos, err := usecases.GetTodos(repo)
		expect(err).ToBe(nil)
		expect(todos).ToEqual(dummyTodos)
	})

}
