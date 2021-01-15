package brackets_test

import (
	"bytes"
	"github.com/k0kubun/pp"
	"github.com/v8platform/brackets"
	"log"
	"os"
)

func ExampleParser_NextNode_bytes() {

	data := []byte("{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},")

	parser := brackets.NewParser(bytes.NewReader(data))

	for node := parser.NextNode(); node != nil; node = parser.NextNode() {

		log.Printf("readed node <%s>: ", node)

	}

}

func ExampleParser_NextNode_file() {

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

func ExampleParser_ReadAllNodes_bytes() {

	data := []byte("{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},")

	parser := brackets.NewParser(bytes.NewReader(data))
	nodes := parser.ReadAllNodes()

	log.Printf("readed nodes %d", len(nodes))
	pp.Println(nodes)

}

func ExampleParser_ReadAllNodes_file() {

	filename := "./20200412130000.lgp"

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		log.Panicln(err)
	}

	parser := brackets.NewParser(file)

	nodes := parser.ReadAllNodes()

	log.Printf("readed nodes %d", len(nodes))
	pp.Println(nodes)
}
