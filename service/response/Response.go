package response

import (
	. "github.com/rew3/rew3-base/app/service/constants"
)

/*
  A model to represent request specific response for GQL (GraphQL Client) consumers for various mutation related
  requests. This response model can be used to return response for GQL request that contains request specific entity
  data along with generic information such as `status`, `message` etc. A request specific response is curated only
  for that particular request.

  @see [[GQLCommonResponse]] for generic response model.

  author rew3 on 2023/05/12

  @field id - Entity ID.
  @field action - Type of action (e.g. create, update, delete etc.)
  @field message - Response message.
  @field status - Status of the operation.
  @field data - Request specific data.
  @tparam T - interface{} type that can hold any type of data. Similar to Generic Type in Scala
*/

type Response struct {
	ID      string      `json:"id"`
	Action  string      `json:"action"`
	Message string      `json:"message"`
	Status  StatusType  `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}
