package configs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/IlyaZh/feedsgram/internal/configs"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/IlyaZh/feedsgram/internal/utils"
	"go.uber.org/zap"
)

//go:generate mockgen -source component.go -package mocks -destination mocks/component.go
type ConfigsCache interface {
	GetValues() configs.Config
}

var cache *Cache

func (c *Cache) loadFromFile(ctx context.Context) error {

	yamlFile, err := os.ReadFile(c.filePath)
	if err != nil {
		return fmt.Errorf("error occured while loading config file: %s", cache.filePath)
	}
	var newValue configs.Config
	err = newValue.Scan(ctx, yamlFile, c.secDist)
	if err != nil {
		return fmt.Errorf("error occured while parsing config file: %s", cache.filePath)
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.value = newValue

	return nil
}

func NewCache(ctx context.Context, configFilePath string, secdistFilePath string, period time.Duration) *Cache {
	if cache != nil {
		return cache
	}
	log := logger.GetLoggerComponent(ctx, name)

	configAbsFilePath, err := filepath.Abs(configFilePath)
	if err != nil {
		panic(err)
	}

	cache = &Cache{filePath: configAbsFilePath, period: period}

	secdistFile, err := filepath.Abs(secdistFilePath)
	if err != nil {
		panic(err)
	}
	cache.secDist = configs.NewSecDist(ctx, secdistFile)

	// check if initial loading is done
	log.Info("Wait for loading file", zap.String("file", name))
	err = cache.loadFromFile(ctx)
	if err != nil {
		panic(err)
	}
	log.Info("Init OK")

	go utils.FileChangedNotify(ctx, configAbsFilePath, cache.loadFromFile)

	log.Info("Initialization has finished")

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
