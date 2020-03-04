package my_map

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Logf("len m1= %d", len(m1))
	t.Log(m1[2])

	m2 := map[int]int{}
	t.Logf("len m2= %d", len(m2))
	t.Log(m2)
	m2[4] = 16
	t.Logf("len m2= %d", len(m2))
	t.Log(m2)

	m3 := make(map[int]int, 10) // 第二个参数是cap
	t.Logf("len m3= %d", len(m3))
	//t.Log(cap(m3))			// 但不能用cap()函数来获取map的值
}

// go里面 map 的key 不存在会返回一个零值(不能通过返回nil来判断元素是否存在) ，这时需要和该key的值为零时做区分
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m2 := map[int]int{}
	m2[1] = 0
	t.Log(m2[1])
	// 区分
	if value, ok := m2[1]; ok { // map某个key对应值的返回值有两个，第一个是该值，第二个是该值是否存在
		t.Logf("it is %d", value)
	} else {
		t.Logf("it is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m { // 类似python的for...in...结构，但go用range来表达
		t.Log(k, v)
	}
}
