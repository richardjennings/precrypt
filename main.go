package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/richardjennings/precrypt/internal/precrypt"
	"os"
	"text/template"
)

//go:embed template/index.html
var indexTemplate string

//go:embed dist/main.js
var loaderJs string

//go:embed dist/style.css
var loaderCss string

//go:embed example/index.html
var exampleHtml []byte

//go:embed example/style.css
var exampleCss []byte

//go:embed example/index.js
var exampleJs []byte

func main() {
	key := os.Args[1]
	htmlEnc, err := precrypt.Encrypt(exampleHtml, []byte(key))
	e(err)
	cssEnc, err := precrypt.Encrypt(exampleCss, []byte(key))
	e(err)
	jsEnc, err := precrypt.Encrypt(exampleJs, []byte(key))
	e(err)
	tmpl := template.Must(template.New("page").Parse(indexTemplate))
	data, err := json.Marshal(map[string]interface{}{
		"a": []string{string(htmlEnc)},
		"b": []string{string(cssEnc)},
		"c": []string{string(jsEnc)},
	})
	e(err)
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"data": string(data), "js": loaderJs, "css": loaderCss})
	e(err)
}

func e(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
