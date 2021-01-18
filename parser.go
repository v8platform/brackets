package brackets

import (
	"bufio"
	"io"
	"strings"
)

type Parser struct {
	rd *bufio.Reader
}

const (
	NullRune           = '\uFFFD'
	OpenBracketRune    = '{'
	CloseBracketRune   = '}'
	QuoteRune          = '"'
	SpaceRune          = ' '
	CommaRune          = ','
	NewLineRune        = '\n'
	CarriageReturnRune = '\r'
	spaceRunes         = " \n\r\t"
)

// NewParser создает новый парсер для скобко файла
func NewParser(r io.Reader) *Parser {
	return &Parser{
		bufio.NewReader(r),
	}
}

// ReadAllNodes выполняет последовательное чтение нод/объектов скобко файла
func (p Parser) NextNode() (Node, int) {

	textNextNode, size := p.nextNodeText()

	if len(textNextNode) == 0 {
		return nil, 0
	}

	node := parseBlock(textNextNode)
	node.size = size

	return node, size

}

// ReadAllNodes выполняет чтение всех нод/объектов скобко файла
func (p Parser) ReadAllNodes() (Nodes, int) {

	var (
		nodes Nodes
		size  int
	)

	for node, s := p.NextNode(); node != nil; node, s = p.NextNode() {
		size += s
		nodes = append(nodes, node)

	}

	return nodes, size
}

func (p Parser) nextNodeText() ([]rune, int) {

	var (
		started                 bool
		index, quotes, brackets int
		nodeText                []rune
		size                    int
	)

	endIndex := -1

	for {

		r, n, err := p.rd.ReadRune()
		size += n
		if err != nil {
			return nodeText, size
		}

		if !started && r == OpenBracketRune {
			started = true
		}

		if started {
			nodeText = append(nodeText, r)
		} else {
			continue
		}

		endIndex = getNodeEndIndex(nodeText, &index, &quotes, &brackets)

		if endIndex != -1 {
			break
		}

	}

	return nodeText, size

}

func getNodeEndIndex(text []rune, index, quotes, brackets *int) int {

	startIdx := *index

	for *index < len(text) {

		prevChar := NullRune

		if *index > startIdx {
			prevChar = text[*index-1]
		}

		curChar := text[*index]

		if prevChar == CommaRune && curChar == QuoteRune {

			textValueEndIndex := getTextValueEndIndex(text, *index)

			if textValueEndIndex == -1 {
				*index = textValueEndIndex
				return textValueEndIndex
			}

			*index = textValueEndIndex
			*index++
			continue
		}

		switch curChar {
		case QuoteRune:
			*quotes++
		case OpenBracketRune:
			*brackets++
		case CloseBracketRune:
			*brackets--
		}
		if *brackets == 0 && (*quotes == 0 || (*quotes != 0 && (*quotes%2) == 0)) {
			return *index
		}

		*index++
	}

	return -1
}

func getValueEndIndex(text []rune, idx int) int {

	for i := idx; i < len(text); i++ {

		if c := text[i]; c == CommaRune || c == CloseBracketRune {
			return i
		}
	}

	return -1
}

func getTextValueEndIndex(text []rune, idx int) int {

	var quotes int

	for i := idx; i < len(text); i++ {

		curChar := text[i]
		nextChar := NullRune

		if len(text) > i+1 {
			nextChar = text[i+1]
		}

		if curChar == QuoteRune {
			quotes++
		}

		if curChar == QuoteRune &&
			(nextChar == CommaRune || nextChar == CloseBracketRune) &&
			(quotes == 0 || quotes%2 == 0) {
			return i
		}
	}

	return -1
}

func parseBlock(text []rune, startEndIdx ...int) bracketsNode {

	node := bracketsNode{}

	startIdx := 0
	endIdx := -1

	if len(startEndIdx) > 0 && len(startEndIdx) == 2 {

		startIdx = startEndIdx[0]
		endIdx = startEndIdx[1]

	}

	if endIdx == -1 {
		endIdx = len(text) - 1
	}

	if text[startIdx] == OpenBracketRune {
		startIdx++
	}
	if text[endIdx] == CloseBracketRune {
		endIdx--
	}

	for i := startIdx; i <= endIdx; i++ {

		curChar := text[i]

		switch {

		case curChar == QuoteRune:

			valueEndIndex := getTextValueEndIndex(text, i)
			if valueEndIndex == -1 {
				return node
			}

			value := text[i+1 : valueEndIndex]
			node.Nodes = append(node.Nodes, newValueNode(string(value)))
			i = valueEndIndex

		case curChar == OpenBracketRune:

			var (
				quotes   int
				brackets int
				idx      int
			)

			idx = i

			valueEndIndex := getNodeEndIndex(text, &idx, &quotes, &brackets)
			if valueEndIndex == -1 {
				return node
			}

			childNode := parseBlock(text, i, valueEndIndex)
			node.Nodes = append(node.Nodes, childNode)
			i = valueEndIndex
		case curChar != QuoteRune &&
			curChar != CloseBracketRune &&
			curChar != CommaRune &&
			!isSpaceRune(curChar) &&
			curChar != SpaceRune:

			valueEndIndex := getValueEndIndex(text, i)
			if valueEndIndex == -1 {
				return node
			}
			node.Nodes = append(node.Nodes, newValueNode(string(text[i:valueEndIndex])))
			i = valueEndIndex
		}

	}

	return node

}

func inSlice(r rune, slice string) bool {

	return strings.ContainsRune(slice, r)

}

func isSpaceRune(r rune) bool {

	return inSlice(r, spaceRunes)

}
