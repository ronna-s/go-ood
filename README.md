# A path to OOD with Go - Workshop


https://github.com/ronna-s/go-ood/ [Clone me!]

This workshop is aimed to clarify the OOP features that Go provides.
It is named A Path to OOD and not OOP because different language features mean different design concepts.  

## Schedule

- 09:00-09:20: Introduction to Object-Oriented Programming [link](#introduction-to-object-oriented-programming)
- 09:20-09:50: Exercise 1 - understanding the benefits [link](#exercise-1---understanding-the-benefits)
- 09:50-10:00: Break
- 10:00-10:10: Object Oriented Fundamentals and Go [link](#oo-fundamentals-and-go)
- 10:20-10:50: Exercise 2 - Interfaces and Embedding [link](#exercise-2---interfaces-and-embedding)
- 10:50-11:00: Break
- 11:00-11:10: Organizing your Packages [link](#organizing-your-packages)
- 11:10-11:20: Code Generation [link](#code-generation-why-when)
- 11:20-11:30: More Theory [link](#generics)
- 11:30-11:50: Generics [link](#generics)
- 11:50-12:00: Break
- 12:00-12:50: Exercise 3 - Generics [link](#exercise-3---generics)
- 12:50-13:00: Conclusion

## Introduction to Object-Oriented Programming

### What is OOP?

What we can all agree on: The central idea behind Object-Oriented is to divide software into "things" or "objects" or "instances" that communicate via "messages" or "methods" or "member functions".
Or in short, combining data and functionality.
This core idea has not changed in the 4-5+ decades since it was conceptualized.
It is meant to allow the developer to build code and separate responsibilities or concerns just like in the real world which is what we are familiar with and how we generally think and solve problems.

It is important to know that in most OOP languages: 
- Objects are instances of a class because only classes can define methods (that's how we support messaging).
- Classes have constructor methods that allow for safe instantiation of objects.
- Classes can inherit methods and fields from other classes as well as override them and sometimes overload them (we will get to that later).
- In case of overriding and overloading methods, the method that will eventually run is decided at runtime. This is called late binding or dynamic binding.

### Is Go an Object-Oriented language?

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
 
### Do you need OOP?
Just like in the real world, wherever there are things, there can be a mess. *__That's why Marie Kondo.__*
Just as you can write insane procedural code, you can write sane OO code. You and your team should define best practices that match your needs.
This workshop is meant to give you the tools to make better design choices. 

## Exercise 1 - Understanding the benefits:

Where we will understand some OO basics using an example of a gopher and a maze.

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

All exercises can be completed using the go tool, docker or a combination docker and the make tool:

If planning to use docker and you don't have the make tool, run:
```bash
docker build . -t go-ood
```

If you have the make tool and docker run:
```bash
make build
```

Run the tests:
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

Let's review the code that made this possible and examine the Go features it uses.

Run:
```bash
make godoc
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
3. We have structs and they can  have fields.
4. You can define a new type out of any underlying type
5. Any type can have methods (except for primitives)
6. That means that any type satisfies the interface{} - an interface with no methods
7. You can alias to any type
8. If you want to add methods to primitives, just define a new type with the desired primitive underlying type
9. Methods are added to types using Receivers (value or pointer receivers).
10. Methods that can change/mutate the value of the type need a pointer receiver (the common practice says not to mix receiver types) 

Speaking of "Receivers", Remember that we said that OO is about objects communicated via messages?
The idea for the receiver was borrowed from Oberon-2 which is an OO version of Oberon.
But the receiver is also just a special function parameter, so there is no receiver.

![https://giphy.com/gifs/the-matrix-there-is-no-
-3o6Zt0hNCfak3QCqsw](https://gifimage.net/wp-content/uploads/2018/06/there-is-no-spoon-gif-10.gif)

Navigate around to see the travel package, then the robot package and finally the main package in `cmd/maze`

That package defines the interface that abstracted away our `robot.Robot` struct into the `Gopher` interface. This is not common.

The common OOP languages approach is that class A must inherit from class B or implement interface I in order to be used as an instance of B or I,
but our Robot type has no idea that Gopher type even exists. Gopher is defined in a completely different package that is not imported by robot.
Go was written for the 21st century and allows you to plug-in types into your code from anywhere on the internet so long that they have the correct method signatures. 
This is done in scripting languages with duck-typing, but in Go it's just type-safe, and you get compile time validation of your code.
Implicit interfaces mean that packages don't have to provide interfaces to the user, the user can define their own interface with the smallest subset of functionality that they need.
In fact our `robot.Robot` has another public method `Steps` that is not part of the `Gopher` interface because we don't need to use it.
This makes plugging-in code and defining and mocking dependencies safely a natural thing in Go and makes the code minimal to its usage.  

>_The problem with object-oriented languages is they've got all this implicit environment that they carry around with them. You wanted a banana but what you got was a gorilla holding the banana and the entire jungle._
(Joe Armstrong)

What did he mean by that?

He likely meant that OO is overcomplicated but in reality those rules that we discussed that apply to common OOP languages cause this complication:

The class Banana will have to extend or inherit from Fruit (or a similar Object class) to be considered a fruit, implement a Holdable interface just in case we ever want it to be held, implement a GrowsOnTree just in case we need to know where it came from. etc.
What happens if the Banana we imported doesn't implement an interface that we need it to like holdable? We have to write a new implementation of Banana that wraps the original Banana.

Now would be a good time to discuss inheritance:
Most OO languages limit inheritance to allow every class to inherit functionality from exactly one other class.
That means that you can't express that an instance of class A is an instance of class B and class C, for example: a truck can't be both a vehicle and also a container of goods.
In the case where you need to express this you will end up doing the same as you would do in Go with interfaces, except as we saw the Go implicit interface implementation is far more powerful.
In addition, common language that offer inheritance often force you to inherit from a common Object class which is why objects can only be class instances (and can't be just values with methods, like in Go).

From the Go FAQ - is Go an object-oriented language?
>_Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes). <br>Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java._

## OO fundamentals and Go

### So no CTORs, huh?

Go doesn't provide us constructors that ensure that users of our types initialize them correctly, but as we saw, we can provide our own ctor function to make our types easy to use.
Developers coming from other language often make types and fields private to ensure that users don't make mistakes.
If your type is not straight-forward, the Go common practices are:
1. Provide a CTOR function.
2. The CTOR should return the type that works with all the methods properly so if the type has methods with pointer receivers it will likely return a pointer.
3. Leave things public, comment clearly that the zero value of the type is not ready to use or should not be copied, etc.  
4. Provide default functionality when a required field is zero value.

### Composition vs. Inheritance

In Go we don't have inheritance. To express that A is I we use interfaces. To express that A is made of B or composed of B we use embedding like so:

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

We see that we can embed both interfaces and structs.

## Exercise 2 - Interfaces and Embedding

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
3. Create new types `Gopher` and `Rubyist` that implement the `pnp.Player` interface, use the failing tests to do this purposefully, see how to run the tests below:
4. Add `NewGopher()` and `NewRubyist()` in cmd/pnp/pnp.go to our list of players.
5. Run the game.
6. We will notice that the Gopher and the Rubyist's names are not properly serialized...

To test our players:
```bash
# make + docker
make test-pnp
# go tool
go test github.com/ronna-s/go-ood/pkg/pnpdev
# any other setup with docker
[docker command from before] go test github.com/ronna-s/go-ood/pkg/pnpdev
```

## Organizing your packages

Whether you choose the common structures with cmd, pkg, etc. you should try to follow certain guidelines:
1. Support multiple binaries: Your packages structure should allow compiling multiple binaries (have multiple main packages that should be easy to find).
2. Don't try to reduce the number of your imports: If you have a problem it's probably the structure and unclear responsibilities, not the amount.
3. An inner package is usually expected to extend the functionality of the upper package and import it (not the other way around), for example:
   - `net/http`
   - `image/draw`
   - and the example in this repo `maze/travel`
4. There are some exceptions to this for instance `image/color` is a dependency for `image`, but it's not the rule. In an application it's very easy to have cyclic imports this way.   
5. We already said this, but just to be clear: A package does not provide interfaces except for those it uses for its dependencies.
6. Use godoc to see what your package looks like without the code. It helps. 
7. Keep your packages' hierarchy relatively flat. Just like your functions, imports don't do spaghetti well. 
8. Try to adhere to open/close principals to reduce the number of changes in your code. It's a good sign if you add functionality but not change everything with every feature.
9. Your packages should describe tangible things that have clear boundaries - domain, app, utils, aren't things.
10. Package path with `internal` cannot be imported. It's for code that you don't want to allow to import, not for your entire application. It's especially useful for anyone using your APIs to be able to import your models for instance. 

## Code generation, why? When?
I like this simple explanation by (Gabriele Tomassetti)[https://tomassetti.me/code-generation/]
> The reasons to use code generation are fundamentally four: 
> - productivity;
> - simplification;
> - portability;
> - consistency

It's about automating a process of writing repetitive error-prone code.
Consider the simple [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
Consider [Mockery](http://github.com/vektra/mockery)

Both were used to generate code for this workshop.

## More Theory

### Emerging patterns:
1. Constructing complex objects with no constructors and overloading [Functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
2. Default variables, exported variables, overrideable and otherwise [net/http](https://pkg.go.dev/net/http) also in this repo - the `pnp.Rand` function

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

## Generics
It was a long time consensus that "real gophers" don't need generics so much so that around the time the generics draft of 2020 was released, many gophers expressed that they will likely never use this feature.
Let's understand first the point that they were trying to make.

Let's look at [this code](https://gist.github.com/Xaymar/7c82ed127c8f1def53075f414a7df153), made using C++.
We see here generic code (templates) that allows an event to add functions (listeners) to its subscribers.
Let's ignore for a second that this code adds functions, not objects and let's assume it did take in objects with the function `Handle(e Event)`. 
We don't need generics in Go to make this work because interfaces are implicit. As we saw already in C++ and object has to be aware of it's implementations, this is why to allow plugging-in of functionality we have to use generics in C++ (and in Java).

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
Those cases are when the behavior is derived from the type or leaks to the type:

For example:
The linked list
```go
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
Of course, you are not likely to use linked lists in your day to day, but you are likely to use:
1. Repositories, databases, data structures that are type specific
2. Event handlers and processors that are type specific
3. The [concurrent map in the sync package](https://pkg.go.dev/sync#Map) which uses the empty interface.
4. [The heap](https://pkg.go.dev/container/heap#example-package-IntHeap) 

The common thread to these examples is that before generics we had to trade certain functionality for type safety.

## Exercise 3 - Generics
Implement a new generic slice Heap OOP style in `pkg/heap` (as usual failing tests provided).
The heap is used by TOP OF THE POP! `cmd/top` to print the top 10 Artists and Songs 

Uncomment the following line in the Dockerfile:

As usual:
```bash
# go tool
go test github.com/ronna-s/go-oog/pkg/heap
# make + docker
make test-heap
# any other setup with docker 
[docker command from before] test github.com/ronna-s/go-ood/pkg/heap
````

## Conclusion
What we've learned today:
1. The value of OOP
2. Defining types that fit our needs
3. Writing methods
4. Interfaces
5. Composition
6. Generics
7. To generate code otherwise
