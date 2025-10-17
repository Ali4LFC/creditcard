# Credit Card Tool

This is a Go-based implementation of a credit card validation and generation tool for Alem School Foundation project. It simulates operations on credit card numbers based on Luhn's Algorithm and specific rules for brands and issuers, with support for command-line inputs and flags.

## Features

- **Validation**: Checks if a credit card number is valid using Luhn's Algorithm.
  - Minimum 13 digits.
  - Outputs "OK" for valid, "INCORRECT" for invalid.
  - Supports multiple numbers and --stdin flag.
- **Generation**: Generates possible card numbers by replacing asterisks (*) with digits.
  - Up to 4 asterisks at the end.
  - Outputs in ascending order or picks a random one with --pick.
- **Information**: Provides details about the card brand and issuer based on prefixes.
  - Uses brands.txt and issuers.txt files.
  - Outputs validity, brand, and issuer.
  - Supports multiple numbers and --stdin.
- **Issue**: Generates a random valid credit card number for a specified brand and issuer.
  - Uses brands.txt and issuers.txt.
- **Flags** (depending on command):
  - --stdin: Read input from standard input.
  - --pick: Pick a random generated number.
  - --brands=X: Path to brands.txt file.
  - --issuers=X: Path to issuers.txt file.
  - --brand=X: Brand name (e.g., VISA).
  - --issuer=X: Issuer name (e.g., "Kaspi Gold").

## Installation

1. Ensure you have [Go](https://golang.org/doc/install) installed (version 1.22 or later recommended).
2. Clone this repository:
   ```bash
   git clone https://platform.alem.school/git/altaberkh/creditcard.git
   cd creditcard

Build the executable:
bashgo build -o creditcard .

Format the code (required for submission):
bashgo install mvdan.cc/gofumpt@latest
gofumpt -w .

Run with commands (examples below).

Usage Examples
Validate
bash./creditcard validate "4400430180300003"  # Output: OK
echo "4400430180300003" | ./creditcard validate --stdin  # Output: OK
Generate
bash./creditcard generate "440043018030****"  # List of valid numbers
./creditcard generate --pick "440043018030****"  # Random valid number
Information
bash./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
# Output:
# 4400430180300003
# Correct: yes
# Card Brand: VISA
# Card Issuer: Kaspi Gold
Issue
bash./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"  # Random valid number
