package config

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"gopkg.in/yaml.v2"
)

type Cache struct {
	value      atomic.Pointer[entities.Config]
	fileName   string
	periodSecs int
	init       bool
}

var cache *Cache

const (
	PRERIX          = "../configs/"
	NAME            = "configs_cache"
	LOADING_TIMEOUT = 5
)

// TODO make chan for reading in goroutine

func loadFromFilePeriodic() {
	for range time.Tick(time.Second * time.Duration(cache.periodSecs)) {
		yamlFile, err := os.ReadFile(fmt.Sprintf("%s%s", PRERIX, cache.fileName))
		if err != nil {
			log.Panicf("Error occured while loading config file: %s%s", PRERIX, cache.fileName)
		}
		newValue := &entities.Config{}
		err = yaml.Unmarshal(yamlFile, newValue)
		if err != nil {
			log.Panicf("Error occured while parsing config file: %s%s", PRERIX, cache.fileName)
		}
		_ = cache.value.Swap(newValue)
		cache.init = true
	}
}

func NewCache(filename string, periodSecs int) *Cache {
	if cache != nil {
		return cache
	}
	log.Printf("[%s] Component start initialization", NAME)

	cache = &Cache{fileName: filename, periodSecs: periodSecs}

	go loadFromFilePeriodic()

	log.Printf("[%s] Component wait for loading file", NAME)
	start := time.Now()
	for !cache.init {
		if time.Since(start) > time.Duration(LOADING_TIMEOUT)*time.Second {
			errorMessage := fmt.Sprintf("[%s] Component hasn't initialized until %d seconds timeout. Panic!", NAME, LOADING_TIMEOUT)
			log.Fatalln(errorMessage)
			panic(errorMessage)
		}
	}
	log.Printf("[%s] Component initialization has finished", NAME)

	return cache

}

func (c *Cache) GetValues() *entities.Config {
	if cache == nil {
		panic("Trying to get config which value is none")
	}
	return cache.value.Load()
}
