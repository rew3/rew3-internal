package entity

/*
   A Type to represent [[Rew3Url]].
   contains scheme, full url i.e. http://www.rew3.com and domain url i.e. www.rew3.com of website
   Only domain url is displayed in UI but filtering is done on full url
   User can provide either domain or full url in Account Form
   So to support filtering and sorting in standard format and to support UI displayable url we use [[Rew3Url]]

   @param scheme  scheme of url. For example, http://, https:// , etc
   @param domain  url that does not contains scheme. For example, www.rew3.com
   @param fullUrl full url. For example. http://www.rew3.com
*/

type Rew3Url struct {
	SCHEME  *string `bson:"scheme,omitempty"`
	DOMAIN  *string `bson:"domain,omitempty"`
	FULLURL *string `bson:"full_url,omitempty"`
}
