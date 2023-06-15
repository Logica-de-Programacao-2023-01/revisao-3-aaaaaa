package q4

//Dado um array não vazio de números inteiros "nums", cada elemento aparece duas vezes, exceto um. Encontre esse único
//elemento.
//
//Você deve implementar uma solução com complexidade de tempo linear e sem memória extra.

func SingleNumber(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 0
	}
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}
	for i, valor := range nums {
		if count[valor] == 1 {
			return i
		}
	}
	return -1
}
