package precrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"os"
	"text/template"
)

//go:embed dist/index.html
var loaderHTML []byte

//go:embed dist/main.js
var loaderJs []byte

//go:embed dist/style.css
var loaderCSS []byte

type RenderOptions struct {
	HtmlFiles  []string
	CssFiles   []string
	JsFiles    []string
	LoaderJS   string
	LoaderCSS  string
	LoaderHTML string
	Key        []byte
	Out        io.Writer
}

func Render(opts RenderOptions) error {
	var err error
	if len(opts.Key) != 32 {
		return errors.New("invalid key length")
	}

	if opts.LoaderJS != "" {
		loaderJs, err = os.ReadFile(opts.LoaderJS)
		if err != nil {
			return err
		}
	}
	if opts.LoaderCSS != "" {
		loaderCSS, err = os.ReadFile(opts.LoaderCSS)
		if err != nil {
			return err
		}
	}
	if opts.LoaderHTML != "" {
		loaderHTML, err = os.ReadFile(opts.LoaderHTML)
		if err != nil {
			return err
		}
	}
	tmpl := template.Must(template.New("page").Parse(string(loaderHTML)))
	html, err := encryptFiles(opts.HtmlFiles, opts.Key)
	if err != nil {
		return err
	}
	css, err := encryptFiles(opts.CssFiles, opts.Key)
	if err != nil {
		return err
	}
	js, err := encryptFiles(opts.JsFiles, opts.Key)
	if err != nil {
		return err
	}
	data, err := json.Marshal(map[string]interface{}{"a": html, "b": css, "c": js})
	if err != nil {
		return err
	}
	return tmpl.Execute(opts.Out, map[string]interface{}{"data": string(data), "js": string(loaderJs), "css": string(loaderCSS)})
}

func encryptFiles(paths []string, key []byte) ([]string, error) {
	var ret []string
	for _, v := range paths {
		b, err := encryptFile(v, key)
		if err != nil {
			return nil, err
		}
		ret = append(ret, string(b))
	}
	return ret, nil
}

func encryptFile(path string, key []byte) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return encrypt(content, key)
}

func encrypt(text []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	nc := gcm.Seal(nonce, nonce, text, nil)
	b64 := make([]byte, base64.RawStdEncoding.EncodedLen(len(nc)))
	base64.RawStdEncoding.Encode(b64, nc)
	return b64, nil
}

func decrypt(b64 []byte, key []byte) ([]byte, error) {
	nc := make([]byte, base64.RawStdEncoding.DecodedLen(len(b64)))
	if _, err := base64.RawStdEncoding.Decode(nc, b64); err != nil {
		return nil, err
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := nc[:nonceSize], nc[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, nil)
}
