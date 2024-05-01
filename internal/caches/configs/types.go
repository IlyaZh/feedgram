package configs

import (
	"github.com/IlyaZh/feedsgram/internal/configs"
	"sync/atomic"
	"time"
)

type Cache struct {
	value    atomic.Pointer[configs.Config]
	secDist  configs.SecDist
	filePath string
	period   time.Duration
	init     bool
}
