package main

import (
	"fmt"
	"github.com/richardjennings/precrypt/precrypt"
	"os"
)

func main() {
	key := os.Args[1]
	if err := precrypt.Render(precrypt.RenderOptions{
		HtmlFiles:  []string{"example/index.html"},
		CssFiles:   []string{"example/style.css"},
		JsFiles:    []string{"example/index.js"},
		LoaderJS:   "precrypt/dist/main.js",
		LoaderCSS:  "precrypt/dist/style.css",
		LoaderHTML: "precrypt/dist/index.html",
		Key:        []byte(key),
		Out:        os.Stdout,
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
