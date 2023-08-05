package slice

import "errors"

// Add 在切片的index位置插入一个元素
func Add[T any](src []T, index int, element T) ([]T, error) {

	// 判断index是否越界
	if index < 0 || index >= len(src) {
		return src, errors.New("index out fo range")
	}

	// 扩容切片，这里采用的方法是给切片追加一个零值元素让append函数帮忙完成动态扩容
	var t T
	src = append(src, t)

	// 将index位置元素和其后面的元素往后移动一位
	for i := len(src) - 1; i > index; i-- {
		src[i] = src[i-1]
	}

	// 最后将index位置插入element
	src[index] = element

	return src, nil
}
