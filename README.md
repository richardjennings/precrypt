# PreCrypt

PreCrypt generates a HTML page which can self load embedded AES GCM encrypted HTML, CSS, Javascript when a correct 
passcode is entered.

## Status

Work in progress currently only supporting embedding content from the `example` directory.

## Example

```
go run main.go passphrasewhichneedstobe32bytes! > index.html
open index.html
```
