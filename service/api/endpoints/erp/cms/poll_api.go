package cms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_POLL                    endpoints.Endpoint = "cms_addPoll"
	UPDATE_POLL                 endpoints.Endpoint = "cms_updatePoll"
	DELETE_POLL                 endpoints.Endpoint = "cms_deletePoll"
	CLONE_POLL                  endpoints.Endpoint = "cms_clonePoll"
	CHANGE_POLL_OWNER           endpoints.Endpoint = "cms_changePollOwner"
	CHANGE_POLL_FAVORITE_STATUS endpoints.Endpoint = "cms_changePollFavoriteStatus"

	ADD_POLL_VOTE    endpoints.Endpoint = "cms_addPollVote"
	UPDATE_POLL_VOTE endpoints.Endpoint = "cms_updatePollVote"
	DELETE_POLL_VOTE endpoints.Endpoint = "cms_deletePollVote"

	// READ Operations.
	LIST_POLLS     endpoints.Endpoint = "cms_listPolls"
	COUNT_POLLS    endpoints.Endpoint = "cms_countPolls"
	GET_POLL_BY_ID endpoints.Endpoint = "cms_getPollById"
)
