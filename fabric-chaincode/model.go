/*

Copyright 2020 Telefónica Digital España. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0

*/
package main

// Request to serialize args
type Request struct {
	Did       string `json:"did,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	Payload   string `json:"payload,omitempty"` // me pasa una firma // el controller lo meto yo
}

// Identity stored in bc
type Identity struct {
	PublicKey  string `json:"publicKey"`
	Controller string `json:"controller"` // issuer's DID
	Access     int    `json:"access,omitempty"`
}

// IdentityRequest to serialize args
type IdentityRequest struct {
	Did        string `json:"did"`
	Controller string `json:"controller,omitempty"`
	PublicKey  string `json:"publicKey,omitempty"`
	Payload    string `json:"payload,omitempty"` // me pasa una firma // el controller lo meto yo
	Access     int    `json:"access,omitempty"`
}

// Service stored in bc
type Service struct {
	Name       string         `json:"name"`
	Controller string         `json:"controller,omitempty"` // issuer's DID
	Access     map[string]int `json:"access,omitempty"`     // mapping did - access type
	Public     bool           `json:"isPublic"`
	Channel    string         `json:"channel"`
}

// ServiceRequest stored in bc
type ServiceRequest struct {
	Name   string `json:"name"`
	Did    string `json:"did"`
	Public bool   `json:"isPublic"`
}

// IdentityUnverifiedRequest to serialize args
type IdentityUnverifiedRequest struct {
	PublicKey string `json:"publicKey"`
	Payload   string `json:"payload,omitempty"` // me pasa una firma // el controller lo meto yo
}

// CcRequest payload from jws
type CcRequest struct {
	Name    string   `json:"name,omitempty"`
	Args    []string `json:"args"`
	Channel string   `json:"channel"`
	Did     string   `json:"did"`
}

// Error responses
// ERROR_XXX occurs when XXX
const (
	ERRORWrongNumberArgs = `Wrong number of arguments. Expecting a JSON with token information.`
	ERRORParsingData     = `Error parsing data `
	ERRORPutState        = `Failed to store data in the ledger.	`
	ERRORGetState        = `Failed to get data from the ledger. `
	ERRORDelState        = `Failed to delete data from the ledger. `
	ERRORChaincodeCall   = `Error calling chaincode`
	IDGATEWAY            = `IDGateway`
	IDREGISTRY           = `IDRegistry`
	ServiceGATEWAY       = `IDGateway`
	ServiceREGISTRY      = `IDRegistry`
)
