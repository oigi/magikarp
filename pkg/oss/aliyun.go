package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/oigi/Magikarp/config"
)

type ClientOss struct {
	Client *oss.Client
	Bucket *oss.Bucket
}

func InitOss() (c *ClientOss, err error) {

	/*	// 设置连接数为10，每个主机的最大闲置连接数为20，每个主机的最大连接数为20。
		conn := oss.MaxConns(10, 20, 20)
		// 设置HTTP连接超时时间为20秒，HTTP读取或写入超时时间为60秒。
		time := oss.Timeout(20, 60)
		// 设置是否支持将自定义域名作为Endpoint，默认不支持。
		cname := oss.UseCname(false)
		// 设置HTTP的User-Agent头，默认为aliyun-sdk-go。
		userAgent := oss.UserAgent("aliyun-sdk-go")
		// 设置是否开启HTTP重定向，默认开启。
		redirect := oss.RedirectEnabled(true)
		// 设置是否开启SSL证书校验，默认关闭。
		verifySsl := oss.InsecureSkipVerify(false)
		// 设置代理服务器地址和端口。
		//proxy := oss.Proxy("yourProxyHost")
		// 设置代理服务器的主机地址和端口，代理服务器验证的用户名和密码。
		authProxy := oss.AuthProxy("yourProxyHost", "yourProxyUserName", "yourProxyPassword")
		// 开启CRC加密。
		crc := oss.EnableCRC(true)
		// 设置日志模式。
		logLevel := oss.SetLogLevel(oss.LogOff)

		// 创建OSSClient实例。
		// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
		client, err := oss.New(config.CONFIG.Oss.Endpoint, config.CONFIG.Oss.AccessKey, config.CONFIG.Oss.Secret, conn, time, cname, userAgent, authProxy, verifySsl, redirect, crc, logLevel)
		if err != nil {
			return nil, err
		}*/

	client, err := oss.New(config.CONFIG.Oss.Endpoint, config.CONFIG.Oss.AccessKey, config.CONFIG.Oss.Secret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(config.CONFIG.Oss.BucketName)
	if err != nil {
		return nil, err
	}

	return &ClientOss{Client: client, Bucket: bucket}, nil
}
