package crawler

import (
	"context"
	"io"
	"net/http"
	"sync"
)

const (
	goroutineMAX = 8
)

func CrawlerWeb(ctx context.Context, urls chan string) (body chan []*string) {
	body = make(chan []*string)

	go func() {
		defer close(body)

		urlsString := []string{}
		for url := range urls {
			urlsString = append(urlsString, url)
		}

		result := make([]*string, len(urlsString))
		wg := sync.WaitGroup{}
		sem := make(chan struct{}, goroutineMAX)

		for i, url := range urlsString {
			wg.Add(1)
			i, url := i, url // Создаем локальные копии переменных

			sem <- struct{}{} // Блокируем, если горутин слишком много

			go func() {
				defer wg.Done()
				defer func() { <-sem }() // Освобождаем семафор по завершению горутины

				res := new(string)
				resp, err := http.Get(url)
				if err != nil {
					*res = err.Error()
					result[i] = res
					return
				}
				defer resp.Body.Close()

				upackedBody, err := io.ReadAll(resp.Body)
				if err != nil {
					*res = err.Error()
				} else {
					*res = string(upackedBody)
				}
				result[i] = res
			}()
		}

		wg.Wait()

		select {
		case <-ctx.Done():
			return
		default:
			body <- result
		}
	}()

	return body
}