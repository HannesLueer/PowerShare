package config

import (
	"PowerShare/helper"
	"log"
	"os"
	"path/filepath"
)

// file and directory paths

var curDir, _ = os.Getwd()
func CurDir() string {
	path := curDir
	return path
}

var certFilePath = filepath.Join("certs", "cert.pem")
var keyFilePath = filepath.Join("certs", "key.pem")

// GetCertFilePath returns the path where the cert.pem file is expected.
// If it does not exist the empty folder will be created.
func GetCertFilePath() string {
	if !helper.Exists(certFilePath) {
		err := helper.CreateDir("certs")
		if err != nil{
			log.Println(err)
		}
	}
	return certFilePath
}

// GetKeyFilePath returns the path where the key.pem file is expected.
// If it does not exist the empty folder will be created.
func GetKeyFilePath() string {
	if !helper.Exists(keyFilePath) {
		err := helper.CreateDir("certs")
		if err != nil{
			log.Println(err)
		}
	}
	return keyFilePath
}

