package motorhead

type Field struct {
	Parent *Field
	Name string
	Value interface{}
}

type Specification struct {
	Fields map[string]Field
}