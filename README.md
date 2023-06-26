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

<hr>

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

### Basic Go for Object-Oriented
This section is meant to highlight most of the functionality that supports OO in Go. 

#### Type definition
The following code defines a new type A with the underlying type bool.
It then defines a new type B with the underlying type A.
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
We can create more complex types using structs.

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
In Go there is no copy constructor but de facto everything is passed by value and something is always copied.
The difference is, that if you use a pointer explicitly (and pointers are explicit in Go), what is copied is the address of the value, so the memory is still accessible.

Teaser: what will the following code do?

```go
// https://go.dev/play/p/iOx0L_p65jz
package main

import "fmt"

type A struct{}

func (a *A) Foo() string {
	return "Hi from foo"
}
func main() {
	var a *A //a is nil
	fmt.Println(a.Foo())
}
```

This code worked because under the hood, the method `a.Foo()` is just sugar syntax to the function `Foo` on the Type level that takes the receiver as a first parameter. 

```go
//https://go.dev/play/p/zVtRx_mX2rq
package main

import "fmt"

type A struct{}

func (a A) Foo() int {
	return 1
}

func main() {
	var a A
	fmt.Println(a.Foo())
	fmt.Println(A.Foo(a)) //exactly the same
}
```
[Run me](https://go.dev/play/p/zVtRx_mX2rq)

#### Aliases
You can define an alias to any type, but what does alias mean?
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

#### Interfaces
One of Go's strongest features is the Interfaces. Go interfaces are implicit and therefore we can plug into them any code that we found on the web anywhere.
This also means that a package does not need to provide interfaces to its clients, only for its own dependencies when necessary.

From ["A Tour of Go"](https://go.dev/tour/methods/10)
```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

```

#### Embedding & Promotion

When type `B` embeds type `A` we say that type `B` is composed of type `A`.
Upon embedding, the methods of `A` become available to `B`, we call this "promotion".

```go
// https://go.dev/play/p/BcNhFRjQ988
type A int //Creates a new type A with an underlying type int

// Foo is now a method of my A
func (a A) Foo() int {
	return int(a)
}

type B struct {
	// B embeds A so B now has method Foo()
	A
}

func (b B) Bar() int {
	return int(b.A)
}

type I interface {
	Foo() int
}

// to implement J we have to provide implementation for Foo() and Bar()
type J interface {
	I
	Bar() int
}

func main() {
	var j J = B{1}
	fmt.Println(j.Foo()) // 1 
	fmt.Println(j.Bar()) // 1
}
```

We can embed as many types as we want.


#### The empty interface (any)
The empty interface `interface{}` (now also comes as the built-in alias `any`), defines an interface with no methods, and therefore requires no methods. 
This is why any type (including primitive types that have no method) can be passed around as `any` or `interface{}`.

#### Interface type assertion
We can check at run time if an interface implements another interface and convert it.

```go
package main

import "fmt"

type A struct{}

func (a *A) Foo() string {
	return "Hi from foo"
}

type I interface {
	Foo() string
}

func main() {
	var a *A //a is nil
	var b interface{} = a

	if val, ok := b.(I); ok {
		fmt.Println(val.Foo())
	} else {
		fmt.Println("val doesn't have Foo() int and doesn't implement I")
	}
}

```

### Exercise 1 - how it all worked
Now that we have the basics we can go back and understand the code.

Let's review the code that made this possible and examine the Go features it uses.

Run:
```bash
# make tool + docker
make godoc
# using docker
docker run --rm -p 8080:8080 go-ood godoc -http=:8080
# or, install godoc and run
go install golang.org/x/tools/cmd/godoc@v0.1.12
godoc -http=:8080 #assuming $GOBIN is part of your path. For help run `go help install`
```
The repo started with one package in the pkg directory called maze which offers a basic maze generator and nothing else.
Go to: http://127.0.0.1:8080/pkg/github.com/ronna-s/go-ood/pkg/maze

The package defines 5 types:
1. Cell - an alias type to int
2. Coords - a new type defined as a pair of integers (an array of 2 ints)
3. Direction - a new type with an underlying type int (enum)
4. Maze - a generated 2D maze that is a struct
5. Wall - a struct that holds 2 neighboring cells

We see that:
1. There are no constructors in Go (since there are no classes), but we can create functions that serve as constructors.
2. The godoc tool identified our constructor function New and added it under the Maze type.
3. We have structs and they can have fields.
4. We can define a new type out of any underlying type.
5. Any type can have methods (except for primitives).
6. That means that any type satisfies the interface{} - an interface with no methods.
7. You can alias to any type.
8. If you want to add methods to primitives, just define a new type with the desired primitive underlying type.
9. Methods are added to types using Receivers (value or pointer receivers).
10. Methods that can change/mutate the value of the type need a pointer receiver (the common practice says not to mix receiver types).

Let's proceed to examine the maze code, navigate around to see the `travel` package, then the `robot` package and finally the `main` package in `cmd/maze`

That package defines the interface that abstracted away our `robot.Robot` struct into the `Gopher` interface. This ability that Go provides is not common.

The common OOP languages approach is that class A must inherit from class B or implement interface I in order to be used as an instance of B or I,
but our Robot type has no idea that Gopher type even exists. Gopher is defined in a completely different package that is not imported by robot.
Go was written for the 21st century and allows you to plug-in types into your code from anywhere on the internet so long that they have the correct method signatures.

Scripting languages achieve this with duck-typing, but Go is type-safe and we get compile time validation of our code.
Implicit interfaces mean that packages don't have to provide interfaces to the user, the user can define their own interface with the smallest subset of functionality that they need.

In fact our `robot.Robot` has another public method `Steps` that is not part of the `Gopher` interface because we don't need to use it.
This makes plugging-in code and defining and mocking dependencies safely a natural thing in Go and makes the code minimal to its usage.

**In conclusion:** before you write code make sure it's necessary. Be lazy. Be minimal. Be Marie Kondo.

### Object-Oriented Fundamentals and Go
#### Do we need OOP?

>_The problem with object-oriented languages is they've got all this implicit environment that they carry around with them. You wanted a banana but what you got was a gorilla holding the banana and the entire jungle._
(Joe Armstrong)

Just like in the real world, wherever there are things, there can be a mess. *__That's why Marie Kondo is rich__*.
Just as you can write insane procedural code, you can write sane OO code. You and your team should define best practices that match your needs.
This workshop is meant to give you the tools to make better design choices.

#### The C++/Java School of OO
It is important to know that in common OOP languages:
- Objects are instances of a class because only classes can define methods (that's how we support messaging).
- Classes have constructor methods that allow for safe instantiation of objects.
- Classes can inherit methods and fields from other classes as well as override them and sometimes overload them (we will get to that later).
- In case of overriding and overloading methods, the method that will eventually run is decided at runtime. This is called late binding or dynamic binding.

#### Is Go an OO language? == Is t an object?

Go doesn't offer classes, which means there are no constructors (or destructors) and no inheritance, etc.
There is also no late or late late or late late late binding in Go (but there's something else, we'll get to that).
These are technical concepts that have become synonymous with Object-Oriented Programming.
Go does have a variety of very strong features for Object-Oriented Programming that enable Gophers to express their code in a manner that follows the OO principals.
In the end, the answer to the question is Go an OOP language depends on the answer to the question "is t an object" in this [sample code](https://go.dev/play/p/ZfWFad7-TyM)

```go
package main

import "fmt"

type MyThing int //Creates a new type MyThing with an underlying type int

// Foo is now a method of my MyThing, in many languages to have a method you have to have a class or a struct
func (t MyThing) Foo() int {
	return int(t)
}
func main() {
	var t MyThing = 1
	fmt.Println(t.Foo()) // Q: is t an object?
}
```

Whether you think t is an object or not, no gopher is complete without all the tools in the gopher toolbox so let's get (re)acquainted with them.

#### Missing CTORs?
Go doesn't provide us constructors that ensure that users of our types initialize them correctly, but as we saw, we can provide our own ctor function to make our types easy to use.
Developers coming from other language often make types and fields private to ensure that users don't make mistakes.
If your type is not straight-forward, the Go common practices are:
1. Provide a CTOR function.
2. The CTOR should return the type that works with all the methods properly so if the type has methods with pointer receivers it will likely return a pointer.
3. Leave things public, comment clearly that the zero value of the type is not ready to use or should not be copied, etc.
4. Provide default functionality when a required field is zero value.

When you do it right, the godoc will [nest your CTOR function inside your type](https://pkg.go.dev/github.com/ronna-s/go-ood/pkg/robot#pkg-index).

#### Missing Inheritance? - Composition vs. Inheritance
In software, if A inherits from B, it's the code equivalent to saying A is B.
On the other hand, we use composition to express that A is made of B.
In Go, we don't have inheritance, but we have unlimited composition. 
Since we don't have inheritance, to express that A is I we use interfaces. To express that A is made of B or composed of B we use embedding.

The difference between inheritance and composition can be seen [here](https://go.dev/play/p/dkJezhyypeh).

Most OO languages limit inheritance to allow every class to inherit functionality from exactly one other class.
That means that you can't express that an instance of class A is an instance of class B and class C, for example: a truck can't be both a vehicle and also a container of goods.
In the case where you need to express this you will end up doing the same as you would do in Go with interfaces, except as we saw the Go implicit interface implementation is far more powerful.
In addition, common language that offer inheritance often force you to inherit from a common Object class which is why objects can only be class instances (and can't be just values with methods, like in Go).

#### The Alan Kay School of OO
Alan Kay is considered to the person who coined the term Object-Oriented
#### Missing Messaging?

Using interface type assertion (or conversion), we can check at runtime if a type has a method (or a set of method) of the exact signature we would like to invoke, and call it.
The result of this check is cached for performance.

In addition, the receiver (used to define methods) is inspired by Oberon-2 which is an OO version of Oberon.

#### Conclusion: is Go OO?
##### The Go FAQ - is Go an object-oriented language?
>_Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes). <br>Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java._

[Source](https://go.dev/doc/faq#Is_Go_an_object-oriented_language)

##### Rob Pike - Go is Object-Oriented
>_Go is object-oriented, even though it doesn't have the notion of a class. The type system is more general. Any type—even basic types such as integers and strings—can have methods. This allows inheritance and other object-oriented techniques to apply more broadly than with classes alone. For instance, Go's formatted printing library, package fmt, uses interfaces and methods to provide a way to use a printf-like API to print any value, ranging from basic types to arbitrary user-defined objects, with perfect type safety._

[Source](https://www.informit.com/articles/article.aspx?p=1623555)



<hr>

## Go OO in Practice
### Exercise 2 - Interfaces and Embedding
We are going to add 2 types of players to the game P&P - Platforms and Programmers who will attempt to take on a Production environment.
The roles that we will implement are `pnpdev.Gopher`, `pnpdev.Rubyist`.
The player roles are going to be composed of the struct `pnpdev.Character` for common traits like XP and Health.
Gopher and Rubyist will also need to implement their own methods for their individual `Skills` and `AsciiArt`.

Run the game with the minion player:

```bash
# go tool
go run cmd/pnp/pnp.go
# make + docker
make run-pnp
# any other setup with docker 
[docker command from before] go run github.com/ronna-s/go-ood/cmd/pnp.go
```

```go
// Player represents a P&P player
type Player interface {
	Alive() bool
	Health() int
	XP() int
	ApplyXPDiff(int) int
	ApplyHealthDiff(int) int
	Skills() []Skill
	Art() string 
}
```

We already have a type Minion in package `pkg/pnpdev` with some implementations for a player.
1. We will extract the methods `Alive`, `ApplyXPDiff`, `ApplyHealthDiff`, `Health` and `XP` to a common type `Character`.
2. We will embed `Character` inside `Minion`.
3. Create new types `Gopher` and `Rubyist` that implement the `pnp.Player` interface, use the failing tests to do this purposefully, see how to run the tests below.
4. Add `NewGopher()` and `NewRubyist()` in cmd/pnp/pnp.go to our list of players.
5. Run the game.
6. We notice that the Gopher and the Rubyist's names are not properly serialized... We will fix that in a moment.

To test our players:
```bash
# make + docker
make test-pnp
# go tool
go test github.com/ronna-s/go-ood/pkg/pnpdev
# any other setup with docker
[docker command from before] go test github.com/ronna-s/go-ood/pkg/pnpdev
```

### Stringers
We are not done with exercise 2.
As we saw our Rubyist and Gopher's name were not displayed properly.
We fix this by adding the `String() string` method to them:

```go
func (r Rubyist) String() string {
	return "Rubyist"
}
func (g Gopher) String() string {
	return "Gopher"
}
```
We run the game and see that it works as expected but what actually happened here? - String() is not part of the `Player` interface?
We can check if a type implements an interface at runtime:

```go
https://go.dev/play/p/6Ia8aGJS7Bc
package main

import "fmt"

type fooer interface {
	Foo() string
}

type A struct{}

func (_ A) Foo() string {
	return "Hello from A"
}

func main() {
	var a interface{} = A{}
	var i interface{} = 5
	if v, ok := a.(fooer); ok {
		fmt.Println(v.Foo())
	} else {
		panic("should not be called")
	}
	if v, ok := i.(fooer); ok {
		panic("should not be called")
	} else {
		fmt.Println("v is nil:", v)
	}
}
```
Go's print function checked at runtime if our types have the method `String() string` by checking if it implements an interface with this method and then invoked it.

Russ Cox compared this to duck typing and explained how it works [here](https://research.swtch.com/interfaces).

It's particularly interesting that this information about what types implement what interfaces is cached at runtime to maintain performance. Even though we achieved this behavior without actual receivers that take in messages and check if they can handle them, from design perspective we achieved a similar goal.

This feature only makes sense when interfaces are implicit because in languages when the interface is explicit there's no way a type can suddenly implement a private interface that is used in our code.

#### Effective interface type assertion
If you are going to add your own type assertions, remember that the code execution becomes unpredictable, you therefore should comply with a couple of rules:
1. The user of your code might not know what interfaces they are expected to implement or might provide them but cause a panic. Provide default behavior and __in addition__ use `defer` and `recover` to prevent crashing the app or return errors if the interface allows it.
2. If your type is expected to implement an interface, to protect against changes add a line to your code that will fail to compile if your type doesn't implement the interface, like so:

```go 
// In the global scope directly
var _ interface{ String() string } = NewGopher()
var _ interface{ String() string } = NewRubyist()
```

### Organizing your packages
Whether you choose the common structures with cmd, pkg, etc. you should probably define some guidelines for your team. Here are a few suggestions:
1. Support multiple binaries: Your packages structure should allow compiling multiple binaries (have multiple main packages that should be easy to find).
2. An inner package is usually expected to extend the functionality of the upper package and import it (not the other way around), for example:
    - `net/http`
    - `image/draw`
    - and the example in this repo `maze/travel`
3. There are some exceptions to this for instance `image/color` is a dependency for `image`, but it's not the rule. In an application it's very easy to have cyclic imports this way.
4. A package does not provide interfaces except for those it uses for its dependencies.
5. Use godoc to see what your package looks like without the code. It helps.
6. Keep your packages' hierarchy relatively flat. Just like your functions, imports don't do spaghetti well.
7. Try to adhere to open/close principals.
8. Your packages should describe tangible things that have clear boundaries - domain, app, utils, aren't those things.
9. Package path with `internal` cannot be imported. It's for code that you don't want to allow to import, not for your entire application. It's especially useful for anyone using your APIs to be able to import your models for instance.

### Code generation, why? When?
I like this simple explanation by (Gabriele Tomassetti)[https://tomassetti.me/code-generation/]
> The reasons to use code generation are fundamentally four:
> - productivity;
> - simplification;
> - portability;
> - consistency

It's about automating a process of writing repetitive error-prone code.
Code generation is similar to meta-programming but we compile it and that makes it safer to run.
Consider the simple [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
Consider [Mockery](http://github.com/vektra/mockery)
Both were used to generate code for this workshop.

Also, as a favour to me, please commit your generated code. A codebase is expected to be complete and runnable.

### Emerging patterns
#### Complex CTORs with no function overloading
[Functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

#### Default variables 
Default variables are very useful Singleton values.
Examples: [net/http](https://pkg.go.dev/net/http).

```go
// https://go.dev/play/p/8hiAeuJ90uz
package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	http.ErrBodyNotAllowed = errors.New("my error")
	fmt.Println(http.ErrBodyNotAllowed)
}
```

### Short Lived Objects vs. Long Lived Objects
[Consider this conversation](https://twitter.com/francesc/status/1293556263196844032)

#### Context keys as types
[Article](https://medium.com/@matryer/context-keys-in-go-5312346a868d)

<hr>

## Generics
Generics are not specifically related to Object-Oriented, but generics do enable more flexible work with types, which benefits OO.   

### Introduction to Generics
It was a long time consensus that "real gophers" don't need generics, so much so that around the time the generics draft of 2020 was released, many gophers still expressed that they are not likely to use them.

Let's understand first the point that they were trying to make.

Consider [this code](https://gist.github.com/Xaymar/7c82ed127c8f1def53075f414a7df153), made using C++.
We see here generic code (templates) that allows an event to add functions (listeners) to its subscribers.
Let's ignore for a second that this code adds functions, not objects and let's assume it did take in objects with the function `Handle(e Event)`.
We don't need generics in Go to make this work because interfaces are implicit. As we saw already in C++ an object has to be aware of it's implementations, this is why to allow plugging-in of functionality we have to use generics in C++ (and in Java).

In Go this code would look something like [this](https://go.dev/play/p/Tqm_Hb0vcZb):

```go
package main

import "fmt"

type Listener interface {
	Handle(Event)
}

type Event struct {
	Lis []Listener
}

func (e *Event) Add(l Listener) {
	e.Lis = append(e.Lis, l)
}

func main() {
	var l Listener
	var e Event
	e.Add(l)
	fmt.Println(e)
}
```

**We didn't need generics at all!**

However, there are cases in Go where we have to use generics and until recently we used code generation for.
Those cases are when the behavior is derived from the type or leaks to the type's behavior:

For example:
The linked list
```go
// https://go.dev/play/p/ZpAqvVFAIDZ
package main

import "fmt"

type Node[T any] struct { // any is builtin for interface{}
  Value T
  Next  *Node[T]
}

func main() {
  n1 := Node[int]{1, nil}
  n2 := Node[int]{3, &n1}
  fmt.Println(n2.Value, n2.Next.Value)
}
```
Example 2 - [Addition](https://go.dev/play/p/dmeQEVxpyAq)
```go
package main

import "fmt"

type A int

// Add takes any type with underlying type int 
func Add[T ~int](i T, j T) T { 
  return i + j
}

func main() {
  var i, j A
  fmt.Println(Add(i, j))
}
```
Of course, you might not be likely to use linked lists in your day to day, but you are likely to use:
1. Repositories, databases, data structures that are type specific, etc.
2. Event handlers and processors that are specific to a type.
3. The [concurrent map in the sync package](https://pkg.go.dev/sync#Map) which uses the empty interface.
4. [The heap](https://pkg.go.dev/container/heap#example-package-IntHeap)

The common thread to these examples is that before generics we had to trade generalizing certain behavior for type safety (or generate code to do so), now we can have both.

## Exercise 3 - Generics
Implement a new generic Heap OOP style in `pkg/heap` (as usual failing tests provided).
The heap is used by TOP OF THE POP! `cmd/top` to print the top 10 Artists and Songs.

Test:
```bash
# go tool
go test github.com/ronna-s/go-ood/pkg/heap
# make + docker
make test-heap
# any other setup with docker 
[docker command from before] go test github.com/ronna-s/go-ood/pkg/heap
````

Run our TOP OF THE POP app:
```bash
# go tool
go run cmd/top/top.go
# make + docker
make run-heap
# any other setup with docker 
[docker command from before] go run cmd/top/top.go
````

## Conclusion
What we've learned today:
1. The value of OOP
2. Defining types that fit our needs
3. Writing methods
4. Value receivers and pointer receivers
5. Organizing packages
6. Interfaces
7. Composition
8. Generics
9. To generate code otherwise

