package main

import "github.com/Equanox/gotron"

func main() {
	window, err := gotron.New()
	if err != nil {
		panic(err)
	}

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	processEvent(window)

	<-done
}

func processEvent(w *gotron.BrowserWindow) {
	w.On(&gotron.Event{Event: "app.host.start"}, func(bin []byte) {
		host := newHost(w)
		go host.run()
	})

	w.On(&gotron.Event{Event: "app.worker.start"}, func(bin []byte) {
		worker := newWorker(w)
		go worker.run()
	})
}

