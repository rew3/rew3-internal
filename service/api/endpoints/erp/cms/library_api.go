package cms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_LIBRARY_DOCUMENT                    endpoints.Endpoint = "cms_addLibraryDocument"
	UPDATE_LIBRARY_DOCUMENT                 endpoints.Endpoint = "cms_updateLibraryDocument"
	DELETE_LIBRARY_DOCUMENT                 endpoints.Endpoint = "cms_deleteLibraryDocument"
	CLONE_LIBRARY_DOCUMENT                  endpoints.Endpoint = "cms_cloneLibraryDocument"
	CHANGE_LIBRARY_DOCUMENT_OWNER           endpoints.Endpoint = "cms_changeLibraryDocumentOwner"
	CHANGE_LIBRARY_DOCUMENT_FAVORITE_STATUS endpoints.Endpoint = "cms_changeLibraryDocumentFavoriteStatus"

	ADD_LIBRARY_FOLDER    endpoints.Endpoint = "cms_addLibraryFolder"
	UPDATE_LIBRARY_FOLDER endpoints.Endpoint = "cms_updateLibraryFolder"
	DELETE_LIBRARY_FOLDER endpoints.Endpoint = "cms_deleteLibraryFolder"

	ADD_LIBRARY_CATEGORY    endpoints.Endpoint = "cms_addLibraryCategory"
	UPDATE_LIBRARY_CATEGORY endpoints.Endpoint = "cms_updateLibraryCategory"
	DELETE_LIBRARY_CATEGORY endpoints.Endpoint = "cms_deleteLibraryCategory"

	// READ Operations.
	LIST_LIBRARY_DOCUMENTS     endpoints.Endpoint = "cms_listLibraryDocuments"
	COUNT_LIBRARY_DOCUMENTS    endpoints.Endpoint = "cms_countLibraryDocuments"
	GET_LIBRARY_DOCUMENT_BY_ID endpoints.Endpoint = "cms_getLibraryDocumentById"

	LIST_LIBRARY_FOLDERS     endpoints.Endpoint = "cms_listLibraryFolders"
	COUNT_LIBRARY_FOLDERS    endpoints.Endpoint = "cms_countLibraryFolders"
	GET_LIBRARY_FOLDER_BY_ID endpoints.Endpoint = "cms_getLibraryFolderById"

	LIST_LIBRARY_CATEGORIES    endpoints.Endpoint = "cms_listLibraryCategories"
	COUNT_LIBRARY_CATEGORIES   endpoints.Endpoint = "cms_countLibraryCategories"
	GET_LIBRARY_CATEGORY_BY_ID endpoints.Endpoint = "cms_getLibraryCategoryById"
)
