package emitter

type Operation interface {
	unwrap() *operation
}

type operation struct {
	eventRegister
}

func (o *operation) unwrap() *operation {
	return o
}

type ArgOf[O Operation] interface {
	IsArgOf(O)
}

func NewOperation() Operation {
	return &operation{
		eventRegister: eventRegister{
			listeners: make(map[string][]any, 2),
		},
	}
}

func StartOperation[O Operation, E ArgOf[O]](name string, op O, args E) O {
	emitEvent(&op.unwrap().eventRegister, name, op, args)
	return op
}
