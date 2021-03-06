// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package actors

// CallRequest is the request object to call an actor
type CallRequest struct {
	ActorType string            `json:"actorType"`
	ActorID   string            `json:"actorId"`
	Method    string            `json:"method"`
	Data      []byte            `json:"data"`
	Metadata  map[string]string `json:"metadata"`
}
