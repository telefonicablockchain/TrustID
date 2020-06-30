/*

Copyright 2020 Telefónica Digital España. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0

*/
package main

import (
	log "coren-identitycc/src/chaincode/log"
	"encoding/json"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (cc *Chaincode) createServiceIdentity(stub shim.ChaincodeStubInterface, did string, args interface{}) (string, error) {
	var err error
	service := make(map[string]interface{})
	service = args.(map[string]interface{})

	log.Debugf("[%s][createServiceIdentity] Calling to registry", ServiceGATEWAY)

	serviceStore := Service{Name: service["name"].(string), Controller: did, Public: service["isPublic"].(bool)}

	res, err := cc.createServiceRegistry(stub, service["did"].(string), serviceStore)
	if err != nil {
		log.Errorf("[%s][createServiceIdentity] Error creating service in registry: %v", ServiceGATEWAY, err.Error())
		return "", err
	}

	log.Infof("[%s][createServiceIdentity] Everything went ok", ServiceGATEWAY)
	return res, nil

}
func (cc *Chaincode) updateServiceAccess(stub shim.ChaincodeStubInterface, args interface{}) (string, error) {
	log.Infof("[%s][updateServiceAccess] Entry in updateServiceAccess", ServiceGATEWAY)

	service := make(map[string]interface{})
	service = args.(map[string]interface{})

	m := make(map[string]interface{}) // parse access to interact
	m = service["access"].(map[string]interface{})

	result, err := cc.updateRegistryAccess(stub, service["did"].(string), m["did"].(string), int(m["type"].(float64)))
	if err != nil {
		log.Errorf("[%s][updateServiceAccess] Error updating registry access: %v", ServiceGATEWAY, err.Error())
		log.Errorf("[%s][updateServiceAccess] Return error", ServiceGATEWAY)
		return "", err

	}
	log.Infof("[%s][updateServiceAccess] Update registry Ok", ServiceGATEWAY)

	return result, nil

}

func (cc *Chaincode) getServiceIdentity(stub shim.ChaincodeStubInterface, args interface{}) (string, error) {
	var err error
	servReq := make(map[string]interface{})
	servReq = args.(map[string]interface{})

	result, err := cc.getServiceRegistry(stub, servReq["did"].(string))

	if err != nil {
		log.Errorf("[%s][getServiceIdentity] Error getting registry access: %v", ServiceGATEWAY, err.Error())
		log.Errorf("[%s][getServiceIdentity] Return error", ServiceGATEWAY)
		return "", err

	}

	log.Infof("[%s][getServiceIdentity]Service to return Name: %s, Controller: %s, is Public %t", ServiceGATEWAY, result.Name, result.Controller, result.Public)
	serviceBytes, err := json.Marshal(*result)
	return string(serviceBytes), nil

}
