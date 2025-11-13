package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func handleValidate(args []string) int {
	fs := flag.NewFlagSet("validate", flag.ExitOnError)
	stdin := fs.Bool("stdin", false, "Read numbers from stdin")
	fs.Parse(args)

	var numbers []string
	if *stdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}
			for _, num := range strings.Fields(line) {
				num = strings.ReplaceAll(num, " ", "")
				if num != "" {
					numbers = append(numbers, num)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
			return 1
		}
	} else {
		numbers = fs.Args()
		for i, num := range numbers {
			numbers[i] = strings.ReplaceAll(num, " ", "")
		}
	}

	if len(numbers) == 0 {
		fmt.Fprintln(os.Stderr, "No numbers provided")
		return 1
	}

	exitCode := 0
	for _, num := range numbers {
		if luhnCheck(num) {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			exitCode = 1
		}
	}
	return exitCode
}

func generateCombinations(pattern string) ([]string, error) {
	if pattern == "" {
		return nil, fmt.Errorf("Empty pattern")
	}
	if !strings.HasSuffix(pattern, "*") && strings.Contains(pattern, "*") {
		return nil, fmt.Errorf("Asterisks must be at the end")
	}
	stars := strings.Count(pattern, "*")
	if stars > 4 {
		return nil, fmt.Errorf("Up to 4 asterisks allowed")
	}
	if stars == 0 {
		if luhnCheck(pattern) {
			return []string{pattern}, nil
		}
		return nil, fmt.Errorf("Pattern is not a valid number")
	}
	prefix := strings.TrimSuffix(pattern, strings.Repeat("*", stars))
	for _, r := range prefix {
		if r < '0' || r > '9' {
			return nil, fmt.Errorf("Invalid character in pattern")
		}
	}
	var results []string
	for i := 0; i < intPow(10, stars); i++ {
		suffix := fmt.Sprintf("%0*d", stars, i)
		num := prefix + suffix
		if luhnCheck(num) {
			results = append(results, num)
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("No valid numbers found")
	}
	return results, nil
}

// intPow -> степень числа
func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func handleGenerate(args []string) int {
	fs := flag.NewFlagSet("generate", flag.ExitOnError)
	pick := fs.Bool("pick", false, "Pick a random valid number")
	fs.Parse(args)

	if len(fs.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Expected one argument: the pattern")
		return 1
	}
	pattern := fs.Args()[0]
	numbers, err := generateCombinations(pattern)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if *pick {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(numbers[rand.Intn(len(numbers))])
	} else {
		sort.Strings(numbers)
		for _, num := range numbers {
			fmt.Println(num)
		}
	}
	return 0
}

func handleInformation(args []string) int {
	fs := flag.NewFlagSet("information", flag.ExitOnError)
	brandsFile := fs.String("brands", "", "File with brands")
	issuersFile := fs.String("issuers", "", "File with issuers")
	stdin := fs.Bool("stdin", false, "Read numbers from stdin")
	fs.Parse(args)

	if *brandsFile == "" || *issuersFile == "" {
		fmt.Fprintln(os.Stderr, "Missing --brands or --issuers")
		return 1
	}

	brands, err := loadMap(*brandsFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading brands:", err)
		return 1
	}
	issuers, err := loadMap(*issuersFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading issuers:", err)
		return 1
	}

	var numbers []string
	if *stdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}
			numbers = append(numbers, strings.Fields(line)...)
		}
	} else {
		numbers = fs.Args()
	}

	if len(numbers) == 0 {
		fmt.Fprintln(os.Stderr, "No numbers provided")
		return 1
	}

	for _, num := range numbers {
		num = strings.ReplaceAll(strings.TrimSpace(num), " ", "")
		if num == "" {
			continue
		}

		fmt.Println(num)

		isValid := luhnCheck(num)
		fmt.Printf("Correct: %s\n", map[bool]string{true: "yes", false: "no"}[isValid])

		if isValid {
			fmt.Printf("Card Brand: %s\n", getBrand(num, brands))
			fmt.Printf("Card Issuer: %s\n", getIssuer(num, issuers))
		} else {
			fmt.Printf("Card Brand: -\n")
			fmt.Printf("Card Issuer: -\n")
		}

		if luhnCheck(num) {
			fmt.Printf("Card Brand: %s\n", getBrand(num, brands))
			fmt.Printf("Card Issuer: %s\n", getIssuer(num, issuers))
		} else {
			fmt.Printf("Card Brand: -\n")
			fmt.Printf("Card Issuer: -\n")
		}
	}
	return 0
}

func handleIssue(args []string) int {
	fs := flag.NewFlagSet("issue", flag.ExitOnError)
	brandsFile := fs.String("brands", "", "File with brands")
	issuersFile := fs.String("issuers", "", "File with issuers")
	brand := fs.String("brand", "", "Brand name")
	issuer := fs.String("issuer", "", "Issuer name")
	fs.Parse(args)

	if *brandsFile == "" || *issuersFile == "" || *brand == "" || *issuer == "" {
		fmt.Fprintln(os.Stderr, "Missing required flags")
		return 1
	}
	if len(fs.Args()) > 0 {
		fmt.Fprintln(os.Stderr, "Unexpected arguments:", fs.Args())
		return 1
	}

	brands, err := loadMap(*brandsFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading brands:", err)
		return 1
	}
	issuers, err := loadMap(*issuersFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading issuers:", err)
		return 1
	}

	var brandPrefix, issuerPrefix string
	for p, b := range brands {
		if strings.EqualFold(b, *brand) {
			brandPrefix = p
			break
		}
	}
	if brandPrefix == "" {
		fmt.Fprintln(os.Stderr, "Brand not found")
		return 1
	}
	for p, i := range issuers {
		if strings.EqualFold(i, *issuer) {
			issuerPrefix = p
			break
		}
	}
	if issuerPrefix == "" {
		fmt.Fprintln(os.Stderr, "Issuer not found")
		return 1
	}

	if !strings.HasPrefix(issuerPrefix, brandPrefix) {
		fmt.Fprintln(os.Stderr, "Issuer prefix does not match brand prefix")
		return 1
	}

	length := 16
	remaining := length - len(issuerPrefix)
	if remaining < 1 {
		fmt.Fprintln(os.Stderr, "Issuer prefix too long")
		return 1
	}

	rand.Seed(time.Now().UnixNano())

	attempts := 0
	for attempts < 1000 {
		attempts++

		suffix := ""
		for i := 0; i < remaining-1; i++ {
			suffix += strconv.Itoa(rand.Intn(10))
		}

		temp := issuerPrefix + suffix + "0"
		digits := make([]int, len(temp))
		for i, r := range temp {
			d, _ := strconv.Atoi(string(r))
			digits[i] = d
		}

		// *2 from the end
		var sum int
		for i := len(digits) - 2; i >= 0; i -= 2 {
			doubled := digits[i] * 2
			if doubled > 9 {
				doubled -= 9
			}
			digits[i] = doubled
		}
		for _, d := range digits {
			sum += d
		}

		checkDigit := (10 - sum%10) % 10
		num := issuerPrefix + suffix + strconv.Itoa(checkDigit)

		if luhnCheck(num) {
			fmt.Println(num)
			return 0
		}
	}

	fmt.Fprintln(os.Stderr, "Failed to generate valid card after 1000 attempts")
	return 1
}
