package main

import (
	"ch8/thumbnail"
	"fmt"
	"log"
	"os"
	"sync"
)

func makeThumbnails(filenames []string) {
	var wg sync.WaitGroup
	for _, file := range filenames {
		wg.Add(1)

		go func(file string) {
			thumbnail.ImageFile(file)
			wg.Done()
		}(file)
	}
	wg.Wait()
}

func makeThumbnails2(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	//使用buffered chan 可以防止当某个goroutine 出错后，下面判断会进行return, 那么其他的g
	//的goroutine想chan发送数据的时候就会阻塞，从而导致goroutine泄露
	ch := make(chan item, len(filenames))
	for _, file := range filenames {

		go func(file string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(file)
			ch <- it
		}(file)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

//参数是一个只接收的channel
func makeThumbnails3(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	//closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	files := []string{"1.png", "2.png", "3.png"}
	c := make(chan string)

	go func() {
		for _, f := range files {
			c <- f
		}
		close(c) //关闭发送channel，这样接收channel才会结束
	}()

	sizes := makeThumbnails3(c)
	fmt.Println(sizes)
}
