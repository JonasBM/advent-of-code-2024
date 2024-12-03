package helpers

import (
	"embed"
	"strings"
)

// func ReadFile(filePath string) []string {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return lines
// }

// ReadFile reads a file from the embedded filesystem and returns its lines
func ReadFile(embeddedFiles embed.FS, filePath string) []string {
	data := Must(embeddedFiles.ReadFile(filePath))
	lines := strings.Split(string(data), "\n")
	return lines
}

// Must panics if err is not nil, otherwise it returns the value.
// To make things easy and allow me to just have fun in AoC
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
