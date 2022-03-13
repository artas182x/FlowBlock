package tokenapi

// Helper func that converts params string array to [][]byte array compatible with Hyperledger
func ParamsToHyperledgerArgs(params []string) [][]byte {
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}
	return queryArgs
}
