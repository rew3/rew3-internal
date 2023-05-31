package repository

import (
	m "github.com/rew3/rew3-base/db/meta"
	s "github.com/rew3/rew3-base/db/security"
)

type RepositoryContext struct {
	MetaDataWriter m.MetaDataWriter
	DataSecurity   s.DataSecurity
}

func NewRepositoryContext() RepositoryContext {
	return RepositoryContext{
		MetaDataWriter: m.MetaDataWriter{},
		DataSecurity:   s.DataSecurity{},
	}
}
