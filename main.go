package main

import (
	"fmt"
	"time"

	"github.com/Pratam-Kalligudda/Expense-Tracker/expense"
)

func main() {
	if err := expense.Add("mobile", 20.4, time.Now()); err != nil {
		fmt.Printf("Error : %v", err)
		return
	}
}
