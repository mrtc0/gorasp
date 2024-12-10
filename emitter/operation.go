package emitter

type Operation interface {
	unwrap() *operation
	SetIsBlock(b bool)
	IsBlocked() bool
}

type operation struct {
	eventRegister
	isBlocked bool
}

func (o *operation) unwrap() *operation {
	return o
}

func (o *operation) SetIsBlock(b bool) {
	o.isBlocked = b
}

func (o *operation) IsBlocked() bool {
	return o.isBlocked
}

type ArgOf[O Operation] interface {
	IsArgOf(O)
}

func NewOperation() Operation {
	return &operation{
		eventRegister: eventRegister{
			listeners: make(map[string][]any, 2),
		},
		isBlocked: false,
	}
}

func StartOperation[O Operation, E ArgOf[O]](name string, op O, args E) O {
	emitEvent(&op.unwrap().eventRegister, name, op, args)
	return op
}
