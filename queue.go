package queue_interface

import (
	"errors"
	"fmt"
	"reflect"
)

// TSlice 为了防止浪费数组的空间，使用环形的队列
type TSlice struct {
	i []interface{}
	//你需要的其他字段
	Head    int //头指针
	Tail    int //位指针
	Size    int //当前队列的大小
	MaxSize int //队列的最大容量
	Value   interface{}
	Type    reflect.Kind
}

var (
	full    = errors.New("队列已经满了,无法加入元素")
	empty   = errors.New("队列是空的，无法取出元素")
	typeErr = errors.New("与加入队列的第一个元素类型不符合")
)

func New(maxSize int) *TSlice {
	//head tail 初始化为0,为了防止未知的错误，通常在环形队列中空一个标志位
	return &TSlice{
		i:       make([]interface{}, maxSize),
		Head:    0,
		Tail:    0,
		Size:    0,
		MaxSize: maxSize,
	}
}
func (t *TSlice) Top() (interface{}, error) {
	if t.empty() {
		return nil, empty
	}
	return t.i[t.Head], nil
}

func (t *TSlice) Push(value interface{}) error {
	//记录第一个push进来的类型

	//判断队列是否已经到最大容量
	if t.full() {
		return full
	}
	if t.isSameType(value) {
		return typeErr
	}
	t.Tail++
	t.i[t.Tail] = value
	t.Size++
	//记录下第一个元素的type
	if t.Size == 1 {
		v := reflect.ValueOf(value)
		t.Type = v.Kind()

	}
	return nil
}
func (t *TSlice) Pop() (interface{}, error) {
	if t.empty() {
		return nil, empty
	}
	value := t.i[t.Head]
	t.Head++
	t.Size--
	return value, nil
}

func (t *TSlice) Len() int {
	return t.Size
}
func (t *TSlice) Show() {
	//获取队列中的所有元素
	fmt.Println("环形列表元素如下")
	if t.Size == 0 {
		fmt.Println("环形列表为空")
	}
	temp := t.Head
	for i := 0; i < t.Size; i++ {
		fmt.Printf("环形队列的[%d]值为%v \n", temp, t.i[temp])
		temp = (temp + 1) % t.MaxSize
	}
}

//内部使用，所以小写
func (t *TSlice) full() bool {
	if (t.Tail+1)%t.MaxSize == t.Head {
		return true
	}
	return false

}
func (t *TSlice) empty() bool {
	if t.Head == t.Tail {
		return true
	}
	return false
}
func (t *TSlice) isSameType(value interface{}) bool {
	v := reflect.ValueOf(value)
	if t.Type == v.Kind() {
		return true
	}
	return false

}
