package cms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_LIBRARY_DOCUMENT                    api.APIOperation = "cms_addLibraryDocument"
	UPDATE_LIBRARY_DOCUMENT                 api.APIOperation = "cms_updateLibraryDocument"
	DELETE_LIBRARY_DOCUMENT                 api.APIOperation = "cms_deleteLibraryDocument"
	CLONE_LIBRARY_DOCUMENT                  api.APIOperation = "cms_cloneLibraryDocument"
	CHANGE_LIBRARY_DOCUMENT_OWNER           api.APIOperation = "cms_changeLibraryDocumentOwner"
	CHANGE_LIBRARY_DOCUMENT_FAVORITE_STATUS api.APIOperation = "cms_changeLibraryDocumentFavoriteStatus"

	ADD_LIBRARY_FOLDER    api.APIOperation = "cms_addLibraryFolder"
	UPDATE_LIBRARY_FOLDER api.APIOperation = "cms_updateLibraryFolder"
	DELETE_LIBRARY_FOLDER api.APIOperation = "cms_deleteLibraryFolder"

	ADD_LIBRARY_CATEGORY    api.APIOperation = "cms_addLibraryCategory"
	UPDATE_LIBRARY_CATEGORY api.APIOperation = "cms_updateLibraryCategory"
	DELETE_LIBRARY_CATEGORY api.APIOperation = "cms_deleteLibraryCategory"

	// READ Operations.
	LIST_LIBRARY_DOCUMENTS     api.APIOperation = "cms_listLibraryDocuments"
	COUNT_LIBRARY_DOCUMENTS    api.APIOperation = "cms_countLibraryDocuments"
	GET_LIBRARY_DOCUMENT_BY_ID api.APIOperation = "cms_getLibraryDocumentById"

	LIST_LIBRARY_FOLDERS     api.APIOperation = "cms_listLibraryFolders"
	COUNT_LIBRARY_FOLDERS    api.APIOperation = "cms_countLibraryFolders"
	GET_LIBRARY_FOLDER_BY_ID api.APIOperation = "cms_getLibraryFolderById"

	LIST_LIBRARY_CATEGORIES    api.APIOperation = "cms_listLibraryCategories"
	COUNT_LIBRARY_CATEGORIES   api.APIOperation = "cms_countLibraryCategories"
	GET_LIBRARY_CATEGORY_BY_ID api.APIOperation = "cms_getLibraryCategoryById"
)
