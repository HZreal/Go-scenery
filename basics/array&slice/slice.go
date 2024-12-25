package main

/**
 * @Author nico
 * @Date 2024-12-25
 * @File: slice.go
 * @Description:
 */

// ---------------------------------------------------------------------------------------
// slice 和 map 都是引用类型，它们包含对底层数据结构的指针。因此，当你将它们作为参数传递、返回或存储时，需要特别小心，以避免不必要的数据共享或修改

type Driver struct {
	name  string
	trips []int
}

func (d *Driver) SetTrips1(trips []int) {
	// 这样赋值会导致 d.trips 和 trips 共享同一个底层数组，修改一个会影响另一个
	d.trips = trips
}

func (d *Driver) SetTrips2(trips []int) {
	// 进行了拷贝，不会影响原数组
	d.trips = make([]int, len(trips))
	copy(d.trips, trips)
}

// ---------------------------------------------------------------------------------------

func main() {

}
