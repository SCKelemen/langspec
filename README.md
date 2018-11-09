# langspec
A family of languages: languages should feel familiar from baremetal to cloud.
 
 ## Motivation
 
 
 ## Architecture
 Languages:
  - [Ash](https://github.com/sckelemen/ash)
  - [Ember](https://github.com/sckelemen/ember)
  - [Blaze](https://github.com/sckelemen/blaze)
  
 Tools:
  - [Pyro](https://github.com/sckelemen/pyro)
  
  ### Ash
  An abstraction over assembly. (asm)
  
  ### Ember
  An embedded systems language. (c++)
  
  ### Blaze
  A full systems language. (c#)
  
  ### Pyro
  A unified build system.
  
 
 ## Syntax
 
 ### Comments
 ```C#
 // line comment
 
 /*
    Multiline comment
 */
 
/* inline comment */  
 ```
 
 ### Package / Import
 ```Go
 package SomePackage        // the name of the package
 
 import "AnotherPackage"    // a single-package import 
 
 import (                   // a multipackage import
    "ThirdPackage",
    "FourthPackage",
    "FifthPackage" 
 )
 ```
If subpackage imports are required, the syntax would be:
```typescript
import { "sub1", "sub2" } from "SixthPackage"    // subpackage imports
```

### Variables
```swift
let name = expression    // type-inferred
SomeType instance = GetSomeType() // statically typed as SomeType
```

### Constants 
```Go
const name = "someName"    // type-inferred
const string someName = "someName"    // explicitly-typed
```

### Booleans
```coffee 
bool isTrue = true
let isNotTrue = false
```

### Numbers
```Go
int   num = 123
int8  num = 123
int16 num = 123
int32 num = 123
int64 num = 123

uint   num = -123
uint8  num = -123
uint16 num = -123
uint32 num = -123
uint64 num = -123

float32 num = 1.23
float64 num = 1.23

complex64  num = 6 + 7i
complex128 num = 6 + 7i
```

### Byte
```Go
byte A = 0x41    // hexadecimal number 
byte a = 97      // decimal number

type byte uint8  // byte type definition => uint8 alias => represents ascii codepoint
```

### Characters
```Go
rune A = 0x41
rune a = 97 

type rune int32 // rune type definition => int32 alias => represents unicode codepoint
```

### Strings
```
```

### Arithmetic Operators
```coffee
let addition       = a + b
let subtraction    = a - b
let division       = a / b
let multiplication = a * b
let modulo         = a % b

let exponent       = a ** b
let negation       = -a 
```

### Logical Operators
```coffee 
let and = expressionA && expressionB
let or  = expressionA || expressionB
let not = !expression
```

### Bitwise Operators
```coffee
let and = a & b
let or  = a | b
let xor = a ^ b
let not = ~a

let lshift = a << b
let rshift = a >> b
```

### Comparison
```swift
let equality = a == b
let notequal = a != b
```

### Conditionals
```Go
// standard if block
if (condition) {
  doSomething()
}

// dangling if
if (condition) 
  doSomething() 

// inline if block
if (condition) { doSomething() }


// standard if-else block
if (condition) {
  doSomething()
} else {
  doSomethingElse()
}


switch (condition) {
  case 1:
    doCase1()
    break
  case 2:
    doCase2()
    break
  default:
    doDefault()
    break
}

```

### Functions
```C#
string GetEmptyString() {
  return ""
}
```

### Multi-Returns
```Go
(bool, item) TryGetItem(id: string) {
  return true, item{id: id}
}
```
Named returns:
```Go
(bool itemExists, item item) TryGetItem(id: string) {
 itemExists = true
 item = item{id: id}
 return 
}
```

### Blank Identifier
ignore extra values:
```Go
_, item = TryGetItem(id: "1")
```

### Type Unions
```Typescript
type uint ( uint8 | uint16 | uint32 | uint64 )

type int ( int8 | int16 | int32 | int64 )

type byte uint8
type rune int32
```

### Nullable Types
```Go
type NullableBool bool?
type NullableBool ( bool | null )
```

### Generics 
```C#
public class Map<TKey, TValue> {

    public TKey[] GetKeys() {
       return new TKey[]{};
    }

    public TValue[] GetValues() {
       return new TValue[]{}
    }
}
```
### Collections
```C# 
let array = [1, 2, 3, 4] // arrays store homogeneous, static data. type can be inferred for value-types
let collection = { 1, 2, 3, 4 } // collection are dynamic. type is inferred
let typedCollection = SomeClass[]{ {}, {}, {} }
let demoTypedCollec = string[]{ "one", "two", "three" }
let associateArray = Hash<TKey, TVal>{ {}:{}, {}:{}, {}:{} }
let demoAsscArr = Hash<string, int>{ "one":1, "two":2, "three":3 }
```

### Composition 
I'm not really sure how this will work with multiple returns.
Haskell has an idea of function composition using the . (dot) operator. F# has a pipe-forward operator (|>). Rust has pipelines (->). I think it would be interesting to consider what could happen with iterators where the pipeline yields to the next operation immediately instead of waiting for the first composed operation to complete, but I think that would have some ramifications with concurrency.

```F#
let F x y z = x (y z)
let F x y z = y z |> x    // using forward pipe
let F x y z = x <| y z    // using backward pipe
```

Proposed: 
```F#
[1..100] 
|> Seq.map (function
    | x when x%5=0 && x%3=0 -> "FizzBuzz"
    | x when x%3=0 -> "Fizz"
    | x when x%5=0 -> "Buzz"
    | x -> string x)
|> Seq.iter (printfn "%s")
```
This uses value-matching, though.

```Go
// iterable vs enumberable
//func Sum( collection: Iterable<int> ) {
// var sum = 0
// yield sum += value
//} 

//func Count( collection: Iterable<int> ) {
//  var count = 0
//  yield count++
//}

//func Average = 
//  |> collection: Iterable<int>
//  |> Sum(collection)
//  |> Su
 
 
 func Square(value: int) {
    return value * value
    // == return value ** 2
 }
 
 func Triple(value: int) {
   return value * 3
 }
 
 // works for iterators
 // I guess pipelining single value functions over a iterable is equivelent to a map
 func TripleSquare = 
 |> collection: Iterable<int>
 |> Square(collection.Iterate())
 |> Triple(^) // get value from previous step
 |> yield ^ // return iterator of previous value
 
 
 
```

``` 

### Example

```
package main

import "some/package"

type Unsigned ( uint8 | uint16 | uint32 | uint64 )

class SomeClass {
    public SomeClass() {
    
    }
    
    fn ( string, error ) ToString() {
      return "somestring"
    }
}

public OtherClass<T> where T : class
{
  private T _item 
  public OtherClass(T item) {
    _item = item
  }
  
  public T Item => _item
  
}
```

# Theory 

## Function Composition and Associativity
- [F#](https://fsharpforfunandprofit.com/posts/function-composition/)
- [Rust](https://docs.rs/pipelines/0.4.0/pipelines/)

## Collections
- Set Thoery

## Graphs
- Graph Theory

# Resources

## Collections

## Grammar / Syntax

## Typesytems
- [Dependent Types](https://en.wikipedia.org/wiki/Dependent_type)
- [Pure Type Systems](https://en.wikipedia.org/wiki/Pure_type_system)
- [Typed Lambda Calculus](https://en.wikipedia.org/wiki/Typed_lambda_calculus)
- [Intuitionistic Logic](https://en.wikipedia.org/wiki/Intuitionistic_logic)
- [Simply Typed Lambda Calculus](https://en.wikipedia.org/wiki/Simply_typed_lambda_calculus)

# Reference
- [C#-lang](https://github.com/dotnet/csharplang)
- [Go-lang](https://github.com/golang/go)
- [Typescript](https://github.com/Microsoft/TypeScript)
- [Pony-lang](https://github.com/ponylang/ponyc)
- [Ghost-lang](https://github.com/jamiebuilds/ghost-lang)
