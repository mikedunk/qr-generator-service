package controller

func (app *App) AddQrRoutes() {

	qr_router := app.Router.PathPrefix("/qr").Subrouter()

	qr_router.HandleFunc("/gen", app.GenerateQr).Methods("POST")
	qr_router.HandleFunc("/user/{code}", app.GetQr).Methods("GET")
	qr_router.HandleFunc("/upload", app.Upload).Methods("POST")
	/**
	SO i need endpoints
	- i need to be able to generate a
		- plain QR
		- qr with logo
		- qr with reference to image on cloudinary
		-
	-

	**/

}
