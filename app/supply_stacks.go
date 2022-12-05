package app

import (
	"math"
	"strconv"
	"strings"
)

func MoveSupplyStacks(input string) (topmostStacks string) {

	stacksInput, instructionsInput := strings.Split(input, "\n\n")[0], strings.Split(input, "\n\n")[1]

	stacksInput = sanitize(stacksInput)
	stacks := buildStacksMatrix(stacksInput)

	instructionsInput = sanitize(instructionsInput)
	if instructionsInput != "" {

		instructionsList := strings.Split(instructionsInput, "\n")
		for _, instruction := range instructionsList {
			amount, from, to := parseInstruction(instruction)

			stacks.Move(amount, from, to)
		}
	}

	return stacks.TopmostElements()
}

func parseInstruction(instruction string) (amount int64, from int64, to int64) {

	split := strings.Split(instruction, " ")
	amountStr, fromStr, toStr := split[1], split[3], split[5]

	amount, _ = strconv.ParseInt(amountStr, 10, 64)
	from, _ = strconv.ParseInt(fromStr, 10, 64)
	to, _ = strconv.ParseInt(toStr, 10, 64)
	return amount, from, to
}

func sanitize(input string) string {

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimPrefix(input, "\n")
	return input
}

func buildStacksMatrix(input string) StacksMatrix {

	rowsList := strings.Split(input, "\n")
	rowsList = rowsList[:len(rowsList)-1]

	nOfStacks := inferNumberOfStacks(rowsList)

	stacks := make([]Stack, nOfStacks)
	for i := 0; i < len(stacks); i++ {
		stacks[i] = make(Stack, 0)
	}

	for nOfStack := 0; nOfStack < nOfStacks; nOfStack++ {

		letterOffset := 4*nOfStack + 1
		for rowIdx := len(rowsList) - 1; rowIdx >= 0; rowIdx-- {

			if string(rowsList[rowIdx][letterOffset]) != " " {

				stacks[nOfStack] = append(stacks[nOfStack], string(rowsList[rowIdx][letterOffset]))
			}
		}
	}

	return stacks
}

func inferNumberOfStacks(rowsList []string) int {

	// being very explicit here, bc it's arbitrary calculation
	arbitraryRow := rowsList[0] // arbitrary bc all are of equal len
	lengthOfRowByFour := float64(len(arbitraryRow)) / 4.0
	ceiledLenghtOfRowByFour := math.Ceil(lengthOfRowByFour)
	nOfStacks := int(ceiledLenghtOfRowByFour)
	return nOfStacks
}

type StacksMatrix []Stack

func (stacks StacksMatrix) TopmostElements() (topmostElements string) {

	for _, stack := range stacks {
		topmostElements += stack[len(stack)-1]
	}

	return
}

func (stacks *StacksMatrix) Move(crates, fromStack, toStack int64) {

	for i := 0; i < int(crates); i++ {

		from, to := stacks.At(fromStack), stacks.At(toStack)
		last := from.Unstack()
		if last != "" {
			to.Stack(last)
		}
	}
}

func (stacks *StacksMatrix) At(position int64) *Stack {

	matrix := *stacks
	return &matrix[position-1]
}

type Stack []string

func (stack *Stack) Unstack() (topmost string) {

	list := *stack

	topmost = list[len(list)-1]

	list = list[:len(list)-1]

	*stack = list

	return
}

func (stack *Stack) Stack(crate string) {

	list := *stack

	list = append(list, crate)

	*stack = list
}
