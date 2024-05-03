package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/richardjennings/precrypt/pkg/precrypt"
	"github.com/spf13/cobra"
	"os"
)

var htmlFiles []string
var cssFiles []string
var jsFiles []string
var key string

var build = &cobra.Command{
	Use:  "",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		out := args[0]
		fh, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		defer func() { _ = fh.Close() }()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		opts := precrypt.RenderOptions{
			HtmlFiles: htmlFiles,
			CssFiles:  cssFiles,
			JsFiles:   jsFiles,
			Out:       fh,
		}
		if len(key) > 0 {
			opts.Key, err = hex.DecodeString(key)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		if len(key) == 0 {
			// generate key
			bytes := make([]byte, 32)
			if _, err := rand.Read(bytes); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			opts.Key = bytes
			fmt.Println(hex.EncodeToString(bytes))
		}
		if err := precrypt.Render(opts); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	build.Flags().StringArrayVar(&htmlFiles, "html", []string{}, "--html <file1.html>,<file2.html>")
	build.Flags().StringArrayVar(&cssFiles, "css", []string{}, "--css <file1.css>,<file2.css>")
	build.Flags().StringArrayVar(&jsFiles, "js", []string{}, "--js <file1.js>,<file2.js>")
	build.Flags().StringVar(&key, "key", "", "--key hexencoded32bytes")
}

func Execute() {
	err := build.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
