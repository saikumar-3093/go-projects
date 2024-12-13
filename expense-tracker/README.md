
GitHub Expense Tracker    


Sample solution for the expense tracker challenge from roadmap.sh.  

How to run  

git clone https://github.com/saikumar-3093/go-projects  
cd go-projects/expense-tracker  

build the cli app
go build -o expense-tracker(bash terminal)

run the app

1. to add new budget  
.\expense-tracker.exe set-budget --month [month number] --amount [budget amount]                          
Example: .\expense-tracker.exe set-budget --month 12 --amount 100                         

2. to update budget  
.\expense-tracker.exe update-budget --month [month number] --amount [budget amount]                          
Example: .\expense-tracker.exe update-budget --month 12 --amount 100

3. to get month budget  
.\expense-tracker.exe month-budget --month [month number]  
Example: .\expense-tracker.exe month-budget --month 12

4. to add new expense  
.\expense-tracker.exe add --description [any description] --amount [cost amount] --category [expnse category]  
Example: .\expense-tracker.exe add --description "Vegetables" --amount 20 --category "Groceries"

5. to delete an expense  
.\expense-tracker.exe delete --id [expense ID]  
Example: .\expense-tracker.exe delete --id 2

6. to update an expense  
 .\expense-tracker.exe update --id [expense ID] --description [any description] --amount [cost amount] --category [expnse category]  
Example:  .\expense-tracker.exe update --id 1 --description "cake" --amount 10 --category "Bakery"

7. to list all expense  
.\expense-tracker.exe list

8. to list expense based on category  
.\expense-tracker.exe list --category [category name]  
Example:  .\expense-tracker.exe list --category "Fruits"

9. to get total expense summary  
.\expense-tracker.exe summary

10. to get a particular month summary  
.\expense-tracker.exe summary --month [month number]  
Example: .\expense-tracker.exe summary --month 12

11. to create a csv file  
.\expense-tracker.exe convert-to-csv [json file to be converted to csv file] [csv file name]  
Example: .\expense-tracker.exe convert-to-csv "expenses.json" "expenses-csv"

Project-Url:   
https://github.com/saikumar-3093/go-projects/edit/main/expense-tracker


