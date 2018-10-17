# langspec

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
