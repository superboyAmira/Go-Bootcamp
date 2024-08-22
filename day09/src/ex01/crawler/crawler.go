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
			sem <- struct{}{} // Блокируем, если горутин слишком много

			go func(i_t int, url_t string) {
				defer wg.Done()
				defer func() { <-sem }() // Освобождаем семафор по завершению горутины

				res := new(string)
				resp, err := http.Get(url_t)
				if err != nil {
					*res = err.Error()
					result[i_t] = res
					return
				}
				defer resp.Body.Close()

				upackedBody, err := io.ReadAll(resp.Body)
				if err != nil {
					*res = err.Error()
				} else {
					*res = string(upackedBody)
				}
				result[i_t] = res
			}(i, url)
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