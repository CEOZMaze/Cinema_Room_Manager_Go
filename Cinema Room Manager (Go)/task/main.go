package main

import (
	"fmt"
	"math"
	"os"
)

func askSize(inputRows, inputSeats int) (int, int) {
	if inputRows <= 0 {
		fmt.Println("Enter the number of rows:")
		fmt.Scan(&inputRows)
	}
	if inputSeats <= 0 {
		fmt.Println("Enter the number of seats in each row:")
		fmt.Scan(&inputSeats)
	}
	return inputRows, inputSeats
}

func createCustomSlice(inputRows int, inputSeats int) [][]string {
	customArray := make([][]string, inputRows)
	for i := range customArray {
		customArray[i] = make([]string, inputSeats)
	}
	return customArray
}
func populateCustom(inputRows int, inputSeats int) [][]string {
	customRoom := createCustomSlice(inputRows, inputSeats)
	for i := range customRoom {
		for j := range customRoom[i] {
			customRoom[i][j] = "S"

		}
	}
	return customRoom
}

func printCustomSeats(inputSeats int) {
	fmt.Println("Cinema:")
	fmt.Printf("  ")
	for j := 1; j <= inputSeats; j++ {
		fmt.Printf("%d ", j)
	}
	fmt.Println()

}

func printCustomRoom(customRoom [][]string) {

	for i := range customRoom {
		fmt.Printf("%d ", i+1)
		for j := range customRoom[i] {

			fmt.Printf("%s ", customRoom[i][j])
		}

		fmt.Println()

	}

}

func selectSeat(inputRows int, customRoom [][]string) (int, int) {
	var seatRow, seatColumn int

	for {
		fmt.Println("Enter a row number:")
		fmt.Scan(&seatRow)

		fmt.Println("Enter a seat number in that row:")
		fmt.Scan(&seatColumn)

		if seatRow <= 0 || seatRow > len(customRoom) {
			fmt.Println("Wrong input! Please enter a valid row.")
			continue
		}

		if seatColumn <= 0 || seatColumn > len(customRoom[0]) {
			fmt.Println("Invalid seat number! Please enter a valid seat.")
			continue
		}

		if customRoom[seatRow-1][seatColumn-1] == "B" {
			fmt.Println("That ticket has already been purchased!")
			continue
		}

		halfRows := math.Floor(float64(inputRows) / 2)
		if float64(seatRow) <= halfRows {
			fmt.Println("Ticket price: $10")
		} else {
			fmt.Println("Ticket price: $8")
		}

		break
	}

	return seatRow, seatColumn
}

func calculateIncome(inputRows, inputSeats int) int {

	totalSeats := inputRows * inputSeats

	if totalSeats <= 60 {
		return totalSeats * 10
	} else {
		halfRows := inputRows / 2
		backRows := inputRows - halfRows

		return (backRows * inputSeats * 8) + (halfRows * inputSeats * 10)
	}

}

func inputSelectedSeat(seatRow, seatColumn int, customRoom [][]string) [][]string {

	customRoom[seatRow-1][seatColumn-1] = "B"

	return customRoom
}

func evalPrice(selectedRow, inputRows, inputSeats int) int {

	halfRows := math.Floor(float64(inputRows) / 2)

	if float64(selectedRow) <= halfRows || inputSeats*inputRows <= 60 {

		return 10
	} else {
		return 8
	}

}

func lazySpaces() {
	fmt.Println(" ")
}

func mainMenu(inputSeats, selectedRow, selectedColumn, inputRows, currentEarnings, maxEarnings int, customRoom [][]string) {
	var ticketSold int
	for {
		lazySpaces()
		fmt.Println(" 1. Show the seats\n", "2. Buy a ticket\n", "3. Statistics\n", "0. Exit")

		choice := -1

		for choice < 0 {
			fmt.Scan(&choice)
		}

		switch {

		case choice == 1:
			lazySpaces()
			printCustomSeats(inputSeats)
			printCustomRoom(customRoom)
			lazySpaces()
		case choice == 2:

			lazySpaces()
			selectedRow = 0
			selectedColumn = 0

			selectedRow, selectedColumn = selectSeat(inputRows, customRoom)
			ticketSold++

			customRoom = inputSelectedSeat(selectedRow, selectedColumn, customRoom)

			currentEarnings += evalPrice(selectedRow, inputRows, inputSeats)
			lazySpaces()

		case choice == 3:
			totalSeats := inputRows * inputSeats

			percent := (float32(ticketSold) / float32(totalSeats)) * 100
			fmt.Printf("Number of purchased tickets: %d\n", ticketSold)
			fmt.Printf("Percentage: %.2f%%\n", percent)
			fmt.Printf("Current income: $%d\n", currentEarnings)
			fmt.Printf("Total income: $%d\n", maxEarnings)

		case choice == 0:
			os.Exit(0)
		}
	}
}

func main() {

	currentEarnings := 0

	var inputRows, inputSeats int

	var selectedRow, selectedColumn int
	inputRows, inputSeats = askSize(inputRows, inputSeats)
	maxEarnings := calculateIncome(inputRows, inputSeats)
	createCustomSlice(inputRows, inputSeats)
	customRoom := populateCustom(inputRows, inputSeats)

	mainMenu(inputSeats, selectedRow, selectedColumn, inputRows, currentEarnings, maxEarnings, customRoom)

}
