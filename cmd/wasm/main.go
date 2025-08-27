//go:build wasm

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Set("formatJSON", jsonWrapper())
	// make server wait on channel for go
	<-make(chan struct{})
}

// Marshal Indent func
func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

// JS wrapper
// func jsonWrapper() js.Func {
// 	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
// 		if len(args) != 1 {
// 			return "Invalid no of arguments passed"
// 		}
// 		inputJSON := args[0].String()
// 		fmt.Printf("input %s\n", inputJSON)
// 		pretty, err := prettyJson(inputJSON)
// 		if err != nil {
// 			fmt.Printf("unable to convert to json %s\n", err)
// 			return err.Error()
// 		}
// 		return pretty
// 	})
// 	return jsonFunc
// }

// JS wrapper direct to DOM
func jsonWrapper() js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return errors.New("Invalid no of arguments passed")
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			return errors.New("Unable to get document object")
		}
		jsonOuputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOuputTextArea.Truthy() {
			return errors.New("Unable to get output text area")
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			errStr := fmt.Sprintf("unable to parse JSON. Error %s occurred\n", err)
			return errors.New(errStr)
		}
		jsonOuputTextArea.Set("value", pretty)
		return nil
	})
	return jsonfunc
}
