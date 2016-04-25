package array

func ReverseArray(input []int) []int {

    for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
        input[i], input[j] = input[j], input[i]
    }

    return input
}

