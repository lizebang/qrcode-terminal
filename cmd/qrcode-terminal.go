package main

import (
	"flag"
	"fmt"

	qr "github.com/lizebang/qrcode-terminal"
)

var (
	frontColor      string
	backgroundColor string
	levelString     string
)

func init() {
	flag.Usage = printHelp
	flag.StringVar(&frontColor, "f", "black", "Front color")
	flag.StringVar(&backgroundColor, "b", "white", "Background color")
	flag.StringVar(&levelString, "l", "m", "Error correction level")
}

func printHelp() {

	helpStr := `QRCode generater terminal edition.

Supported background colors: [black, red, green, yellow, blue, magenta, cyan, white]
Supported front colors: [black, red, green, yellow, blue, magenta, cyan, white]
Supported error correction levels: [L, M, Q, H]
`
	fmt.Println(helpStr)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	content := flag.Arg(0)
	if content == "" {
		return
	}

	qr.QRCode(content, frontColor, backgroundColor, levelString)
}
