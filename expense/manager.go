package expense

import (
	"fmt"
	"time"
)

const File = "Expense.json"

func Add(desc string, amount float64, date time.Time) error {
	expenses, err := ReadFromFile()
	if err != nil {
		return fmt.Errorf("error while reading from file : %w", err)
	}

	expense := Expense{
		ID:          len(expenses) + 1,
		Description: desc,
		Date:        date.Format("2006-01-02"),
		Amount:      amount,
	}

	expenses = append(expenses, expense)

	return WriteToFile(expenses)

}
