package chat

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// READ APIs.
	ME             api.APIOperation = "users_me"
	GET_TEAM_USERS api.APIOperation = "users_getTeamUsers"
)
