package service

import (
)

type Service struct {
	cfg        *conf.Config
	goodsDao   *goods.Dao
	collectDao *collect.Dao
	es         *es.Dao
}

// New create service instance and return.
func New(cfg *conf.Config) *Service {
	return &Service{
		cfg:        cfg,
		goodsDao:   goods.New(cfg),
		collectDao: collect.New(cfg),
		es:         es.New(cfg),
	}
}

// Close close the resource.
func (s *Service) Close() {
	s.goodsDao.Close()
	s.collectDao.Close()
}
