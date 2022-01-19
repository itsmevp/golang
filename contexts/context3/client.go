package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

/*
Two rules related to contexts:
- Whatever in the context has to be request specific
- Whatever in the context should add some extra information but it should not affect the flow
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:9999", nil); err == nil {
		if response, err := http.DefaultClient.Do(request); err == nil {
			defer response.Body.Close()
			if response.StatusCode != http.StatusOK {
				log.Fatal("received non-ok status code")
			} else {
				io.Copy(os.Stdout, response.Body)
			}
		} else {
			log.Fatal(err.Error()) // 2022/01/19 13:42:39 Get "http://localhost:9999": context deadline exceeded
		}
	} else {
		log.Fatal(err.Error())
	}
}
