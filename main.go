package main

import (
	"fmt"
	"gotasks/api"
	"gotasks/cache"
	"gotasks/channel"
	msg "gotasks/message"
	"gotasks/slices"
)

func main() {

	// Slices

	var sl slices.Slices
	sl.Fill(7)
	fmt.Println(sl.FromToRight(1)) // [1 - 7]
	fmt.Println(sl.LeftTo(3))      // [0 - 3]
	sl.PrintOrigin()

	fmt.Println("-------------------")

	// Interfaces & structs

	var channels []msg.MessageService
	seven := msg.SevenNewsService{Messages: make([]msg.Message, 0)}
	channels = append(channels, &seven)
	message := channels[0].NewMessage("fsdbdsfh", "test", "info")
	channels[0].Add(message)
	channels[0].Print()
	fmt.Println("-------------------")

	// Cache

	lru := cache.New()
	lru.Add(5)
	lru.Add(7)
	lru.Add(11)
	lru.Add(8)
	lru.Add(9)  // 9 8 11 7 5
	lru.Add(3)  // 3 9 8 11 7
	lru.Add(8)  // 8 3 9 11 7
	lru.Add(11) // 11 8 3 9 7
	lru.Add(6)  // 6 11 8 3 9

	lru.PrintQueue()

	fmt.Println("-------------------")

	// Slow API

	api.RequestSlowApiRequest()

	fmt.Println("-------------------")

	channel.RunChannels()

	fmt.Println("-------------------")
}
