package q3

//Você está subindo uma escada. Leva n passos para chegar ao topo.
//
//A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

func ClimbStairs(n int) int {
	if n == 1 {
		return 1 // apenas um jeito
	}
	//
	if n == 2 {
		return 2 // apenas dois jeitos
	}

	var degraus = make([]int, n+1)
	degraus[0], degraus[1] = 1, 1

	for i := 2; i <= n; i++ {
		degraus[i] = degraus[i-1] + degraus[i-2]
	}
	return degraus[n]
}
