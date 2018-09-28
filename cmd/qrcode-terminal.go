package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/skip2/go-qrcode"

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

func main() {
	flag.Parse()
	content := flag.Arg(0)
	if content == "" {
		return
	}

	front, err := parseColor(frontColor)
	if err != nil {
		fmt.Println(err)
		return
	}
	back, err := parseColor(backgroundColor)
	if err != nil {
		fmt.Println(err)
		return
	}
	level, err := parseLevel(levelString)
	if err != nil {
		fmt.Println(err)
		return
	}

	qr.QRCode(content, front, back, level)
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

func parseColor(str string) (color qr.Color, err error) {
	s := strings.ToUpper(str)
	switch s {
	case "BLACK":
		color = qr.BrightBlack
	case "RED":
		color = qr.BrightRed
	case "GREEN":
		color = qr.BrightGreen
	case "YELLOW":
		color = qr.BrightYellow
	case "BLUE":
		color = qr.BrightBlue
	case "MAGENTA":
		color = qr.BrightMagenta
	case "CYAN":
		color = qr.BrightCyan
	case "WHITE":
		color = qr.BrightWhite
	default:
		err = fmt.Errorf("'%s' is not support.", str)
	}

	return
}

func parseLevel(str string) (level qrcode.RecoveryLevel, err error) {
	s := strings.ToUpper(str)
	switch s {
	case "L":
		level = qrcode.Low
	case "M":
		level = qrcode.Medium
	case "Q":
		level = qrcode.High
	case "H":
		level = qrcode.Highest
	default:
		err = fmt.Errorf("'%s' is not support.", str)
	}

	return
}
