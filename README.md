# langspec
 A modern language.
 
 ## Motivation
 
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
const SomeType someName = "someName"    // explicitly-typed
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

type rune uint32 // rune type definition => uint32 alias => represents unicode codepoint
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
```swift
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


# Reference
- [C#-lang](https://github.com/dotnet/csharplang)
- [Go-lang](https://github.com/golang/go)
- [Typescript](https://github.com/Microsoft/TypeScript)
- [Pony-lang](https://github.com/ponylang/ponyc)
- [Ghost-lang](https://github.com/jamiebuilds/ghost-lang)
