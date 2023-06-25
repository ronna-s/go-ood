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

TBD
## Introduction to Object-Oriented Programming
### What is OOP?
### Exercise 1 - Understanding the benefits
Help a gopher get out of a maze without knowing anything about how the maze is implemented (or the gopher).
### Basic Go for OO
#### Type definition
#### Adding Methods
#### Aliases
#### Structs
#### Embedding & Promotion
#### Overriding
#### Interfaces
#### Interface type assertion
### Exercise 1 - how it all worked
Now that we have the basics we can go back and understand the code.
### Object-Oriented fundamentals and Go
#### Do we need OOP?
#### Is Go an OO language? == Is t an object?
#### The C++/Java School
#### Missing CTORs?
#### Missing Inheritance?
#### Missing Inheritance? and, Composition vs. Inheritance
#### The Alan Kay School
#### Missing Messaging?
#### the Go FAQ - is Go an object-oriented language?
#### Rob pike - Go is Object-Oriented
#### Conclusion: Is t an object?

## Go OO in Practice
### Exercise 2 - Interfaces and Embedding
### Stringers
#### Effective interface type assertion
### The empty interface (any)
### Organizing your packages
### Code generation, why? When?
### More Theory
### Emerging patterns
#### Complex CTORs with no function overloading
#### Default variables, exported variables, overrideable and otherwise
#### Short Lived Objects vs. Long Lived Objects
#### Context keys as types

## Generics
### Introduction to Generics
### Exercise 3 - Generics
## Conclusion
