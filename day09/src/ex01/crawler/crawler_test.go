package crawler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCrawlerWeb_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer ts.Close()

	urls := make(chan string, 1)
	urls <- ts.URL
	close(urls)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bodyChan := CrawlerWeb(ctx, urls)

	select {
	case bodies := <-bodyChan:
		if len(bodies) != 1 {
			t.Fatalf("expected 1 result, got %d", len(bodies))
		}
		if *bodies[0] != "Hello, World!" {
			t.Fatalf("expected 'Hello, World!', got %s", *bodies[0])
		}
	case <-time.After(5 * time.Second):
		t.Fatal("test timed out")
	}
}

func TestCrawlerWeb_Failure(t *testing.T) {
	urls := make(chan string, 1)
	urls <- "http://nonexistent.url"
	close(urls)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bodyChan := CrawlerWeb(ctx, urls)

	select {
	case bodies := <-bodyChan:
		if len(bodies) != 1 {
			t.Fatalf("expected 1 result, got %d", len(bodies))
		}
		if !strings.Contains(*bodies[0], "no such host") {
			t.Fatalf("expected an error containing 'no such host', got %s", *bodies[0])
		}
	case <-time.After(5 * time.Second):
		t.Fatal("test timed out")
	}
}

func TestCrawlerWeb_ContextCancel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer ts.Close()

	urls := make(chan string, 1)
	urls <- ts.URL
	close(urls)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	bodyChan := CrawlerWeb(ctx, urls)

	select {
	case <-bodyChan:
		t.Fatal("expected context cancellation, but got a result")
	case <-ctx.Done():
	}
}

func TestCrawlerWeb_MultipleURLs(t *testing.T) {
	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Response from server 1"))
	}))
	defer ts1.Close()

	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Response from server 2"))
	}))
	defer ts2.Close()

	urls := make(chan string, 2)
	urls <- ts1.URL
	urls <- ts2.URL
	close(urls)

	// Создаем контекст
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Запускаем веб-краулер
	bodyChan := CrawlerWeb(ctx, urls)

	// Проверяем результат
	select {
	case bodies := <-bodyChan:
		if len(bodies) != 2 {
			t.Fatalf("expected 2 results, got %d", len(bodies))
		}
		expected := map[string]bool{
			"Response from server 1": false,
			"Response from server 2": false,
		}
		for _, body := range bodies {
			if _, ok := expected[*body]; ok {
				expected[*body] = true
			} else {
				t.Fatalf("unexpected response: %s", *body)
			}
		}
		for k, v := range expected {
			if !v {
				t.Fatalf("expected response '%s' not received", k)
			}
		}
	case <-time.After(5 * time.Second):
		t.Fatal("test timed out")
	}
}
