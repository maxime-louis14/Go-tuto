package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessLine Chercher les ancien mots pour les remplacer par les nouveau.
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line

	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

func FindReplaceFile(src string, old string, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()

	old = old + " "
	new = new + " "
	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += 0
			lines = append(lines, lineIdx)
		}

		fmt.Println(res)
		lineIdx++
	}
	return occ, lines, nil
}

func main() {
	old := "Go"
	new := "Python"
	occ, lines, err := FindReplaceFile("wikigo.txt", old, new)

	if err != nil {
		fmt.Printf("Error while executing find replace: %v\n", err)
	}

	fmt.Println("== Summary ==")
	defer fmt.Println(" End of Summary == ")
	fmt.Printf("Number of occurences of %v: %v\n", old, new)
	fmt.Printf("Number of lines %d\n", len(lines))

}
