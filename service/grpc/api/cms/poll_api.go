package cms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_POLL                    api.APIOperation = "cms_addPoll"
	UPDATE_POLL                 api.APIOperation = "cms_updatePoll"
	DELETE_POLL                 api.APIOperation = "cms_deletePoll"
	CLONE_POLL                  api.APIOperation = "cms_clonePoll"
	CHANGE_POLL_OWNER           api.APIOperation = "cms_changePollOwner"
	CHANGE_POLL_FAVORITE_STATUS api.APIOperation = "cms_changePollFavoriteStatus"

	ADD_POLL_VOTE    api.APIOperation = "cms_addPollVote"
	UPDATE_POLL_VOTE api.APIOperation = "cms_updatePollVote"
	DELETE_POLL_VOTE api.APIOperation = "cms_deletePollVote"

	// READ Operations.
	LIST_POLLS     api.APIOperation = "cms_listPolls"
	COUNT_POLLS    api.APIOperation = "cms_countPolls"
	GET_POLL_BY_ID api.APIOperation = "cms_getPollById"
)
