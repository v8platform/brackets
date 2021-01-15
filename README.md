# Brackets-parser
Библиотека для парсинга внутрисистемного формата файлов 1С. Предприятие

## Функционал

* `brackets.NewParser(reader io.Reader)` Создание парсера скобко файла
* `parser.NextNode() brackets.Node` Выполняет парсинг следующей ноды/объекта
* `parser.ReadAllNodes() brackets.Nodes` Выполняет парсинг всего скобко файла и возвращает все ноды/объекты 

## Примеры

### Чтение последовательное скобко файла следующею ноду/объект
```go
package main

import (
	"log"
	"os"
	"v8platform/brackets"
)

func main() {

	filename := "./20200412130000.lgp"

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		log.Panicln(err)
	}

	parser := brackets.NewParser(file)

	for node := parser.NextNode(); node != nil; node = parser.NextNode() {

		log.Printf("readed node <%s>", node)

	}
}

```

### Работы с нодой и адресами других нод
```go
package main

import (
	"fmt"
	"v8platform/brackets"
)

func main() {

	data := []byte("{20200412134348,N,\n{0,13},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n}")
	//                      ^       ^    ^  ^     ....                 ^
	// Адрес ноды           0       1   2,0 2,1   ....                11,0
	parser := brackets.NewParser(bytes.NewReader(data))

	node := parser.NextNode()

	node0, _ := node.GetNode(0)
	fmt.Printf("node <%s>\n", node0.String()) // 20200412134348

	node1, _ := node.GetNode(1)
	fmt.Printf("node <%s>\n", node1.String()) // N

	node21, _ := node.GetNode(2, 1)
	fmt.Printf("node <%d>\n", node21.Int()) // {0,13} -> 13

	node11, _ := node.GetNode(11)
	fmt.Printf("node <%s>\n", node11) // {"U"}

	_, err := node.GetNode(1, 2) // Отсутствующий адрес ноды
	fmt.Printf("err <%s>\n", err) // {0,13} -> 13

	// Output:
	// node <20200412134348>
	// node <N>
	// node <13>
	// node <{U}>
	// err <address node is broken>
}

```

### Полное последовательное чтение скобко файла в массив нод/объектов
```go
package main

import (
	"log"
	"os"
	"v8platform/brackets"
)

func main() {

	filename := "./20200412130000.lgp"

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		log.Panicln(err)
	}

	parser := brackets.NewParser(file)

	nodes := parser.ReadAllNodes()

	log.Printf("readed nodes %d", len(nodes))
}

```