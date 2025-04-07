package econf

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

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

// SetFieldByName set field of struct from environment variable or file by name
func SetFieldByName(s interface{}, name string) {
	elem := reflect.ValueOf(s).Elem()
	if elem.Kind() != reflect.Struct {
		return
	}

	if fld := elem.FieldByName(name); fld.IsValid() && fld.CanSet() {
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
			if fld.Kind() == reflect.Int ||
				fld.Kind() == reflect.Int8 ||
				fld.Kind() == reflect.Int16 ||
				fld.Kind() == reflect.Int32 ||
				fld.Kind() == reflect.Int64 {
				fld.SetInt(parseInt(v))
			} else if fld.Kind() == reflect.String {
				fld.SetString(v)
			} else if fld.Kind() == reflect.Slice &&
				(fld.Type().Elem().Kind() == reflect.Int ||
					fld.Type().Elem().Kind() == reflect.Int8 ||
					fld.Type().Elem().Kind() == reflect.Int16 ||
					fld.Type().Elem().Kind() == reflect.Int32 ||
					fld.Type().Elem().Kind() == reflect.Int64) {
				strValues := strings.Split(v, ",")
				sliceType := fld.Type()
				intValues := reflect.MakeSlice(sliceType, len(strValues), len(strValues))
				for i, sv := range strValues {
					intValues.Index(i).SetInt(parseInt(sv))
				}
				fld.Set(intValues)
			} else if fld.Kind() == reflect.Slice && fld.Type().Elem().Kind() == reflect.String {
				values := strings.Split(v, ",")
				fld.Set(reflect.ValueOf(values))
			} else {
				panic(fmt.Sprintf("unsupported field type: %v for field %s", fld.Type(), name))
			}
		}
	}
}

// SetFields looping set all field of struct
func SetFields(s interface{}) {
	t := reflect.TypeOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		SetFieldByName(s, name)
	}
}
