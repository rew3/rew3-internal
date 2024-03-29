package constants

type Entity string

const (

	// CRM Entities
	CRM_CONTACT              Entity = "crm_contact"
	CRM_ACCOUNT              Entity = "crm_account"
	CRM_LEAD                 Entity = "crm_lead"
	CRM_OPPORTUNITY          Entity = "crm_opportunity"
	CRM_OPPORTUNITY_PIPELINE Entity = "crm_opportunity_pipeline"
	CRM_CASE                 Entity = "crm_case"
	CRM_CAMPAIGN             Entity = "crm_campaign"

	// RMS Entities
	RMS_LISTING     Entity = "rms_listing"
	RMS_PROPERTY    Entity = "rms_property"
	RMS_SHOWING     Entity = "rms_showing"
	RMS_REQUEST     Entity = "rms_request"
	RMS_TRANSACTION Entity = "rms_transaction"
	RMS_OFFER       Entity = "rms_offer"

	// DMS Entities
	DMS_FILE                               Entity = "dms_file"
	DMS_FOLDER                             Entity = "dms_folder"
	DMS_DOCUMENT                           Entity = "dms_document"
	DMS_DOCUMENT_TYPE                      Entity = "dms_document_type"
	DMS_VERSIONED_DOCUMENT                 Entity = "dms_versioned_document"
	DMS_DOCUMENT_ACCESS                    Entity = "dms_document_access"
	DMS_DOSSIER                            Entity = "dms_dossier"
	DMS_DOSSIER_TYPE                       Entity = "dms_dossier_type"
	DMS_DOSSIER_FOLDER                     Entity = "dms_dossier_folder"
	DMS_DOSSIER_DOCUMENT_GROUP             Entity = "dms_dossier_document_group"
	DMS_DOSSIER_USER_DOCUMENT_GROUP        Entity = "dms_dossier_user_document_group"
	DMS_DOSSIER_DOCUMENT_CATEGORY_DOCUMENT Entity = "dms_dossier_document_category_document"
	DMS_DOSSIER_TEMPLATE                   Entity = "dms_dossier_template"
	DMS_DOSSIER_TEMPLATE_DOCUMENT          Entity = "dms_dossier_template_document"
	DMS_DOSSIER_TEMPLATE_GROUP             Entity = "dms_dossier_template_group"
	DMS_DOSSIER_COLLABORATOR               Entity = "dms_dossier_collaborator"
	DMS_DOSSIER_CATEGORY                   Entity = "dms_dossier_category"

	// Project.
	PROJECT_PROJECT      Entity = "project_project"
	PROJECT_PROJECTGROUP Entity = "project_projectgroup"
	PROJECT_WORKSPACE    Entity = "project_workspace"

	// CMS
	CMS_AVPOST             Entity = "cms_avpost"
	CMS_BLOG               Entity = "cms_blog"
	CMS_BLOG_CATEGORY      Entity = "cms_blog_category"
	CMS_FORM               Entity = "cms_form"
	CMS_FORM_CATEGORY      Entity = "cms_form_category"
	CMS_FORUM              Entity = "cms_forum"
	CMS_FORUM_CATEGORY     Entity = "cms_forum_category"
	CMS_KNOWLEDGE          Entity = "cms_knowledge"
	CMS_KNOWLEDGE_CATEGORY Entity = "cms_knowledge_category"
	CMS_LIBRARY_DOCUMENT   Entity = "cms_library_document"
	CMS_LIBRARY_FOLDER     Entity = "cms_library_folder"
	CMS_LIBRARY_CATEGORY   Entity = "cms_library_category"
	CMS_NEWS               Entity = "cms_news"
	CMS_NEWS_CATEGORY      Entity = "cms_news_category"
	CMS_POST               Entity = "cms_post"
	CMS_POST_CATEGORY      Entity = "cms_post_category"
	CMS_POLL               Entity = "cms_poll"

	// Collaboration

	COLLABORATION_ASSET              Entity = "collaboration_asset"
	COLLABORATION_CALENDAR           Entity = "collaboration_calendar"
	COLLABORATION_CALENDARCATEGORY   Entity = "collaboration_calendar_category"
	COLLABORATION_CALL               Entity = "collaboration_call"
	COLLABORATION_CHECKLIST          Entity = "collaboration_checklist"
	COLLABORATION_COMMENT            Entity = "collaboration_comment"
	COLLABORATION_EVENT              Entity = "collaboration_event"
	COLLABORATION_EVENTCATEGORY      Entity = "collaboration_event_category"
	COLLABORATION_EVENTINVITATION    Entity = "collaboration_event_invitation"
	COLLABORATION_EVENTREQUEST       Entity = "collaboration_event_request"
	COLLABORATION_APPOINTMENT        Entity = "collaboration_appointment"
	COLLABORATION_GROUP              Entity = "collaboration_group"
	COLLABORATION_GROUPINVITATION    Entity = "collaboration_group_invitation"
	COLLABORATION_MULTIMEDIA         Entity = "collaboration_multimedia"
	COLLABORATION_MULTIMEDIAALBUM    Entity = "collaboration_multimedia_album"
	COLLABORATION_NOTE               Entity = "collaboration_note"
	COLLABORATION_ORGANIZATION       Entity = "collaboration_organization"
	COLLABORATION_ORGANIZATIONINVITE Entity = "collaboration_organization_invite"
	COLLABORATION_SOCIALINTEREST     Entity = "collaboration_social_interest"
	COLLABORATION_STATUSUPDATE       Entity = "collaboration_statu_supdate"
	COLLABORATION_STREAMGROUP        Entity = "collaboration_stream_group"
	COLLABORATION_TASK               Entity = "collaboration_task"
	COLLABORATION_TASKTEMPLATE       Entity = "collaboration_task_template"
	/*COLLABORATION_CONVERSATION       Entity = "collaboration_conversation"
	COLLABORATION_EVENT              Entity = "collaboration_event"
	COLLABORATION_FACEBOOK           Entity = "collaboration_facebook"
	COLLABORATION_GROUP              Entity = "collaboration_group"
	COLLABORATION_MEDIA              Entity = "collaboration_media"
	COLLABORATION_ORGANIZATION       Entity = "collaboration_organization"
	COLLABORATION_PROJECT            Entity = "collaboration_project"
	COLLABORATION_TASK               Entity = "collaboration_task"
	COLLABORATION_CALL               Entity = "collaboration_call"
	COLLABORATION_NOTE               Entity = "collaboration_note"
	COLLABORATION_COMMENT            Entity = "collaboration_comment"
	COLLABORATION_TAG                Entity = "collaboration_tag"
	COLLABORATION_TAG_LINK           Entity = "collaboration_tag_link"
	COLLABORATION_SECTION            Entity = "collaboration_section"
	COLLABORATION_STATUS_UPDATE      Entity = "collaboration_status_update"
	COLLABORATION_STREAM_GROUP       Entity = "collaboration_stream_group"
	COLLABORATION_EVENT_TYPE         Entity = "collaboration_event_type"
	COLLABORATION_EVENT_CATEGORY     Entity = "collaboration_event_category"
	COLLABORATION_FOLLOW_LIKE_STATUS Entity = "collaboration_follow_like_status"
	COLLABORATION_CHAT               Entity = "collaboration_chat"
	COLLABORATION_TICKET             Entity = "collaboration_ticket"
	COLLABORATION_EMAIL              Entity = "collaboration_email"
	COLLABORATION_PROJECT_GROUP      Entity = "collaboration_project_group"
	COLLABORATION_WORKSPACE          Entity = "collaboration_workspace"
	COLLABORATION_MULTIMEDIA         Entity = "collaboration_multimedia"
	COLLABORATION_SOCIAL_INTEREST    Entity = "collaboration_social_interest"
	COLLABORATION_MULTIMEDIA_ALBUM   Entity = "collaboration_multimedia_album"
	COLLABORATION_CHECKLIST          Entity = "collaboration_checklist"
	COLLABORATION_ASSET              Entity = "collaboration_asset"
	COLLABORATION_MEETING            Entity = "collaboration_meeting"
	COLLABORATION_APPOINTMENT        Entity = "collaboration_appointment"
	COLLABORATION_CALENDAR           Entity = "collaboration_calendar"
	COLLABORATION_CALENDAR_TYPE      Entity = "collaboration_calendar_type"
	COLLABORATION_CALENDAR_CATEGORY  Entity = "collaboration_calendar_category"*/

	// Financial entities.
	//------------------------------------------------------------
	FINANCIAL_ACCOUNTING_ACCOUNT              Entity = "financial_accounting_account"
	FINANCIAL_ACCOUNTING_ACCOUNTPERIOD        Entity = "financial_accounting_account_period"
	FINANCIAL_ACCOUNTING_JOURNAL              Entity = "financial_accounting_journal"
	FINANCIAL_BROKERAGE_ACP                   Entity = "financial_brokerage_acp"
	FINANCIAL_BROKERAGE_SINGLERATEACP         Entity = "financial_brokerage_single_rate_acp"
	FINANCIAL_BROKERAGE_TIREDACP              Entity = "financial_brokerage_tired_acp"
	FINANCIAL_BROKERAGE_TIREDSTAGE            Entity = "financial_brokerage_tired_stage"
	FINANCIAL_BROKERAGE_AGENT                 Entity = "financial_brokerage_agent"
	FINANCIAL_BROKERAGE_AGENTCLOSINGFEE       Entity = "financial_brokerage_agent_closing_fee"
	FINANCIAL_BROKERAGE_AGENTPRECOMMISSIONS   Entity = "financial_brokerage_agent_precommissions"
	FINANCIAL_BROKERAGE_AGENTPOSTCOMMISSIONS  Entity = "financial_brokerage_agent_postcommissions"
	FINANCIAL_BROKERAGE_AGENTPLAN             Entity = "financial_brokerage_agentplan"
	FINANCIAL_BROKERAGE_ASSOCIATE             Entity = "financial_brokerage_associate"
	FINANCIAL_BROKERAGE_COMMISSIONPLAN        Entity = "financial_brokerage_commissionplan"
	FINANCIAL_BROKERAGE_DEDUCTION             Entity = "financial_brokerage_deduction"
	FINANCIAL_BROKERAGE_DISBURSEMENT          Entity = "financial_brokerage_disbursement"
	FINANCIAL_BROKERAGE_GROSSCOMMISSIONPLAN   Entity = "financial_brokerage_gross_commissionplan"
	FINANCIAL_BROKERAGE_OFFICE                Entity = "financial_brokerage_office"
	FINANCIAL_BROKERAGE_TRUSTACCOUNT          Entity = "financial_brokerage_trustaccount"
	FINANCIAL_CATALOG_PRODUCT                 Entity = "financial_catalog_product"
	FINANCIAL_PAYMENT_BILLPAYMENT             Entity = "financial_payment_bill_payment"
	FINANCIAL_PAYMENT_INVOICEPAYMENT          Entity = "financial_payment_invoice_payment"
	FINANCIAL_PAYMENT_RECURRINGINVOICEPAYMENT Entity = "financial_payment_recurring_invoice_payment"
	FINANCIAL_PAYMENT_PAYMENTTERM             Entity = "financial_payment_payment_term"
	FINANCIAL_PURCHASE_BILL                   Entity = "financial_purchase_bill"
	FINANCIAL_PURCHASE_DEBITNOTE              Entity = "financial_purchase_debitnote"
	FINANCIAL_PURCHASE_EXPENSE                Entity = "financial_purchase_expense"
	FINANCIAL_PURCHASE_VENDOR                 Entity = "financial_purchase_vendor"
	FINANCIAL_SALE_CREDITNOTE                 Entity = "financial_sale_creditnote"
	FINANCIAL_SALE_CUSTOMER                   Entity = "financial_sale_customer"
	FINANCIAL_SALE_ESTIMATE                   Entity = "financial_sale_estimate"
	FINANCIAL_SALE_INVOICE                    Entity = "financial_sale_invoice"
	FINANCIAL_SALE_INVOICEREFERENCE           Entity = "financial_sale_invoice_reference"
	FINANCIAL_SALE_RECEIPT                    Entity = "financial_sale_receipt"
	FINANCIAL_SALE_RECURRINGINVOICE           Entity = "financial_sale_recurring_invoice"
	FINANCIAL_SALE_RECURRINGSCHEDULE          Entity = "financial_sale_recurring_schedule"
	FINANCIAL_SALE_RECURRINGTEMPLATE          Entity = "financial_sale_recurring_template"
	FINANCIAL_SALE_SALETAX                    Entity = "financial_sale_saletax"

	/*FINANCIAL_ACCOUNTING_CLASS           Entity = "financial_accounting_class"
	FINANCIAL_ACCOUNTING_CODE            Entity = "financial_accounting_code"
	FINANCIAL_ACCOUNTING_PERIOD          Entity = "financial_accounting_period"
	FINANCIAL_ACCOUNTING_PERIOD_REQUEST  Entity = "financial_accounting_period_request"
	FINANCIAL_BANK_DEPOSIT_SLIP          Entity = "financial_bank_deposit_slip"
	FINANCIAL_BANK_RECONCILIATION        Entity = "financial_bank_reconciliation"
	FINANCIAL_BANK_TRANSACTION           Entity = "financial_bank_transaction"
	FINANCIAL_BILLING_ACCOUNT            Entity = "financial_billing_account"
	FINANCIAL_DEPOSIT_ITEM               Entity = "financial_deposit_item"
	FINANCIAL_INVOICE                    Entity = "financial_invoice"
	FINANCIAL_INVOICE_ITEM               Entity = "financial_invoice_item"
	FINANCIAL_INVOICE_REQUEST            Entity = "financial_invoice_request"
	FINANCIAL_NORMAL_USER                Entity = "financial_normal_user"
	FINANCIAL_PAYMENT_RECEIPT            Entity = "financial_payment_receipt"
	FINANCIAL_PAYMENT_OPTION             Entity = "financial_payment_option"
	FINANCIAL_PAYMENT_RECEIPT_ITEM       Entity = "financial_payment_receipt_item"
	FINANCIAL_PRODUCT                    Entity = "financial_product"
	FINANCIAL_PRODUCT_CATEGORY           Entity = "financial_product_category"
	FINANCIAL_PRODUCT_CATEGORY_LINK      Entity = "financial_product_category_link"
	FINANCIAL_PRODUCT_FEATURE            Entity = "financial_product_feature"
	FINANCIAL_PRODUCT_FEATURE_LINK       Entity = "financial_product_feature_link"
	FINANCIAL_PRODUCT_RATE_PLAN          Entity = "financial_product_rate_plan"
	FINANCIAL_PRODUCT_RATE_PLAN_CHARGE   Entity = "financial_product_rate_plan_charge"
	FINANCIAL_PRODUCT_RATE_LINK          Entity = "financial_product_rate_link"
	FINANCIAL_SUB_ACCOUNTING_HEAD        Entity = "financial_sub_accounting_head"
	FINANCIAL_TERMS                      Entity = "financial_terms"
	FINANCIAL_TRANSACTION                Entity = "financial_transaction"
	FINANCIAL_PAYMENT_TERM               Entity = "financial_payment_term"
	FINANCIAL_ASSOCIATE_COMISSION_PLAN   Entity = "financial_associate_comission_plan"
	FINANCIAL_BILL                       Entity = "financial_bill"
	FINANCIAL_BILL_PAYMENT               Entity = "financial_bill_payment"
	FINANCIAL_CUSTOMER                   Entity = "financial_customer"
	FINANCIAL_ESTIMATE                   Entity = "financial_estimate"
	FINANCIAL_EXPENSES                   Entity = "financial_expenses"
	FINANCIAL_RECURRING_INVOICE          Entity = "financial_recurring_invoice"
	FINANCIAL_RECURRING_INVOICE_PAYMENT  Entity = "financial_recurring_invoice_payment"
	FINANCIAL_RECURRING_SCHEDULE         Entity = "financial_recurring_schedule"
	FINANCIAL_RECURRING_TEMPLATE         Entity = "financial_recurring_template"
	FINANCIAL_VENDOR                     Entity = "financial_vendor"
	FINANCIAL_SALES_TAX                  Entity = "financial_sales_tax"
	FINANCIAL_INVOICE_CUSTOMERS          Entity = "financial_invoice_customers"
	FINANCIAL_INVOICE_CREDIT_NOTES       Entity = "financial_invoice_credit_notes"
	FINANCIAL_INVOICE_PAYMENTS           Entity = "financial_invoice_payments"
	FINANCIAL_INVOICE_RECURRING_INVOICES Entity = "financial_invoice_recurring_invoices"
	FINANCIAL_DEBIT_NOTE                 Entity = "financial_debit_note"
	FINANCIAL_CREDIT_NOTE                Entity = "financial_credit_note"
	FINANCIAL_OFFICE                     Entity = "financial_office"
	FINANCIAL_AGENT                      Entity = "financial_agent"
	FINANCIAL_TRUST_ACCOUNT              Entity = "financial_trust_account"*/

	// Activity
	ACTIVITY_ACTIVITY_STREAMS Entity = "activity_activity_streams"
	ACTIVITY_USER_ACTIVITY    Entity = "activity_user_activity"
	ACTIVITY_HISTORY_LOG      Entity = "activity_history_log"

	// Customization
	CUSTOMIZATION_LIST_VIEW           Entity = "customization_listview"
	CUSTOMIZATION_VIEW_PREFERENCE     Entity = "customization_view_preference"
	CUSTOMIZATION_TAGS                Entity = "customization_tags"
	CUSTOMIZATION_ENTITY_RELATIONSHIP Entity = "customization_entity_relationship"

	// Notifications.
	NOTIFICATION_NOTIFICATIONS Entity = "notification_notifications"

	// Chat
	CHAT_USERS   Entity = "chat_user"
	CHAT_MESSAGE Entity = "chat_message"
	CHAT_CHANNEL Entity = "chat_channel"
	CHAT_TEAM    Entity = "chat_team"

	ACCOUNT_AVENUEACCOUNT Entity = "account_avenueaccount"
	ACCOUNT_COMPANY       Entity = "account_company"
	ACCOUNT_USER          Entity = "account_user"

	// Account Service
	ACCOUNT_SERVICE_SUBSCRIPTION             Entity = "account_service_subscription"
	ACCOUNT_SERVICE_ACCOUNT                  Entity = "account_service_account"
	ACCOUNT_SERVICE_COMPANY                  Entity = "account_service_company"
	ACCOUNT_SERVICE_AVENUE_ACCOUNT           Entity = "account_service_avenue_account"
	ACCOUNT_SERVICE_AVENU_ACCOUNT            Entity = "account_service_avenu_account"
	ACCOUNT_SERVICE_CONNECTION               Entity = "account_service_connection"
	ACCOUNT_SERVICE_ORG_ACCESS_SETTING       Entity = "account_service_org_access_setting"
	ACCOUNT_SERVICE_EXTERNAL_USER_VISIBILITY Entity = "account_service_external_user_visibility"
	ACCOUNT_SERVICE_EXTERNAL_USER_INVITATION Entity = "account_service_external_user_invitation"
	ACCOUNT_SERVICE_PARTNERSHIP_SETTING      Entity = "account_service_partnership_setting"

	// Security
	SECURITY_USER                         Entity = "security_user"
	SECURITY_ASSOCIATE_USER               Entity = "security_associate_user"
	SECURITY_CSRF                         Entity = "security_csrf"
	SECURITY_ME                           Entity = "security_me"
	SECURITY_ROLE                         Entity = "security_role"
	SECURITY_USER_INVITATION              Entity = "security_user_invitation"
	SECURITY_USER_FAVOURITE               Entity = "security_user_favourite"
	SECURITY_USER_GROUP                   Entity = "security_user_group"
	SECURITY_USER_TEAM                    Entity = "security_user_team"
	SECURITY_USER_PROFILE                 Entity = "security_user_profile"
	SECURITY_COLLABORATOR_ROLE            Entity = "security_collaborator_role"
	SECURITY_ORGANIZATION_ROLE            Entity = "security_organization_role"
	SECURITY_USER_ACTIVENESS              Entity = "security_user_activeness"
	SECURITY_USER_PERSONALIZED_VIEWS      Entity = "security_user_personalized_views"
	SECURITY_SHARING_RULE                 Entity = "security_sharing_rule"
	SECURITY_DEFAULT_ORGANIZATION_SETTING Entity = "security_default_organization_setting"
	SECURITY_TEAM_VISIBILITY              Entity = "security_team_visibility"
	APP_SECURITY                          Entity = "app_security"
	APP_SECURITY_PROFILE                  Entity = "app_security_profile"

	// Platform - todo later.
	// Messaging.
	//MESSAGING_CONVERSATION Entity = "messaging_conversation"
	//MESSAGING_MESSAGE      Entity = "messaging_message"

	//EMAIL                   Entity = "email"
	//AVENUE_EMAIL            Entity = "avenue_email"
	//AVENUE_EMAIL_SIGNATURE  Entity = "avenue_email_signature"
	//AVENUE_EMAIL_SCHEDULE   Entity = "avenue_email_schedule"
	//AVENUE_EMAIL_TEMPLATE   Entity = "avenue_email_template"
	//ADDRESS                 Entity = "address"
	//PHONE                   Entity = "phone"
	//LOOKUP                  Entity = "lookup"
	//REW3_LOOKUP             Entity = "rew3_lookup"
	//RTF                     Entity = "rtf"
	//LIST_VIEW               Entity = "list_view"
	//CONTEXT_USER_PREFERENCE Entity = "context_user_preference"

	// Common Lookup
	//COMMON_LOOKUP Entity = "common_lookup"
)
