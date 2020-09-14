package worker

import (
	"runtime"
)

type AbstractWorker interface {
	InstantiateResources()
	WorkersToCreate() int
	HeavyWork(chan int) bool
}

func StartWorking(worker AbstractWorker) {

	worker.InstantiateResources()

	var workersToCreate int = 1

	if worker.WorkersToCreate() == 0 {
		workersToCreate = runtime.NumCPU()
	} else {
		workersToCreate = worker.WorkersToCreate()
	}

	done := make(chan int)
	for i := 0; i < workersToCreate; i++ {

		go worker.HeavyWork(done)
		<-done
	}

}
