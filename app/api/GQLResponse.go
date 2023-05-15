package app

import (
	. "github.com/rew3/app-core/app/_constants"
)

/*
  A model to represent request specific response for GQL (GraphQL Client) consumers for various mutation related
  requests. This response model can be used to return response for GQL request that contains request specific entity
  data along with generic information such as `status`, `message` etc. A request specific response is curated only
  for that particular request.

  @see [[GQLCommonResponse]] for generic response model.

  Created on 2023/05/12

  @author rew3
  @version 1.0.0

  @param id - Entity ID.
  @param action - Type of action (e.g. create, update, delete etc.)
  @param message - Response message.
  @param status - Status of the operation.
  @param data - Request specific data.
  @tparam T - interface{} type that can hold any type of data. Similar to Generic Type in Scala
*/

type GQLResponse struct {
	ID      string      `json:"id"`
	Action  string      `json:"action"`
	Message string      `json:"message"`
	Status  StatusType  `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}
