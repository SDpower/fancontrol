package main

import (
	"fmt"
	"os"

	"strconv"

	"strings"

	"github.com/olekukonko/tablewriter"
)

func printUsage() {
	fmt.Println("Usage: fancontrol [options] <command> [arguments...]")
	fmt.Println("More help: fancontrol help")
}

func printHelp() {
	fmt.Println("Usage: fancontrol [options] <command> [arguments...]")

	//fmt.Println("Options:")

	fmt.Println(`Commands:
help
	Dispalys this help
ls, list
	Shows a list with information about cards installed
pls, plainlist
	Same as above but with ugly format. Thought for being used by other programs.
set <card> <"auto"|0-100>
	Setting the fan speed
	`)
}

func plainPrintListCards(command []string) {
	if len(command) == 1 {
		list := listCardsS()
		for _, f := range list {
			fmt.Printf("%s\t%s\t%s\t%s\n", f.card, f.name, f.temp, f.fan)
		}
		fmt.Println(len(list))
	} else {
		newCard(command[1])
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

func setFan(command []string) {
	if len(command) != 3 {
		fmt.Println("Wrong number of arguments")
		return
	}

	card := command[1]
	speed := command[2]

	if strings.EqualFold(speed, "auto") {
		fmt.Println("Setting speed of " + card + " to: auto")
		setFanMode(card, 2)
	} else {
		speedInt, _ := strconv.ParseInt(speed, 10, 32)
		if speedInt >= 0 && speedInt <= 100 {
			fmt.Println("Setting speed of " + card + " to: " + speed)
			setFanMode(card, 1)
			setFanSpeed(card, int(speedInt))
		} else {
			fmt.Println("Wrong speed argument")
		}
	}

}
