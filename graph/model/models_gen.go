// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ABUnion interface {
	IsABUnion()
}

type A struct {
	Common *CommonA `json:"common"`
}

func (A) IsABUnion() {}

type B struct {
	Common *CommonB `json:"common"`
}

func (B) IsABUnion() {}

type CommonA struct {
	Test int `json:"test"`
}

type CommonB struct {
	Broken int `json:"broken"`
}

type Query struct {
}

type Unions struct {
	AbUnion ABUnion `json:"abUnion"`
}
