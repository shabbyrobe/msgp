package _generated

//go:generate msgp

//msgp:tuple AllowExtra
//msgp:allowextra AllowExtra

type AllowExtra struct {
	Foo string
	Bar string
}

//msgp:tuple AllowExtraEmpty
//msgp:allowextra AllowExtraEmpty

type AllowExtraEmpty struct {
}

//msgp:tuple AllowExtraChild
//msgp:allowextra AllowExtraChild

type AllowExtraChild struct {
	Before string
	Extra  *AllowExtra
	After  string
}
