package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "To track expenses and budget",
}

func Execute() {
	rootCmd.AddCommand(addExpenseCmd())
	rootCmd.AddCommand(deleteExpenseCmd())
	rootCmd.AddCommand(updateExpenseCmd())
	rootCmd.AddCommand(listAllExpenses())
	rootCmd.AddCommand(expensesSummary())
	rootCmd.AddCommand(addBudget())
	rootCmd.AddCommand(updateBudget())
	rootCmd.AddCommand(monthBudget())
	rootCmd.Execute()
}
