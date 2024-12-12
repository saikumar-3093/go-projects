package main

import (
	"github.com/saikumar-3093/go-projects/expense-tracker/cmd"
)

func main() {
	// monthNames := map[int]string{
	// 	1:  "January",
	// 	2:  "February",
	// 	3:  "March",
	// 	4:  "April",
	// 	5:  "May",
	// 	6:  "June",
	// 	7:  "July",
	// 	8:  "August",
	// 	9:  "September",
	// 	10: "October",
	// 	11: "November",
	// 	12: "December",
	// }
	// fmt.Println("building expense tracker")

	// filesystem.AddBudget("December", 100)

	// filesystem.UpdateExpense(7, "Buy vegetables", 10, "Veg")
	// filesystem.MonthSummary(monthNames[12])
	// filesystem.CategoryExpenses("others")

	cmd.Execute()

}
