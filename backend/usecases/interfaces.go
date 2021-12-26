package usecases

import "github.com/alifurkangokce/GoWithReactExample/backend/entities"

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}
