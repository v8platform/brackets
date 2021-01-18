package brackets_test

import (
	"bytes"
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/v8platform/brackets"
	"log"
	"os"
)

func ExampleParser_NextNode_bytes() {

	data := []byte("{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},")

	parser := brackets.NewParser(bytes.NewReader(data))

	for node, size := parser.NextNode(); node != nil; node, size = parser.NextNode() {

		log.Printf("readed node <%s> size<%d>", node, size)

	}

}

func ExampleParser_NextNode_file() {

	filename := "./20200412130000.lgp"

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		log.Panicln(err)
	}

	parser := brackets.NewParser(file)

	for node, size := parser.NextNode(); node != nil; node, size = parser.NextNode() {

		log.Printf("readed node <%s> size<%d>", node, size)

	}

}

func ExampleParser_ReadAllNodes_bytes() {

	data := []byte("{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},")

	parser := brackets.NewParser(bytes.NewReader(data))
	nodes, size := parser.ReadAllNodes()

	pp.Println("readed nodes", len(nodes), "size", size)
	pp.Println(nodes)

}

func ExampleParser_ReadAllNodes_file() {

	filename := "./20200412130000.lgp"

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		log.Panicln(err)
	}

	parser := brackets.NewParser(file)

	nodes, size := parser.ReadAllNodes()

	log.Printf("readed nodes %d, size: %d", len(nodes), size)
	pp.Println(nodes)
}

func ExampleBracketsNode_GetNode() {

	data := []byte("{20200412134348,N,\n{0,13},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},")
	//                      ^       ^    ^  ^     ....                 ^
	// Адрес ноды           0       1   2,0 2,1   ....                11,0
	parser := brackets.NewParser(bytes.NewReader(data))

	node, _ := parser.NextNode()

	node0, _ := node.GetNodeE(0)
	fmt.Printf("node <%s>\n", node0.String()) // 20200412134348

	node1 := node.Get(1)
	fmt.Printf("node <%s>\n", node1) // N

	node21 := node.Int64(2, 1)
	fmt.Printf("node <%d>\n", node21) // {0,13} -> 13

	node11 := node.GetNode(11)
	fmt.Printf("node <%s>\n", node11) // {"U"}

	_, err := node.GetNodeE(1, 2) // Отсутствующий адрес ноды
	fmt.Printf("err <%s>\n", err) // {0,13} -> 13

	// Output:
	// node <20200412134348>
	// node <N>
	// node <13>
	// node <{U}>
	// err <address node is broken>

}
