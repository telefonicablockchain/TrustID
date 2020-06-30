/*

Copyright 2020 Telefónica Digital España. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0

*/
package main

import (
	log "coren-identitycc/src/chaincode/log"
	"encoding/json"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
)

// Chaincode struct
type Chaincode struct {
}

const logLevel string = "DEBUG"

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	log.Init("DEBUG")
	log.Infof("[IdentityCC][Init] Initializing identity root")
	idReq := IdentityRequest{}
	_, args := stub.GetFunctionAndParameters()

	err := json.Unmarshal([]byte(args[0]), &idReq)

	if err != nil {
		log.Errorf("[IdentityGateway][CreateIdentity] Error parsing: %v", err.Error())
	}
	identityStore := Identity{PublicKey: idReq.PublicKey, Controller: idReq.Controller, Access: 0}
	_, err = cc.createIDRegistry(stub, idReq.Did, identityStore)
	log.Infof("[IdentityCC][Init] Chaincode initialized")

	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	var err error
	var result string

	if fcn == "proxy" {
		result, err = cc.checkArgs(stub, params)
	}

	if err != nil {
		log.Errorf("[IdentityCC][Init] Errror %v", err)
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		panic(err)
	}
}
