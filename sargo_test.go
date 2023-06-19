package sargo

import (
	"io"
	"os"
	"strings"
	"testing"
)

type error interface {
	Error() string
}

var firstOptionForStringTest = Option{
	Name:         "first",
	ShortCut:     "f",
	DefaultValue: "first default value",
	Description:  "first text description",
}

var secondOptionForStringTest = Option{
	Name:         "second",
	ShortCut:     "s",
	DefaultValue: "second default value",
	Description:  "second text description",
}

var thirdOptionForStringTest = Option{
	Name:         "third",
	ShortCut:     "t",
	DefaultValue: "third default value",
	Description:  "third text description",
}

var firstOptionForIntTest = Option{
	Name:         "first",
	ShortCut:     "f",
	DefaultValue: 10,
	Description:  "first text description",
}

var secondOptionForIntTest = Option{
	Name:         "second",
	ShortCut:     "s",
	DefaultValue: "20",
	Description:  "second text description",
}

var thirdOptionForIntTest = Option{
	Name:         "third",
	ShortCut:     "t",
	DefaultValue: "just a string",
	Description:  "third text description",
}

var firstOptionForFloatTest = Option{
	Name:         "first",
	ShortCut:     "f",
	DefaultValue: 10.1,
	Description:  "first text description",
}

var secondOptionForFloatTest = Option{
	Name:         "second",
	ShortCut:     "s",
	DefaultValue: "20.2",
	Description:  "second text description",
}

var thirdOptionForFloatTest = Option{
	Name:         "third",
	ShortCut:     "t",
	DefaultValue: "just a string",
	Description:  "third text description",
}

var greenOptionForBoolTest = Option{
	Name:         "green",
	ShortCut:     "g",
	DefaultValue: true,
	Description:  "green text description",
}

var redOptionForBoolTest = Option{
	Name:         "red",
	ShortCut:     "r",
	DefaultValue: "false",
	Description:  "red text description",
}

var invalidOptionForBoolTest = Option{
	Name:         "invalid",
	ShortCut:     "i",
	DefaultValue: "Not a boolean",
	Description:  "none text description",
}

func TestGet(t *testing.T) {
	options = nil

	Set(firstOptionForStringTest)
	Set(secondOptionForStringTest)
	Set(thirdOptionForStringTest)

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
	Set(firstOptionForStringTest)
	Set(secondOptionForStringTest)
	Set(thirdOptionForStringTest)

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
	Set(firstOptionForIntTest)
	Set(secondOptionForIntTest)
	Set(thirdOptionForIntTest)

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
	Set(firstOptionForIntTest)
	Set(secondOptionForIntTest)
	Set(thirdOptionForIntTest)

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
	Set(firstOptionForIntTest)
	Set(secondOptionForIntTest)
	Set(thirdOptionForIntTest)

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
	Set(firstOptionForIntTest)
	Set(secondOptionForIntTest)
	Set(thirdOptionForIntTest)

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
	Set(firstOptionForIntTest)
	Set(secondOptionForIntTest)
	Set(thirdOptionForIntTest)

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
	Set(firstOptionForIntTest)
	Set(secondOptionForIntTest)
	Set(thirdOptionForIntTest)

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
	Set(firstOptionForFloatTest)
	Set(secondOptionForFloatTest)
	Set(thirdOptionForFloatTest)

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
	Set(firstOptionForFloatTest)
	Set(secondOptionForFloatTest)
	Set(thirdOptionForFloatTest)

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
	Set(firstOptionForFloatTest)
	Set(secondOptionForFloatTest)
	Set(thirdOptionForFloatTest)

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
	Set(greenOptionForBoolTest)
	Set(redOptionForBoolTest)
	Set(invalidOptionForBoolTest)

	// green option = TRUE
	value, err := GetBool("green")
	if value {
		t.Log("[PASSED] default value for the option \"green\" is TRUE")
	} else {
		t.Error("[FAILED] default value for the option \"green\" expect to be TRUE, but got FALSE")
	}

	if err == nil {
		t.Log("[PASSED] error is NIL for option \"green\"")
	} else {
		t.Error("[FAILED] error expected to be nil for option \"green\", but got NOT NIL")
	}

	value, _ = GetBool("g")
	if value {
		t.Log("[PASSED] default value for the shortcut option \"g\" is TRUE")
	} else {
		t.Error("[FAILED] default value for the shortcut option \"g\" expect to be TRUE, but got FALSE")
	}

	// red option = FALSE
	value, err = GetBool("red")
	if !value {
		t.Log("[PASSED] default value for the option \"red\" is FALSE")
	} else {
		t.Error("[FAILED] default value for the option \"red\" expect to be FALSE, but got TRUE")
	}

	if err == nil {
		t.Log("[PASSED] error is NIL for a option \"red\"")
	} else {
		t.Error("[FAILED] error expected to be nil for a option \"red\", but got NOT NIL")
	}

	value, _ = GetBool("r")
	if !value {
		t.Log("[PASSED] default value for the shortcut option \"r\" is FALSE")
	} else {
		t.Error("[FAILED] default value for the shortcut option \"r\" expect to be FALSE, but got TRUE")
	}

	// invalid option - FALSE
	value, err = GetBool("invalid")
	if !value {
		t.Log("[PASSED] default value for the option \"invalid\" is FALSE")
	} else {
		t.Error("[FAILED] default value for the option \"invalid\" expect to be FALSE, but got TRUE")
	}

	if err != nil {
		t.Log("[PASSED] error expected to be nil for a option \"invalid\", but got NOT NIL")
	} else {
		t.Error("[FAILED] error is NIL for a option \"invalid\"")
	}

	// option does not exist
	value, err = GetBool("fourth")
	if !value {
		t.Log("[PASSED] default value for an option that does not exist is FALSE")
	} else {
		t.Errorf("[FAILED] default value for an option that does not exist expected to be FALSE, but got TRUE")
	}

	if err != nil {
		t.Log("[PASSED] error is NOT NIL for a option does not exist")
	} else {
		t.Error("[FAILED] error expected to be NOT NIL for an option that does not exist, but got NIL")
	}

	if err.Error() == "Option \"fourth\" was not found" {
		t.Log("[PASSED] when option is not in Set, the error message is 'Option \"option-name\" was not found'")
	} else {
		t.Errorf("[FAILED] when option is not in Set, the error message is 'Option \"option-name?\" was not found' but got '%s'", err.Error())
	}

}

func TestPrintHelp(t *testing.T) {
	options = nil
	Set(firstOptionForStringTest)
	Set(secondOptionForStringTest)
	Set(thirdOptionForStringTest)

	SetUsage("whatever usage")

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

	out, _ := io.ReadAll(r)

	os.Stdout = rescueStdout

	if strings.Trim(helpContent, " \n") == strings.Trim(string(out), " \n") {
		t.Log("[PASSED] Print help is OK")
	} else {
		t.Errorf("[FAILED] Print help error: \nExpected:\n\"%s\"\n\nGot:\n\"%s\"", strings.Trim(helpContent, " \n"), strings.Trim(string(out), " \n"))
	}
}

func TestGetArgsWithEqualSymbol(t *testing.T) {
	os.Args = []string{"program", "-f=foo", "--second=bar"}

	options = nil
	Set(firstOptionForStringTest)
	Set(secondOptionForStringTest)
	Set(thirdOptionForStringTest)

	value, _ := Get("f")
	if value == "foo" {
		t.Log("[PASSED] value of param \"f\" is \"foo\"")
	} else {
		t.Errorf("[FAILED] value of params \"f\" excpected be \"foo\", but got \"%s\"", value)
	}

	value, _ = Get("second")
	if value != "bar" {
		t.Errorf("value should be equal in args")
	}
}

func TestGetArgsWithoutEqualSymbol(t *testing.T) {
	os.Args = []string{"program", "-f", "bar", "--second", "foo"}

	options = nil
	Set(firstOptionForStringTest)
	Set(secondOptionForStringTest)
	Set(thirdOptionForStringTest)

	value, _ := Get("f")
	if value == "bar" {
		t.Log("[PASSED] option \"f\" is \"bar\"")
	} else {
		t.Errorf("[FAILED] option \"f\" expected to be \"bar\", but got %s", value)
	}

	value, _ = Get("second")
	if value == "foo" {
		t.Log("[PASSED] option \"second\" is \"foo\"")
	} else {
		t.Errorf("[FAILED] option \"second\" expected to be \"foo\", but got %s", value)
	}
}

func TestGetExplicityBooleanArgs(t *testing.T) {
	os.Args = []string{"program", "-r", "--second", "foo"}

	options = nil
	Set(redOptionForBoolTest)
	Set(secondOptionForStringTest)

	value, _ := GetBool("r")
	if value {
		t.Log("[PASSED] option \"r\" is TRUE")
	} else {
		t.Error("[FAILED] option \"r\" expected to be TRUE, but got FALSE")
	}
}

func TestGetHiddenBooleanArgs(t *testing.T) {
	os.Args = []string{"program", "--second", "foo"}

	options = nil
	Set(redOptionForBoolTest)
	Set(secondOptionForStringTest)

	value, _ := GetBool("f")
	if !value {
		t.Log("[PASSED] option \"f\" is FALSE")
	} else {
		t.Error("[FAILED] option \"f\" expected to be FALSE, but got TRUE")
	}
}
