package sargo

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Option struct {
	Name          string
	ShortCut      string
	DefaultValue  interface{}
	Description   string
	MatcherRegexp *regexp.Regexp
}

var options []Option
var args []string
var usage string

func SetUsage(tUsage string) {
	usage = tUsage
}

func GetUsage() string {
	return usage
}

func Set(name string, shortCut string, defaultValue interface{}, description string) {
	var newOption Option

	newOption.Name = name
	newOption.ShortCut = shortCut
	newOption.DefaultValue = defaultValue
	newOption.Description = description
	newOption.MatcherRegexp = regexp.MustCompile(`\A((\-\-` + name + `)|(\-` + shortCut + `))(=|\z)`)

	options = append(options, newOption)
}

func Get(name string) (string, error) {
	var checkForEqualChar = regexp.MustCompile(`=`)
	var checkForNext bool
	var interfaceValue string

	option, err := searchOptionByName(name)
	if err != nil {
		return "", err
	}

	interfaceValue = fmt.Sprintf("%v", option.DefaultValue)
	checkForNext = false

	for _, value := range args {
		if checkForNext {
			interfaceValue = value
			checkForNext = false
			break
		} else if option.MatcherRegexp.MatchString(value) && checkForEqualChar.MatchString(value) {
			s := checkForEqualChar.Split(value, 2)
			interfaceValue = s[1]
			break
		} else if option.MatcherRegexp.MatchString(value) && !checkForEqualChar.MatchString(value) {
			checkForNext = true
		}
	}

	return interfaceValue, nil
}

func GetString(name string) (string, error) {
	return Get(name)
}

func GetInt(name string) (int, error) {
	return GetInt32(name)
}

func GetInt32(name string) (int, error) {
	intValue, err := parseInt(name, 32)
	return int(intValue), err
}

func GetInt64(name string) (int64, error) {
	intValue, err := parseInt(name, 64)
	return intValue, err
}

func GetUint(name string) (uint32, error) {
	return GetUint32(name)
}

func GetUint32(name string) (uint32, error) {
	intValue, err := parseUint(name, 32)
	return uint32(intValue), err
}

func GetUint64(name string) (uint64, error) {
	intValue, err := parseUint(name, 64)
	return intValue, err
}

func GetFloat(name string) (float32, error) {
	return GetFloat32(name)
}

func GetFloat32(name string) (float32, error) {
	floatValue, err := parseFloat(name, 32)
	return float32(floatValue), err
}

func GetFloat64(name string) (float64, error) {
	floatValue, err := parseFloat(name, 64)
	return floatValue, err
}

func GetBool(name string) (bool, error) {
	value, err := Get(name)
	if err != nil {
		return false, err
	} else {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return false, err
		} else {
			return boolValue, nil
		}
	}
}

func PrintHelp() {
	biggestShortCut := 0
	biggestName := 0

	for _, value := range options {
		if len(value.ShortCut) > biggestShortCut {
			biggestShortCut = len(value.ShortCut)
		}

		if len(value.Name) > biggestName {
			biggestName = len(value.Name)
		}
	}

	fmt.Println("Usage:")
	fmt.Println("  " + usage)
	fmt.Println("")
	fmt.Println("Options:")

	for _, value := range options {
		shortCut := `-` + value.ShortCut + `,`
		shortCutFormatter := `%-` + strconv.Itoa(biggestShortCut+2) + `v`

		name := `[--` + value.Name + `]`
		nameFormatter := `%-` + strconv.Itoa(biggestName+4) + `v`

		formatter := `  -` + shortCutFormatter + ` ` + nameFormatter + `  # %s` + "\n"

		fmt.Printf(formatter, shortCut, name, value.Description)
	}
}

func PrintHelpAndExit() {
	PrintHelp()
	os.Exit(0)
}

func searchOptionByName(name string) (Option, error) {
	for _, value := range options {
		if value.Name == name || value.ShortCut == name {
			return value, nil
		}
	}

	return Option{}, errors.New("Option \"" + name + "\" was not found")
}

func parseInt(name string, bitSite int) (int64, error) {
	value, err := Get(name)
	if err != nil {
		return 0, err
	} else {
		intValue, err := strconv.ParseInt(value, 10, bitSite)
		if err != nil {
			return 0, err
		} else {
			return intValue, nil
		}
	}
}

func parseUint(name string, bitSize int) (uint64, error) {
	value, err := Get(name)
	if err != nil {
		return 0, err
	} else {
		intValue, err := strconv.ParseUint(value, 10, bitSize)
		if err != nil {
			return 0, err
		} else {
			return intValue, nil
		}
	}
}

func parseFloat(name string, bitSite int) (float64, error) {
	value, err := Get(name)
	if err != nil {
		return 0, err
	} else {
		floatValue, err := strconv.ParseFloat(value, bitSite)
		if err != nil {
			return 0, err
		} else {
			return floatValue, nil
		}
	}
}

func init() {
	args = os.Args
	usage = "go run main.go [options]"
}
