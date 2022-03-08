package tokenapi

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// In debug mode for example file checksums are skipped, because example data contains invalid ones
const DEBUG_MODE = true

type Method struct {
	Name        string `json:"name"`
	Args        string `json:"args"`
	RetType     string `json:"retType"`
	Description string `json:"description"`
}

type Ret struct {
	RetValue string `json:retValue`
	RetType  string `json:retType`
}

type Token struct {
	ID             string    `json:"ID"` //Must be string
	UserRequested  string    `json:userRequested`
	ChaincodeName  string    `json:chaincodeName`
	Method         string    `json:method`
	Arguments      string    `json:arguments`
	Ret            Ret       `json:"ret,omitempty" metadata:"ret,optional"`
	TimeRequested  time.Time `json:timeRequested`
	ExpirationTime time.Time `json.expirationTime`
}

// Check whether nonce provided in parameter is equal to actual nonce. Can be used to check
// whether method has been executed by other method/chaincode. We should use this function
// to forbid users from direct access to algorithm function. Each algorithm should use this
// function before computation starts
// Nonce is in fact ctx.GetStub().GetCreator() converted to base64
func IsNonceValid(ctx contractapi.TransactionContextInterface, nonceStr string) (bool, error) {
	creatorByte, err := ctx.GetStub().GetCreator()
	if err != nil {
		return false, err
	}

	creator := base64.StdEncoding.EncodeToString(creatorByte)

	fmt.Printf("tokenapi:isNonceValid: Comparing GetCreator(): %s vs nonce: %s\n", creator, nonceStr)
	return creator == nonceStr, nil
}

// Downloads file from S3 to given directory
func DownloadFromS3(ctx contractapi.TransactionContextInterface, fileName string, sha256sum string, baseDir string) error {

	mspId, err := ctx.GetClientIdentity().GetMSPID()

	if err != nil {
		fmt.Println("Failed to get MSPID", err)
		return err
	}

	fmt.Printf("Organisation: %+q\n", mspId)

	fmt.Printf("Downloading: %+q\n", fileName)

	orgNum := getStringInBetween(mspId, "Org", "MSP")

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

	os.MkdirAll(baseDir, 0755)

	file, err := os.Create(baseDir + fileName)
	if err != nil {
		fmt.Println("Failed to create file", err)
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
		fmt.Println("Failed to download file", err)
		return err
	}

	sha256sumCalculated, err := calculateSha256(baseDir + fileName)

	if err != nil {
		fmt.Println("Failed to calculate checksum", err)
		return err
	}

	if sha256sum != sha256sumCalculated {
		if !DEBUG_MODE {
			return fmt.Errorf("security error: Invalid checksum")
		}
	}

	fmt.Println("Downloaded file", file.Name(), numBytes, "bytes")

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
