package counter

import "github.com/marcoshuck/holago/grpc/counter/proto"

type Counter interface {
	Increment(input proto.OperationInput) (proto.OperationOutput, error)
	Decrement(input proto.OperationInput) (proto.OperationOutput, error)
	Result() (proto.ResultOutput, error)
}

type counter struct {
	result int32
}

func (c *counter) Increment(input proto.OperationInput) (proto.OperationOutput, error) {
	value := input.GetValue()
	c.result += value
	output := proto.OperationOutput{
		Value:  value,
		Result: c.result,
	}
	return output, nil
}

func (c *counter) Decrement(input proto.OperationInput) (proto.OperationOutput, error) {
	value := input.GetValue()
	c.result -= value
	output := proto.OperationOutput{
		Value:  value,
		Result: c.result,
	}
	return output, nil
}

func (c counter) Result() (proto.ResultOutput, error) {
	return proto.ResultOutput{Result: c.result}, nil
}

func NewCounter() Counter {
	return &counter{}
}