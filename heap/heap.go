// Created by reyza on 23/12/2024.
package heap

func Heapify[T any](list []T, less func(i, j int) bool) {
	for i := len(list) - 1; i >= 0; i -= 1 {
		ShiftDown(list, less, i)
	}
}

func ShiftDown[T any](list []T, less func(i, j int) bool, i int) {
	if i < 0 || i >= len(list) {
		return
	}

	if list == nil {
		return
	}

	left := -1
	right := -1

	if idx := i*2 + 1; idx < len(list) {
		left = idx
	}

	if idx := i*2 + 2; idx < len(list) {
		right = idx
	}

	if left == -1 && right == -1 {
		return
	}

	if right == -1 {
		if less(left, i) {
			list[i], list[left] = list[left], list[i]
			ShiftDown(list, less, left)
		}

		return
	}

	smaller := left
	if less(right, left) {
		smaller = right
	}

	if less(smaller, i) {
		list[i], list[smaller] = list[smaller], list[i]
		ShiftDown(list, less, smaller)
	}
}

func Push[T any](list *[]T, less func(i, j int) bool, val T) {
	if list == nil {
		return
	}

	arr := *list

	arr = append(arr, val)
	*list = arr

	ShiftUp(arr, less, len(arr)-1)
}

func ShiftUp[T any](list []T, less func(i, j int) bool, i int) {
	if i == 0 {
		return
	}

	if list == nil {
		return
	}

	arr := list
	if len(arr) == 0 {
		return
	}

	parent := (i - 1) / 2

	if less(i, parent) {
		arr[i], arr[parent] = arr[parent], arr[i]
		ShiftUp(list, less, parent)
	}
}

func Pop[T any](list *[]T, less func(i, j int) bool) (T, bool) {
	var resp T

	if list == nil {
		return resp, false
	}

	arr := *list

	if len(arr) <= 0 {
		return resp, false
	}

	resp = arr[0]
	arr[0] = arr[len(arr)-1]
	*list = arr[1:]

	ShiftDown(arr, less, 0)

	return resp, true
}

func PopPush[T any](list *[]T, less func(i, j int) bool, val T) (T, bool) {
	var resp T

	if list == nil {
		return resp, false
	}

	arr := *list

	if len(arr) <= 0 {
		*list = []T{val}
		return resp, false
	}

	resp = arr[0]
	arr[0] = val
	ShiftDown(arr, less, 0)

	return resp, true
}
