package octopus


// каналы только для чтения
func Multiplex(toMultiplex ...<-chan any) <-chan any {
	resultChan := make(chan any)
	
	for _, ch := range toMultiplex {
		go func(c <-chan any) {
			for v := range c {
				resultChan <- v
			}
		}(ch)
	}

	return resultChan	
}
