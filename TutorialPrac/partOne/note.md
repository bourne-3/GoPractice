记录重要的知识点

### part1
- Go有两个很重要的环境变量GOROOT和GOPATH,这是两个定义路径的环境变量，GOROOT是安装Go的路径，比如/usr/local/go；GOPATH是我们自己定义的开发者个人的工作空间，比如/home/flysnow/go。
- 对于包的查找，是有优先级的，编译器会优先在GOROOT里搜索，其次是GOPATH,一旦找到，就会马上停止搜索
  - go get工具可以递归获取依赖包，如果github.com/spf13/cobra也引用了其他的远程包，该工具可以一并下载下来。
  - 今天在复习的时候，解决了昨天的问题，就是go get一直下载不下来的问题。原来是GOPATH没有配置好导致的


### part3
- Slices are like references to arrays 
  - A slice does not store any data, it just describes a section of an underlying array.
    - 注意这里的not store 以及 underlying
  - Changing the elements of a slice modifies the corresponding elements of its underlying array.
- len 以及 cap
  - The length of a slice is the number of elements it contains.
  - The capacity of a slice is the number of elements in the **underlying array**, counting from the first element in the slice.
- 切片索引默认值
  - The default is zero for the low bound and the length of the slice for the high bound.
    - 注意这里的high bound 是切片的len，而已这个len是可变的
    - 具体参考 nums3[2:] 以及 nums3[2:5]这两个的区别 （此时nums3的长度为2）
- Function values
  - 函数也可以充当参数进行传递。既可以作为参数，也可以作为返回值
    - Function values may be used as function arguments and return values.
  - 