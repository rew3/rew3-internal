package db

import (
	m "github.com/rew3/rew3-base/repository/meta"
	s "github.com/rew3/rew3-base/repository/security"
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
