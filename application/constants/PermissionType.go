package constants

type PermissionType string

const (
	VIEW   PermissionType = "privateOpen"
	CREATE PermissionType = "publicClose"
	EDIT   PermissionType = "privateClose"
	DELETE PermissionType = "publicOpen"

	IMPORTOWNRECORDS          PermissionType = "importOwnRecords"
	IMPORTORGANIZATIONRECORDS PermissionType = "importOrganizationRecords"
	EXPORT                    PermissionType = "export"

	//superAdminPrivileges TODO use this from AdminLevelPermissions later
	ACTIVATEUSER     PermissionType = "activateUser"
	DEACTIVATEUSER   PermissionType = "deactivateUser"
	SETISSYSTEMADMIN PermissionType = "setIsSystemAdmin"
	SETISSUPERADMIN  PermissionType = "setIsSuperAdmin"

	MANAGEDOSSIERS PermissionType = "manageDossiers"
)
