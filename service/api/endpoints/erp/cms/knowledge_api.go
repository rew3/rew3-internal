package cms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_KNOWLEDGE                    endpoints.Endpoint = "cms_addKnowledge"
	UPDATE_KNOWLEDGE                 endpoints.Endpoint = "cms_updateKnowledge"
	DELETE_KNOWLEDGE                 endpoints.Endpoint = "cms_deleteKnowledge"
	CLONE_KNOWLEDGE                  endpoints.Endpoint = "cms_cloneKnowledge"
	CHANGE_KNOWLEDGE_OWNER           endpoints.Endpoint = "cms_changeKnowledgeOwner"
	CHANGE_KNOWLEDGE_FAVORITE_STATUS endpoints.Endpoint = "cms_changeKnowledgeFavoriteStatus"

	ADD_KNOWLEDGE_CATEGORY    endpoints.Endpoint = "cms_addKnowledgeCategory"
	UPDATE_KNOWLEDGE_CATEGORY endpoints.Endpoint = "cms_updateKnowledgeCategory"
	DELETE_KNOWLEDGE_CATEGORY endpoints.Endpoint = "cms_deleteKnowledgeCategory"

	// READ Operations.
	LIST_KNOWLEDGES     endpoints.Endpoint = "cms_listKnowledges"
	COUNT_KNOWLEDGES    endpoints.Endpoint = "cms_countKnowledges"
	GET_KNOWLEDGE_BY_ID endpoints.Endpoint = "cms_getKnowledgeById"

	LIST_KNOWLEDGE_CATEGORIES    endpoints.Endpoint = "cms_listKnowledgeCategories"
	COUNT_KNOWLEDGE_CATEGORIES   endpoints.Endpoint = "cms_countKnowledgeCategories"
	GET_KNOWLEDGE_CATEGORY_BY_ID endpoints.Endpoint = "cms_getKnowledgeCategoryById"
)
