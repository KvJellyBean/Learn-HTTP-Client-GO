package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	returnedData := make([][]byte, 0, len(items))

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		returnedData = append(returnedData, data)
	}

	return returnedData, nil
}
