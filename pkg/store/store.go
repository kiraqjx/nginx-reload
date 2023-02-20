package store

import (
	"github.com/kiraqjx/nginx-reload/pkg/entity"
)

type Store interface {
	AllConfig() []entity.ProxyConfig
}
