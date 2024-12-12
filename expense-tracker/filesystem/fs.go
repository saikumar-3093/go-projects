package filesystem

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/saikumar-3093/go-projects/expense-tracker/models"
)

func GetFilePath() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	dir = path.Join(dir, "expenses.json")
	return dir, nil
}

func checkBudget() (int64, error) {
	MonthBudget, err := MonthBudget(time.November.String())

	if err != nil {
		return 0, err

	}
	return MonthBudget, nil
}
func ReadExpensesFile() ([]models.Expense, error) {
	fileName, err := GetFilePath()

	if err != nil {
		fmt.Println("Error getting file path")
		return []models.Expense{}, err
	}
	fileinfo, err := os.Stat(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(fileName)
			if err != nil {
				fmt.Println("Error creating file")
				return []models.Expense{}, err
			}
			fmt.Println("created file!")
			fileinfo, err = os.Stat(fileName)
			if err != nil {
				fmt.Println("Error getting fileinfo after creating")
				return []models.Expense{}, err
			}
		} else {
			fmt.Println("Error getting file information")
			return []models.Expense{}, err
		}

	}

	if fileinfo.Size() == 0 {
		fmt.Println("File is empty")
		return []models.Expense{}, nil
	}
	_, err = os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file")
		return []models.Expense{}, err
	}
	fileData, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading file", err)
		return []models.Expense{}, err
	}
	var expensesList []models.Expense

	err = json.Unmarshal(fileData, &expensesList)

	if err != nil {
		fmt.Println("Error Unmarshaling file data", err)
		return []models.Expense{}, err
	}
	return expensesList, nil
}

func WritetoFile(expenses []models.Expense, file string) error {

	jsondata, err := json.Marshal(expenses)

	if err != nil {
		fmt.Println("Error Marshaling data")
		return err
	}

	err = os.WriteFile(file, jsondata, 0666)

	return err

}

func New(desc string, amount int64, category string) *models.Expense {
	return &models.Expense{
		Date:        time.Now(),
		Description: desc,
		Amount:      amount,
		Category:    category,
	}

}
func AddExpense(desc string, amount int64, category string) {
	totalBudget, err := MonthBudget(time.Now().Month().String())
	if err != nil {
		fmt.Println("Error getting month Budget")
		return
	}

	totalExpense, err := MonthSummary(time.Now().Month().String())
	if err != nil {
		fmt.Println("Error getting month Summary")
		return
	}

	if totalExpense+amount > totalBudget {
		fmt.Println("\033[31mSpending out of the Budget\033[0m")
		fmt.Printf("%s Month Budget: \033[33m%v\033[0m\n", time.Now().Month().String(), totalBudget)
		fmt.Printf("%s Month Expense Summary till now: \033[33m%v\033[0m\n", time.Now().Month().String(), totalExpense)
		fmt.Printf("New Expense: \033[33m%v\033[0m\n", amount)
		return
	}

	newExpense := New(desc, amount, category)
	file, err := GetFilePath()
	if err != nil {
		fmt.Println("Error getting file")
		return
	}
	ExpenseList, err := ReadExpensesFile()

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	if len(ExpenseList) == 0 {
		newExpense.ID = 1
	} else {
		newExpense.ID = ExpenseList[len(ExpenseList)-1].ID + 1
	}
	ExpenseList = append(ExpenseList, *newExpense)
	WritetoFile(ExpenseList, file)
}

func DeleteExpense(id int64) error {
	file, err := GetFilePath()
	if err != nil {
		fmt.Println("Error getting file")
		return err
	}
	allExpenses, err := ReadExpensesFile()

	if err != nil {
		fmt.Println("Error reading file", err)
		return err
	}
	if len(allExpenses) == 0 {
		return fmt.Errorf("No Expense with Id = %v to Delete", id)

	}
	if id > allExpenses[len(allExpenses)-1].ID || id <= 0 {
		return fmt.Errorf("Invalid ID")

	}
	var expId = int64(-1)
	for index, exp := range allExpenses {
		if exp.ID == id {
			expId = int64(index)
		}
	}
	if expId == -1 {
		return fmt.Errorf("No expense with the ID: %v", id)
	}

	allExpenses = append(allExpenses[:expId], allExpenses[expId+1:]...)
	return WritetoFile(allExpenses, file)
}

func UpdateExpense(id int64, desc string, amount int64, category string) error {
	file, err := GetFilePath()
	if err != nil {
		fmt.Println("Error getting file")
		return err
	}
	allExpenses, err := ReadExpensesFile()

	for index, exp := range allExpenses {
		if exp.Category == "" {
			allExpenses[index].Category = "other"
		}
	}

	if err != nil {
		fmt.Println("Error reading file", err)
		return err
	}
	if len(allExpenses) == 0 {
		return fmt.Errorf("No Expense with ID = %v to Update", id)

	}
	if id > allExpenses[len(allExpenses)-1].ID || id <= 0 {
		return fmt.Errorf("Invalid ID")

	}
	var expId = int64(-1)
	for index, exp := range allExpenses {
		if exp.ID == id {
			expId = int64(index)
		}
	}
	if expId == -1 {
		return fmt.Errorf("No expense with the ID: %v", id)
	}

	allExpenses[expId].Description = desc
	allExpenses[expId].Amount = amount
	allExpenses[expId].Category = category

	return WritetoFile(allExpenses, file)

}

func List() {
	allExpenses, err := ReadExpensesFile()
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	if len(allExpenses) == 0 {
		fmt.Println("No expenses to Show.")
	}
	// Create a new tab writer
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Print header
	fmt.Fprintln(writer, "#\tID\tDate\tDescription\tAmount")

	// Print each expense
	for _, expense := range allExpenses {
		fmt.Fprintf(writer, "#\t%d\t%s\t%s\t$%v\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount)
	}

	// Flush the writer
	writer.Flush()

}

func Summary() (int64, error) {
	_, err := GetFilePath()
	if err != nil {
		fmt.Println("Error getting file")
		return 0, err
	}
	allExpenses, err := ReadExpensesFile()

	if err != nil {
		fmt.Println("Error reading file", err)
		return 0, err
	}
	var summaryCount int64 = 0
	for _, exp := range allExpenses {
		summaryCount += exp.Amount
	}
	fmt.Println("Summary:", summaryCount)
	return summaryCount, nil
}

func MonthSummary(month string) (int64, error) {
	_, err := GetFilePath()
	if err != nil {
		fmt.Println("Error getting file")
		return 0, err
	}
	allExpenses, err := ReadExpensesFile()

	if err != nil {
		fmt.Println("Error reading file", err)
		return 0, err
	}
	var monthSummaryCount int64 = 0
	for _, exp := range allExpenses {
		if exp.Date.Month().String() == month {
			monthSummaryCount += exp.Amount
		}

	}
	fmt.Printf("%s month summary: %v\n", month, monthSummaryCount)
	return monthSummaryCount, nil
}

func CategoryExpenses(category string) {
	allExpenses, err := ReadExpensesFile()
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	var categoriedExpenses []models.Expense

	for index, exp := range allExpenses {
		if strings.TrimSpace(exp.Category) == strings.TrimSpace(category) {
			categoriedExpenses = append(categoriedExpenses, allExpenses[index])
		}
	}

	if len(categoriedExpenses) == 0 {
		fmt.Printf("No Expenses for the Category: \033[33m%s\033[0m\n", category)
		return
	}
	fmt.Printf("\033[33m%s Expenses\033[0m\n", category)
	// Create a new tab writer
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Print header
	fmt.Fprintln(writer, "#\tID\tDate\tDescription\tAmount")

	// Print each expense
	for _, expense := range categoriedExpenses {
		fmt.Fprintf(writer, "#\t%d\t%s\t%s\t$%v\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount)
	}

	// Flush the writer
	writer.Flush()

}
