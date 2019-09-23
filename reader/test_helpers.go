package reader

import (
	"bufio"
	"os"
)

func setupFile(testPath string, lines []string) {
	file, _ := os.Create(testPath)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line)
		writer.WriteString("\n")
	}
	writer.Flush()
}
