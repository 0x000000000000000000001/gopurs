package main

type TypeNode struct {
	Type string      `json:"type"` // "Named", "Func", "Array", "Map", "Unknown"
	Name string      `json:"name,omitempty"`
	Args []*TypeNode `json:"args,omitempty"`
	Ret  *TypeNode   `json:"ret,omitempty"`
	Elem *TypeNode   `json:"elem,omitempty"`
	Key  *TypeNode   `json:"key,omitempty"`
	Val  *TypeNode   `json:"val,omitempty"`
}

type FFIDecl struct {
	Name       string      `json:"name"`
	IsVar      bool        `json:"isVar"`
	TypeParams []string    `json:"typeParams"`
	Args       []*TypeNode `json:"args"`
	Ret        *TypeNode   `json:"ret"`
}
