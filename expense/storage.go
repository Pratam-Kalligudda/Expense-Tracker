package expense

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFromFile() ([]Expense, error) {
	var expenses []Expense
	data, err := os.ReadFile(File)
	if err != nil {
		if os.IsNotExist(err) {
			if file, err := os.Create(File); err != nil {
				return []Expense{}, nil
			} else {
				file.Close()
			}
		} else {
			return nil, fmt.Errorf("error reading from file : %w", err)
		}
	}
	if len(data) == 0 {
		return []Expense{}, nil
	}
	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling json : %w", err)
	}
	if len(expenses) == 0 {
		return []Expense{}, nil
	}
	return expenses, nil
}

func WriteToFile(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return fmt.Errorf("error while marshaling data: %w", err)
	}

	if err = os.WriteFile(File, data, 0664); err != nil {
		return fmt.Errorf("error while writing data to file : %w", err)
	}
	return nil
}
