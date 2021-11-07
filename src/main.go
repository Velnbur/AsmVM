package main

import (
    "bufio"
    "fmt"
    "github.com/Velnbur/AsmVM/src/machine"
    "log"
    "os"
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
        var tmp string
        tmp = scanner.Text()
        if tmp == "" {
            continue
        }
        lines = append(lines, tmp)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return lines
}

func startMachine(lines []string) {
    mach := machine.New()

    for _, line := range lines {
        mach.SetCommand(line)
        mach.PrintStatus()
        err := mach.HandleCommand()
        if err != nil {
            log.Fatal(err.Error())
        }
        mach.PrintStatus()
        fmt.Println()
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
