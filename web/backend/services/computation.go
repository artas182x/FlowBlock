package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/models"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/vars"
	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-computationtoken/tokenapi"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

var cli *gocelery.CeleryClient
var backendRedis gocelery.RedisCeleryBackend

const TASK_NAME = "worker.computeOnBlockchain"

var computations []models.ComputationData

func runTask(Certificate string, PrivateKey string, MspID string, TokenId string) models.TaskResult {

	log.Printf("Computation %s started\n", TokenId)

	login := models.Login{Certificate: Certificate,
		PrivateKey: PrivateKey,
		MspID:      MspID}

	out, err := SubmitTransaction(login, vars.ComputationTokenChaincodeName, vars.ComputationTokenSmartContractName, "Compute", TokenId)

	result := models.TaskResult{}
	result.Finished = true

	if err != nil {
		result.Error = err
		return result
	}

	var token *tokenapi.Token
	err = json.Unmarshal(out, &token)
	if err != nil {
		result.Error = err
		return result
	}

	result.Result = *token

	log.Printf("Computation %s finished\n", TokenId)

	return result
}

func InitCelery() {
	redisHost := os.Getenv("REDIS_URL")
	if redisHost == "" {
		redisHost = "redis://localhost:6379"
	}
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisHost)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	backendRedis = gocelery.RedisCeleryBackend{Pool: redisPool}

	cli, _ = gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&backendRedis,
		5,
	)

	cli.Register(TASK_NAME, runTask)

	//computations := make([]ComputationData)

	cli.StartWorker()
}

func QueueComputation(Login models.Login, TokenId string) (*models.ComputationData, error) {
	asyncResult, err := cli.Delay(TASK_NAME, Login.Certificate, Login.PrivateKey, Login.MspID, TokenId)
	if err != nil {
		return nil, err
	}
	computation := models.ComputationData{
		TokenId:         TokenId,
		UserCertificate: Login.Certificate,
		TaskResult:      asyncResult,
	}
	computations = append(computations, computation)

	return &computation, nil
}

func HasTaskFinished(TaskID string) (bool, error) {
	for _, element := range computations {
		if element.TaskResult.TaskID == TaskID {
			task := element.TaskResult
			return task.Ready()
		}
	}

	return false, fmt.Errorf("[HasTaskFinished] Task with id %s not found\n", TaskID)
}

func GetUsersRunningComputations(user models.Login) []models.TaskResult {
	var tasksRet []models.TaskResult
	for _, element := range computations {
		if element.UserCertificate == user.Certificate {

			taskResult, _ := GetTaskResult(element.TaskResult.TaskID, user)

			if taskResult != nil && !taskResult.Finished {
				tasksRet = append(tasksRet, *taskResult)
			}

		}
	}
	return tasksRet
}

func GetTaskResult(TaskID string, User models.Login) (*models.TaskResult, error) {
	for _, element := range computations {
		if element.TaskResult.TaskID == TaskID {
			task := element.TaskResult

			ready, _ := task.Ready()

			result := models.TaskResult{}
			result.TaskID = TaskID

			out, err := EvaluateTransaction(User, vars.ComputationTokenChaincodeName, vars.ComputationTokenSmartContractName, "ReadToken", element.TokenId)

			if err != nil {
				return nil, fmt.Errorf("[GetTaskResult] Error during fetching task from Blockchain", TaskID)
			}

			var token tokenapi.Token
			err = json.Unmarshal(out, &token)
			if err != nil {
				return nil, fmt.Errorf("[GetTaskResult] Error during parsing task result from Blockchain")
			}

			result.Result = token

			if !ready {
				result.Finished = false
				return &result, nil
			}

			res, err := task.Get(10)

			if err != nil {
				result.Finished = true
				return &result, nil
			}

			result = res.(models.TaskResult)

			return &result, nil

		}
	}

	return nil, fmt.Errorf("[GetTaskResult] Task with id %s not found", TaskID)
}

func DeinitCelery() {
	cli.StopWorker()
}
