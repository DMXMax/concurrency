package generator

var Repeat = func(done <-chan any, values ...any) <-chan any {
	valueStream := make(chan any, 2)
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

var Take = func(done <-chan any, valueStream <-chan any, num int) <-chan any {
	takeStream := make(chan any)
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}
