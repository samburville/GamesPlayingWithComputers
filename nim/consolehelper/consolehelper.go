package consolehelper

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// AskPlayerForInt ask a player a question and expect an integer response
func AskPlayerForInt(question string, reader *bufio.Reader) (int, error) {
	fmt.Println(question)
	userInput, _ := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\n", "", -1)

	fmt.Println()

	return strconv.Atoi(userInput)
}
