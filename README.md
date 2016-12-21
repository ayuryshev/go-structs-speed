# Speed of golang structures

## Description

This collection of tests answers questions "What is faster in the best case? And how much?"

Some of the results are unexpected for me.

### __Example 1. What is faster?__

This:
```go
func int64ArrayAtomicInsertor() {
	var int64Array [TestQty]int64
	for i := int64(0); i < TestQty; i++ {
		atomic.StoreInt64(&int64Array[i], i)
	}
}
``` 
or this:

```go
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
```
or that:

```go
func int64ChanInsertor() {
	int32Chan := make(chan int32, TestQty)
	for i := int32(0); i < TestQty; i++ {
		int32Chan <- i
	}
	close(int32Chan)
}
```

?

### Example 2. How much **this** is slower than **that**?

**This**:

```go
func structArrayInt32Insertor() {
	type StructInt32 struct{ x int32 }
	var structInt32Array [TestQty]StructInt32
	for i := int32(0); i < TestQty; i++ {
		structInt32Array[i].x = i
	}
}
```

and **that**:

```go
func int32ArrayInsertor() {
	var intArray [TestQty]int32
	for i := int32(0); i < TestQty; i++ {
		intArray[i] = i
	}
}
```

----------

### __Results__


<table>
<tr><th>Structure</th><th>Operation</th><th>Summary Duration(Ns)</th><th>Op Cost(ns)</th></tr>
<tr>
        <td><pre>type StructInt32 struct{ x int32 }
var structInt32Array [TestQty]StructInt32</pre></td><td><pre>structInt32Array[i].x = i</pre></td><td>13280408</td>
        <td>1.26ns</td>
    </tr><tr>
        <td><pre>boolArray [testQty]bool</pre></td><td><pre>boolArray[i]=(i%2==0)</pre></td><td>15358327</td>
        <td>1.46ns</td>
    </tr><tr>
        <td><pre>int64Array [testQty]int64</pre></td><td><pre>int64Array[i]=i</pre></td><td>16585783</td>
        <td>1.58ns</td>
    </tr><tr>
        <td><pre>int32Array [testQty]int32</pre></td><td><pre>int32Array[i]=i</pre></td><td>20111850</td>
        <td>1.91ns</td>
    </tr><tr>
        <td><pre>byteArray [testQty]byte</pre></td><td><pre>byteArray[i]=byte(i % 255)</pre></td><td>29764109</td>
        <td>2.83ns</td>
    </tr><tr>
        <td><pre>int32Slice []int32</pre></td><td><pre>*int32Slice = append(*int32Slice, i)</pre></td><td>106798203</td>
        <td>10.18ns</td>
    </tr><tr>
        <td><pre>int32Array [testQty]int32</pre></td><td><pre>atomic.StoreInt32(&amp;intArray[i], i)</pre></td><td>118543330</td>
        <td>11.3ns</td>
    </tr><tr>
        <td><pre>int64Array [testQty]int64</pre></td><td><pre>atomic.StoreInt64(&amp;intArray[i], i)</pre></td><td>137920576</td>
        <td>13.15ns</td>
    </tr><tr>
        <td><pre>int64Slice []int64</pre></td><td><pre>*int64Slice = append(*int64Slice, i)</pre></td><td>170328162</td>
        <td>16.24ns</td>
    </tr><tr>
        <td><pre>type MutexedInt32Array struct {
	sync.Mutex
	items [testQty]int32
}</pre></td><td><pre>mtxInt32Array.Lock()
mtxInt32Array.items[i] = i
mtxInt32Array.Unlock()</pre></td><td>310796749</td>
        <td>29.63ns</td>
    </tr><tr>
        <td><pre>type MutexedInt64Array struct {
	sync.Mutex
	items [testQty]int64
}</pre></td><td><pre>mtxInt64Array.Lock()
mtxInt64Array.items[i] = i
mtxInt64Array.Unlock()</pre></td><td>323613671</td>
        <td>30.86ns</td>
    </tr><tr>
        <td><pre>int8Chan := make(chan int8, TestQty)</pre></td><td><pre>int8chan &lt;- i</pre></td><td>461275292</td>
        <td>43.99ns</td>
    </tr><tr>
        <td><pre>int64Chan := make(chan int64, TestQty)</pre></td><td><pre>int64chan &lt;- i</pre></td><td>476790665</td>
        <td>45.47ns</td>
    </tr><tr>
        <td><pre>int32Chan := make(chan int32, TestQty)</pre></td><td><pre>int32chan &lt;- i</pre></td><td>484103547</td>
        <td>46.16ns</td>
    </tr><tr>
        <td><pre>intMap := map[int32]bool{}</pre></td><td><pre>intMap[i] = true</pre></td><td>3152781487</td>
        <td>300.67ns</td>
    </tr>
</table>


