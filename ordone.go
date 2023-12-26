package concurrency

var OrDone = func(done <-chan any, c <-chan any) <-chan any {
	valStream := make(chan any)
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case <-done:
				case valStream <- v:
				}
			}
		}
	}()
	return valStream
}
