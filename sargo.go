package sargo

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Option struct {
	Name          string
	ShortCut      string
	DefaultValue  interface{}
	Description   string
	rawValue      string
	matcherRegexp *regexp.Regexp
}

var options []Option

var usage string = "go run main.go [options]"

func GetUsage() string {
	return usage
}

func SetUsage(message string) {
	usage = message
}

func Set(newOption Option) {
	newOption.matcherRegexp = regexp.MustCompile(`\A((\-\-` + newOption.Name + `)|(\-` + newOption.ShortCut + `))(=|\z)`)
	newOption.rawValue = getRawValue(newOption)

	options = append(options, newOption)
}

func getRawValue(option Option) string {
	var checkForEqualChar = regexp.MustCompile(`=`)

	args := os.Args[1:]
	for index, value := range args {
		if option.matcherRegexp.MatchString(value) {
			if checkForEqualChar.MatchString(value) {
				s := checkForEqualChar.Split(value, 2)
				return s[1]
			} else if index < len(args)-1 && !strings.HasPrefix(args[index+1], "-") {
				return args[index+1]
			} else {
				return "TRUE"
			}
		}
	}

	return ""
}

func Get(name string) (string, error) {
	option, err := searchOptionByName(name)
	if err != nil {
		return "", err
	}

	if option.rawValue == "" {
		return fmt.Sprintf("%v", option.DefaultValue), nil
	} else {
		return option.rawValue, nil
	}
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

		formatter := `  ` + shortCutFormatter + ` ` + nameFormatter + `  # %s` + "\n"

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

func parseInt(name string, bitSize int) (int64, error) {
	value, err := Get(name)
	if err != nil {
		return 0, err
	} else {
		intValue, err := strconv.ParseInt(value, 10, bitSize)
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

func parseFloat(name string, bitSize int) (float64, error) {
	value, err := Get(name)
	if err != nil {
		return 0, err
	} else {
		floatValue, err := strconv.ParseFloat(value, bitSize)
		if err != nil {
			return 0, err
		} else {
			return floatValue, nil
		}
	}
}
