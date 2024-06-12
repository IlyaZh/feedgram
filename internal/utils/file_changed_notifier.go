package utils

import (
	"context"
	"path/filepath"
	"time"

	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

func FileChangedNotify(ctx context.Context, filePath string, execFunc func(ctx context.Context) error) {
	ctx = logger.CreateTrace(ctx)
	log := logger.GetLogger(ctx)
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
				log.Debug("file has modified", zap.String("file_path", filePath))
				err = execFunc(ctx)
				if err != nil {
					log.Error("error while execute function for file watche", zap.String("file_path", filePath), zap.Error(err))
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
