//go:build js && wasm

package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"syscall"
	"syscall/js"
	"time"
)

func getFileAttr(path string) (string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Mode: %s\nOwner: %d\nGroup: %d\nSize: %d\nModified: %s\nIsDir: %t",
		fi.Mode(),
		fi.Sys().(*syscall.Stat_t).Uid,
		fi.Sys().(*syscall.Stat_t).Gid,
		fi.Size(),
		fi.ModTime().String(),
		fi.IsDir(),
	), nil
}

func main() {
	// **ERROR**: fsnotify is not supported in WASM
	/*
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				panic(err)
			}
			defer watcher.Close()

			go func() {
				for {
					select {
					case event, ok := <-watcher.Events:
						if !ok {
							println("not ok")
							return
						}
						if event.Op.Has(fsnotify.Remove) {
							println("remove", event.Name)
						}
						if event.Op.Has(fsnotify.Create) {
							println("create", event.Name)
						}
						if event.Op.Has(fsnotify.Rename) {
							println("rename", event.Name)
						}
						if event.Op.Has(fsnotify.Chmod) {
							println("chmod", event.Name)
						}
						if event.Op.Has((fsnotify.Write)) {
							println("write", event.Name)
						}
					case err, ok := <-watcher.Errors:
						if !ok {
							println("error")
							return
						}
						println("error:", err)
					}
				}
			}()

			err = watcher.Add("/Users/edonymurphy/Library/Mobile Documents/iCloud~md~obsidian/Documents/anthelion/4.permanent/permanent-diary/permanent-2024/Diary-2024-05-04.md")
			if err != nil {
				println("add watch failed, error: ", err.Error())
			}


		fsInfo, err := os.Lstat("/Users/edonymurphy/Library/Mobile Documents/iCloud~md~obsidian/Documents/anthelion/4.permanent/permanent-diary/permanent-2024/Diary-2024-05-04.md")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(fsInfo.ModTime().String())
		fmt.Println(fsInfo.Mode().String())

		tmp, err := os.Open("./xxxxxx.md")
		if err != nil {
			panic(err.Error())
		}
		defer tmp.Close()

		content := fsInfo.ModTime().String() + "\n" + fsInfo.Mode().String() + "\n"
		_, err = tmp.Write([]byte(content))
		if err != nil {
			panic(err.Error())
		}
	*/
	fmt.Println("testing...........")
	_, err := os.Lstat("/Users/edonymurphy/Library/Mobile Documents/iCloud~md~obsidian/Documents/anthelion/4.permanent/permanent-diary/permanent-2024/Diary-2024-05-04.md")
	if err != nil {
		// as expected, lstat is not supported in WASM
		fmt.Println(err.Error())
	}

	/*
		attFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			path := args[0].String()
			attr, err := getFileAttr(path)
			if err != nil {
				js.ValueOf(fmt.Errorf("error getting file attributes: %w", err))
				return nil
			}
			return js.ValueOf(attr)
		})
		js.Global().Call(attFunc.String(), "/Users/edonymurphy/Library/Mobile Documents/iCloud~md~obsidian/Documents/anthelion/4.permanent/permanent-diary/permanent-2024/Diary-2024-05-04.md")
	*/
	ret, err := getFileAttr("/Users/edonymurphy/Library/Mobile Documents/iCloud~md~obsidian/Documents/anthelion/4.permanent/permanent-diary/permanent-2024/Diary-2024-05-04.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret)

	js.Global().Set("jsPI", jsPI())

	js.Global().Call("alert", "this is an alerting!")
	v := js.Global().Get("app")
	fmt.Println(v.Get("title").String())
	fmt.Println(v.Call("getAppTitle", "").String())
	select {}
}

func pi(samples int) float64 {
	cpus := runtime.NumCPU()

	threadSamples := samples / cpus
	results := make(chan float64, cpus)

	for j := 0; j < cpus; j++ {
		go func() {
			var inside int
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < threadSamples; i++ {
				x, y := r.Float64(), r.Float64()

				if x*x+y*y <= 1 {
					inside++
				}
			}
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var total float64
	for i := 0; i < cpus; i++ {
		total += <-results
	}

	return total / float64(cpus)
}

func jsPI() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		samples := args[0].Int()

		return pi(samples)
	})
}
