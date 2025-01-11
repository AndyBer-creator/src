package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

func main() {
	defaulLinks := []string{
		"https://ya.ru",
		"https://yandex.ru",
		"https://wikipedia.ru",
	}

	query := "1"
	result, err := SiteSearch(query, defaulLinks)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Println(strings.Join(result, "\n"))
}

func SiteSearch(needle string, urls []string) ([]string, error) {

	group := struct {
		errgroup.Group
		sync.Mutex
		urls []string
	}{
		urls: make([]string, 0, len(urls)),
	}
	for _, u := range urls {
		url := u
		group.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}

			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			if strings.Contains(string(body), needle) {
				group.Lock()
				group.urls = append(group.urls, url)
				group.Unlock()
			} else {
				return fmt.Errorf("Не найдена подстрока \"%s\" на сайте %s", needle, url)
			}

			return nil
		})
	}
	err := group.Wait()
	return group.urls, err
}
