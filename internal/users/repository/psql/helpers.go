package repository

import (
	"EM-Api-testTask/internal/models"
)

func filtersToValues(input models.FiltersDto) []interface{} {
	var output []interface{}
	output = append(output, input.Name)
	output = append(output, input.Surname)
	output = append(output, input.Patronymic)
	output = append(output, input.TaskId)
	return output
}
