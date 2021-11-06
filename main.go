package main

import (
    "os"
    "log"
    "bufio"
    "github.com/Velnbur/AsmVM/machine"
)

func ParseSrc(path string) []string {
    var lines []string
    srcFile, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }

    defer srcFile.Close()

    scanner := bufio.NewScanner(srcFile)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return lines
}

func startMachine(lines []string) {
    machine := machine.New()

    for _, line := range lines {
        err := machine.HandleCommand(line)
        if err != nil {
            log.Fatal(err.Error())
        }
        machine.PrintStatus()
    }
}

func main() {
    if len(os.Args) < 2 {
        panic("No path to file")
    }

    pathToFile := os.Args[1]
    lines := ParseSrc(pathToFile)
    startMachine(lines)
}
