package cmd

import (
	"fmt"

	"github.com/saikumar-3093/go-projects/expense-tracker/filesystem"
	"github.com/spf13/cobra"
)

func csvFileCmd() *cobra.Command {
	csvCmd := &cobra.Command{
		Use:   "convert-to-csv",
		Short: "to create csv file from json data",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("please provide inputfile name and outputfile parameters\nExample: convert-to-csv 'expenses.json' 'expenses.csv'")
			}
			inputFile := args[0]
			outFile := args[1]

			return filesystem.ConvertToCsv(inputFile, outFile)
		},
	}
	return csvCmd
}
