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

//go:generate mockgen -source component.go -package mocks -destination mocks/component.go
type ConfigsCache interface {
	GetValues() configs.Config
}

var cache *Cache

func (c *Cache) loadFromFile() error {

	yamlFile, err := os.ReadFile(c.filePath)
	if err != nil {
		return fmt.Errorf("error occured while loading config file: %s", cache.filePath)
	}
	var newValue configs.Config
	err = newValue.Scan(yamlFile, c.secDist)
	if err != nil {
		return fmt.Errorf("error occured while parsing config file: %s", cache.filePath)
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.value = newValue

	return nil
}

func NewCache(ctx context.Context, configFileName string, secdistFilePath string, period time.Duration) *Cache {
	if cache != nil {
		return cache
	}
	log.Infof("[%s] Component start initialization", name)

	cache = &Cache{filePath: utils.MakePath(nil, configFileName), period: period}

	cache.secDist = configs.NewSecDist(utils.MakePath(nil, secdistFilePath))

	// check if initial loading is done
	log.Infof("[%s] Component wait for loading file", name)
	err := cache.loadFromFile()
	if err != nil {
		panic(err)
	}
	log.Infof("[%s] Init OK", name)

	ticker := time.NewTicker(period)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return // exit properly on cancellation
			case <-ticker.C:
				err := cache.loadFromFile()
				if err != nil {
					log.Errorf("error while loading config: %s", err.Error())
				}
			}
		}
	}()

	log.Infof("[%s] Component initialization has finished", name)

	return cache

}

func (c *Cache) GetValues() configs.Config {
	if cache == nil {
		panic("Trying to get config which value is none")
	}
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.value
}
