package helpers

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Function to create a file with given content
func CreateFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// Function to check and create a directory if it doesn't exist
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Printf("Created directory: %s\n", dir)
	}
	return nil
}

// formate module names to camelCase format
func FormatModuleName(name string) string {
	// Split the name into words
	words := strings.Fields(name)

	// If there are no words, return an empty string
	if len(words) == 0 {
		return ""
	}

	words[0] = strings.ToLower(words[0])

	// Format words to start with an uppercase letter
	for i := 1; i < len(words); i++ {
		words[i] = cases.Title(language.English).String(words[i])
	}

	// Join the words back into a single string
	return strings.Join(words, "")
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

func LowercaseFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]+32) + s[1:]
}

//camel case to kebab case convertor

func CamelToKebab(input string) string {
	var result []rune

	for i, r := range input {
		// If the rune is an uppercase letter, add a hyphen before it
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, '-')
		}
		// Append the lowercase version of the rune
		result = append(result, unicode.ToLower(r))
	}

	return string(result)
}
