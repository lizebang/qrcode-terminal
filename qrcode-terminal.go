package qrcode

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

// Color is the color string in terminal.
type Color string

const (
	NormalBlack   Color = "\033[38;5;0m  \033[0m"
	NormalRed     Color = "\033[38;5;1m  \033[0m"
	NormalGreen   Color = "\033[38;5;2m  \033[0m"
	NormalYellow  Color = "\033[38;5;3m  \033[0m"
	NormalBlue    Color = "\033[38;5;4m  \033[0m"
	NormalMagenta Color = "\033[38;5;5m  \033[0m"
	NormalCyan    Color = "\033[38;5;6m  \033[0m"
	NormalWhite   Color = "\033[38;5;7m  \033[0m"

	BrightBlack   Color = "\033[48;5;0m  \033[0m"
	BrightRed     Color = "\033[48;5;1m  \033[0m"
	BrightGreen   Color = "\033[48;5;2m  \033[0m"
	BrightYellow  Color = "\033[48;5;3m  \033[0m"
	BrightBlue    Color = "\033[48;5;4m  \033[0m"
	BrightMagenta Color = "\033[48;5;5m  \033[0m"
	BrightCyan    Color = "\033[48;5;6m  \033[0m"
	BrightWhite   Color = "\033[48;5;7m  \033[0m"
)

// QRCode generates a QR code and prints it on the command line.
func QRCode(content string, fcolor, bcolor Color, level qrcode.RecoveryLevel) {
	qr, err := qrcode.New(content, level)
	if err != nil {
		fmt.Println(err)
		return
	}

	for ir, row := range qr.Bitmap() {
		lr := len(row)
		if ir == 0 || ir == 1 || ir == 2 ||
			ir == lr-1 || ir == lr-2 || ir == lr-3 {
			continue
		}
		for ic, col := range row {
			lc := len(qr.Bitmap())
			if ic == 0 || ic == 1 || ic == 2 ||
				ic == lc-1 || ic == lc-2 || ic == lc-3 {
				continue
			}
			if col {
				fmt.Print(fcolor)
			} else {
				fmt.Print(bcolor)
			}
		}
		fmt.Println()
	}
}
