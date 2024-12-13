package filesystem

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/saikumar-3093/go-projects/expense-tracker/models"
)

// type Expense struct {
// 	ID          int64
// 	Date        time.Time
// 	Description string
// 	Amount      int64
// 	Category    string
// }

func ConvertToCsv(dataFile, csvFileName string) error {
	file, err := os.Open(dataFile)

	if err != nil {
		return fmt.Errorf("error opening file")
	}
	defer file.Close()

	var ExpenseTable []models.Expense

	if err := json.NewDecoder(file).Decode(&ExpenseTable); err != nil {

		return fmt.Errorf("error decoding data from file to csv")

	}

	csvFile, err := os.Create(csvFileName)

	if err != nil {
		return fmt.Errorf("error creating csv file")
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	defer writer.Flush()

	headers := []string{"ID", "Date", "Description", "Amount", "Category"}

	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("error writing header to file")
	}

	for _, record := range ExpenseTable {
		var csvRow []string
		csvRow = append(csvRow, fmt.Sprint(record.ID), fmt.Sprint(record.Date.Format("2006-01-02")), record.Description, fmt.Sprint(record.Amount), record.Category)
		if err := writer.Write(csvRow); err != nil {
			return fmt.Errorf("error writing records to file")
		}
	}
	return nil
}
