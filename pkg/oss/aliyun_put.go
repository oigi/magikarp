package oss

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/oigi/Magikarp/config"
	"io"
)

func (c *ClientOss) Put(objectType string, filename string, data io.Reader) error {
	err := c.Bucket.PutObject(
		objectType+"/"+filename,
		data,
		// oss.Callback(callback),
	)

	if err != nil {
		return err
	}

	return nil
}

func GetCallBackMap() string {
	callbackMap := make(map[string]string)
	callbackMap["callbackUrl"] = config.CONFIG.Oss.Callback
	callbackMap["callbackBody"] = "filename=${object}&size=${size}&mimeType=${mimeType}"
	callbackMap["callbackBodyType"] = "application/x-www-form-urlencoded"

	callbackBuffer := bytes.NewBuffer([]byte{})
	callbackEncoder := json.NewEncoder(callbackBuffer)
	callbackEncoder.SetEscapeHTML(false)
	callbackEncoder.Encode(callbackMap)

	callbackVal := base64.StdEncoding.EncodeToString(callbackBuffer.Bytes())
	return callbackVal
}
