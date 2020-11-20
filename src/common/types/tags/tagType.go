package tags

type TagType byte

const (
	BasicTag   TagType = 0 // tag format -> <keyLength:uint16><valueLength:uint16><Key:string>
	BooleanTag TagType = 1 // tag format -> <keyLength:uint16><valueLength:uint16><Key:string><value:boolean>
	StringTag  TagType = 2 // tag format -> <keyLength:uint16><valueLength:uint16><Key:string><value:string>
	IntegerTag TagType = 3 // tag format -> <keyLength:uint16><valueLength:uint16><Key:string><value:integer>
	Float64Tag TagType = 4 // tag format -> <keyLength:uint16><valueLength:uint16><Key:string><value:float64>

)
