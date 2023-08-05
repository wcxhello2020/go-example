package slice

import "errors"

func Delete[T any](src []T, index int) ([]T, error) {
	// 判断index是否越界
	if index < 0 || index >= len(src) {
		return src, errors.New("index out of range")
	}

	// 删除index位置的元素，从index位置开始将后面的元素向前移动一位
	for i := index; i < len(src)-1; i++ {
		src[i] = src[i+1]
	}
	src = src[:len(src)-1]

	// 缩容，如果切片的长度小于等于容量的1/4时则考虑缩容，缩容后的切片大小为原来切片长度的2倍
	length := len(src)
	capacity := cap(src)
	if length <= capacity {
		newSlice := make([]T, length, 2*length)
		copy(newSlice, src)
		src = newSlice
	}

	return src, nil
}
