## Schedule

- 13:00-13:10: Introduction to Object-Oriented Programming [link](#introduction-to-object-oriented-programming)
- 13:10-13:40: Exercise 1 - understanding the benefits [link](#exercise-1---understanding-the-benefits)
- 13:40-13:50: Exercise 1 - How does it all work? [link](#exercise-1---how-does-it-all-work)
- 13:50-14:00: Break
- 14:00-14:10: Object Oriented Fundamentals and Go [link](#oo-fundamentals-and-go)
- 14:10-14:50: Exercise 2 - Interfaces and Embedding [link](#exercise-2---interfaces-and-embedding)
- 14:50-15:00: Break
- 15:00-15:10: Organizing your Packages [link](#organizing-your-packages)
- 15:10-15:20: Code Generation [link](#code-generation-why-when)
- 15:20-15:30: More Theory [link](#generics)
- 15:30-15:50: Generics [link](#generics)
- 15:50-16:00: Break
- 16:00-16:50: Exercise 3 - Generics [link](#exercise-3---generics)
- 16:50-17:00: Conclusion

## Introduction to Object-Oriented Programming

### What is OOP?


### Is Go an Object-Oriented language?

   
 
### Do you need OOP?

## Exercise 1 - Understanding the benefits:


### Exercise 1 - how does it all work?
[sort of there]
Speaking of "Receivers", Remember that we said that OO is about objects communicated via messages?
The idea for the receiver was borrowed from Oberon-2 which is an OO version of Oberon.
But the receiver is also just a special function parameter, so **"there is no spoon"** (receiver) but from a design perspective there is.

![https://giphy.com/gifs/the-matrix-there-is-no-
-3o6Zt0hNCfak3QCqsw](https://gifimage.net/wp-content/uploads/2018/06/there-is-no-spoon-gif-10.gif)


How do we know that there's no actual receiver? [Run this code](https://go.dev/play/p/iOx0L_p65jz)

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


### Missing Inheritance?

What did he mean by that?

He likely meant that OO is overcomplicated but in reality those rules that we discussed that apply to common OOP languages cause this complication:

The class Banana will have to extend or inherit from Fruit (or a similar Object class) to be considered a fruit, implement a Holdable interface just in case we ever want it to be held, implement a GrowsOnTree just in case we need to know where it came from. etc.
What happens if the Banana we imported doesn't implement an interface that we need it to like holdable? We have to write a new implementation of Banana that wraps the original Banana.


### Finally, from the Go FAQ - is Go an object-oriented language?

## OO fundamentals and Go

### So no CTORs, huh?

### Composition vs. Inheritance



We see that we can embed both interfaces and structs.



### The empty interface{} (any):
- Since all types can have methods, all types implement the empty interface (`interface {}`) which has no methods.
- The empty interface has a built-in alias `any`. So you can now use `any` as a shorthand for `interface{}`

## Organizing your packages

## Code generation, why? When?

## More Theory

## Generics

## Exercise 3 - Generics
## Conclusion

