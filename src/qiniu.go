package main

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"golang.org/x/net/context"
	"log"
	"os"
	"path/filepath"
)

type Qiniu struct {
	Bucket    string
	AccessKey string
	SecretKey string

	bucketManager *storage.BucketManager
}

func NewQiniu() Qiniu {
	var q = Qiniu{
		Bucket:    os.Getenv("bucket"),
		AccessKey: os.Getenv("accessKey"),
		SecretKey: os.Getenv("secretKey"),
	}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	q.bucketManager = storage.NewBucketManager(mac, &cfg)
	return q
}

func (q *Qiniu) Upload() {
	log.Println("Qiniu uploading...")
	// 获取当前云上文件
	fileList, err := q.getFileList()
	if err != nil {
		log.Fatalln(err)
	}
	if len(fileList) != 0 {
		// 删除文件
		err = q.cleanBucket(fileList)
		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, filePath := range GetAllFiles("./public/") {
		absPath, _ := filepath.Abs(filePath)
		_, fileName := filepath.Split(filePath)
		// 忽略文件
		if fileName[0] == '.' || fileName[0] == '_'{
			continue
		}

		filePath = filePath[7:] // 去除 public/

		log.Printf("Upload file: %s\n", filePath)
		err = q.uploadFile(absPath, filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("Upload successfully!")
}

func (q *Qiniu) uploadFile(absPath string, key string) error {
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket,
	}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, absPath, nil)
	if err != nil {
		return err
	}
	return nil
}

func (q *Qiniu) getFileList() ([]string, error) {
	limit := 1000
	marker := ""
	file := make([]string, 0)
	for {
		entries, _, nextMarker, hashNext, err := q.bucketManager.ListFiles(q.Bucket, "", "", marker, limit)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			file = append(file, entry.Key)
		}
		if hashNext {
			marker = nextMarker
		} else {
			break
		}
	}
	return file, nil
}

func (q *Qiniu) cleanBucket(file []string) error {
	deleteOps := make([]string, 0, len(file))
	for _, key := range file {
		deleteOps = append(deleteOps, storage.URIDelete(q.Bucket, key))
	}
	_, err := q.bucketManager.Batch(deleteOps)
	if err != nil {
		return err
	} else {
		return nil
	}
}
