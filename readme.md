# Task Tracker

Solution for the [expense-tracker](https://github.com/robertd2000/expense-tracker-go) challenge from [roadmap.sh](https://roadmap.sh/) (without any external libraries).

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/robertd2000/expense-tracker-go.git
cd expense-tracker-go
```

Run the following command to build and run the project:

```bash
# To add a expense
go run . add --description "Dinner" --amount 10
# Expense added successfully (ID: 1)
go run . add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

# To update a expense
go run . update 1 update --id 1   --amount 20
# Expense updated successfully (ID: 1)

# To delete a expense
go run . delete --id 2
# Expense deleted successfully

# To list all expenses
go run . list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

# To show summary
go run . summary
# Total expenses: $30
go run . summary --month 3
# Total expenses for March: $30

```

[https://roadmap.sh/projects/expense-tracker](https://roadmap.sh/projects/expense-tracker)
