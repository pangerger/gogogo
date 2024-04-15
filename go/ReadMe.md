# Basic

## Packages


Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.


```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}


```

## Imports


Imports
This code groups the imports into a parenthesized, "factored" import statement.

You can also write multiple import statements, like:

import "fmt"
import "math"
But it is good style to use the factored import statement.


```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

```


## Exported Names


In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Run the code. Notice the error message.

To fix the error, rename math.pi to math.Pi and try it again.



## Constants

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.



## for loop

Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.


## IF

Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.



## Map


```
elem, ok = m[key]
```

If `key` is in `m`, `ok` is true. If not, `ok` is false


Note: If key or ok have not yet been declared you could use a short declaration form:

`elem, ok := m[key]`


## Pointer

value of copy

pointer


## Methods

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package(which includes the built-in types such as int)



### methods and pointer

```
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```

Functions with a pointer argument must take a pointer
```
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

While mehtods with pointer receivers take either a value or a pointer as the recevier when they are called:
```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is , as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver. 


The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

```
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

while methods with value receivers take either a value or a pointer as the receiver when they are called

```
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

in this case, the method call p.Abs() is interpreted as (*p).Abs().

### Choosing a value or pointer receiver

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call.This can be more efficient if the receiver is a large struct, even though the method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.


## Interfaces

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

### Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

### Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
`(value, type)`

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

### Interface values with nil underlying values

If the concrete value inside the interface itself is nil, the method will be called with a nil receiverl

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver.

Note that an interface value that holds a nil concrete value is itself non-nil.

### Nil interface values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error bacasuse there is no type inside the interface tuple to indicate which concrete method to call.

### The empty interface

The interface type that specifies zero methods is known as the empty interface:
`interface {}`

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}

### Type assertions

A type assertion provides access to an interface value's underlying concrete value.

`t := i.(T)`

This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

If not, ok will false and t will be the zero value of the type T, and no pani occurs.

Note the similarity between this syntax and that of reading from a map.

### Type switches

A type switch is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```
switch v := i.(type) {
	case T:
		// here v has type T
	case S:
		// here v has type S
	default:		
}
```

The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.

## Stringers

One of the most ubiquitous interfaces is Stringer defined by the fmt package.

```
type Stringer interface {
	String() string
}
```

A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

## Error

Go programs express error state with error values

The error type is a built-in interface similar to fmt.Stringer

```
type error interface {
	Error() string
}
```

As with fmt.Stringer, the fmt package looks for the error interface when printing values.

Function often return an error value, and calling code should handle errors by testing whether the error equals nil.

```
i, err := strconv.Atoi("42")
if err != nil {
	fmt.Printf("couldn't convert number: %v\n", err)
	return
}
fmt.Println("Converted interger: ", i)
```

A nil error denotes success; a non-nil error denotes failure.



# others

```go

%T

%v

%q

```

## 赋值

:=

=
















# Reference


https://go.dev/tour/list


https://go.dev/blog/slices-intro
















