### part1
- Go does not have classes. However, you can define methods on types.
- The receiver appears in its own argument list between the func keyword and the method name.
  - 注意这里是between
  - Remember: a method is just a function with a receiver argument.
  - You can only declare a method with a receiver whose type is defined in the same package as the method.
- pointer receiver （指针receiver） （下面是为什么要使用）
  - The first is so that the method can modify the value that its receiver points to.
  - The second is to avoid copying the value on each method call. 
- 是用方法的时候会自动适应参数（方法调用）
  - As a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver

### part2
- 在内部，接口值可以看做包含值和具体类型的元组
  - (value, type)
  - 接口值保存了一个具体底层类型的具体值。 
  - 接口值调用方法时会执行其底层类型的同名方法。
- 未初始化的两个情况
  - Interface values with nil underlying values
  - Nil interface values （这个会报错）