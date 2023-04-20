package main

import (
  "fmt"
  "github.com/dop251/goja"
  "os"
)

func main() {
  script, err := os.ReadFile("script/json-to-go.js")
  if err != nil {
    panic(err)
  }
  vm := goja.New()
  _, err = vm.RunString(string(script))
  if err != nil {
    panic(err)
  }
  jsonToGo, ok := goja.AssertFunction(vm.Get("jsonToGo"))
  if !ok {
    panic("AssertFunction jsonToGo error")
  }
  exampleCard, err := os.ReadFile("example_json/example_card.json")
  if err != nil {
    panic(err)
  }
  res, err := jsonToGo(goja.Undefined(), vm.ToValue(string(exampleCard)), vm.ToValue("Card"))
  if err != nil {
    panic(err)
  }
  fmt.Println(res.Export())

}
