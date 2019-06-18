package events

type Dispatcher interface {
	Dispatch(e Event) error
}

type Event interface {
	Body() interface{}
}
