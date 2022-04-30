package storage

type CloudStorageService interface {
	UploadToCloud()
	GetFromCloud()
}
type Options struct {
	SaveLocal bool   `json:"save_local"`
	CloudName string `json:"cloud_name"` //use this option if you want the resource to be called sth else on cloud

}
