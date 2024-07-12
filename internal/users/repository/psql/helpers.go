package repository

import (
	"EM-Api-testTask/internal/models"
	"errors"
)

func filtersToValues(input models.FiltersDto) ([]interface{}, error) {
	var output []interface{}

	if input.Name == "" && input.Surname == "" && input.Patronymic == "" && input.Id == 0 {
		return nil, errors.New("required at least one filter")
	}
	output = append(output, input.Id)
	output = append(output, input.Name)
	//println(input.Name)
	output = append(output, input.Surname)
	//println(input.Surname)
	output = append(output, input.Patronymic)
	//println(input.Patronymic)
	output = append(output, input.TaskId)
	//println(input.TaskId)
	output = append(output, input.Limit)
	//println(input.Limit)
	output = append(output, input.Limit*(input.Page-1))
	//println(input.Page)

	return output, nil
}
