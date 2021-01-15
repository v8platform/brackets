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

### Чтение полное последовательное чтение скобко файла в массив нод/объектов
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