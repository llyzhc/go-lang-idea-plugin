package main

import "fmt"
import "go/ast"

type Ref struct {
	Expr    *ast.Expr
}

type Closer interface {
    Close()
}

type CloseImpl struct {}

func (impl *CloseImpl) Close() {
    fmt.Println("closed")
}

type File struct {
    name string
    Closer
}

func getFile() *File {
    return &File{"hello", &CloseImpl{}}
}

func main() {
    f1 := &File{"hello", &CloseImpl{}}
    f1.Close()

    f2 := getFile()
    f2.Close() // Unresolved reference Close
}

type User interface {
    Name() string
}

func <warning>Add0</warning>(users ...User)  { users[0].Name() }
func <warning>Add1</warning>(users ...*User) { users[0].<error>Name</error>() } 
func <warning>Add2</warning>(users User)     { users[0].<error>Name</error>() }
func <warning>Add3</warning>(users ...User)  { users.<error>Name</error>() }

func <warning>Add4</warning>(users *User)    { (*users).Name() }
func <warning>Add5</warning>(users User)     { users.Name() }

func (r *Ref) Pos() {
	l1 := *r.Expr
	l1.Pos()

	l2 := (*r.Expr)
	l2.Pos()

	*r.Expr.<error>Pos</error> ()
	(*r.Expr).Pos()
}