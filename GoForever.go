package my

func GoForever(f func()) {
	go func() {
		for {
			f()
		}
	}()
}
