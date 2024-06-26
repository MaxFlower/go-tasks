package channel

import "fmt"

func traverse(c chan int) {
	for i := 0; i < 10; i++ {
		select {
		case c <- i + 5:
			fmt.Println("process: ", i)
		default:
			fmt.Print("No space for ", i, " ")
		}
	}
}

func RunChannels() {
	ch := make(chan int, 10)

	traverse(ch)

	for {
		select {
		case res := <-ch:
			fmt.Println(res)
		default:
			fmt.Println("Nothing left to read!")
			close(ch)
			return
		}
	}
}
