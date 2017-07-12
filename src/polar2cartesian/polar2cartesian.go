package main

import (
    "runtime"
    "fmt"
    "math"
    "bufio"
    "os"
)

type polar struct {
    radius float64
    O      float64
}

type cartesian struct {
    x   float64
    y   float64
}

var promt = "Enter a radius and angle (in degrees), e.g., 12.5 90, or %s to quit."
const result = "Polar radius=%.02f q=%.02f° ® Cartesian x=%.02f y=%.02f\n"

func init()  {
    if runtime.GOOS == "windows" {
        promt = fmt.Sprintf(promt, "Ctrl+Z, Enter")
    } else {
        promt = fmt.Sprintf(promt, "Ctrl+D")
    }
}

func main()  {
    questions := make(chan polar)
    defer close(questions)
    
    answers := createSolver(questions)
    defer close(answers)
    
    interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
    answers := make(chan cartesian)
    go func() {
        for {
            polarCoords := <-questions
            O := polarCoords.O * math.Pi / 180.0
            x := polarCoords.radius * math.Cos(O)
            y := polarCoords.radius * math.Sin(O)
            answers <- cartesian{x, y}
        }
    }()
    return answers
}

func interact(questions chan polar, answers chan cartesian) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println(promt)
    for {
        fmt.Printf("Radius and angle: ")
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        var radius, O float64
        if _, err := fmt.Sscanf(line, "%f %f", &radius, &O); err != nil {
            fmt.Fprintln(os.Stderr, "invalid input")
            continue
        }
        questions <- polar{radius, O}
        coord := <-answers
        fmt.Printf(result, radius, O, coord.x, coord.y)
    }
    fmt.Println()
}
