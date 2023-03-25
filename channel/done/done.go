package main

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int,
	w worker) {
	for n := range w.in {
		println("worker", id, "received", n)
		w.done <- true
	}
}

func createWorker(
	id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	// 但这样纸是顺序打印的
	// 顺序打印就没必要用channel了
	for i, worker := range workers {
		worker.in <- 'a' + i
		<-worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		<-worker.done
	}

}

func main() {
	chanDemo()
}
