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
	if list1.Head == nil {
		return list2
	}
	if list2.Head == nil {
		return list1
	}
	listaMesclada := LinkedList{}
	ponteiro1 := list1.Head
	ponteiro2 := list2.Head

	if ponteiro1.Value <= ponteiro2.Value {
		listaMesclada.Head = ponteiro1
		ponteiro1 = ponteiro1.Next
	} else {
		listaMesclada.Head = ponteiro2
		ponteiro2 = ponteiro2.Next
	}
	ponteiroAtual := listaMesclada.Head

	for ponteiro1 != nil && ponteiro2 != nil {
		if ponteiro1.Value <= ponteiro2.Value {
			ponteiroAtual.Next = ponteiro1
			ponteiro1 = ponteiro1.Next
		} else {
			ponteiroAtual.Next = ponteiro2
			ponteiro2 = ponteiro2.Next
		}
		ponteiroAtual = ponteiroAtual.Next
	}
	if ponteiro1 != nil {
		ponteiroAtual.Next = ponteiro1
	} else if ponteiro2 != nil {
		ponteiroAtual.Next = ponteiro2
	}
	return &listaMesclada
}
