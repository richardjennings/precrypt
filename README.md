# PreCrypt

PreCrypt generates a HTML page which can self load embedded AES GCM encrypted HTML, CSS, Javascript when a correct 
passcode is entered.

An example is available on [Github Pages](https://richardjennings.github.io/precrypt/).


## CLI Usage

Key is hex encoded and is generated if not supplied.

```
precrypt --html example/index.html --css example/style.css --js example/index.js --key 329625b9767075c799e90499c59f4e775c0c0ca8c8320b99fc485ba68add025b index.html
```

## Library Usage

```
err := precrypt.Render(precrypt.RenderOptions{
    HtmlFiles:  []string{"example/index.html"},
    CssFiles:   []string{"example/style.css"},
    JsFiles:    []string{"example/index.js"},
    Key:        []byte{"passphrasewhichneedstobe32bytes!"},
    Out:        os.Stdout,
}
```



