package chat

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// READ APIs.
	LIST_TEAM_CHANNELS api.APIOperation = "message_listTeamChannels"
	GET_CHANNEL_BY_ID  api.APIOperation = "message_GetChannelById"
)
