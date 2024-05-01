package configs

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/IlyaZh/feedsgram/internal/configs"
	"github.com/IlyaZh/feedsgram/internal/utils"
	"github.com/labstack/gommon/log"
)

type ConfigsCache interface {
	GetValues() *configs.Config
}

var cache *Cache

// TODO make chan for reading in goroutine

func (c *Cache) loadFromFilePeriodic(ctx context.Context) {
	for range time.Tick(cache.period) {
		select {
		case <-ctx.Done():
			return // exit properly on cancellation
		default:
			yamlFile, err := os.ReadFile(c.filePath)
			if err != nil {
				log.Panicf("Error occured while loading config file: %s", cache.filePath)
			}
			var newValue configs.Config
			err = newValue.Scan(yamlFile, c.secDist)
			if err != nil {
				log.Panicf("Error occured while parsing config file: %s", cache.filePath)
			}

			_ = cache.value.Swap(&newValue)
			cache.init = true
		}
	}
}

func NewCache(ctx context.Context, configFileName string, secdistFilePath string, period time.Duration) *Cache {
	if cache != nil {
		return cache
	}
	log.Infof("[%s] Component start initialization", name)

	cache = &Cache{filePath: utils.MakePath(nil, configFileName), period: period}

	cache.secDist = configs.NewSecDist(utils.MakePath(nil, secdistFilePath))

	ctx, cancel := context.WithTimeout(ctx, time.Duration(3*time.Second))
	defer cancel()

	go cache.loadFromFilePeriodic(ctx)
	log.Infof("[%s] Component wait for loading file", name)
	start := time.Now()
	for !cache.init {
		if time.Since(start) > time.Duration(timeout)*time.Second {
			errorMessage := fmt.Sprintf("[%s] Component hasn't initialized until %d seconds timeout. Panic!", name, timeout)
			log.Errorf(errorMessage)
			panic(errorMessage)
		}
	}

	log.Infof("[%s] Component initialization has finished", name)

	return cache

}

func (c *Cache) GetValues() *configs.Config {
	if cache == nil {
		panic("Trying to get config which value is none")
	}
	return cache.value.Load()
}
