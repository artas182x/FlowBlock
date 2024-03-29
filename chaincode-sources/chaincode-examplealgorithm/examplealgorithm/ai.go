package examplealgorithm

import (
	"archive/zip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/artas182x/FlowBlock/chaincode-sources/chaincode-computationtoken/tokenapi"
	"github.com/artas182x/FlowBlock/chaincode-sources/chaincode-medicaldata/medicaldatastructs"
	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nfnt/resize"
)

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func imageToTensor(img image.Image) (*tf.Tensor, error) {
	var image [1][150][150][3]float32
	for i := 0; i < 150; i++ {
		for j := 0; j < 150; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			image[0][j][i][0] = convertColor(r)
			image[0][j][i][1] = convertColor(g)
			image[0][j][i][2] = convertColor(b)
		}
	}
	return tf.NewTensor(image)
}
func convertColor(value uint32) float32 {
	return (float32(value>>8) - float32(127.5)) / float32(127.5)
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	log.Printf("Opening image: %+q\n", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func downloadPneumoniaModel(ctx contractapi.TransactionContextInterface, baseDir string) error {
	modelFilename := "pneumonia_model.zip"
	modelChecksum := "64c779ec892aab4c70bc7a81d8c072c0c7c0c2e78d52a112ec43f9c7f99c6d85"
	err := tokenapi.DownloadFromS3(ctx, modelFilename, modelChecksum, baseDir)
	if err != nil {
		return err
	}
	err = unzip(baseDir+modelFilename, baseDir+"model")
	if err != nil {
		return err
	}
	return nil
}

func classifyPneumonia(ctx contractapi.TransactionContextInterface, imageFilename string, checksum string, baseDir string) (float32, error) {
	err := tokenapi.DownloadFromS3(ctx, imageFilename, checksum, baseDir)
	if err != nil {
		return 0.0, err
	}

	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	img, err := getImageFromFilePath(baseDir + imageFilename)

	if err != nil {
		return 0.0, err
	}

	log.Printf("Resizing image\n")

	img = resize.Resize(150, 150, img, resize.Lanczos3)

	log.Printf("Image to tensor\n")

	tensor, err := imageToTensor(img)

	if err != nil {
		return 0.0, err
	}

	log.Printf("Loading model\n")

	model := tg.LoadModel(baseDir+"model", []string{"serve"}, nil)

	results := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_input_4", 0): tensor,
	})

	predictions := results[0]
	var predictionsVal = predictions.Value().([][]float32)

	return predictionsVal[0][0], nil
}

// Classify pneumonia based on XRay image
func (s *ExampleAlgorithmSmartContract) PneumoniaImageClassification(ctx contractapi.TransactionContextInterface, nonce string, medicalEntryId string) (*tokenapi.Ret, error) {

	medicalEntryIdDecoded, err := base64.StdEncoding.DecodeString(medicalEntryId)

	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: invalid medical entry id"), RetType: "string"}
		return &ret, nil
	}

	params := []string{"MedicalDataSmartContract:ReadMedicalEntry", string(medicalEntryIdDecoded), nonce}
	queryArgs := tokenapi.ParamsToHyperledgerArgs(params)

	log.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't read data from Blockchain"), RetType: "string"}
		return &ret, nil
	}

	var medicalEntry medicaldatastructs.MedicalEntry
	err = json.Unmarshal(response.Payload, &medicalEntry)
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't parse data from Blockchain"), RetType: "string"}
		return &ret, nil
	}

	if medicalEntry.MedicalEntryType != "s3img" {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: medical entry is not an image"), RetType: "string"}
		return &ret, nil
	}

	var baseDir = fmt.Sprintf("%s%d/", BASE_DIRECTORY, rand.Intn(100000))

	log.Printf("Cleanup: %+q\n", baseDir)
	os.RemoveAll(baseDir)

	err = downloadPneumoniaModel(ctx, baseDir)

	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't download model"), RetType: "string"}
		return &ret, nil
	}

	medicalEntryVals := strings.Split(medicalEntry.MedicalEntryValue, "?")
	fileName := medicalEntryVals[0]
	checksum := medicalEntryVals[1]

	log.Printf("Classifying: %+q\n", fileName)
	result, err := classifyPneumonia(ctx, fileName, checksum, baseDir)
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: error during classification"), RetType: "string"}
		return &ret, nil
	}

	if result > 0.5 {
		ret := tokenapi.Ret{RetValue: fmt.Sprintf("Pneumonia. Result: %f", result), RetType: "string"}
		return &ret, nil
	} else {
		ret := tokenapi.Ret{RetValue: fmt.Sprintf("No pneumonia. Result: %f", result), RetType: "string"}
		return &ret, nil
	}

}

// Calculates pneumonia cases over a time
func (s *ExampleAlgorithmSmartContract) XRayPneumoniaCases(ctx contractapi.TransactionContextInterface, nonce string, startDateTimestamp string, endDateTimestamp string) (*tokenapi.Ret, error) {

	params := []string{"MedicalDataSmartContract:GetMedicalEntries", "ChestXRay", startDateTimestamp, endDateTimestamp, nonce}
	queryArgs := tokenapi.ParamsToHyperledgerArgs(params)

	log.Printf("Starting computing: %+q\n", params)

	response := ctx.GetStub().InvokeChaincode("medicaldata", queryArgs, "")

	if response.Status != shim.OK {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't read data from Blockchain"), RetType: "string"}
		return &ret, nil
	}

	var medicalEntries []medicaldatastructs.MedicalEntry
	err := json.Unmarshal(response.Payload, &medicalEntries)
	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't parse data from Blockchain"), RetType: "string"}
		return &ret, nil
	}

	var baseDir = fmt.Sprintf("%s%d/", BASE_DIRECTORY, rand.Intn(100000))

	log.Printf("Cleanup: %+q\n", baseDir)
	os.RemoveAll(baseDir)

	err = downloadPneumoniaModel(ctx, baseDir)

	if err != nil {
		ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: can't download model"), RetType: "string"}
		return &ret, nil
	}

	cases := 0

	for _, arg := range medicalEntries {
		medicalEntryVals := strings.Split(arg.MedicalEntryValue, "?")
		fileName := medicalEntryVals[0]
		checksum := medicalEntryVals[1]
		log.Printf("Classifying: %+q\n", fileName)
		result, err := classifyPneumonia(ctx, fileName, checksum, baseDir)
		if err != nil {
			ret := tokenapi.Ret{RetValue: fmt.Sprintln("Error: error during classification"), RetType: "string"}
			return &ret, nil
		}
		if result > 0.5 {
			cases += 1
		}
	}

	ret := tokenapi.Ret{RetValue: fmt.Sprint(cases), RetType: "int64"}
	return &ret, nil
}
