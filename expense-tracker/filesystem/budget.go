package filesystem

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/saikumar-3093/go-projects/expense-tracker/models"
)

func GetBudgetFilePath() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	dir = path.Join(dir, "budget.json")
	return dir, nil
}

func NewBudget(month string, amount int64) *models.Budget {
	return &models.Budget{
		Month:  month,
		Amount: amount,
	}
}

func ReadBudgetFile() ([]models.Budget, error) {
	fileName, err := GetBudgetFilePath()

	if err != nil {
		fmt.Println("Error getting budget file path")
		return []models.Budget{}, err
	}
	fileinfo, err := os.Stat(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(fileName)
			if err != nil {
				fmt.Println("Error creating Budget file")
				return []models.Budget{}, err
			}

			fileinfo, err = os.Stat(fileName)
			if err != nil {
				fmt.Println("Error getting Budget fileinfo after creating")
				return []models.Budget{}, err
			}
		} else {
			fmt.Println("Error getting budget file information")
			return []models.Budget{}, err
		}

	}

	if fileinfo.Size() == 0 {
		return []models.Budget{}, nil
	}
	_, err = os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening budget file")
		return []models.Budget{}, err
	}
	fileData, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading budget file", err)
		return []models.Budget{}, err
	}
	var budgetList []models.Budget

	err = json.Unmarshal(fileData, &budgetList)

	if err != nil {
		fmt.Println("Error Unmarshaling budget file data", err)
		return []models.Budget{}, err
	}
	return budgetList, nil
}

func WritetoBudgetFile(budget []models.Budget, file string) error {

	jsondata, err := json.Marshal(budget)

	if err != nil {
		fmt.Println("Error Marshaling data")
		return err
	}

	err = os.WriteFile(file, jsondata, 0666)

	return err

}

func AddBudget(month string, amount int64) error {
	fileName, err := GetBudgetFilePath()

	if err != nil {
		return fmt.Errorf("error getting Budget file path")

	}
	newBudget := NewBudget(month, amount)

	allBudgets, err := ReadBudgetFile()

	if err != nil {
		return fmt.Errorf("error reading Budget file")

	}

	for _, budget := range allBudgets {
		if budget.Month == month {
			fmt.Println("Budget already created for month:", month)
			return nil
		}
	}

	allBudgets = append(allBudgets, *newBudget)

	return WritetoBudgetFile(allBudgets, fileName)
}

func UpdateBudget(month string, amount int64) error {
	fileName, err := GetBudgetFilePath()

	if err != nil {
		return fmt.Errorf("error getting Budget file path")

	}

	allBudgets, err := ReadBudgetFile()

	if err != nil {
		return fmt.Errorf("error reading Budget file")

	}
	budgetIsPresent := false

	for index, budget := range allBudgets {
		if budget.Month == month {
			budgetIsPresent = true
			allBudgets[index].Amount = amount
		}
	}
	if !budgetIsPresent {
		fmt.Printf("No Budget for the month %s\n", month)
	}
	return WritetoBudgetFile(allBudgets, fileName)

}

func MonthBudget(month string) (int64, error) {

	allBudgets, err := ReadBudgetFile()

	if err != nil {
		return 0, fmt.Errorf("error reading Budget file")

	}
	var monthBudget int64 = 0
	for index, budget := range allBudgets {
		if budget.Month == month {
			monthBudget = allBudgets[index].Amount
		}
	}
	if monthBudget == 0 {
		fmt.Printf("%s does not have any Budget\n", month)
		return 0, nil
	}
	return monthBudget, err
}
