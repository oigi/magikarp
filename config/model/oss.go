package model

type Oss struct {
	Endpoint   string `json:"endpoint"yaml:"endpoint"`
	AccessKey  string `json:"accessKey"yaml:"accessKey"`
	Secret     string `json:"secret"yaml:"secret"`
	BucketName string `json:"bucketName"yaml:"bucketName"`
	Callback   string `json:"callback"yaml:"callback"`
}
