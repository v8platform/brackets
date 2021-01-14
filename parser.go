package raw_parser

import (
	"bufio"
)

type BucketsParser struct {
	rd *bufio.Reader
}

func (p BucketsParser) NextNode() *BucketsNode {

	textNextNode := p.nextNodeText()

	if len(textNextNode) == 0 {
		return nil
	}

	node := ParseBlock(textNextNode)
	return &node

}

func (p BucketsParser) nextNodeText() string {

	var (
		started                 bool
		index, quotes, brackets int
		nodeText                string
	)

	endIndex := -1

	for {

		r, _, err := p.rd.ReadRune()

		if err != nil {
			return nodeText
		}

		if !started && r == '{' {
			started = true
		}

		if started {
			nodeText += string(r)
		} else {
			continue
		}

		endIndex = GetNodeEndIndex(nodeText, &index, &quotes, &brackets)

		if endIndex != -1 {
			break
		}

	}

	return nodeText

}

func GetNodeEndIndex(text string, index, quotes, brackets *int) int {

	startIdx := *index

	for *index < len(text) {

		prevChar := uint8(0)

		if *index > startIdx {
			prevChar = text[*index-1]
		}

		curChar := text[*index]

		if prevChar == ',' && curChar == '"' {

			textValueEndIndex := GetTextValueEndIndex(text, *index)

			if textValueEndIndex == -1 {
				*index = textValueEndIndex
				return textValueEndIndex
			}

			*index = textValueEndIndex
			*index++
			continue
		}

		switch curChar {
		case '"':
			*quotes++
		case '{':
			*brackets++
		case '}':
			*brackets--
		}
		if *brackets == 0 && (*quotes == 0 || (*quotes != 0 && (*quotes%2) == 0)) {
			return *index
		}

		*index++
	}

	return -1
}

func GetValueEndIndex(text string, idx int) int {

	for i := idx; i < len(text); i++ {

		if c := text[i]; c == ',' || c == '}' {
			return i
		}
	}

	return -1
}

func GetTextValueEndIndex(text string, idx int) int {

	var quotes int

	for i := idx; i < len(text); i++ {

		curChar := text[i]
		nextChar := uint8(0)

		if len(text) > i+1 {
			nextChar = text[i+1]
		}

		if curChar == '"' {
			quotes++
		}

		if curChar == '"' &&
			(nextChar == ',' || nextChar == '}') &&
			(quotes == 0 || quotes%2 == 0) {
			return i
		}
	}

	return -1
}

func ParseBlock(text string, startEndIdx ...int) BucketsNode {

	node := BucketsNode{}

	startIdx := 0
	endIdx := -1

	if len(startEndIdx) > 0 && len(startEndIdx) == 2 {

		startIdx = startEndIdx[0]
		endIdx = startEndIdx[1]

	}

	if endIdx == -1 {
		endIdx = len(text) - 1
	}

	if text[startIdx] == '{' {
		startIdx++
	}
	if text[endIdx] == '}' {
		endIdx--
	}

	for i := startIdx; i <= endIdx; i++ {

		curChar := text[i]

		switch {

		case curChar == '"':

			valueEndIndex := GetTextValueEndIndex(text, i)
			value := text[i : valueEndIndex-1]
			node.Nodes = append(node.Nodes, NewValueNode(value))
			i = valueEndIndex

		case curChar == '{':

			var (
				quotes   int
				brackets int
				idx      int
			)

			idx = i

			valueEndIndex := GetNodeEndIndex(text, &idx, &quotes, &brackets)
			blockText := text[i+1 : valueEndIndex]
			childNode := ParseBlock(blockText)
			node.Nodes = append(node.Nodes, childNode)
			i = valueEndIndex
		case curChar != '"' &&
			curChar != '}' &&
			curChar != ',' &&
			curChar != '\n' &&
			curChar != ' ':

			valueEndIndex := GetValueEndIndex(text, i)
			node.Nodes = append(node.Nodes, NewValueNode(text[i:valueEndIndex]))
			i = valueEndIndex
		}

	}

	return node

}
