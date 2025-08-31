package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetUserInput reads input from the console
func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
