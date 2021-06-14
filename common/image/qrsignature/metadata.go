package qrsignature

import "fmt"

const (
	positionQRSize     = 185
	positionQRCapacity = QRCodeCapacityAlphanumeric
)

// Metadata represents service information about QR codes such as their names, coordinates on the canvas.
type Metadata struct {
	writer *QRCodeWriter
	qrCode *QRCode
}

// Encode encodes meta data to qrCode representation.
func (pos *Metadata) Encode(payloads []*Payload) error {
	var positionVector string

	for _, payload := range payloads {
		positionVector += fmt.Sprintf("%v:", payload.name.String())
		for _, qrCode := range payload.qrCodes {
			positionVector += fmt.Sprintf("%v,%v,%v;", qrCode.X, qrCode.Y, qrCode.Bounds().Size().X)
		}
	}

	data := []byte(positionVector)
	if err := positionQRCapacity.Validate(data); err != nil {
		return err
	}

	img, err := pos.writer.Encode(data)
	if err != nil {
		return err
	}
	pos.qrCode.Image = img
	return nil
}

// NewMetadata returns a new Metadata instance.
func NewMetadata() *Metadata {
	return &Metadata{
		writer: NewQRCodeWriter(positionQRSize),
		qrCode: NewEmptyQRCode(positionQRSize, positionQRSize),
	}
}
