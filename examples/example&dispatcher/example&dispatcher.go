package main

import (
	"fmt"
	"github.com/czaorz/dispatcher"
	"net/http"
)

const (
	signal_1 = dispatcher.Signal(iota)
	signal_2
	signal_3
)

type Task struct {
	Name string
}

func (t Task) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.RequestURI {
	case "/1":
		_ = dispatcher.Dispatcher(signal_1, "this", "is", "1")
	case "/2":
		_ = dispatcher.Dispatcher(signal_2, "this", "is", "2")
	case "/3":
		_ = dispatcher.Dispatcher(signal_3, "this", "is", "3")
	}
	_, _ = writer.Write([]byte("hello world"))
}

func (t Task) Dis(signal dispatcher.Signal, args ...interface{}) {
	fmt.Println(t.Name, signal, args)
}

func main() {
	dispatcher.AddDisWithSignal(Task{}, signal_1, signal_2, signal_3)
	dispatcher.AddDisWithSignal(Task{Name: "test1"}, signal_1)
	dispatcher.AddDisWithSignal(Task{Name: "test11"}, signal_1)
	dispatcher.AddDisWithSignal(Task{Name: "test2"}, signal_2)
	dispatcher.AddDisWithSignal(Task{Name: "test3"}, signal_3)
	err := http.ListenAndServe(":8080", Task{})
	if err != nil {
		panic(err)
	}
}
