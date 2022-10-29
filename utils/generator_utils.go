package utils

type Generator func(string) interface{}

func NewGenerator[T any]() Generator {
	return func(dataType string) interface{} {
		switch dataType {
		case "instance":
			var instance T
			return instance
		case "list":
			var instance []T
			return instance
		default:
			var instance T
			return instance
		}
	}
}
