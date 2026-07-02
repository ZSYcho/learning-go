package main

import (
	"errors"
	"fmt"
)

type ConfigItem struct {
	Key   string
	Value interface{}
	IsSet bool
}

/*
%v  - the default formatting (e.g. 5000, true, "hello")
%+v - default formatting, but adds field names for structs (e.g. {Key:API_URL Value:... IsSet:true})
%#v - Go-syntax representation (e.g. main.ConfigItem{Key:"API_URL", Value:"...", IsSet:true})
%T  - type name (e.g. main.ConfigItem, int, string)
%s  - string (or value's String() method); wrong type prints %!s(int=5000)
%d  - integer in base 10 (e.g. 5000)
%f  - floating point (%.2f for 2 decimals, e.g. 1.20)
%t  - for boolean (true / false)
%q  - for double-quoted string (e.g. "hello")
%%  - present % itself (a literal percent sign)
%w  - wrap an error (only works with fmt.Errorf, not Printf)
*/
func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s, Value: %s, IsSet:, %t", c.Key, c.Value, c.IsSet)
}

func main() {

	appName := "EnvParser"
	version := 1.2
	prot := 8080
	isEnabled := true

	status := fmt.Sprintf("Application: %s (Version: %.2f) running on port %d. Enabled: %t",
		appName, version, prot, isEnabled)
	fmt.Println(status)

	item1 := ConfigItem{Key: "API_URL", Value: "http://localhost:3000/api", IsSet: true}
	item2 := ConfigItem{Key: "TIMEOUT_MS", Value: 5000, IsSet: true}
	item3 := ConfigItem{Key: "DEBUG_MODE", Value: false, IsSet: false}

	fmt.Printf("Item 1 (%%v): %v\n", item1)

	fmt.Printf("Item 2 (%%+v): %+v\n", item2)

	fmt.Printf("Item 3 (%%#v): %#v\n", item3)

	err := errors.New("test")

	fmt.Errorf("here is an error on port %d: %w", prot, err) // %w - wrapping the error into this fmt

}
