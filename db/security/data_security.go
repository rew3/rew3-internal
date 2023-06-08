package security

import (
	c "github.com/rew3/rew3-internal/db/security/constants"
	f "github.com/rew3/rew3-internal/db/security/filter"
	s "github.com/rew3/rew3-internal/service/common/request"
	"go.mongodb.org/mongo-driver/bson"
)

type DataSecurity struct {
	commonDataFilter        f.CommonDataFilter
	securityFilterGenerator f.SecurityFilterGenerator
}

func NewDataSecurity() DataSecurity {
	return DataSecurity{
		commonDataFilter:        f.CommonDataFilter{},
		securityFilterGenerator: f.SecurityFilterGenerator{},
	}
}

func (ds DataSecurity) GenerateFilter(
	scope c.DataAccessScope,
	context *s.RequestContext,
	transitionalFilters ...*bson.D) bson.D {

	// generate security filter from commonDataFilter and security Filter generator.
	return bson.D{}
}

func (ds DataSecurity) GenerateExtendedFilter(
	scope c.DataAccessScope,
	context *s.RequestContext,
	param *s.RequestParam,
	isApplyArchiveFilter bool,
	transitionalFilters ...*bson.D) bson.D {

	return bson.D{}
}
