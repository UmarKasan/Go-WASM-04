# Go-WASM-04

```bash
GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm
```

```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ~/Documents/webassembly/assets/
```

```console input
formatJSON('{"website":"golangbot.com", "tutorials": [{"title": "Strings", "url":"/strings/"}]}')
```

```input
{"website":"golangbot.com", "tutorials": [{"title": "Strings", "url":"/strings/"}, {"title":"maps", "url":"/maps/"}, {"title": "goroutines","url":"/goroutines/"}]}
```

## Marshal Indent Func

```go
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
```

## Go to Js type

```go
func jsonWrapper() js.Func {
        jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
                if len(args) != 1 {
                        return "Invalid no of arguments passed"
                }
                inputJSON := args[0].String()
                fmt.Printf("input %s\n", inputJSON)
                pretty, err := prettyJson(inputJSON)
                if err != nil {
                        fmt.Printf("unable to convert to json %s\n", err)
                        return err.Error()
                }
                return pretty
        })
        return jsonFunc
}
```

## make server wait on channel for go

```go
        <-make(chan struct{})
```
