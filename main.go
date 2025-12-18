package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nurashi/advanced-programming-1/Bank"
	"github.com/nurashi/advanced-programming-1/Company"
	"github.com/nurashi/advanced-programming-1/Library"
	"github.com/nurashi/advanced-programming-1/Shapes"
)

func libraryMenu() {
	libra := Library.Library{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1) Add Book")
		fmt.Println("2) Borrow Book")
		fmt.Println("3) Return Book")
		fmt.Println("4) List Available Books")
		fmt.Println("5) Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())

			id := libra.GetNextID()
			book := Library.Book{
				ID:         id,
				Title:      title,
				Author:     author,
				IsBorrowed: false,
			}
			libra.AddBook(book)
			fmt.Printf("Book added successfully with ID: %d\n", id)

		case "2":
			fmt.Print("Enter Book ID to borrow: ")
			scanner.Scan()
			id, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			book, err := libra.BorrowBook(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Successfully borrowed: %s\n", book.Title)
			}

		case "3":
			fmt.Print("Enter Book ID to return: ")
			scanner.Scan()
			id, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			err := libra.ReturnBook(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully!")
			}

		case "4":
			books, err := libra.ListAvailableBooks()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("\nAvailable Books:")
				for _, book := range books {
					fmt.Printf("ID: %d, Title: %s, Author: %s\n",
						book.ID, book.Title, book.Author)
				}
			}

		case "5":
			fmt.Println("Exiting Library Management System...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func shapesDemo() {
	shapes := []Shapes.Shape{
		Shapes.NewRectangle(5, 3),
		Shapes.NewCircle(4),
		Shapes.NewSquare(6),
		Shapes.NewTriangle(3, 4, 5),
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d - Area: %.2f, Perimeter: %.2f\n",
			i+1, shape.Area(), shape.Perimeter())
	}
}

func companyDemo() {
	comp := Company.NewCompany("linkedin")

	fullTime1 := Company.FullTimeEmployee{
		ID:         1,
		Name:       "John Doe",
		Department: "Engineering",
		Salary:     75000.00,
	}
	fullTime2 := Company.FullTimeEmployee{
		ID:         2,
		Name:       "Jane Smith",
		Department: "Marketing",
		Salary:     65000.00,
	}

	partTime1 := Company.PartTimeEmployee{
		ID:          3,
		Name:        "Bob Wilson",
		Department:  "Support",
		Salary:      25.00,
		HoursOfWork: 20.0,
	}
	partTime2 := Company.PartTimeEmployee{
		ID:          4,
		Name:        "Alice Brown",
		Department:  "Design",
		Salary:      30.00,
		HoursOfWork: 15.0,
	}

	comp.AddEmployee(fullTime1)
	comp.AddEmployee(fullTime2)
	comp.AddEmployee(partTime1)
	comp.AddEmployee(partTime2)
	employees := comp.ListEmployees()
	for _, emp := range employees {
		fmt.Println(emp.GetDetails())
	}
}

func bankMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Account Number: ")
	scanner.Scan()
	accountNumber := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Account Holder Name: ")
	scanner.Scan()
	holderName := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Initial Balance: ")
	scanner.Scan()
	initialBalance, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

	account := Bank.NewBankAccount(accountNumber, holderName, initialBalance)
	fmt.Printf("\nAccount created for %s with balance: $%.2f\n", holderName, initialBalance)

	for {
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter deposit amount: ")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

			err := account.Deposit(amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Deposited %.2f. New balance: %.2f\n", amount, account.GetBalance())
			}

		case "2":
			fmt.Print("Enter withdrawal amount: ")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

			err := account.Withdraw(amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Withdrew %.2f. New balance: %.2f\n", amount, account.GetBalance())
			}

		case "3":
			fmt.Printf("Current balance: %.2f\n", account.GetBalance())

		case "4":
			account.PrintTransactionHistory()

		case "5":
			fmt.Println("Exiting Bank Account System...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}


// accually this functions are used(and can be used), but to test other one's I have commented some of them.
func main() {
	// libraryMenu()
	// shapesDemo()
	// companyDemo()
	bankMenu()
}
