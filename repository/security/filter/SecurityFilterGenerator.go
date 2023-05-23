package filter

import (
	s "github.com/rew3/rew3-base/data/service"
	c "github.com/rew3/rew3-base/repository/security/constants"
	"go.mongodb.org/mongo-driver/bson"
)

type SecurityFilterGenerator struct{}

func (cb *SecurityFilterGenerator) build(scope *c.DataAccessScope, context s.RequestContext) bson.D {

	// Build core data filter. based on scope.
	return bson.D{}
}
