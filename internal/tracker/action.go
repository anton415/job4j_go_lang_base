package tracker

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}
