package tokenapi

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func getS3Session(orgNum string) *session.Session {
	s3Endpoint := fmt.Sprintf("http://minio%s:9000", orgNum)

	// Configure to use MinIO Server
	s3Config := &aws.Config{
		// TODO: Change auth
		Credentials:      credentials.NewStaticCredentials("admin", "adminadmin", ""),
		Endpoint:         aws.String(s3Endpoint),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, _ := session.NewSession(s3Config)

	return newSession
}

// Uploads file to S3 storage
func UploadToS3(ctx contractapi.TransactionContextInterface, filePath string) (string, error) {
	mspId, err := ctx.GetClientIdentity().GetMSPID()

	if err != nil {
		log.Printf("Failed to get MSPID %v", err)
		return "", err
	}

	file, err := os.Open(filePath)

	if err != nil {
		log.Printf("failed to open file %s %v", filePath, err)
		return "", err
	}
	defer file.Close()

	log.Printf("Organisation: %+q\n", mspId)

	orgNum := getStringInBetween(mspId, "Org", "MSP")

	newSession := getS3Session(orgNum)

	uploader := s3manager.NewUploader(newSession)

	bucket := aws.String("input-files")

	fileName := filepath.Base(filePath)

	sha256sumCalculated, err := calculateSha256(filePath)

	namePlusSum := fmt.Sprintf("%s?%s", fileName, sha256sumCalculated)

	if err != nil {
		log.Printf("Failed to calculate checksum %v", err)
		return "", err
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: bucket,
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		log.Printf("Failed to upload file to S3 %v", err)
		return "", err
	}

	return namePlusSum, nil
}

// Downloads file from S3 to given directory
func DownloadFromS3(ctx contractapi.TransactionContextInterface, fileName string, sha256sum string, baseDir string) error {

	mspId, err := ctx.GetClientIdentity().GetMSPID()

	if err != nil {
		log.Printf("Failed to get MSPID %v", err)
		return err
	}

	log.Printf("Organisation: %+q\n", mspId)

	log.Printf("Downloading: %+q\n", fileName)

	orgNum := getStringInBetween(mspId, "Org", "MSP")

	newSession := getS3Session(orgNum)

	os.MkdirAll(baseDir, 0755)

	file, err := os.Create(baseDir + fileName)
	if err != nil {
		log.Printf("Failed to create file %v", err)
		return err
	}
	defer file.Close()

	bucket := aws.String("input-files")
	key := aws.String(fileName)

	downloader := s3manager.NewDownloader(newSession)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: bucket,
			Key:    key,
		})
	if err != nil {
		log.Printf("Failed to download file %v", err)
		return err
	}

	sha256sumCalculated, err := calculateSha256(baseDir + fileName)

	if err != nil {
		log.Printf("Failed to calculate checksum %v", err)
		return err
	}

	if sha256sum != sha256sumCalculated {
		if !DEBUG_MODE {
			return fmt.Errorf("security error: Invalid checksum")
		}
	}

	log.Printf("Downloaded file %s % bytes", file.Name(), numBytes)

	return nil
}

func calculateSha256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func getStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	e += s + e - 1
	return str[s:e]
}
