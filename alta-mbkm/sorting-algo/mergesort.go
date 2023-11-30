package main

// func main() {
// 	arr := []int{87, 57, 370, 110, 90, 610, 2, 710, 140, 203, 150}

// 	fmt.Println(arr)

// 	mergeSort(arr, 0, len(arr) - 1)

// 	fmt.Println(arr)

// }

//* Mergesort algorithm
func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	L := make([]int, 0)
	R := make([]int, 0)

	for i := 0; i < n1; i++ {
		L = append(L, arr[l + i]) 
	}
    for j := 0; j < n2; j++ {
        R = append(R, arr[m + 1 + j]) 
	}

	i := 0
	j := 0
	k := l
	for i < n1 && j < n2 {
        if (L[i] <= R[j]) {
            arr[k] = L[i]
            i++
        } else {
            arr[k] = R[j]
            j++
        }
        k++
    }

	for i < n1 {
        arr[k] = L[i]
        i++
        k++
    }

	for j < n2 {
        arr[k] = R[j]
        j++
        k++
    }

}

func mergeSort(arr []int, start, end int) {
	if start < end {
        // Same as (l+r)/2, but avoids overflow for
        // large l and h
        m := start + (end - start) / 2

        // Sort first and second halves
        mergeSort(arr, start, m);
        mergeSort(arr, m + 1, end);

        merge(arr, start, m, end);
    }
}
//! End