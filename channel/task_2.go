package main

/*func doWork(size int, capacity int) int {
	start := time.Now()
	jobs := make(chan *Job, capacity)
	results := make(chan *Job, capacity)
	sem := make(chan struct{}, capacity)
	go chanWorker(jobs, results, sem)
	for i := 0; i < size; i++ {
		jobs <- &Job{id: i}
	}
	close(jobs)
	successCount := 0
	for i := 0; i < size; i++ {
		item := <-results
		if item.result {
			successCount++
		}
		fmt.Printf("Job %d completed %v\n", item.id, item.result)
	}
	close(results)
	close(sem)
	fmt.Printf("Time taken to execute %d jobs with %d capacity = %v\n", size, capacity, time.Since(start))
	return successCount
}
func chanWorker(jobs <-chan *Job, results chan<- *Job, sem chan struct{}) {

	for item := range jobs {
		it := item
		sem <- struct{}{}
		fmt.Printf("Job %d started\n", it.id)
		go func() {
			timeOutCtx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
			defer cancel()
			time.Sleep(time.Duration(it.id) * 100 * time.Millisecond)
			select {
			case <-timeOutCtx.Done():
				fmt.Printf("Job %d timed out\n", it.id)
				it.result = false
				results <- it
				<-sem
				return
			default:
				fmt.Printf("Total number of routines %d\n", runtime.NumGoroutine())
				it.result = true
				results <- it
				<-sem
			}
		}()
	}
}
*/
