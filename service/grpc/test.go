package grpc

import (
	"encoding/json"
	"errors"
	"fmt"
)

type APIName string

const (
	CRM_ADD_CONTACT    APIName = "addContact"    // grpc method name.
	CRM_DELETE_CONTACT APIName = "deleteContact" // grpc method name.
	CRM_UPDATE_CONTACT APIName = "updateContact" // grpc method name.
)

type RequestPayloads struct {
	API     APIName
	request string // json ?
	Context RequestContext
}

type ResponsePayloads struct {
	API      APIName
	response ExecutionResult // json? // raw.JsonMessage , in grpc Any type.
}

type AddContactRequest struct {
	// some fields // or command ???
}
type AddContactResponse struct {
	// some fields // or command ???
}

// apiName -> data type mapping. or def method.
// apiName -> { input : typeName, output: TypeName }
// switch // handler.
var apiRequestMap = map[APIName]interface{}{
	CRM_ADD_CONTACT: AddContactRequest{},
}

var apiResponseMap = map[APIName]interface{}{
	CRM_ADD_CONTACT: AddContactResponse{},
}

// I need mapping for data in request and response payload?
//

// generate json schema for each model or api request and response?

// define in platform, provide a partial function in param to handle
// project/module specific apis e.g. contacct, etc. this function take channel to send result,
// and interface{} type request payload. the channel must accept execution result or resposne payload.
func HandleRequest(reqPayload RequestPayload) error {
	// Check if the APIName is valid and exists in the map
	requestType, found := apiRequestMap[reqPayload.API]
	if !found {
		return fmt.Errorf("unsupported API: %s", reqPayload.API)
	}

	// Unmarshal the raw request data into the appropriate request type
	err := json.Unmarshal([]byte(reqPayload.Request), &requestType)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request data: %v", err)
	}

	// Now, you can validate the request data as needed for each specific API.
	switch reqPayload.API {
	case APIAddContact:
		// Perform validation for AddContactRequest
		if contactReq, ok := requestType.(AddContactRequest); ok {
			// Perform validation logic for AddContactRequest
			if contactReq.Name == "" {
				return errors.New("name is required for AddContact")
			}
			// Additional validation for other fields, if needed.
		}
		// Handle other APIs similarly, if needed.
	}

	// Perform any common validation or checks here.
	// For example, you can validate the RequestContext or other common data.

	return nil
}

// as chunky, do we need validation for now, or assume we get directly valid data?

// should we dfine a common struct/model, in platform for type safe e.g AddContactPayload? but this is not good?
