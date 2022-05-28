package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	for _, value := range input {
		defer func(elem int64) {
			result = append(result, elem)
		}(value)
	}
	return
}
