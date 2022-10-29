package utils

type Generator func(DataType) interface{}

type DataType uint

const (
	INSTANCE DataType = iota + 1
	PTR
	LIST
)

func NewGenerator[T any]() Generator {
	return func(dataType DataType) interface{} {
		switch dataType {
		case INSTANCE:
			var instance T
			return instance
		case PTR:
			var instance T
			return &instance
		case LIST:
			var instance []T
			return instance
		default:
			var instance T
			return instance
		}
	}
}
