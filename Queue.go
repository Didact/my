package my

type T interface{}
type Pipe <-chan T

//Shamelessly stolen from https://github.com/kylelemons/iq
func (p Pipe) Queue() Pipe {
	next := make(chan T)
	pending := []T{}
	go func() {
		defer close(next)
	L:
		for {
			if len(pending) == 0 {
				v, ok := <-p
				if !ok {
					break
				}
				pending = append(pending, v)
			}
			select {
			case v, ok := <-p:
				if !ok {
					break L
				}
				pending = append(pending, v)
			case next <- pending[0]:
				pending = pending[1:]
			}
		}
		for _, v := range pending {
			next <- v
		}
	}()
	return next
}
