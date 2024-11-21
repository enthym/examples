package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	threads_count := 5
	c := make(chan int, 100)

	wg.Add(1)
	go commiter(c, threads_count)

	for i := range threads_count {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, c)
	}
	wg.Wait()
	close(c)
}

func commiter(c chan int, threads_count int) {
	for val := range c {
		fmt.Println(val, " - callback: ", time.Now())
		threads_count--
		if threads_count == 0 {
			wg.Done()
			break
		}
	}
}

func worker(i int, c chan int) {
	defer wg.Done() // Decrement the WaitGroup counter when done
	fmt.Println(i, " - start: ", time.Now())
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	c <- i
	fmt.Println(i, " - end: ", time.Now())
}

// func main() {

// 	f, err := os.OpenFile("async.output", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	writer := bufio.NewWriter(f)
// 	cb := make(chan int)

// 	for range 10 {
// 		wg.Add(1)
// 		go runner(*writer, cb)
// 	}
// 	wg.Wait()
// 	fmt.Println(<-cb)
// 	close(cb)
// }

// func runner(writer bufio.Writer, callback chan int) {
// 	defer wg.Done()
// 	fmt.Println(time.Now())
// 	writer.WriteString(time.Now().String())
// 	writer.WriteString("\n")
// 	// time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
// 	for range rand.Intn(1000) {
// 		fmt.Print(".")
// 	}
// 	fmt.Println(time.Now())
// 	writer.WriteString(time.Now().String())
// 	writer.WriteString("\n")
// 	writer.Flush()
// 	callback <- 0
// 	return
// }
