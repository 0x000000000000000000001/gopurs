package main

// ffiSchemaVersion changes whenever the shared native/WASM signature contract changes.
const ffiSchemaVersion = 2

type TypeNode struct {
	Type        string      `json:"type"`
	Name        string      `json:"name,omitempty"`
	Identity    string      `json:"identity,omitempty"`
	PackagePath string      `json:"packagePath,omitempty"`
	Underlying  *TypeNode   `json:"underlying,omitempty"`
	Args        []*TypeNode `json:"args,omitempty"`
	ArgNames    []string    `json:"argNames"`
	// Ret is the compatibility view: only a single result can inhabit it.
	Ret         *TypeNode     `json:"ret,omitempty"`
	Results     []*TypeNode   `json:"results"`
	ResultNames []string      `json:"resultNames"`
	Elem        *TypeNode     `json:"elem,omitempty"`
	Key         *TypeNode     `json:"key,omitempty"`
	Val         *TypeNode     `json:"val,omitempty"`
	Length      string        `json:"length,omitempty"`
	Fields      []StructField `json:"fields,omitempty"`
}

// Struct fields retain Go identifiers and tags separately. No serialization
// naming policy is inferred from a json/sql tag by the signature parser.
type StructField struct {
	Name     string    `json:"name"`
	Type     *TypeNode `json:"type"`
	Embedded bool      `json:"embedded"`
	Exported bool      `json:"exported"`
	Tag      string    `json:"tag"`
}

type FFIDecl struct {
	SchemaVersion int         `json:"schemaVersion"`
	Name          string      `json:"name"`
	IsVar         bool        `json:"isVar"`
	TypeParams    []string    `json:"typeParams"`
	Args          []*TypeNode `json:"args"`
	ArgNames      []string    `json:"argNames"`
	Ret           *TypeNode   `json:"ret"`
	Results       []*TypeNode `json:"results"`
	ResultNames   []string    `json:"resultNames"`
	ValueType     *TypeNode   `json:"valueType,omitempty"`
	Receiver      *TypeNode   `json:"receiver,omitempty"`
	Signature     string      `json:"signature,omitempty"`
	Line          int         `json:"line"`
	Column        int         `json:"column"`
}
