package upload

import (
	"project/setting"
	"strings"
	"sync"

	"github.com/mlogclub/simple/common/strs"
	"github.com/mlogclub/simple/common/urls"
)

type uploader interface {
	PutImage(data []byte, contentType string) (string, error)
	PutObject(key string, data []byte, contentType string) (string, error)
	CopyImage(originUrl string) (string, error)
}

var (
	qiniu = &qiniuOssUpload{
		once:   sync.Once{},
		bucket: nil,
		auth:   nil,
		domain: "",
	}
	local = &localUpload{}
)

func PutImage(data []byte, contentType string) (string, error) {
	return getUploader().PutImage(data, contentType)
}

func PutObject(key string, data []byte, contentType string) (string, error) {
	return getUploader().PutObject(key, data, contentType)
}

func CopyImage(url string) (string, error) {
	u1 := urls.ParseUrl(url).GetURL()
	u2 := urls.ParseUrl(setting.Conf.LocalConfig.Host).GetURL()
	// 本站host，不下载
	if u1.Host == u2.Host {
		return url, nil
	}
	return getUploader().CopyImage(url)
}

func getUploader() uploader {
	if IsEnabledOss() {
		return qiniu
	}

	return local
}

// IsEnabledOss 是否启用七牛云oss
func IsEnabledOss() bool {
	enable := setting.Conf.QiNiuOssConfig.Enable
	return strs.EqualsIgnoreCase(enable, "qiniuyun") || strs.EqualsIgnoreCase(enable, "oss") ||
		strs.EqualsIgnoreCase(enable, "qiniuyunOss")
}

// IsOssImageUrl 是否是存放在七牛云oss中的图片
func IsOssImageUrl(url string) bool {
	host := urls.ParseUrl(setting.Conf.QiNiuOssConfig.ImgPath).GetURL().Host
	return strings.Contains(url, host)
}
