package screenshot

import (
	"fmt"
	"strconv"
)

func GenerateFilename(prefix string, counter int) string {
	counterStr := strconv.Itoa(counter)
	paddedCounter := fmt.Sprintf("%04s", counterStr) // Добавляем ведущие нули
	filename := fmt.Sprintf("%s_%s.png", prefix, paddedCounter)
	return filename
}
