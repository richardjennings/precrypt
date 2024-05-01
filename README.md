# PreCrypt

PreCrypt generates a HTML page which can self load embedded AES GCM encrypted HTML, CSS, Javascript when a correct 
passcode is entered.

An example is available on [Github Pages](https://richardjennings.github.io/precrypt/).


## CLI Usage

```
make build-deps
go run main.go passphrasewhichneedstobe32bytes! > index.html
open index.html
```

## Library Usage

```
key := precrypt.Render(precrypt.RenderOptions{
    HtmlFiles:  []string{"example/index.html"},
    CssFiles:   []string{"example/style.css"},
    JsFiles:    []string{"example/index.js"},
    Out:        os.Stdout,
}
```



