package filter

import (
	s "github.com/rew3/rew3-internal/service/request"
	"go.mongodb.org/mongo-driver/bson"
)

type CommonDataFilter struct{}

func (filter *CommonDataFilter) ConstructArchivedFilter(param s.RequestParam) bson.D {
	/*requestParam.filters.get.\(s"${ApiConstants.archiveParamName}") match {
		case JsDefined(value: JsValue) => value match {
		  case JsBoolean(bool) => bool
		  case _ => false
		}
		case _ => false
	  }
	  isArchiveFilterApplied(requestParam) match {
		case true => showOnlyArchived
		case false => hideArchivedFilter
	  }*/
	return bson.D{}
}

func (filter *CommonDataFilter) RemoveArchivedFilter(param s.RequestParam) bson.D {
	/*val updatedFilters = requestParam.filters.get.-(s"${ApiConstants.archiveParamName}")
	  requestParam.copy(filters = Some(updatedFilters))*/
	return bson.D{}
}

func (filter *CommonDataFilter) HideSoftDeletedFilter(param s.RequestParam) bson.D {
	// DBQueryBuilder.existsNot("meta._deleted")
	return bson.D{}
}

func (filter *CommonDataFilter) ShoArchivedFilter(param s.RequestParam) bson.D {
	// DBQueryBuilder.exists("meta._archived_at")
	return bson.D{}
}

func (filter *CommonDataFilter) HideArchivedFilter(param s.RequestParam) bson.D {
	// DBQueryBuilder.existsNot("meta._archived_at")
	return bson.D{}
}
