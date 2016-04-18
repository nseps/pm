package pm

import "time"

type Type int

const (
	Bool Type = iota
	Number
	Text
	TextLong
	DateTime
	Enum
)

var types = [...]string {
	"Bool",
	"Number",
	"Text",
	"TextLong",
	"DateTime",
	"Enum",
}

func (t Type) String() string {
	return types[t]
}

type Fielder interface  {
	MarshalJSON() (json string, err error)
}

type fieldBase struct {
	Name string
	IsMandatory bool
}

type BoolField struct {
	fieldBase
	Value bool
}

type TextField string

type NumberField float64

type DateTimeField time.Time