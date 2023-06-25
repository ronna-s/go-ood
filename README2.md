# A path to OOD with Go - Workshop

https://github.com/ronna-s/go-ood/ [Clone me!]

This workshop is aimed to clarify the OOP features that Go provides.
It is named A Path to OOD and not OOP because different language features mean different design concepts.

## Logistics:
All exercises can be completed using the go tool, docker or a combination docker and the make tool.
If you are planning to use the go tool directly you can skip this step.

If planning to use docker and you don't have the make tool, run:
```bash
docker build . -t go-ood
```

If you have the make tool and docker, run:
```bash
make build
```

## Schedule
TBA


## Introduction to Object-Oriented Programming
### What is OOP?
The term Object-Oriented Programming means different things to different people (and we will learn about that), but first, here's what we can all agree on: 
The central idea behind Object-Oriented is to divide software into "things" or "objects" or "instances" that communicate via "messages" or "methods" or "member functions".
Or in short, combining data and functionality.
This core idea has not changed in the 4-5+ decades since it was conceptualized.
It is meant to allow the developer to build code and separate responsibilities or concerns just like in the real world which is what we are familiar with and how we generally think and solve problems.

### Exercise 1 - Understanding the benefits
To understand the benefits of OO, we are going to help a gopher get out of a maze, without knowing anything about how the maze is implemented (or the gopher).

*This exercise is heavily inspired by the Intro to CS first home assignment that [Prof. Jeff Rosenschein](https://scholar.google.com/citations?user=YO7cKNMAAAAJ&hl=en) gave my Intro to CS class in 2003.

To get a sense of what strong OOP can do for us, solve a maze given a Gopher that can perform 4 actions:
```go
// Gopher is an interface to an object that can move around a maze
type Gopher interface {
	Finished() bool // Has the Gopher reached the target cell?
	Move() error    // The Gopher moves one step in its current direction
	TurnLeft()      // The Gopher will turn left
	TurnRight()     // The Gopher will turn right
}
```

Find the function `SolveMaze(g Gopher)` in cmd/maze/maze.go and implement it.

#### Run the tests:
```bash
# go tool
go test github.com/ronna-s/go-ood/cmd/maze 
# make + docker (linux, mac)
make test-maze 
# docker directly (linux, mac)
docker run -v $(pwd):/root --rm -it go-ood go test github.com/ronna-s/go-ood/cmd/maze
# docker on windows + powershell
docker run -v $(PWD):/root --rm -it go-ood go test github.com/ronna-s/go-ood/cmd/maze
# docker on windows without powershell
docker run -v %cd%:/root --rm -it go-ood go test github.com/ronna-s/go-ood/cmd/maze
```

The test checks for very basic navigation. You can also check what your code is doing by running:
```bash
# go tool
go run cmd/maze/maze.go > tmp/maze.html
# make + docker
make run-maze > tmp/maze.html
# any other setup with docker 
[docker command from before] go run cmd/maze/maze.go > tmp/maze.html 
```

Open tmp/maze.html file in your browser to see the results of your code.
You can run the app multiple times to see your gopher running through different mazes.

Done? If not, don't worry. You have the entire conference ;)

### Basic Go for OO
This section is meant to highlight most of the functionality that supports OO in Go. 

#### Type definition
The following code defines a new type A with the underlying type bool.
It then defines a new type B with the underlying type A
We can instantiate the variable `a` (of type `A`) with false (a constant) but to convert `a` to type `B`, we have to be explicit because they are different types.

```go
package main

import "fmt"

type A bool
type B A

func main() {
	var a A = false
	var b B = !B(a) //try changing this to: var b B = a
	fmt.Println(a, b)
}
```
[Run me!](https://go.dev/play/p/2qIxPbBc5QD)

#### Adding Methods - Value Receivers
We can add methods to any type in Go using receivers.
This is how we add a method using a value receiver.
```go
package main

import "fmt"

type A int

func (a A) Zero() bool {
	return a == 0
}

func main() {
	var a A
	fmt.Println(a.Zero())
}
```
[Run me](https://go.dev/play/p/QKj_67aRtMl)

#### Structs
We can create more complex types using structs

```go
package main

import "fmt"

type Person struct {
	Name    string
	Hobbies []Hobby
}

type Hobby string

func main() {
	p := Person{Name: "Batman", Hobbies: []Hobby{"Vigilantism"}}
	fmt.Println("Name:", p.Name, "Hobbies :", p.Hobbies)
}
```
[Run me](https://go.dev/play/p/zyQGSFxLJDU)

#### Adding Methods - Pointer Receivers

If your method requires changing the value of the receiver (outside the method),
we must use a pointer.
From [A Tour of Go](https://go.dev/tour/methods/4)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```
This is NOT what is referred to as pass by value and pass by reference. 
That terminology is used to describe languages with constructors and particularly copy constructors. 
In Go there is no copy constructor but de facto everything is passed by value and something is always copy.
The difference is, that if you use a pointer explicitly (and pointers are explicit in Go), what is copied is the address of the value, so the memory is still accessible.

#### Aliases
You can alias to any type, but what does alias mean?
```go
package main

import "fmt"

type A int
type B = A

func (b B) Foo() int {
	return int(b)
}
func main() {
	var a A = 5
	fmt.Println(a.Foo())

	var b B
	fmt.Printf("a:%T b:%T", a, b)
}
```
[Run me](https://go.dev/play/p/ZqVJbl-2E0m)

#### Embedding & Promotion
#### Overriding methods
#### Interfaces
### The empty interface (any)
#### Interfaces embedding
#### Interface type assertion
### Exercise 1 - how it all worked
Now that we have the basics we can go back and understand the code.

### Object-Oriented fundamentals and Go
#### Do we need OOP?
#### Is Go an OO language? == Is t an object?
#### The C++/Java School of OO
#### Missing CTORs?
#### Missing Inheritance? and, Composition vs. Inheritance
#### The Alan Kay School of OO
#### Missing Messages?
#### Conclusion: is Go OO?
##### the Go FAQ - is Go an object-oriented language?
##### Rob Pike - Go is Object-Oriented
##### Conclusion: Is t an object?

## Go OO in Practice
### Exercise 2 - Interfaces and Embedding
### Stringers
#### Effective interface type assertion
### Organizing your packages
### Code generation, why? When?
### Emerging patterns
#### Complex CTORs with no function overloading
#### Default variables, exported variables, overrideable and otherwise
#### Short Lived Objects vs. Long Lived Objects
#### Context keys as types

## Generics
### Introduction to Generics
### Exercise 3 - Generics
## Conclusion
