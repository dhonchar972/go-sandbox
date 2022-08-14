package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		err := SomeFunc(ctx)
		if err != nil {
			cancel()
		}
	}()

	doRequest(ctx, "https://www.youtube.com/")
}

func SomeFunc(ctx context.Context) error {
	time.Sleep(1500 * time.Millisecond)
	return fmt.Errorf("error")
}

func doRequest(ctx context.Context, requestStr string) {
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("response completed, status code=%d", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("request too long")
	}
}
