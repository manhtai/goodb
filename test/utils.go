package test

import "fmt"

func printStruct(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}
