package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// [
// 	{
// 	"createdAt": "2024-06-23T14:03:29.066Z",
// 	"name": "Don Kemmer",
// 	"avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/192.jpg",
// 	"id": "1"
// 	},
// 	...
// ]

type User struct {
	CreatedAt string
	Name      string
	Avatar    string
	Id        string
}

func ApiRequest() []User {
	resp, err := http.Get("https://6678ecbe0bd4525056204c1f.mockapi.io/api/v1/users")

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Printf("Error: %s", readErr.Error())
		os.Exit(1)
	}

	var users []User

	parseErr := json.Unmarshal(body, &users)
	if parseErr != nil {
		fmt.Printf("Error(parseErr): %s", parseErr.Error())
		os.Exit(1)
	}

	return users
}

func ApiRequestWithCtx(ctx context.Context) ([]User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	ch := make(chan []User)

	go func() {
		// Slow api request
		ch <- ApiRequest()
	}()

	select {
	case <-ctx.Done():
		return []User{}, ctx.Err()
	case result := <-ch:
		return result, nil
	}

}

func RequestSlowApiRequest() {
	ctx := context.Background()

	fmt.Println("Run ApiRequestWithCtx request...")

	users, err := ApiRequestWithCtx(ctx)
	if err != nil {
		fmt.Printf("Error ApiRequestWithCtx: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("Result: ", users)

	fmt.Println("ApiRequestWithCtx Completed!")
}
