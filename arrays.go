package main

import (
    "fmt"
    "math/rand"
    "time"
)

var (
    l int = 20 // array length
    m int = 30 // array max int 
    s int = 8 // search for 
)

func GenSlice(size int, maxInt int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(maxInt)
	}
	return slice
}

func GenShuffledSequence(size int) []int {
    seq := make([]int, size, size)
    for i := range seq {
        seq[i] = i+1
    }
    rand.Seed(time.Now().UnixNano())
    for i := len(seq) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        seq[i], seq[j] = seq[j], seq[i]
    }
    return seq
}

func SortArray(nums []int) []int {
    if len(nums)< 20 {
        return MergeSort(nums)
    }
    QuickSort(nums, 0, len(nums)-1)
    return nums
}

func MergeSort(s []int) []int {
    if len(s) < 2 {
        return s
    }
    l := s[:len(s)/2]
    r := s[len(s)/2:]
    return Merge(MergeSort(l),MergeSort(r))
}

func Merge(l, r []int) (result []int) {
    result = make([]int, 0, len(r)+len(l))
    for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(result, r...)
		}
		if len(r) == 0 {
			return append(result, l...)
		}
		if l[0] <= r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
    }
    return
}

func QuickSort(nums []int, low, high int) []int {
    if low < high {
        pivot := Partition(nums, low, high)
        QuickSort(nums, low, pivot - 1)
        QuickSort(nums, pivot + 1, high)
    }
    return nums
}

func Partition(nums []int, low, high int) int {
    pivot := nums[low]
    for low < high {
        for low < high && nums[high] >= pivot {
            high--
        }
        nums[low], nums[high] = nums[high], nums[low]
        for low < high && nums[low] <= pivot {
            low++
        }
        nums[low], nums[high] = nums[high], nums[low]
    }
    return low
}

func BinarySearch(a []int, search int) int {
    var result int
    mid := len(a) / 2
    switch {
    case len(a) == 0:
        result = -1
    case a[mid] > search:
        result = BinarySearch(a[:mid], search)
    case a[mid] < search:
        result = BinarySearch(a[mid+1:], search)
        if result >= 0 { // if found in high slice
            result += mid + 1
        }
    default: // a[mid] == search
        result = mid // found
    }
    return result
}

func ReverseArray(a []int) []int {
    var result []int
    if len(a) < 2 {
        return a
    }
    size := len(a) - 1
    for size >= 0 {
        result = append(result, a[size])
        size--
    }
    return result
}

func ReverseRecursive(a []int) []int{
    if len(a)<2 {
        return(a)
    }
    return append(ReverseRecursive(a[1:]),a[0])
}

func main() {
    // create sequence of length l with max int of m
    a := GenSlice(l,m)
    fmt.Printf("Unsorted: %v\n", a)
    fmt.Printf("Sorted: %v\n", SortArray(a))
    i := BinarySearch(a, s)
    if i > 0 {
        fmt.Printf("%d found at index %d\n", s, i)
    }
    fmt.Println("Reversed:", ReverseRecursive(a))
}
