package concurrency

var Bridge = func(done <-chan any, chanStream <-chan <-chan any) <-chan any {
	valStream := make(chan any)
	go func() {
		defer close(valStream)
		for {
			var stream <-chan any
			for {
				select {
				case maybeStream, ok := <-chanStream:
					if !ok {
						return
					}
					stream = maybeStream
				case <-done:
					return

				}
				for val := range OrDone(done, stream) {
					select {
					case <-done:
					case valStream <- val:
					}
				}
			}
		}
	}()
	return valStream
}
