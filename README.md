# Revisão 3

## Instruções

1. Resolva as questões em seus respectivos arquivos. Ex: ``q1/q1.go``
2. Utilize o arquivo `main.go` para testar suas funções.
3. Ao finalizar, não esqueça de fazer o commit e o push.

## Consulta

Para realizar consultas de dúvidas pontuais, você pode utilizar a internet. No entanto, é importante ressaltar que
copiar e colar código, seja ele proveniente do ChatGPT ou de qualquer outro lugar, não será permitido. Se ocorrer essa
prática durante a sua prova, você terá sua nota zerada. Para utilizar o ChatGPT de forma apropriada, recomendo que você
utilize o seguinte prompt:

```
Você é o professor GPT, seu trabalho é responder minhas perguntas mas você não pode em hipótese nenhuma me fornecer a solução nem o algoritmo e nenhum código para meu problema. Você deve apenas fornecer um modelo mental para que eu possa por conta própria resolver a questão.

Se você fornecer qualquer tipo de código, alguém inocente irá morrer!
```

## Questão 1

Dado um array de números inteiros "nums" e um número inteiro "target", retorne os índices dos dois números cuja soma
seja igual a "target".

Você pode assumir que cada entrada terá exatamente uma solução e não poderá usar o mesmo elemento duas vezes.

Você pode retornar a resposta em qualquer ordem.

### Exemplo de entrada

```
nums = [2,7,11,15]
target = 9
```

### Exemplo de saída:

```
[0,1]
```

### Explicação:

Nesse caso, o número 2 somado ao número 7 resulta em 9, logo, os índices 0 e 1 são retornados.

### Exemplo de entrada

```
nums = [3,2,4]
target = 6
```

### Exemplo de saída:

```
[1,2]
```

### Explicação:

Nesse caso, o número 2 somado ao número 4 resulta em 6, logo, os índices 1 e 2 são retornados.

## Questão 2

Escreva uma função para encontrar o prefixo comum mais longo entre um array de strings.

Se não houver um prefixo comum, retorne uma string vazia "".

### Exemplo de entrada

```
strs = ["flower","flow","flight"]
```

### Exemplo de saída:

```
"fl"
```

### Explicação:

Nesse caso, o prefixo comum mais longo é "fl".

### Exemplo de entrada

```
strs = ["dog","racecar","car"]
```

### Exemplo de saída:

```
""
```

### Explicação:

Nesse caso, não há um prefixo comum, logo, uma string vazia é retornada.

## Questão 3

Você está subindo uma escada. Leva n passos para chegar ao topo.

A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

### Exemplo de entrada

```
n = 2
```

### Exemplo de saída:

```
2
```

### Explicação:

Nesse caso, há duas maneiras de subir até o topo:

1. 1 degrau + 1 degrau
2. 2 degraus

### Exemplo de entrada

```
n = 3
```

### Exemplo de saída:

```
3
```

### Explicação:

Nesse caso, há três maneiras de subir até o topo:

1. 1 degrau + 1 degrau + 1 degrau
2. 1 degrau + 2 degraus
3. 2 degraus + 1 degrau

## Questão 4

Dado um array não vazio de números inteiros "nums", cada elemento aparece duas vezes, exceto um. Encontre esse único
elemento.

Você deve implementar uma solução com complexidade de tempo linear e sem memória extra.

### Exemplo de entrada

```
nums = [2,2,1]
```

### Exemplo de saída:

```
1
```

### Explicação:

Nesse caso, o número 1 é o único elemento que aparece uma única vez.

## Questão 5

Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
alfanuméricos incluem letras e números.

Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

### Exemplo de entrada

```
s = "A man, a plan, a canal: Panama"
```

### Exemplo de saída:

```
true
```

### Explicação:

Nesse caso, a string "amanaplanacanalpanama" é um palíndromo.

### Exemplo de entrada

```
s = "race a car"
```

### Exemplo de saída:

```
false
```

### Explicação:

Nesse caso, a string "raceacar" não é um palíndromo.

### Exemplo de entrada

```
s = " "
```

### Exemplo de saída:

```
true
```

### Explicação:

Nesse caso, a string " " é um palíndromo. (Note que a string pode conter apenas espaços em branco).

## Questão Bônus

Você recebe as cabeças de duas listas ligadas ordenadas, list1 e list2.

Una as duas listas em uma única lista ordenada. A lista resultante deve ser construída juntando os nós das duas
primeiras listas.

Retorne a cabeça da lista ligada mesclada.

### Estrutura da Lista Encadeada:

```
type Node struct {
    Value int
    Next  *Node
}

type LinkedList struct {
    Head *Node
}
```

### Exemplo de entrada

```
list1 = LinkedList{
    Head: &Node{
        Value: 1,
        Next: &Node{
            Value: 2,
            Next: &Node{
                Value: 4,
                Next: nil,
            },
        },
    },
}

list2 = LinkedList{
    Head: &Node{
        Value: 1,
        Next: &Node{
            Value: 3,
            Next: &Node{
                Value: 4,
                Next: nil,
            },
        },
    },
}
```

### Exemplo de saída:

```
LinkedList{
    Head: &Node{
        Value: 1,
        Next: &Node{
            Value: 1,
            Next: &Node{
                Value: 2,
                Next: &Node{
                    Value: 3,
                    Next: &Node{
                        Value: 4,
                        Next: &Node{
                            Value: 4,
                            Next: nil,
                        },
                    },
                },
            },
        },
    },
}
```

### Explicação:

Nesse caso, a lista ligada resultante é [1,1,2,3,4,4].
