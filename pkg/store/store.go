package store

import (
	"nginx-reload/pkg/entity"
)

type Store interface {
	AllConfig() []entity.ProxyConfig
}
