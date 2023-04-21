package main

import (
  "errors"
  "fmt"
  "github.com/dop251/goja"
  "os"
)

func main() {
  code, err := generate("example_json/example_card.json")
  if err != nil {
    panic(err)
  }
  fmt.Println(code)
}

func generateStructModel(jsonFile string) (string, error) {
  script, err := os.ReadFile("script/json-to-go.js")
  if err != nil {
    return "", err
  }
  vm := goja.New()
  _, err = vm.RunString(string(script))
  if err != nil {
    return "", err
  }
  jsonToGo, ok := goja.AssertFunction(vm.Get("jsonToGo"))
  if !ok {
    return "", errors.New("AssertFunction jsonToGo error")
  }
  exampleCard, err := os.ReadFile(jsonFile)
  if err != nil {
    return "", err
  }
  scriptReturn, err := jsonToGo(goja.Undefined(), vm.ToValue(string(exampleCard)), vm.ToValue("Card"))
  if err != nil {
    return "", err
  }
  exported := scriptReturn.Export()
  exportedMap, ok := exported.(map[string]interface{})
  if !ok {
    return "", errors.New("type assert error, except map[string]interface{}")
  }
  code, ok := exportedMap["go"].(string)
  if !ok {
    return "", errors.New("type assert error, except string")
  }
  return code, nil
}

func generate(jsonFile string) (string, error) {
  code := "package feishu_cards\n\n"
  structCode, err := generateStructModel(jsonFile)
  if err != nil {
    return "", err
  }
  code += structCode
  return code, nil
}
