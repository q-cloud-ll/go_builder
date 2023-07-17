package upload

import (
	"bytes"
	"context"
	"project/setting"
	"sync"

	"github.com/mlogclub/simple/common/strs"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

type qiniuOssUpload struct {
	once   sync.Once
	bucket *storage.BucketManager
	auth   *qbox.Mac
	domain string
}

func (q *qiniuOssUpload) PutImage(data []byte, contentType string) (string, error) {
	if strs.IsBlank(contentType) {
		contentType = "image/jpeg"
	}
	key := generateImageKey(data, contentType)

	return q.PutObject(key, data, contentType)
}

func (q *qiniuOssUpload) PutObject(key string, data []byte, contentType string) (string, error) {
	bucket := q.getBucket()
	putPolicy := storage.PutPolicy{
		Scope: bucket + ":" + key,
	}

	upToken := putPolicy.UploadToken(q.auth)
	cfg := storage.Config{
		Zone: &storage.ZoneHuadong,
	}
	formUploader := storage.NewFormUploader(&cfg)
	var ret storage.PutRet
	var extra = &storage.PutExtra{}
	if strs.IsNotBlank(contentType) {
		extra.MimeType = contentType
	}
	if err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), int64(len(data)), extra); err != nil {
		return "", err
	}
	return q.domain + "/" + key, nil
}

func (q *qiniuOssUpload) CopyImage(originUrl string) (string, error) {
	data, contentType, err := download(originUrl)
	if err != nil {
		return "", err
	}
	return q.PutImage(data, contentType)
}

func (q *qiniuOssUpload) getBucket() string {
	q.once.Do(func() {
		qConfig := setting.Conf.QiNiuOssConfig
		q.auth = qbox.NewMac(qConfig.AccessKey, qConfig.SecretKey)
		sc := qiniuConfig()
		q.domain = qConfig.ImgPath
		q.bucket = storage.NewBucketManager(q.auth, sc)
	})

	return setting.Conf.QiNiuOssConfig.Bucket
}

func qiniuConfig() *storage.Config {
	qConfig := setting.Conf.QiNiuOssConfig
	cfg := storage.Config{
		UseHTTPS:      qConfig.UseHttps,
		UseCdnDomains: qConfig.UseCdnDomains,
	}

	switch qConfig.Zone {
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}

	return &cfg
}
