package main

import (
	"sync"
	"sync/atomic"
)

func mapInt32Insertor() {
	intMap := map[int32]bool{}
	for i := int32(0); i < TestQty; i++ {
		intMap[i] = true
	}
}

func structArrayInt32Insertor() {
	type StructInt32 struct{ x int32 }
	var structInt32Array [TestQty]StructInt32
	for i := int32(0); i < TestQty; i++ {
		structInt32Array[i].x = i
	}
}

func int32ArrayInsertor() {
	var intArray [TestQty]int32
	for i := int32(0); i < TestQty; i++ {
		intArray[i] = i
	}
}

func int64ArrayInsertor() {
	var intArray [TestQty]int64
	for i := int64(0); i < TestQty; i++ {
		intArray[i] = i
	}
}

func int8ArrayInsertor() {
	var intArray [TestQty]byte
	for i := int32(0); i < TestQty; i++ {
		intArray[i] = byte(i % 255)
	}
}

func boolArrayInsertor() {
	var boolArray [TestQty]bool
	for i := int32(0); i < TestQty; i++ {
		boolArray[i] = (i%2 == 0)
	}
}

func int32ChanInsertor() {
	int32Chan := make(chan int32, TestQty)
	for i := int32(0); i < TestQty; i++ {
		int32Chan <- i
	}
	close(int32Chan)
}

func int64ChanInsertor() {
	int32Chan := make(chan int32, TestQty)
	for i := int32(0); i < TestQty; i++ {
		int32Chan <- i
	}
	close(int32Chan)
}

func int8ChanInsertor() {
	int8Chan := make(chan int8, TestQty)
	for i := int32(0); i < TestQty; i++ {
		int8Chan <- int8(i % 255)
	}
	close(int8Chan)
}

func int32SliceInsertor() {
	intSlice := &([]int32{})

	for i := int32(0); i < TestQty; i++ {
		*intSlice = append(*intSlice, i)

	}
}

func int64SliceInsertor() {
	intSlice := &([]int64{})

	for i := int64(0); i < TestQty; i++ {
		*intSlice = append(*intSlice, i)
	}
}

func int32ArrayAtomicInsertor() {
	var int32Array [TestQty]int32
	for i := int32(0); i < TestQty; i++ {
		atomic.StoreInt32(&int32Array[i], i)
	}
}

func int64ArrayAtomicInsertor() {
	var int64Array [TestQty]int64
	for i := int64(0); i < TestQty; i++ {
		atomic.StoreInt64(&int64Array[i], i)
	}
}

func mutexedInt32ArrayInsertor() {
	type MutexedInt32Array struct {
		sync.Mutex
		items [TestQty]int32
	}
	var mtxInt32Array MutexedInt32Array
	for i := int32(0); i < TestQty; i++ {
		mtxInt32Array.Lock()
		mtxInt32Array.items[i] = i
		mtxInt32Array.Unlock()
	}
}

func mutexedInt64ArrayInsertor() {
	type MutexedInt64Array struct {
		sync.Mutex
		items [TestQty]int64
	}
	var mtxInt64Array MutexedInt64Array
	for i := int64(0); i < TestQty; i++ {
		mtxInt64Array.Lock()
		mtxInt64Array.items[i] = i
		mtxInt64Array.Unlock()
	}
}
