package etcd_test

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type Mutex struct {
	ctx     context.Context
	client  *clientv3.Client
	session *concurrency.Session
	mutex   *concurrency.Mutex
	key     string
}

func (m *Mutex) Lock() error {
	s, err := concurrency.NewSession(m.client)
	if err != nil {
		return err
	}
	m.session = s

	m.mutex = concurrency.NewMutex(m.session, m.key)
	err = m.mutex.Lock(m.ctx)
	if err != nil {
		return err
	}
}

func (m *Mutex) UnLock() error {
	err := m.mutex.Unlock(m.ctx)
	if err != nil {
		return err
	}
	return m.session.Close()
}
