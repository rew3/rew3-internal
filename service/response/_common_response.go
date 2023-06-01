package response

import (
	. "github.com/rew3/rew3-base/app/service/constants"
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

  @field id - Entity ID.
  @field action - Type of action (e.g. create, update, delete etc.)
  @field message - Response message.
  @field status - Status of the operation.
*/

type CommonResponse struct {
	Id      string     `json:"id"`
	Action  string     `json:"action"`
	Message string     `json:"message"`
	Status  StatusType `json:"status"`
}
