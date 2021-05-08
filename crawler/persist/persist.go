package persist

import (
	"log"
)

type Worker struct {
	ItemChan chan interface{}
}

func (w *Worker)CreateWorker() {
	w.ItemChan = make(chan interface{})

	go func() {
		for {
			item := <-w.ItemChan
			log.Printf("[Persist worker] got item: %v\n", item)
		}
	}()
}

func (w *Worker)Save(item interface{}) {
	go func() {
		w.ItemChan <- item
	}()
}
