package main

/**
QR Geneartor
This is a dynamic QR generator.
What this generator does is, you give it data (preferably a link to a resource) and it encodes it
This way, you can always update the data on your end but the qr link would always point to it. very simple
**/

import "github.com/mikedunk/qr-generator-service/controller"

func main() {

	app := controller.NewApp()
	app.StartApplication()

	// mum := "{\"destination\":\"07016182391\",\"source\":\"08067036022\",\"radix\":\"08118765106\",\"email\":\"missgbosi@gmail.com\",\"scanner\":\"08118765106\",\"expiry\":\"Jun 20 2021 00:00:00 GMT+0100\"}"

	// raw := json.RawMessage(mum)
	// fmt.Println(raw)

	// man := string(raw)

	// fmt.Println("Man")

	// fmt.Println(mum)
	// fmt.Println(man)

	//qrParams4 := controller.NewQRParameters("1942a43b-28c8-424b-b4cc-ec833f8cf838", "#31576e", "#f2fcfa")

	//	GenerateDefaulltQr(*qrParams4)

	//GenerateBlackPazaQr(*qrParams3)
	//GenerateQr(*qrParams4)

}
