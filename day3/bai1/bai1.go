package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const calendarColumnWidth = 3
const calendarPaddingLeftToRight = false
const calendarActiveDayColor = "\033[32m"
const calendarColorReset = "\033[0m"
const calendarIsPrintColorInActiveDay = true
const calendarPaddingCharacter = " "
const calendarNumberColumnInARow = 7

func main() {
	printCalendarTitle()
	printCalendarHeader()
	printCalendarBody()
}

/*
print calendar title
*/
func printCalendarTitle() {
	year, month, _ := time.Now().Date()
	fmt.Println(addPadding(""), month.String(), year)
}

/*
print calendar header
*/
func printCalendarHeader() {
	header := "Su Mo Tu We Th Fr Sa"
	for _, word := range strings.Fields(header) {
		fmt.Print(addPadding(word))
	}
	fmt.Println()
}

/*
print calendar body
*/
func printCalendarBody() {
	calendarBody := []string{}
	calendarBody = formatPrintingToFirstDayOfTheMonth(calendarBody, getWeekDayOfTheFirstDayOfTheMonth())
	calendarBody = formatPrintingDayInTheMonth(calendarBody)

	for i := 0; i < len(calendarBody); i++ {
		printDayInTheMonth(calendarBody[i])
		printBreakLineOneCalendarRow(i)
	}
	fmt.Println("")
}

/*
get the first and last day of the month
*/
func getFirstDayAndLastDayOfTheMonth() (int, int) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstDayOfTheMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastDayOfTheMonth := firstDayOfTheMonth.AddDate(0, 1, -1)

	return firstDayOfTheMonth.Day(), lastDayOfTheMonth.Day()
}

/*
get the week day of the first day in the month
*/
func getWeekDayOfTheFirstDayOfTheMonth() int {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstDayOfTheMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	return int(firstDayOfTheMonth.Weekday())
}

/*
format from beginning of week to first day of the month, sunday is 0, saturday is 6
*/
func formatPrintingToFirstDayOfTheMonth(calendarBody []string, firstWeekDay int) []string {
	for weekDay := 0; weekDay <= 6; weekDay++ {
		if firstWeekDay == weekDay {
			break
		}
		calendarBody = append(calendarBody, addPadding(""))
	}
	return calendarBody
}

/*
format each day in the month before printing
*/
func formatPrintingDayInTheMonth(calendarBody []string) []string {
	firstDay, lastDay := getFirstDayAndLastDayOfTheMonth()

	for i := firstDay; i <= lastDay; i++ {
		calendarBody = append(calendarBody, addPadding(strconv.Itoa(i)))
	}
	return calendarBody
}

/*
print the day with or without format active day
*/
func printDayInTheMonth(dayInTheMonth string) {
	_, _, activeDay := time.Now().Date()

	if calendarIsPrintColorInActiveDay && addPadding(strconv.Itoa(activeDay)) == dayInTheMonth {
		fmt.Print(string(calendarActiveDayColor), dayInTheMonth, string(calendarColorReset))
		return
	}
	fmt.Print(dayInTheMonth)
}

/*
print break line after each row of calendar
*/
func printBreakLineOneCalendarRow(i int) {
	if (i+1)%calendarNumberColumnInARow == 0 {
		fmt.Println("")
	}
}

/*
add padding character into string to make string has length that is equal calendar width
*/
func addPadding(s string) string {
	if calendarColumnWidth-len(s) < 0 {
		panic("the string is greater than column width")
	}

	if calendarPaddingLeftToRight {
		return s + strings.Repeat(calendarPaddingCharacter, calendarColumnWidth-len(s))
	}

	return strings.Repeat(calendarPaddingCharacter, calendarColumnWidth-len(s)) + s
}
