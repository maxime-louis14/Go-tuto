package main

import (

	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessLine recherche l'ancien dans la ligne pour le remplacer par le nouveau.
// Il retourne found = true si le motif a été trouvé, res avec la chaîne de caractères résultante.
// et occ avec le nombre d'occurrences de old.

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

// FindReplaceFile recherche et remplace l'ancien modèle par le nouveau dans le fichier src.
// Il stocke le contenu dans le fichier dst.
// Il retourne le nombre d'occurrences de l'ancien motif, une tranche de lignes et une erreur s'il y en a une.
func FindReplaceFile(src string, dst string, old string, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, nil, err
	}
	defer srcFile.Close()

	// ouvrir le fichier dst
	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, nil, err
	}
	defer dstFile.Close()

	old = old + " "
	new = new + " "
	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}
		fmt.Fprintln(writer, res)
		lineIdx++
	}
	return occ, lines, nil
}

func main() {
	old := "Go"
	new := "Python"
	occ, lines, err := FindReplaceFile("wikigo.txt", "wikipython.txt", old, new)
	if err != nil {
		fmt.Printf("Error while executing find replace: %v\n", err)
		return
	}

	fmt.Println("== Summary ==")
	defer fmt.Println("== End of Summary ==")
	fmt.Printf("Number of occurrences of %v: %v\n", old, occ)
	fmt.Printf("Number of lines: %d\n", len(lines))
	fmt.Print("Lines: [ ")
	len := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < len-1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ]")

}
