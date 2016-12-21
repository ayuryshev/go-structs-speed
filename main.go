package main

import (
	"fmt"
	"log"
	"sort"
)

const (
	//TestQty - size of test structure
	TestQty = 1024 * 1024
)

func RunSpeedTests() {

	speedTests := SpeedTests{
		{mapInt32Insertor, "intMap := map[int32]bool{}", "intMap[i] = true", nil},

		{structArrayInt32Insertor, `	type StructInt32 struct{ x int32 }
	var structInt32Array [TestQty]StructInt32`, "structInt32Array[i].x = i", nil},

		{int32ArrayInsertor, "int32Array [testQty]int32", "int32Array[i]=i", nil},
		{int64ArrayInsertor, "int64Array [testQty]int64", "int64Array[i]=i", nil},
		{int8ArrayInsertor, "byteArray [testQty]byte", "byteArray[i]=byte(i % 255)", nil},
		{boolArrayInsertor, "boolArray [testQty]bool", "boolArray[i]=(i%2==0)", nil},

		{int32ArrayAtomicInsertor, "int32Array [testQty]int32", "atomic.StoreInt32(&intArray[i], i)", nil},
		{int64ArrayAtomicInsertor, "int64Array [testQty]int64", "atomic.StoreInt64(&intArray[i], i)", nil},

		{int32SliceInsertor, "int32Slice []int32", "*int32Slice = append(*int32Slice, i)", nil},
		{int64SliceInsertor, "int64Slice []int64", "*int64Slice = append(*int64Slice, i)", nil},

		{mutexedInt32ArrayInsertor, `type MutexedInt32Array struct {
	sync.Mutex
	items [testQty]int32
}`, `	mtxInt32Array.Lock()
	mtxInt32Array.items[i] = i

	mtxInt32Array.Unlock()`, nil},
		{mutexedInt64ArrayInsertor, `type MutexedInt64Array struct {
	sync.Mutex
	items [testQty]int64
}`, `	mtxInt64Array.Lock()
	mtxInt64Array.items[i] = i
	mtxInt64Array.Unlock()`, nil},

		{int32ChanInsertor, "int32Chan := make(chan int32, TestQty)", "int32chan <- i", nil},
		{int64ChanInsertor, "int64Chan := make(chan int64, TestQty)", "int64chan <- i", nil},
		{int8ChanInsertor, "int8Chan := make(chan int8, TestQty)", "int8chan <- i", nil},
	}

	for i := 0; i < len(speedTests); i++ {
		if err := speedTests[i].MeasureDurations(10); err != nil {
			log.Printf("%v. %v. Error: %v\n", speedTests[i].ItemsDefinition, speedTests[i].OperationDefinition, err)
		}
	}

	sort.Sort(speedTests)
	for i := 0; i < len(speedTests); i++ {
		fmt.Printf("%v\n", speedTests[i])
	}

	// speedTests.HTML()

}

func main() {
	RunSpeedTests()
}
