package utils

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFileChangedNotify(t *testing.T) {
	ch := make(chan struct{})
	currentPath, err := os.Getwd()
	if err != nil {
		t.Fatalf("error while getting WD: %s", err.Error())
	}

	type args struct {
		done      chan struct{}
		execFunc  func() error
		wantError bool
	}
	tests := []struct {
		name string
		args args
	}{
		// {
		// 	name: "OK",
		// 	args: args{
		// 		done: ch,
		// 		execFunc: func() error {
		// 			ch <- struct{}{}
		// 			return nil
		// 		},
		// 		wantError: false,
		// 	},
		// },
		{
			name: "error in watcher function",
			args: args{
				done: ch,
				execFunc: func() error {
					return errors.New("some error")
				},
				wantError: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.CreateTemp(currentPath, "_test_*.txt")
			if err != nil {
				panic(err)
			}
			defer func() {
				t.Logf("Remove file: %s", file.Name())
				os.Remove(file.Name())
			}()
			defer func() {
				e := recover()
				if e != nil {
					if !tt.args.wantError {
						t.Fatalf("unexpected panic: %s", e)
					}
				} else if tt.args.wantError {
					t.Fatalf("expected panic hasn't appeared")
				}
			}()
			t.Logf("file: %s", file.Name())

			if tt.args.wantError {
				require.Panics(t, func() {
					go FileChangedNotify(file.Name(), tt.args.execFunc)
				})
			} else {
				go FileChangedNotify(file.Name(), tt.args.execFunc)
			}
			time.Sleep(time.Duration(1) * time.Second)
			tmr := time.NewTimer(time.Duration(1) * time.Second)
			// for debounce checker write 3 times as fast as possible
			for i := 0; i < 3; i++ {
				_, err = file.WriteString(fmt.Sprintf("row %d", i))
				if err != nil {
					t.Fatalf("error while writing file: %s", err.Error())
				}
				err = file.Sync()
				if err != nil {
					t.Fatalf("error while syncing file: %s", err.Error())
				}
			}
			file.Close()
			if tt.args.wantError {
				return
			}
			select {
			case <-ch:
				return
			case <-tmr.C:
				t.Fatalf("timeout occured")
			}
		})
	}
	close(ch)
}
