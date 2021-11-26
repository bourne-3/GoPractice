### part1
- Go does not have classes. However, you can define methods on types.
- The receiver appears in its own argument list between the func keyword and the method name.
  - 注意这里是between
  - Remember: a method is just a function with a receiver argument.
  - You can only declare a method with a receiver whose type is defined in the same package as the method.
- pointer receiver （指针receiver）
  - The first is so that the method can modify the value that its receiver points to.
  - The second is to avoid copying the value on each method call. 
- 