package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/mikedunk/qr-generator-service/lib"
	"github.com/mikedunk/qr-generator-service/model"
	"gorm.io/gorm"
)

func (app *App) GenerateQr(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error("Error Parsing Body :", err)
		lib.ERROR(w, http.StatusBadRequest, err)
		return
	}

	log.Println(string(body))

	qr_record := &model.Qr{}

	err = json.Unmarshal(body, qr_record)

	fmt.Println("unmarshalled")
	fmt.Println(qr_record.Data)

	if err != nil {
		log.Error("Error Unmarshalling Body :", err)
		lib.ERROR(w, http.StatusBadRequest, err)
		return
	}

	parsingComplete := time.Now()

	if err := app.Con.Create(qr_record).Error; err != nil {
		log.Error(err)
		lib.ERROR(w, http.StatusInternalServerError, lib.ErrUnableToCreateRecord)
		return
	}
	saveUser := time.Now()

	GenerateAndReturnPazaQr(w, qr_record.Code.String())

	generateTime := time.Now()

	log.Info("New Created record: ", qr_record.Code)

	log.Info("TIME SPENT PARSING BODY FROM REQUEST :", parsingComplete.Sub(startTime))
	log.Info("TIME SPENT SAVING RECORD TO DB :", saveUser.Sub(parsingComplete))
	log.Info("TIME SPENT GENERATING QR:", generateTime.Sub(saveUser))
	log.Info("TIME SPENT LOGGING:", time.Since(generateTime))

}

func (app *App) GetQr(w http.ResponseWriter, r *http.Request) {

	code := mux.Vars(r)["code"]
	log.Info("searched code", code)
	qr := &model.Qr{}

	if err := app.Con.First(qr, "code = ?", code).Error; err != nil {
		log.Error(err)
		lib.ERROR(w, http.StatusInternalServerError, lib.ErrUnableToFindRecord)
		return
	}
	//qr.Temp = temp
	lib.JSON(w, http.StatusOK, lib.Response{
		Success: true,
		Error:   "",
		Data:    qr,
	})

}

func (app *App) Upload(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		lib.ERROR(w, http.StatusMethodNotAllowed, errors.New("method isnt allowed"))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, app.Config.MaxUploadSize)

	if r.ContentLength > app.Config.MaxUploadSize {
		ErrMaxSizeEXceededc := fmt.Sprintf("the file you're uploading is too large, please upload a file less than %vMB", app.Config.MaxUploadSize/1000000)

		err := errors.New(ErrMaxSizeEXceededc)
		lib.ERROR(w, http.StatusBadRequest, err)
		return

	}
	if err := r.ParseMultipartForm(app.Config.MaxUploadSize); err != nil {
		lib.ERROR(w, http.StatusBadRequest, err)
		return
	}

	file, handler, err := r.FormFile("image")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		lib.ERROR(w, http.StatusUnprocessableEntity, errors.New("your upload can't be processed at the moment"))
		return
	}

	defer file.Close()

	fmt.Println(handler.Filename)

	//upload to cloud service provider //name can be supplied dynamically
	resp, er := app.UploadToCloud(file, handler.Filename)

	if er != nil {
		fmt.Println("Error Uploading to Cloud")
		fmt.Println(er)
		lib.ERROR(w, http.StatusUnprocessableEntity, errors.New("error uploading to cloud"))
		return
	}
	//register record in db

	GenerateAndReturnPazaQr(w, resp.SecureURL)

}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, err := strconv.Atoi(r.FormValue("page"))

		if err != nil {
			page = 1
		}

		if page == 0 {
			page = 1
		}

		pageSize, err := strconv.Atoi(r.FormValue("pageSize"))
		if err != nil {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
