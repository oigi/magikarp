package oss

//type QiNiu struct {
//	putPolicy storage.PutPolicy
//	mac       *qbox.Mac
//}
//
//func NewBucket(bucket, returnBody, CallbackURL, CallbackBody, CallbackBodyType string) *QiNiu {
//	accessKey := config.CONFIG.Oss.AccessKey
//	secretKey := config.CONFIG.Oss.SecretKey
//	return &QiNiu{
//		mac: qbox.NewMac(accessKey, secretKey),
//		putPolicy: storage.PutPolicy{
//			Scope:            bucket,
//			ReturnBody:       returnBody,
//			CallbackURL:      CallbackURL,
//			CallbackBody:     CallbackBody,
//			CallbackBodyType: CallbackBodyType,
//		},
//	}
//}
//
//func (q *QiNiu) PutFile(key string, reader io.Reader) error {
//	cfg := storage.Config{
//		Region:        &storage.ZoneHuanan, // 七牛云的存储区域，这个是华南区
//		UseHTTPS:      true,                // 使用 HTTPS 协议
//		UseCdnDomains: false,               // 使用 CDN 加速域名
//	}
//
//	upToken := q.putPolicy.UploadToken(q.mac)
//	formUploader := storage.NewFormUploader(&cfg)
//	ret := storage.PutRet{}
//
//	// 可选配置
//	putExtra := storage.PutExtra{
//		Params: map[string]string{
//			"x:name": "github logo",
//		},
//	}
//
//	return formUploader.Put(context.Background(), &ret, upToken, key, reader, -1, nil)
//}
//
//func (q *QiNiu) MakeUrl(key string)
