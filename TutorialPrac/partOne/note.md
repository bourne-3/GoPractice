记录重要的知识点

### part3
- Slices are like references to arrays 
  - A slice does not store any data, it just describes a section of an underlying array.
    - 注意这里的not store 以及 underlying
  - Changing the elements of a slice modifies the corresponding elements of its underlying array.
- len 以及 cap
  - The length of a slice is the number of elements it contains.
  - The capacity of a slice is the number of elements in the **underlying array**, counting from the first element in the slice.
- 默认值
  - The default is zero for the low bound and the length of the slice for the high bound.
    - 注意这里的high bound 是切片的len，而已这个len是可变的
    - 具体参考 nums3[2:] 以及 nums3[2:5]这两个的区别 （此时nums3的长度为2）
