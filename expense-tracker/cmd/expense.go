package cmd

import (
	"fmt"
	"time"

	"github.com/saikumar-3093/go-projects/expense-tracker/filesystem"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	desc     string
	amount   int64
	category string
	id       int64
	month    int64
)

func addExpenseCmd() *cobra.Command {

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "to add a new expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			desc, err := cmd.Flags().GetString("description")
			if err != nil {
				return fmt.Errorf("error getting description: %s", err)
			}
			amount, err := cmd.Flags().GetInt64("amount")
			if err != nil {
				return fmt.Errorf("error getting amount: %s", err)
			}
			category, err := cmd.Flags().GetString("category")
			if err != nil {
				return fmt.Errorf("error getting category: %s", err)
			}
			filesystem.AddExpense(desc, amount, category)
			return nil
		},
	}

	addCmd.Flags().StringVar(&desc, "description", "", "add description of the expense")
	addCmd.Flags().Int64Var(&amount, "amount", 0, "add amount of the expense")
	addCmd.Flags().StringVar(&category, "category", "", "add category of the expense")

	return addCmd
}

func deleteExpenseCmd() *cobra.Command {

	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "to delete an expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cmd.Flags().GetInt64("id")
			if err != nil {
				return fmt.Errorf("error getting id: %s", err)
			}
			err = filesystem.DeleteExpense(id)
			if err != nil {
				fmt.Println(err)
			}
			return nil
		},
	}

	deleteCmd.Flags().Int64Var(&id, "id", 0, "enter id of the expense")

	return deleteCmd
}

func updateExpenseCmd() *cobra.Command {

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "to update an expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cmd.Flags().GetInt64("id")
			if err != nil {
				return fmt.Errorf("error getting id: %s", err)
			}
			desc, err := cmd.Flags().GetString("description")
			if err != nil {
				return fmt.Errorf("error getting description: %s", err)
			}
			amount, err := cmd.Flags().GetInt64("amount")
			if err != nil {
				return fmt.Errorf("error getting amount: %s", err)
			}
			category, err := cmd.Flags().GetString("category")
			if err != nil {
				return fmt.Errorf("error getting category: %s", err)
			}
			err = filesystem.UpdateExpense(id, desc, amount, category)
			if err != nil {
				fmt.Println(err)
			}
			return nil
		},
	}

	updateCmd.Flags().Int64Var(&id, "id", 0, "id of the expense to be deleted")
	updateCmd.Flags().StringVar(&desc, "description", "", "add description of the expense")
	updateCmd.Flags().Int64Var(&amount, "amount", 0, "add amount of the expense")
	updateCmd.Flags().StringVar(&category, "category", "", "add category of the expense")

	return updateCmd
}

func listAllExpenses() *cobra.Command {
	listExpenses := &cobra.Command{
		Use:   "list",
		Short: "to list all expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			count := 0
			cmd.Flags().Visit(func(f *pflag.Flag) {
				count++
			})
			if count == 0 {
				filesystem.List()
			}
			if count == 1 {
				filesystem.CategoryExpenses(category)
			}

			if count > 1 {
				return fmt.Errorf("Please provide only one category")
			}
			return nil
		},
	}
	listExpenses.Flags().StringVar(&category, "category", "", "to list the expenses based on category")
	return listExpenses
}

func expensesSummary() *cobra.Command {
	allExpensesSummary := &cobra.Command{
		Use:   "summary",
		Short: "to summarize month expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			count := 0
			cmd.Flags().Visit(func(f *pflag.Flag) {
				count++
			})
			if count == 0 {

				filesystem.Summary()
			}
			if count == 1 {
				if month < 1 || month > 12 {
					return fmt.Errorf("Invalid month")
				}
				month := time.Month(month).String()

				monthexp, err := filesystem.MonthSummary(month)
				if err != nil {
					return err
				}
				fmt.Printf("%s month summary: %v\n", month, monthexp)
			}

			if count > 1 {
				return fmt.Errorf("please provide only one category")
			}
			return nil
		},
	}

	allExpensesSummary.Flags().Int64Var(&month, "month", 0, "to Summarize the month expenses")
	return allExpensesSummary
}
