package controller

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yeqown/go-qrcode"
)

type HexColor string

type QrParameters struct {
	fgColor HexColor
	bgColor HexColor
	data    string
}

func NewQRParameters(fgc, bgc HexColor, data string) *QrParameters {
	return &QrParameters{
		fgColor: fgc,
		bgColor: bgc,
		data:    data,
	}
}

// func (pqr *pazaQrShape) DrawFinder(ctx *qrcode.DrawContext) {

// 	w, h := ctx.Edge()
// 	color := ctx.Color()
// 	ctx.DrawRectangle(float64(ctx.UpperLeft().X), float64(ctx.UpperLeft().Y),
// 		float64(w), float64(h))
// 	ctx.SetColor(color)
// 	ctx.Fill()
// }

// func (pqr *pazaQrShape) Draw(ctx *qrcode.DrawContext) {

// 	w, h := ctx.Edge()
// 	upperLeft := ctx.UpperLeft()
// 	color := ctx.Color()
// 	smallerPercent := 0.60
// 	// choose a proper radius values
// 	radius := w / 2
// 	r2 := h / 2
// 	if r2 <= radius {
// 		radius = r2
// 	}

// 	// 80 percent smaller
// 	radius = int(float64(radius) * smallerPercent)

// 	cx, cy := upperLeft.X+w/2, upperLeft.Y+h/2 // get center point
// 	ctx.DrawCircle(float64(cx), float64(cy), float64(radius))
// 	ctx.SetColor(color)
// 	ctx.Fill()

// }

// func newShape() qrcode.IShape {
// 	return &pazaQrShape{}
// }

func GenerateQrWithParams(qrp QrParameters) {
	startTime := time.Now()
	qrc, err := qrcode.New(string(qrp.data),
		qrcode.WithQRWidth(20),
		qrcode.WithBgColorRGBHex(string(qrp.bgColor)),
		qrcode.WithFgColorRGBHex(string(qrp.fgColor)),
	)

	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	generatedTime := time.Now()
	// save file
	if err := qrc.Save(fmt.Sprintf("repo-qrcode%d.jpeg", time.Now().Nanosecond())); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
	endTime := time.Now()

	fmt.Printf("generation time :%-4s", generatedTime.Sub(startTime))
	fmt.Printf("  save time :%-4s", endTime.Sub(generatedTime))
	fmt.Printf("  total time taken :%-4s", endTime.Sub(startTime))
	fmt.Println()
}
func GeneratePazaQr(data string) {
	startTime := time.Now()
	qrc, err := qrcode.New(string(data),
		qrcode.WithQRWidth(20),
		qrcode.WithFgColorRGBHex("#34558a"),
		qrcode.WithLogoImageFilePNG("logo_2.png"))

	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	generatedTime := time.Now()
	// save file
	if err := qrc.Save(fmt.Sprintf("repo-qrcode%d.jpeg", time.Now().Nanosecond())); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
	endTime := time.Now()

	fmt.Printf("generation time :%-4s", generatedTime.Sub(startTime))
	fmt.Printf("  save time :%-4s", endTime.Sub(generatedTime))
	fmt.Printf("  total time taken :%-4s", endTime.Sub(startTime))
	fmt.Println()
}
func GenerateDefaulltQr(data string) {
	startTime := time.Now()
	qrc, err := qrcode.New(string(data),
		qrcode.WithQRWidth(20),
	)

	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	generatedTime := time.Now()
	// save file
	if err := qrc.Save(fmt.Sprintf("repo-qrcode%d.jpeg", time.Now().Nanosecond())); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
	endTime := time.Now()

	fmt.Printf("generation time :%-4s", generatedTime.Sub(startTime))
	fmt.Printf("  save time :%-4s", endTime.Sub(generatedTime))
	fmt.Printf("  total time taken :%-4s", endTime.Sub(startTime))
	fmt.Println()
}

func GenerateAndReturnPazaQr(w http.ResponseWriter, data string) {
	genmethod := time.Now()

	qrc, err := qrcode.New(string(data),
		qrcode.WithQRWidth(19),
		qrcode.WithFgColorRGBHex("#34558a"),
		qrcode.WithLogoImageFilePNG("logo_2.png"))

	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}
	made := time.Now()

	log.Info("TIME SPENT MAKING QR:", made.Sub(genmethod))

	qrc.SaveTo(w)

	log.Info("TIME SPENT SAVING TO WRITER:", time.Since(made))

}

//	generatedTime := time.Now()
// save file
// if err := qrc.Save(fmt.Sprintf("repo-qrcode%d.jpeg", time.Now().Nanosecond())); err != nil {
// 	fmt.Printf("could not save image: %v", err)
// }
// endTime := time.Now()
// fmt.Printf("generation time :%-4s", generatedTime.Sub(startTime))
// fmt.Printf("  save time :%-4s", endTime.Sub(generatedTime))
// fmt.Printf("  total time taken :%-4s", endTime.Sub(startTime))
// fmt.Println()
