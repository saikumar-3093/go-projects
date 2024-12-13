package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "expense-tracker",
	Short: `To track expenses and budget
	project-repo https://github.com/saikumar-3093/go-projects/blob/main/expense-tracker`,
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
	rootCmd.AddCommand(csvFileCmd())
	rootCmd.Execute()
}
