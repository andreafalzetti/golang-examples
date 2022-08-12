package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	b64 "encoding/base64"

	"github.com/fsnotify/fsnotify"
)

type AdditionalContent struct {
	GitpodYml string `json:".gitpod.yml"`
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					log.Println("modified file:", event.String())

					data, err := os.ReadFile(".gitpod.yml")
					if err != nil {
						panic(err)
					}

					fmt.Println(string(data))

					additionalContentObj := &AdditionalContent{GitpodYml: string(data)}

					b, err := json.Marshal(additionalContentObj)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println(string(b))

					additionalContent := b64.StdEncoding.EncodeToString([]byte(string(b)))
					fmt.Println(additionalContent)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(".gitpod.yml")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
