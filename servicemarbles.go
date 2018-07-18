/*
Package main is the entry point of the hyperledger fabric chaincode and implements the shim.ChaincodeStubInterface
*/
/*
Created by Valerio Mattioli @ HES-SO (valeriomattioli580@gmail.com
*/
package main

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	a "github.com/pavva91/servicemarbles/assets"
	gen "github.com/pavva91/servicemarbles/generalcc"
	invoke "github.com/pavva91/servicemarbles/invokeapi"
)

// ==== CHAINCODE RUN (CHAINCODE CONTAINER) ==================

// CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=scc:0 ./servicemarbles

// ==== IMPORT PACKAGE (CLI) ==================
// go get github.com/hyperledger/fabric/protos/ledger/queryresult

// ==== CHAINCODE INSTALLATION (CLI) ==================

// peer chaincode install -p chaincodedev/chaincode/servicemarbles -n scc -v 0

// ==== CHAINCODE INSTANTIATION (CLI) ==================

// peer chaincode instantiate -n scc -v 0 -c '{"Args":[]}' -C ch2

// ==== CHAINCODE EXECUTION SAMPLES (CLI) ==================

// ==== Invoke servicemarbles ====
// peer chaincode invoke -C ch2 -n scc -c '{"function": "HelloWorld", "Args":[]}'
// ==== INITIALIZATION FUNCTIONS ==================
// peer chaincode invoke -C ch2 -n scc -c '{"function": "InitLedger", "Args":[]}'

// ==== GENERAL FUNCTIONS ==================
// peer chaincode invoke -C ch2 -n scc -c '{"function": "Read", "Args":["idagent1"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "GetValue", "Args":["idagent2"]}' -v 0
// peer chaincode invoke -C ch2 -n scc -c '{"function": "ReadEverything", "Args":[]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "AllLedger", "Args":[]}' //TODO: ALL STATE WORLD

// ==== CREATE ASSET FUNCTIONS ==================
// peer chaincode invoke -C ch2 -n scc -c '{"function": "InitService", "Args":["idservice5","service1","description1asdfasdf"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "InitAgent", "Args":["idagent10","agent10","address10"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "InitServiceAgentRelation", "Args":["idservice1","idagent1","2","6","8"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "InitServiceAndServiceAgentRelation", "Args":["idservice10", "service10","description10","idagent2","2","6","8"]}'

// ==== GET ASSET ==================
// peer chaincode invoke -C ch2 -n scc -c '{"function": "GetService", "Args":["idservice1"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "GetAgent", "Args":["idagent10"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "GetServiceRelationAgent", "Args":["idservice1idagent1"]}'

// ==== GET HISTORY ==================
// peer chaincode invoke -C ch2 -n scc -c '{"function": "GetServiceHistory2", "Args":["idagent2"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "GetHistory", "Args":["idservice10"]}'

// ==== RANGE QUERY (USING COMPOSITE INDEX) ==================
// peer chaincode invoke -C ch2 -n scc -c '{"function": "byService", "Args":["idservice1"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "byAgent", "Args":["idAgent10"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "GetAgentsByService", "Args":["idservice1"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "getServicesByAgent", "Args":["idagent1"]}'

// ==== DELETE ASSET ==================
// peer chaincode invoke -C ch2 -n scc -c '{"function": "DeleteService", "Args":["idservice1"]}'
// peer chaincode invoke -C ch2 -n scc -c '{"function": "DeleteAgent", "Args":["idagent1"]}'

// ==== CALLS IN THE REAL PROJECT ====
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "helloWorld", "Args":[]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "InitLedger", "Args":[]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "allLedger", "Args":[]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "GetServiceHistory2", "Args":["service5"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "InitAgent", "Args":["idagent10","agent10","address10"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "InitService", "Args":["idservice10","service10","description10"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "GetService", "Args":["idservice1"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "GetAgent", "Args":["idagent1"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "GetServiceRelationAgent", "Args":["idservice1idagent1"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "InitServiceAgentRelation", "Args":["idservice1","idagent2","3","5","7"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "GetAgentsByService", "Args":["CIAO"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "byAgent", "Args":["a1"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "GetService", "Args":["idservice5"]}'
// peer chaincode invoke -C servicech -n servicemarbles -c '{"function": "DeleteService", "Args":["idservice5"]}'

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode - %s", err)
	}
}

// Init initialize the chaincode
// The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
// Best practice is to have any Ledger initialization in separate function -- see InitLedger()
//======================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// TODO: General Refactor the "Not found asset" from throwing error to get back null payload
	function, args := stub.GetFunctionAndParameters()
	fmt.Println(" ")
	fmt.Println("starting invoke, for - " + function)

	// Route to the appropriate handler function to interact with the ledger appropriately
	switch function {
	case "InitLedger":
		response := a.InitLedger(stub)
		return response
	case "InitService":
		return invoke.InitService(stub, args)
	case "InitAgent":
		return invoke.InitAgent(stub, args)
	case "InitServiceAgentRelation":
		// Already with reference integrity controls (service already exist, agent already exist, relation don't already exist)
		return invoke.InitServiceAgentRelation(stub, args)
	case "InitServiceAndServiceAgentRelation":
		// If service doesn't exist it will create
		return invoke.InitServiceAndServiceAgentRelation(stub, args)
	case "GetHistory":
		// TODO: Refacoring GetServiceHistory2 da generalcc
		return gen.GetGeneralHistory(stub, args)
	case "GetServiceHistory":
		return a.GetServiceHistory(stub, args)
	case "GetService":
		return invoke.QueryService(stub, args)
	case "GetAgent":
		return invoke.QueryAgent(stub, args)
	case "GetServiceRelationAgent":
		return invoke.QueryServiceRelationAgent(stub, args)
	case "byService":
		return invoke.QueryByServiceAgentRelation(stub, args)
	case "byAgent":
		return invoke.QueryByAgentServiceRelation(stub, args)
	case "GetAgentsByService":
		// also with only one record result return always a JSONArray
		return invoke.GetServiceRelationAgentByServiceWithCostAndTime(stub, args)
	case "GetServicesByAgent":
		// also with only one record result return always a JSONArray
		return invoke.GetServiceRelationAgentByAgentWithCostAndTime(stub, args)
	case "DeleteService":
		return a.DeleteService(stub, args)
	case "DeleteAgent":
		return a.DeleteAgent(stub, args)
	case "Write":
		return gen.Write(stub, args)
	case "Read":
		return gen.Read(stub, args)
	case "ReadEverything":
		return a.ReadEverything(stub)
	case "AllLedger":
		return gen.ReadAllLedger(stub)
	case "GetValue":
		return gen.GetValue(stub, args)
	case "HelloWorld":
		fmt.Println("Ciao")
		// in := []byte(`{"Hello":"HelloWorld"}`)
		// var raw map[string]interface{}
		// json.Unmarshal(in, &raw)
		// out, _ := json.Marshal(raw)
		var buffer bytes.Buffer

		buffer.WriteString("[{\"Hello\":\"HelloWorld\"}]")

		return shim.Success(buffer.Bytes())
	default:
		return shim.Error("Invalid Smart Contract function Name.")
	}

	// error out
	fmt.Println("Received unknown invoke function Name - " + function)
	return shim.Error("Received unknown invoke function Name - '" + function + "'")
}

// ============================================================================================================================
// Query - legacy function
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("Unknown supported call - Query()")
}
