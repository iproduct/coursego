package semaphor

type Semaphor chan struct{}

func New(numTokens int) Semaphor {
	return make(chan struct{}, numTokens)
}

func (s Semaphor) Acquire() {
	s <- struct{}{} //  acquire a  token
}

func (s Semaphor) Release() {
	<-s // release a  token
}
