package bonus

//Você recebe as cabeças de duas listas ligadas ordenadas, list1 e list2.
//
//Una as duas listas em uma única lista ordenada. A lista resultante deve ser construída juntando os nós das duas
//primeiras listas.
//
//Retorne a cabeça da lista ligada mesclada.

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func MergeTwoLists(list1 *LinkedList, list2 *LinkedList) *LinkedList {
	return nil
}
