type DynamicArray struct {
    cap int
    len int
    arr []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
        cap: capacity,
        arr: make([]int,capacity),
        len:0,
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.arr[i]

}

func (da *DynamicArray) Set(i int, n int) {
    da.arr[i]=n

}

func (da *DynamicArray) Pushback(n int) {
    
    if da.len == da.cap {
		da.resize()
	}
    da.arr[da.len]=n
    da.len++
}

func (da *DynamicArray) Popback() int {
    da.len--
    return da.arr[da.len]
    
}

func (da *DynamicArray) resize() {
    da.cap= da.cap *2
    newArr := make([]int,da.cap)

    for i:=0;i<da.len;i++{
        newArr[i] = da.arr[i]
    }
    da.arr =newArr

}

func (da *DynamicArray) GetSize() int {
    return da.len
}

func (da *DynamicArray) GetCapacity() int {
    return da.cap
}
