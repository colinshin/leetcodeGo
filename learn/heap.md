## 堆的使用
关于堆和优先队列，标准库已经实现了核心部分，详见container/heap
基本自定义集合（一般为一个切片）实现了heap.Interface的5个函数，就可以用了。
详见标准库两个example开头的测试文件。
### 需求
假设我们要同时用到大顶堆和小顶堆，怎么办？简单起见，假设元素都是int。
### 初步实现
参考标准库example_intheap_test.go,很容易写出以下代码：
```go
type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
```
### 改进实现
重复代码很多，大顶堆和小顶堆只有比较逻辑不通，即那个Less函数，其他实现没有差别。
想一想，还是Go不支持泛型惹的祸。
在支持泛型前，有没有办法减少代码呢？
还真想到一个，在Go里，函数是一等公民，可以当一般变量传递使用，我们不妨给自定义Heap增加一个属性Cmp，类型是func(i, j int) bool，也就是上边的Less函数的类型
```go
type Heap struct {
	Slice []int
	Cmp   func(i, j int) bool
}

func (h Heap) Len() int            { return len(h.Slice) }
func (h Heap) Less(i, j int) bool  { return h.Cmp(i, j) }
func (h Heap) Swap(i, j int)       { h.Slice[i], h.Slice[j] = h.Slice[j], h.Slice[i] }
func (h *Heap) Push(x interface{}) { h.Slice = append(h.Slice, x.(int)) }
func (h *Heap) Pop() interface{} {
	x := h.Slice[len(h.Slice)-1]
	h.Slice = (h.Slice)[:len(h.Slice)-1]
	return x
}
```
代码一下子减少一半!使用时是这样：
```go
	minHeap := &Heap{}
	minHeap.Cmp = func(i, j int) bool {
		return minHeap.Slice[i] < minHeap.Slice[j]
	}
	maxHeap := &Heap{}
	maxHeap.Cmp = func(i, j int) bool {
		return maxHeap.Slice[i] > maxHeap.Slice[j]
	}
```
要注意几点：
首先在刚刚初始化完要立即赋予其Cmp，Cmp为nil的话后边程序会崩溃给我们看~
其次Cmp后续不能被修改，不然堆的逻辑会混乱
问题：Cmp和Slice都对外暴露了，其实只是内部实现，不该暴露，怎么办呢？
### 进一步改进
目前想到的是这样：
```go
type Heap struct {
	slice []int
	cmp   func(i, j int) bool
}

func (h *Heap) InitWithCmp(cmp func(i, j int) bool)  {
	h.cmp = cmp
}

func (h Heap) Len() int            { return len(h.slice) }
func (h Heap) Less(i, j int) bool  { return h.cmp(i, j) }
func (h Heap) Swap(i, j int)       { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }
func (h *Heap) Push(x interface{}) { h.slice = append(h.slice, x.(int)) }
func (h *Heap) Pop() interface{} {
	x := h.slice[len(h.slice)-1]
	h.slice = (h.slice)[:len(h.slice)-1]
	return x
}
func (h *Heap) Get(i int) int {
	return h.slice[i]
}
```
使用变成了这样：
```go
	minHeap := &Heap{}
	minHeap.InitWithCmp(func(i, j int) bool {
		return minHeap.Get(i) < minHeap.Get(j)
	})
	maxHeap := &Heap{}
	maxHeap.InitWithCmp(func(i, j int) bool {
		return maxHeap.Get(i) > maxHeap.Get(j)
	})
```
基本解决了问题，虽然不够完美~
### 丰富Api，思考题
最后，为了让我们的堆更易用，可以增加Peek、Remove和Fix等方法
比如Peek，就很容易：
```go
func (h *Heap) Peek() int  { return h.slice[0] }
```
### 应用
可以参考[[1438]绝对差不超过限制的最长连续子数组](../solutions/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/d.go)的实现
