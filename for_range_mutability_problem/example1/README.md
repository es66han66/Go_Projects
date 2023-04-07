What we actually see is the result of mutability of the loop variable:

In every iteration, we get a struct instance to work with. Structs are value types — they are copied to the for iteration variable in each iteration. The key word here is copy. To avoid large memory print, instead of creating a new instance of the variable in each iteration, a single instance is created at the beginning of the loop, and in each iteration, the data is copied on it.

Closures are the other part of the equation: Closures in Go (like in most languages) hold a reference of the objects in the closure (not copying the data), so the inner go routine takes a reference of the iterated object, meaning the all the go routines get the same reference to the same instance.