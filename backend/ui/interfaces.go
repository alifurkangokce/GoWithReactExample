package ui

import "github.com/alifurkangokce/GoWithReactExample/backend/entities"

type Service interface {
	GetAllTodos() ([]entities.Todo, error)
}
