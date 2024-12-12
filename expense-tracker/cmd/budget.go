package cmd

import (
	"fmt"
	"time"

	"github.com/saikumar-3093/go-projects/expense-tracker/filesystem"
	"github.com/spf13/cobra"
)

var mon = month
var amt = amount

func addBudget() *cobra.Command {
	newBudget := &cobra.Command{
		Use:   "set-budget",
		Short: "to set budget for the month",
		RunE: func(cmd *cobra.Command, args []string) error {
			if mon < 1 || mon > 12 {
				return fmt.Errorf("Invalid month")
			}
			mon := time.Month(mon).String()

			return filesystem.AddBudget(mon, amt)
		},
	}

	newBudget.Flags().Int64Var(&mon, "month", 0, "month to set budget")
	newBudget.Flags().Int64Var(&amt, "amount", 0, "month amount to set budget")
	return newBudget
}

func updateBudget() *cobra.Command {
	editBudget := &cobra.Command{
		Use:   "update-budget",
		Short: "to update budget for the month",
		RunE: func(cmd *cobra.Command, args []string) error {
			if mon < 1 || mon > 12 {
				return fmt.Errorf("Invalid month")
			}

			mon := time.Month(mon).String()
			return filesystem.UpdateBudget(mon, amt)

		},
	}
	editBudget.Flags().Int64Var(&mon, "month", 0, "month to set budget")
	editBudget.Flags().Int64Var(&amt, "amount", 0, "month amount to set budget")
	return editBudget
}

func monthBudget() *cobra.Command {
	monthBudget := &cobra.Command{
		Use:   "month-budget",
		Short: "to get budget for the month",
		RunE: func(cmd *cobra.Command, args []string) error {
			if mon < 1 || mon > 12 {
				return fmt.Errorf("Invalid month")
			}

			mon := time.Month(mon).String()
			budget, err := filesystem.MonthBudget(mon)

			if err != nil {
				return err
			}
			fmt.Printf("%s month budget: %v\n", mon, budget)
			return nil

		},
	}
	monthBudget.Flags().Int64Var(&mon, "month", 0, "month to set budget")

	return monthBudget
}
