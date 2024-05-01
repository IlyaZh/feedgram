package configs

import (
	"sync"
	"time"

	"github.com/IlyaZh/feedsgram/internal/configs"
)

type Cache struct {
	value    configs.Config
	mtx      sync.RWMutex
	secDist  configs.SecDist
	filePath string
	period   time.Duration
}
