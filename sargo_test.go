package sargo

import (
	// "fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type error interface {
	Error() string
}

func TestUsage(t *testing.T) {
	// deafault
	if GetUsage() != "go run main.go [options]" {
		t.Errorf("invalid default usage")
	}

	SetUsage("whatever usage")

	if GetUsage() != "whatever usage" {
		t.Errorf("invalid usage")
	}
}

func TestGet(t *testing.T) {
	options = nil
	Set("first", "f", "first default value", "first text description")
	Set("second", "s", "second default value", "second text description")
	Set("third", "t", "third default value", "third text description")

	// valid
	value, err := Get("first")
	if value != "first default value" {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = Get("f")
	if value != "first default value" {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	value, err = Get("fourth")
	if value != "" {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = Get("o")
	if value != "" {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetString(t *testing.T) {
	options = nil
	Set("first", "f", "first default value", "first text description")
	Set("second", "s", "second default value", "second text description")
	Set("third", "t", "third default value", "third text description")

	// valid
	value, err := GetString("first")
	if value != "first default value" {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetString("f")
	if value != "first default value" {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	value, err = GetString("fourth")
	if value != "" {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetString("o")
	if value != "" {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetInt(t *testing.T) {
	options = nil
	Set("first", "f", 10, "first text description")
	Set("second", "s", "20", "second text description")
	Set("third", "t", "just a string", "third text description")

	// valid
	// int 10
	value, err := GetInt("first")
	if value != 10 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetInt("f")
	if value != 10 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetInt("second")
	if value != 20 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetInt("s")
	if value != 20 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetInt("third")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetInt("fourth")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetInt("o")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetInt32(t *testing.T) {
	options = nil
	Set("first", "f", 10, "first text description")
	Set("second", "s", "20", "second text description")
	Set("third", "t", "just a string", "third text description")

	// valid
	// int 10
	value, err := GetInt32("first")
	if value != 10 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetInt32("f")
	if value != 10 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetInt32("second")
	if value != 20 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetInt32("s")
	if value != 20 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetInt32("third")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetInt32("fourth")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetInt32("o")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetInt64(t *testing.T) {
	options = nil
	Set("first", "f", 10, "first text description")
	Set("second", "s", "20", "second text description")
	Set("third", "t", "just a string", "third text description")

	// valid
	// int 10
	value, err := GetInt64("first")
	if value != 10 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetInt64("f")
	if value != 10 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetInt64("second")
	if value != 20 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetInt64("s")
	if value != 20 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetInt64("third")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetInt64("fourth")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetInt64("o")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetUint(t *testing.T) {
	options = nil
	Set("first", "f", 10, "first text description")
	Set("second", "s", "20", "second text description")
	Set("third", "t", "just a string", "third text description")

	// valid
	// int 10
	value, err := GetUint("first")
	if value != 10 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetUint("f")
	if value != 10 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetUint("second")
	if value != 20 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetUint("s")
	if value != 20 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetUint("third")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetUint("fourth")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetUint("o")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetUint32(t *testing.T) {
	options = nil
	Set("first", "f", 10, "first text description")
	Set("second", "s", "20", "second text description")
	Set("third", "t", "just a string", "third text description")

	// valid
	// int 10
	value, err := GetUint32("first")
	if value != 10 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetUint32("f")
	if value != 10 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetUint32("second")
	if value != 20 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetUint32("s")
	if value != 20 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetUint32("third")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetUint32("fourth")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetUint32("o")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetUint64(t *testing.T) {
	options = nil
	Set("first", "f", 10, "first text description")
	Set("second", "s", "20", "second text description")
	Set("third", "t", "just a string", "third text description")

	// valid
	// int 10
	value, err := GetUint64("first")
	if value != 10 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetUint64("f")
	if value != 10 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetUint64("second")
	if value != 20 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetUint64("s")
	if value != 20 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetUint64("third")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetUint64("fourth")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetUint64("o")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetFloat(t *testing.T) {
	options = nil
	Set("first", "f", 10.1, "first text description")
	Set("second", "s", "20.2", "second text description")
	Set("third", "t", "just a string", "third text description")

	// valid
	// int 10
	value, err := GetFloat("first")
	if value != 10.1 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetFloat("f")
	if value != 10.1 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetFloat("second")
	if value != 20.2 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetFloat("s")
	if value != 20.2 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetFloat("third")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetFloat("fourth")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetFloat("o")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetFloat32(t *testing.T) {
	options = nil
	Set("first", "f", 10.1, "first text description")
	Set("second", "s", "20.2", "second text description")
	Set("third", "t", "just a string", "third text description")

	// valid
	// int 10
	value, err := GetFloat32("first")
	if value != 10.1 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetFloat32("f")
	if value != 10.1 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetFloat32("second")
	if value != 20.2 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetFloat32("s")
	if value != 20.2 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetFloat32("third")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetFloat32("fourth")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetFloat32("o")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}
}

func TestGetFloat64(t *testing.T) {
	options = nil
	Set("first", "f", 10.1, "first text description")
	Set("second", "s", "20.2", "second text description")
	Set("third", "t", "just a string", "third text description")

	// valid
	// int 10
	value, err := GetFloat64("first")
	if value != 10.1 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetFloat64("f")
	if value != 10.1 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetFloat64("second")
	if value != 20.2 {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetFloat64("s")
	if value != 20.2 {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetFloat64("third")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetFloat64("fourth")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetFloat64("o")
	if value != 0 {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}

}

func TestGetBool(t *testing.T) {
	options = nil
	Set("true", "t", true, "true text description")
	Set("false", "f", "false", "false text description")
	Set("none", "n", "Not a boolean", "none text description")

	// valid
	// int 10
	value, err := GetBool("true")
	if !value {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetBool("t")
	if !value {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// valid
	// string to int 20
	value, err = GetBool("false")
	if value {
		t.Errorf("could not return the option value by name")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option name")
	}

	value, err = GetBool("f")
	if value {
		t.Errorf("could not return the option value by short cut")
	}

	if err != nil {
		t.Errorf("error should be nil for an existing option short cut")
	}

	// invalid
	// in options
	value, err = GetBool("none")
	if value {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	// not in options
	value, err = GetBool("fourth")
	if value {
		t.Errorf("return value should be empty for not existing option name")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option name")
	}

	if err.Error() != "Option \"fourth\" was not found" {
		t.Errorf("error message is invalid")
	}

	value, err = GetBool("o")
	if value {
		t.Errorf("return value should be empty for not existing option short cut")
	}

	if err == nil {
		t.Errorf("error should not be nil for not existent option short cut")
	}

	if err.Error() != "Option \"o\" was not found" {
		t.Errorf("error message is invalid")
	}

}

func TestPrintHelp(t *testing.T) {
	options = nil
	Set("first", "f", "first default value", "first text description")
	Set("second", "s", "second default value", "second text description")
	Set("third", "t", "third default value", "third text description")

	helpContent := `Usage:
  whatever usage

Options:
  -f, [--first]   # first text description
  -s, [--second]  # second text description
  -t, [--third]   # third text description
`

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintHelp()

	w.Close()

	out, _ := ioutil.ReadAll(r)

	os.Stdout = rescueStdout

	if strings.Trim(helpContent, " \n") != strings.Trim(string(out), " \n") {
		t.Errorf("Version error. \nShould \"%s\" .\nGot: \"%s\"", strings.Trim(helpContent, " \n"), strings.Trim(string(out), " \n"))
	}
}

func TestGetArgs(t *testing.T) {
	options = nil
	Set("first", "f", "first default value", "first text description")
	Set("second", "s", "second default value", "second text description")
	Set("third", "t", "third default value", "third text description")

	// testint args with equal
	args = []string{"program", "-f=foo", "--second=bar"}

	value, _ := Get("f")
	if value != "foo" {
		t.Errorf("value should be equal in args")
	}

	value, _ = Get("second")
	if value != "bar" {
		t.Errorf("value should be equal in args")
	}

	// testint args without equal
	args = nil
	args = []string{"program", "-f", "bar", "--second", "foo"}

	value, _ = Get("f")
	if value != "bar" {
		t.Errorf("value should be equal in args")
	}

	value, _ = Get("second")
	if value != "foo" {
		t.Errorf("value should be equal in args")
	}
}
