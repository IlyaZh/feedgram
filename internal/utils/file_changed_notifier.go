package utils

import (
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
)

func FileChangedNotify(filePath string, execFunc func() error) {
	lastSyncTime := time.Now()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	err = watcher.Add(filepath.Dir(filePath))
	if err != nil {
		panic(err)
	}

	defer watcher.Close()
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				log.Error("error whilw getting events from directory watcher")
				return
			}
			syncTime := time.Now()
			if event.Has(fsnotify.Write) && event.Name == filePath && syncTime.After(lastSyncTime.Add(time.Duration(1)*time.Second)) { // debounce
				lastSyncTime = syncTime
				log.Debugf("modified file: %s", filePath)
				err = execFunc()
				if err != nil {
					log.Errorf("error while execute function for file watcher. File = '%s', error = '%s'", filePath, err.Error())
					panic(err)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			panic(err)
		}
	}
}
