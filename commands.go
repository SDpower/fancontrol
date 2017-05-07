package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/olekukonko/tablewriter"
)

func printUsage() {
	fmt.Println("Usage: fanctl [options] <command> [arguments...]")
	fmt.Println("More help: fanctl help")
}

func printHelp() {
	fmt.Println("Usage: fanctl [options] <command> [arguments...]")

	//fmt.Println("Options:")

	fmt.Println("Commands:")
	fmt.Println("help \t Prints this help")
}

func plainPrintListCards(command []string) {
	if len(command) == 1 {
		list := listCardsS()
		for _, f := range list {
			fmt.Printf("%s\t%s\t%s\t%s\n", f.card, f.name, f.temp, f.fan)
		}
		fmt.Println(len(list))
	}
}

func prettyPrintListCards(command []string) {
	list := listCardsS()

	data := make([][]string, 0)
	for _, f := range list {
		data = append(data, []string{f.card, f.name, f.temp, f.fan, strconv.FormatInt(int64(f.fanm), 10)})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Card", "Name", "Temp (°C)", "Fan Speed (%)", "Fan mode"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func printGetTemperature(command []string) {
	if len(command) == 1 {
		fmt.Println("Printing all card temps")
		list := listCards()
		for _, c := range list {
			fmt.Printf(c+"\t%.1f°C\n", getTemperature(c))
		}
	} else if len(command) == 2 {
		fmt.Printf("%.1f°C\n", getTemperature(command[1]))
	}
}

func printGetFanSpeed(command []string) {
	if len(command) == 1 {
		fmt.Println("Printing all fan speeds")
		list := listCards()
		for _, c := range list {
			fmt.Println(c + "\t" + getFanSpeedAsString(c) + "%")
		}
	} else if len(command) == 2 {
		fmt.Println(getFanSpeedAsString(command[1]) + "%")
	}
}
