package econf

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/liuchong/econf/internal/snake"
)

// check error to panic when read configurations
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// parseInt parse string to int64, panic on error
func parseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	check(err)
	return i
}

// parseFloat parse string to float64, panic on error
func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	check(err)
	return f
}

// parseBool parse string to bool, panic on error
func parseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	check(err)
	return b
}

// read file contains and convert to string, panic on error
func read(f string) string {
	dat, err := ioutil.ReadFile(f)
	check(err)
	return strings.TrimSpace(string(dat))
}

func toSnake(s string) string {
	return snake.ToSnake(s, '_', true)
}

// envStr convert string to required environment name in snake case
func envStr(dat interface{}, str string) string {
	datName := reflect.TypeOf(dat).Elem().Name()
	return toSnake(datName) + "_" + toSnake(str)
}

// envFileStr append _FILE to environment name from func envStr
func envFileStr(dat interface{}, str string) string {
	return envStr(dat, str) + "_FILE"
}

// isIntKind checks if the kind is any integer type
func isIntKind(k reflect.Kind) bool {
	return k == reflect.Int ||
		k == reflect.Int8 ||
		k == reflect.Int16 ||
		k == reflect.Int32 ||
		k == reflect.Int64
}

// isFloatKind checks if the kind is any float type
func isFloatKind(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
}

// handleIntValue handles setting an integer value
func handleIntValue(fld reflect.Value, v string) {
	fld.SetInt(parseInt(v))
}

// handleFloatValue handles setting a float value
func handleFloatValue(fld reflect.Value, v string) {
	fld.SetFloat(parseFloat(v))
}

// handleIntSlice handles setting an integer slice
func handleIntSlice(fld reflect.Value, strValues []string) {
	sliceType := fld.Type()
	intValues := reflect.MakeSlice(sliceType, len(strValues), len(strValues))
	for i, sv := range strValues {
		intValues.Index(i).SetInt(parseInt(sv))
	}
	fld.Set(intValues)
}

// handleFloatSlice handles setting a float slice
func handleFloatSlice(fld reflect.Value, strValues []string) {
	sliceType := fld.Type()
	floatValues := reflect.MakeSlice(sliceType, len(strValues), len(strValues))
	for i, sv := range strValues {
		floatValues.Index(i).SetFloat(parseFloat(sv))
	}
	fld.Set(floatValues)
}

// SetFieldByName set field of struct from environment variable or file by name
func SetFieldByName(s interface{}, name string) {
	SetFieldByNameWithSep(s, name, ",")
}

// SetFieldByNameWithSep set field with custom separator for slice types
func SetFieldByNameWithSep(s interface{}, name string, sep string) {
	elem := reflect.ValueOf(s).Elem()
	if elem.Kind() != reflect.Struct {
		return
	}

	// Get the field value
	fld := elem.FieldByName(name)
	if !fld.IsValid() {
		return
	}

	// Make private fields settable
	if !fld.CanSet() {
		fld = reflect.NewAt(fld.Type(), unsafe.Pointer(fld.UnsafeAddr())).Elem()
	}

	var v string
	// set up with environment variable
	if env := os.Getenv(envStr(s, name)); env != "" {
		v = env
	}
	// set up with environment file
	if env := os.Getenv(envFileStr(s, name)); env != "" {
		v = read(env)
	}
	if v != "" {
		if isIntKind(fld.Kind()) {
			handleIntValue(fld, v)
		} else if isFloatKind(fld.Kind()) {
			handleFloatValue(fld, v)
		} else if fld.Kind() == reflect.Bool {
			fld.SetBool(parseBool(v))
		} else if fld.Kind() == reflect.String {
			fld.SetString(v)
		} else if fld.Kind() == reflect.Slice {
			elemKind := fld.Type().Elem().Kind()
			strValues := strings.Split(v, sep)

			if isIntKind(elemKind) {
				handleIntSlice(fld, strValues)
			} else if isFloatKind(elemKind) {
				handleFloatSlice(fld, strValues)
			} else if elemKind == reflect.String {
				fld.Set(reflect.ValueOf(strValues))
			} else if elemKind == reflect.Bool {
				sliceType := fld.Type()
				boolValues := reflect.MakeSlice(sliceType, len(strValues), len(strValues))
				for i, sv := range strValues {
					boolValues.Index(i).SetBool(parseBool(sv))
				}
				fld.Set(boolValues)
			} else {
				panic(fmt.Sprintf("unsupported slice element type: %v for field %s", elemKind, name))
			}
		} else {
			panic(fmt.Sprintf("unsupported field type: %v for field %s", fld.Type(), name))
		}
	}
}

// SetFields looping set all field of struct
func SetFields(s interface{}) {
	SetFieldsWithSep(s, ",")
}

// SetFieldsWithSep looping set all field of struct with custom separator
func SetFieldsWithSep(s interface{}, sep string) {
	t := reflect.TypeOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		SetFieldByNameWithSep(s, name, sep)
	}
}
