package vcard

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// TextToBase64QRCodeImage creates a PNG QR code Base64 encoded for browser rendering
func TextToBase64QRCodeImage(text string, width, height int) string {

	// Create and scale the barcode
	qrCode, _ := qr.Encode(text, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, width, height)

	// create a writer into a byte buffer
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	png.Encode(writer, qrCode)
	writer.Flush()

	// convert the buffer bytes to base64 string
	imgBase64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())

	// Format base64 value into browser readable PNG image string
	return fmt.Sprintf("data:image/png;base64,%s", imgBase64Str)
}
