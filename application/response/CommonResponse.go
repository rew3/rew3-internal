package response

import (
	. "github.com/rew3/app-core/application/constants"
)

/*
  A model to represent common response for GQL (GraphQL Client) consumers for various mutation related requests.
  This response model can be used to return response for GQL request that contains only generic information about
  response such as `status`, `message` etc. Use this response model for request that does not require curated
  response.

  @see [[GQLResponse]] for request specific response.

  Created on 2023/05/12

  @author rew3
  @version 1.0.0

  @param id - Entity ID.
  @param action - Type of action (e.g. create, update, delete etc.)
  @param message - Response message.
  @param status - Status of the operation.
*/

type CommonResponse struct {
	ID      string     `json:"id"`
	Action  string     `json:"action"`
	Message string     `json:"message"`
	Status  StatusType `json:"status"`
}
