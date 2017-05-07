package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func readSysFile(path string) string {
	f, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(f))
}

func readSysFileToInt(path string) int {
	val, _ := strconv.ParseInt(readSysFile(path), 10, 64)

	return int(val)
}

func overrideSysFile(path, data string) {
	err := ioutil.WriteFile(path, []byte(data), 0644)

	if err != nil {
		panic(err.Error())
	}
}

func overrideSysFileWithInt(path string, data int) {
	overrideSysFile(path, strconv.FormatInt(int64(data), 10))
}

func percentage(x, a, b int) int {

	fx := float64(x)
	fa := float64(a)
	fb := float64(b)

	val := (fx - fa) / (fb - fa)
	val *= 100
	return int(val)
}

type card struct {
	card string
	name string
	temp string
	fan  string
	fanm int
}

func newCard(c string) card {

	return card{c, getCardName(c), getTemperatureAsString(c), getFanSpeedAsString(c), getFanMode(c)}

}

// List card information

func listCards() []string {
	list := make([]string, 0)
	files, _ := ioutil.ReadDir("/sys/class/drm")
	for _, f := range files {
		matched, _ := regexp.MatchString("^card\\d$", f.Name())
		if matched {
			list = append(list, f.Name())
		}
	}
	return list
}

func listCardsS() []card {
	list := make([]card, 0)
	files, _ := ioutil.ReadDir("/sys/class/drm")
	for _, f := range files {
		matched, _ := regexp.MatchString("^card\\d$", f.Name())
		if matched {
			list = append(list, newCard(f.Name()))
		}
	}
	return list
}

// Sysfs getters

func getFanSpeed(card string) int {
	min := readSysFileToInt("/sys/class/drm/" + card + "/device/hwmon/hwmon1/pwm1_min")
	max := readSysFileToInt("/sys/class/drm/" + card + "/device/hwmon/hwmon1/pwm1_max")
	val := readSysFileToInt("/sys/class/drm/" + card + "/device/hwmon/hwmon1/pwm1")
	return percentage(val, min, max)
}

func getFanSpeedAsString(card string) string {
	return strconv.FormatInt(int64(getFanSpeed(card)), 10)
}

func getCardName(card string) string {
	name := readSysFile("/sys/class/drm/" + card + "/device/hwmon/hwmon1/name")
	return name
}

func getTemperature(card string) float64 {
	tempStr := readSysFile("/sys/class/drm/" + card + "/device/hwmon/hwmon1/temp1_input")
	tempStr = tempStr[:2] + "." + tempStr[2:4]

	temp, _ := strconv.ParseFloat(tempStr, 64)
	return temp
}

func getTemperatureAsString(card string) string {
	tempStr := readSysFile("/sys/class/drm/" + card + "/device/hwmon/hwmon1/temp1_input")
	tempStr = tempStr[:2] + "." + tempStr[2:3]
	return tempStr
}

func getFanMode(card string) int {
	tempStr := readSysFileToInt("/sys/class/drm/" + card + "/device/hwmon/hwmon1/pwm1_enable")
	return tempStr
}

func setFanMode(card string, m int) {
	overrideSysFileWithInt("/sys/class/drm/"+card+"/device/hwmon/hwmon1/pwm1_enable", m)
}

func setFanSpeed(card string, p int) {
	min := readSysFileToInt("/sys/class/drm/" + card + "/device/hwmon/hwmon1/pwm1_min")
	max := readSysFileToInt("/sys/class/drm/" + card + "/device/hwmon/hwmon1/pwm1_max")

	speed := int(((float64(p)/100.0)*(float64(max)-float64(min)) + float64(min)))

	overrideSysFileWithInt("/sys/class/drm/"+card+"/device/hwmon/hwmon1/pwm1", speed)
}
