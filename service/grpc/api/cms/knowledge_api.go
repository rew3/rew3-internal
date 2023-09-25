package cms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_KNOWLEDGE                    api.APIOperation = "cms_addKnowledge"
	UPDATE_KNOWLEDGE                 api.APIOperation = "cms_updateKnowledge"
	DELETE_KNOWLEDGE                 api.APIOperation = "cms_deleteKnowledge"
	CLONE_KNOWLEDGE                  api.APIOperation = "cms_cloneKnowledge"
	CHANGE_KNOWLEDGE_OWNER           api.APIOperation = "cms_changeKnowledgeOwner"
	CHANGE_KNOWLEDGE_FAVORITE_STATUS api.APIOperation = "cms_changeKnowledgeFavoriteStatus"

	ADD_KNOWLEDGE_CATEGORY    api.APIOperation = "cms_addKnowledgeCategory"
	UPDATE_KNOWLEDGE_CATEGORY api.APIOperation = "cms_updateKnowledgeCategory"
	DELETE_KNOWLEDGE_CATEGORY api.APIOperation = "cms_deleteKnowledgeCategory"

	// READ Operations.
	LIST_KNOWLEDGES     api.APIOperation = "cms_listKnowledges"
	COUNT_KNOWLEDGES    api.APIOperation = "cms_countKnowledges"
	GET_KNOWLEDGE_BY_ID api.APIOperation = "cms_getKnowledgeById"

	LIST_KNOWLEDGE_CATEGORIES    api.APIOperation = "cms_listKnowledgeCategories"
	COUNT_KNOWLEDGE_CATEGORIES   api.APIOperation = "cms_countKnowledgeCategories"
	GET_KNOWLEDGE_CATEGORY_BY_ID api.APIOperation = "cms_getKnowledgeCategoryById"
)
