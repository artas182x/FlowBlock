package tokenchaincode

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/artas182x/hyperledger-fabric-master-thesis/chaincode-sources/chaincode-computationtoken/tokenapi"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Request a workflowflow (list of tasks that can be depedent on each other)
func (s *ComputationTokenSmartContract) RequestFlow(ctx contractapi.TransactionContextInterface, tokenJSON string) (*tokenapi.Flow, error) {
	var flow tokenapi.Flow

	err := json.Unmarshal([]byte(tokenJSON), &flow)

	if err != nil {
		return nil, err
	}

	firstNode, err := getNodeWithoutOuput(&flow)
	firstNode.DirectlyExecutable = true

	if err != nil {
		return nil, err
	}

	_, err = s.submitToken(ctx, firstNode, &flow)

	if err != nil {
		return nil, err
	}

	return &flow, nil

}

// Submit token to world state
func (s *ComputationTokenSmartContract) submitToken(ctx contractapi.TransactionContextInterface, node *tokenapi.Node, flow *tokenapi.Flow) (string, error) {

	if node.TokenId != "" {
		return node.TokenId, nil
	}

	var params []string

	fmt.Printf("[ComputationTokenSmartContract:submitToken] submitting token: %s\n", node.MethodName)

	dependentTokenParam := ""

	for _, intf := range node.Interfaces {
		if strings.HasPrefix(intf.Name, "Input") {
			node, err := findParentNode(intf.ID, flow)
			node.DirectlyExecutable = false

			if err != nil {
				return "", err
			}

			dependentId, err := s.submitToken(ctx, node, flow)

			if err != nil {
				return "", err
			}

			dependentTokenParam += dependentId
			dependentTokenParam += "|"

		}
	}

	if dependentTokenParam != "" {
		params = append(params, strings.TrimSuffix(dependentTokenParam, "|"))
	}

	description := ""

	for _, option := range node.Options {
		if option.Name == "Description" {
			description = option.Value
		} else if option.Name == "Add Input" || option.Name == "Remove Input" {
			continue
		} else {
			params = append(params, option.Value)
		}
	}

	var argsString string

	for index, arg := range params {
		argsString += arg
		if index < len(params)-1 {
			argsString += ";"
		}
	}

	directlyExecutableStr := "false"

	if node.DirectlyExecutable {
		directlyExecutableStr = "true"
	}

	token, err := s.RequestToken(ctx, node.ChaincodeName, node.MethodName, argsString, description, directlyExecutableStr)

	if err != nil {
		return "", err
	}

	node.TokenId = token.ID

	return node.TokenId, nil

}

// Finds a node that has output connected with noting. It will mean that it is the last node that produces result
func getNodeWithoutOuput(flow *tokenapi.Flow) (*tokenapi.Node, error) {

	var foundNode *tokenapi.Node

	for nodeId, _ := range flow.Nodes {
		for _, intf := range flow.Nodes[nodeId].Interfaces {

			if strings.HasPrefix(intf.Name, "Output") {
				conn := getOutputConnection(&flow.Connections, intf.ID)
				if conn == nil {

					if foundNode != nil {
						return nil, fmt.Errorf("there can't be 2 nodes without connected output")
					}

					foundNode = &flow.Nodes[nodeId]
				}
			}

		}
	}
	return foundNode, nil
}

// Gets connection where specified output is
func getOutputConnection(connections *[]tokenapi.Connection, outputId string) *tokenapi.Connection {
	for _, connection := range *connections {
		if outputId == connection.From {
			return &connection
		}
	}

	return nil
}

// Find a parent node of specified node
func findParentNode(nodeId string, flow *tokenapi.Flow) (*tokenapi.Node, error) {

	dependentNodeId := ""

	for _, connection := range flow.Connections {
		if nodeId == connection.To {
			dependentNodeId = connection.From
			break
		}
	}

	if dependentNodeId == "" {
		return nil, fmt.Errorf("dependent node of %s not found", nodeId)
	}

	for nodeId, _ := range flow.Nodes {
		for _, intf := range flow.Nodes[nodeId].Interfaces {

			if strings.HasPrefix(intf.Name, "Output") {
				conn := getOutputConnection(&flow.Connections, intf.ID)
				if conn.From == dependentNodeId {
					return &flow.Nodes[nodeId], nil
				}

			}
		}

		if flow.Nodes[nodeId].ID == dependentNodeId {
			return &flow.Nodes[nodeId], nil
		}
	}

	return nil, fmt.Errorf("dependent node of %s not found (but connection was found)", nodeId)
}
