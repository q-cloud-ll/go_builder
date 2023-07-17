package upload

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"project/setting"
	"project/utils/urls"
)

type localUpload struct{}

func (local *localUpload) PutImage(data []byte, contentType string) (string, error) {
	key := generateImageKey(data, contentType)
	return local.PutObject(key, data, contentType)
}

func (local *localUpload) PutObject(key string, data []byte, contentType string) (string, error) {
	if err := os.MkdirAll("/", os.ModeDir); err != nil {
		return "", err
	}
	c := setting.Conf.LocalConfig
	filename := filepath.Join(c.Path, key)
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(filename, data, os.ModePerm); err != nil {
		return "", err
	}
	return urls.UrlJoin(c.Host, key), nil
}

func (local *localUpload) CopyImage(originUrl string) (string, error) {
	data, contentType, err := download(originUrl)
	if err != nil {
		return "", err
	}
	return local.PutImage(data, contentType)
}
