package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func parseShipFile() []string {
	file, err := os.Open(".ship")
	if err != nil {
		log.Fatal(".ship file does not exist")
	}
	defer file.Close()

	var output []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}

	return output
}

func output(line []byte) {
	t := time.Now()
	// Format based on default reference time "Mon Jan 2 15:04:05 MST 2006"
	timeString := t.Format("15:04:05")
	fmt.Printf("%s %s\n", timeString, line)
}

func scan(pipe io.ReadCloser) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		s := scanner.Bytes()
		output(s)
	}

	err := scanner.Err()
	if err != nil {
		panic(err)
	}
}

func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic("os.Getwd failed")
	}

	return dir
}

func printTimeElapsed(start time.Time) {
	elapsed := time.Since(start)
	line := fmt.Sprintf("Shipped in %s seconds", elapsed)
	output([]byte(line))
}

func runShipCommands(commands []string) {
	startTime := time.Now()

	for _, command := range commands {
		line := fmt.Sprintf("Executing %s", command)
		output([]byte(line))
		cmd := exec.Command("sh", "-c", command)
		cmd.Dir = getCurrentDirectory()
		cmd.Env = os.Environ()

		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			panic("stdoutpipe failed")
		}

		stderrPipe, err := cmd.StderrPipe()
		if err != nil {
			panic("stderrpipe failed")
		}

		err = cmd.Start()
		if err != nil {
			fmt.Println(err)
			panic("exec failed")
		}

		// We need this to run in go-routines otherwise it would listen
		// sequentially, and we want to see the output right away
		go scan(stderrPipe)
		go scan(stdoutPipe)

		err = cmd.Wait()
		if err != nil {
			fmt.Printf("Wait error: %#s\n", err)
		}
	}

	printTimeElapsed(startTime)
}

func main() {
	shipCommands := parseShipFile()
	runShipCommands(shipCommands)
}
