package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	days_in_months := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	scanner := bufio.NewScanner(os.Stdin)
	projectName := "dateDifferenceInDays"
	var dates []string
	fmt.Println("Welcome to the", projectName, "GO application")
	fmt.Println("Please Enter the two dates:")
	scanner.Scan()
	dateinput := scanner.Text()
	dates = strings.Split(dateinput, " ")
	fmt.Printf("Entered two dates are %s and %s\n", dates[0], dates[1])
	if dates[0] == dates[1] {
		fmt.Println("Enter two different dates but not the same")
		return
	}
	days1, errd1 := strconv.Atoi(strings.Split(dates[0], "/")[0])
	days2, errd2 := strconv.Atoi(strings.Split(dates[1], "/")[0])
	months1, errm1 := strconv.Atoi(strings.Split(dates[0], "/")[1])
	months2, errm2 := strconv.Atoi(strings.Split(dates[1], "/")[1])
	years1, erry1 := strconv.Atoi(strings.Split(dates[0], "/")[2])
	years2, erry2 := strconv.Atoi(strings.Split(dates[1], "/")[2])
	if years1 < 1900 || years2 > 2999 || dates[0] == "01/01/1900" || dates[1] == "31/12/2999" {
		fmt.Println("Invalid Input. Please enter date range between 01/01/1900 and 31/12/2999")
		return
	}
	if errd1 != nil || errd2 != nil || errm1 != nil || errm2 != nil || erry1 != nil || erry2 != nil {
		panic("error")
	}
	var monthDays1 int
	var monthDays2 int
	for i := 0; i < months1; i++ {
		monthDays1 += days_in_months[i+1]
	}
	for j := 0; j < months2; j++ {
		monthDays2 += days_in_months[j+1]
	}
	n1 := years1*365 + days1 + monthDays1 + number_of_leap_years(days1, months1, years1)
	n2 := years2*365 + days2 + monthDays2 + number_of_leap_years(days2, months2, years2)
	fmt.Printf("Counting only the number of days between the two given dates yields us : %v days\n", (n2-n1)-1)
}

func number_of_leap_years(d int, m int, y int) int {
	var res int
	if m > 2 {
		res = y/4 - y/100 + y/400
	}
	if m <= 2 {
		res = (y-1)/4 - (y-1)/100 + (y-1)/400
	}
	return res
}
