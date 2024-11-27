package model

type ReliabilityLevel int64

const (
	Safe ReliabilityLevel = iota
	Probable
	LessPropable
	Unsafe
)

func (rl ReliabilityLevel) String() string {
	switch rl {
	case Safe:
		return "safe"
	case Probable:
		return "probable"
	case LessPropable:
		return "less probable"
	case Unsafe:
		return "unsafe"
	}
	return "unkown"
}
