// En kode som demonstrerer hvordan channel fungerer

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // Lager en channel med "make()" funksjonen. Kan bare overføre data av samme type (int)

	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			time.Sleep(1000 * time.Millisecond)
			ch <- i
		}
		close(ch) // Bruker "close()" funksjonen til å stenge kanalen. Den setter i praksis ett flagg som indikerer stopp.
	}(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
