package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	endpint        = "oss-cn-chengdu.aliyuncs.com"
	ak             = "LTAI5tMvG5NA51eiH3ENZDaa"
	sk             = "vWOGbrPKQGmLVo4CKSgmAB62vdum10"
	bucketName     = "cloud-station"
	uploadFilePath = "go.mod"
)

var (
	help bool
)

func loadParam() error {
	flag.Parse()
	return nil
}

//检查程序的入参
func validate() error {
	if endpint == "" {
		return fmt.Errorf("endpint参数不能为空！")
	}
	if ak == "" || sk == "" {
		return fmt.Errorf("宝贝，乖哦！你有参数ak or sk未传呢！")
	}
	if uploadFilePath == "" {
		return fmt.Errorf("upload file path missied")
	}
	return nil
}

func main() {
	//读取用户输入的参数
	loadParam()
	//校验程序的入参
	if err := validate(); err != nil {
		fmt.Printf("上传的参数格式有问题，error:%s\n", err)
		os.Exit(2)
	}
	//执行文件的上传
	if err := upload(uploadFilePath); err != nil {
		fmt.Printf("upload file load errorm, %s\n", err)
		//退出时返回的状态码，0为成功，非0为失败
		os.Exit(1)
	}
	fmt.Printf("upload file :%s success\n", uploadFilePath)
}

func upload(filePath string) error {
	client, err := oss.New(endpint, ak, sk)
	if err != nil {
		//HandleError(err)
		return err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		//HandleError(err)
		return err
	}
	//第一个参数：上传到oss里面的文件key（路径），go.sum --> 2021/7/21/go.sum
	//第二个参数：需要上传的文件路径
	err = bucket.PutObjectFromFile(filePath, filePath)
	if err != nil {
		//HandleError(err)
		return err
	}

	//打印下载URL
	//sts,临时授权token(有效期1天)
	signedURL,err :=bucket.SignURL(filePath,oss.HTTPGet,60*60*24)
	if err!=nil{
		//
		return fmt.Errorf("sign file download url error,%s",err)
	}
	fmt.Printf("下载链接：%s\n",signedURL)
	fmt.Println("\n注意：文件下载有效期为1天，中转站保存时间为3天，请技术下载")
	return nil
}

func usage() {
	fmt.Fprint(os.Stderr, `cloud-station version:0.0.1
Usage:cloud-station [-h] -f <upload_file_path>
Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&help,"h",false,"help usage")
    flag.StringVar(&uploadFilePath,"f","","upload file path")

}
