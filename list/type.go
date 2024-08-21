package list

// List 泛型接口
// 定义各个方法的参数和表现
type List[T any] interface {
	Get(index int) (T, error)
	Append(ts ...T) error
	Add(index int, t T) error

	// Set 重置index位置的值
	Set(index int, t T) error
	Delete(index int) (T, error)
	Len() int
	Cap() int

	// Range遍历List的所有元素, 执行fn(index, t)
	Range(fn func(index int, t T) error) error

	// AsSlice将List转化为切片，不返回nil
	// 没有元素，返回len和cap为0的切片
	// 每次调用返回一个新切片
	AsSlice() []T
}
