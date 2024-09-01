package utils

import (
	"time"

	"github.com/DaniloFaraum/go-crud-api/schemas"
)

func ConvertToPersonResponse(p schemas.Person) schemas.PersonResponse {
	var deletedAt *time.Time
	if p.DeletedAt.Valid {
		deletedAt = &p.DeletedAt.Time
	}

	return schemas.PersonResponse{
		ID:        p.ID,
		Name:      p.Name,
		Age:       p.Age,
		Gender:    p.Gender,
		Ocupation: p.Ocupation,
		Salary:    p.Salary,
		DeletedAt: deletedAt,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ConvertToPersonResponses(persons []schemas.Person) []schemas.PersonResponse {
	var responses []schemas.PersonResponse
	for _, person := range persons {
		responses = append(responses, ConvertToPersonResponse(person))
	}
	return responses
}
