package main

/**
 * @Author elastic·H
 * @Date 2024-09-19
 * @File: fileRead.go
 * @Description:
 */

import (
	"bufio"
	"io"
	"os"
)

// ReadLines reads all lines of the file.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadLines2 reads all lines of the file.
func ReadLines2(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	for {
		// ReadString reads until the first occurrence of delim in the input,
		// returning a string containing the data up to and including the delimiter.
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			lines = append(lines, line)
			break
		}
		if err != nil {
			return lines, err
		}
		lines = append(lines, line[:len(line)-1])
	}
	return lines, nil
}

// ReadLines3 reads all lines of the file.
func ReadLines3(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	r := bufio.NewReader(f)
	for {
		// ReadLine is a low-level line-reading primitive.
		// Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return lines, err
		}
		lines = append(lines, string(bytes))
	}
	return lines, nil
}
