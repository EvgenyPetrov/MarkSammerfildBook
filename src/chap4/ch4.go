package main

import (
    "fmt"
    "strings"
    "sort"
)

func main() {
    var ins = []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
    ins2 := make([]int, len(ins))
    ins2 = UniqueInts(ins)
    fmt.Printf("%v\n%v\n", ins, ins2)
    
    irregularMatrix := [][]int{{1, 2, 3, 4},
                               {5, 6, 7, 8},
                               {9, 10, 11},
                               {12, 13, 14, 15},
                               {16, 17, 18, 19, 20}}
    slice := Flatten(irregularMatrix)
    fmt.Printf("1x%d: %v\n", len(slice), slice)
    
    var ins3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
    fmt.Printf("3: %v\n", Make2D(ins3, 3))
    fmt.Printf("4: %v\n", Make2D(ins3, 4))
    fmt.Printf("5: %v\n", Make2D(ins3, 5))
    fmt.Printf("6: %v\n", Make2D(ins3, 6))
    
    iniData := []string{
        "; Cut down copy of Mozilla application.ini file",
        "",
        "[App]",
        "Vendor=Mozilla",
        "Name=Iceweasel",
        "Profile=mozilla/firefox",
        "Version=3.5.16",
        "[Gecko]",
        "MinVersion=1.9.1",
        "MaxVersion=1.9.1.*",
        "[XRE]",
        "EnableProfileMigrator=0",
        "EnableExtensionManager=1",
    }
    
    ini := ParseIni(iniData)
    PrintIni(ini)
}

func ParseIni(ini []string) map[string]map[string]string {
    const separator = "="
    group := "General"
    result := make(map[string]map[string]string)
    
    for _, line := range ini {
        line = strings.TrimSpace(line)
        if strings.HasPrefix(line, ";") || line == "" {
            continue
        }
        if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
            fkey := line[1:len(line)-1]
            if _, found := result[fkey]; !found {
                result[fkey] = make(map[string]string)
                group = fkey
            }
        } else if ind := strings.Index(line, separator); ind > -1 {
            key := line[0:ind]
            value := line[ind+1:]
            if _, found := result[group]; !found {
                result[group] = make(map[string]string)
            }
            result[group][key] = " " + value
        }
    }
    return result
}

func Make2D(slice []int, columns int) [][]int {
    matrix := make([][]int, neededRows(slice, columns))
    for i, x := range slice {
        row := i / columns
        column := i % columns
        if matrix[row] == nil {
            matrix[row] = make([]int, columns)
        }
        matrix[row][column] = x
    }
    return matrix
}

func neededRows(slice []int, columns int) int {
    rows := len(slice) / columns
    if len(slice)%columns != 0 {
        rows++
    }
    return rows
}

func Flatten(input [][]int) (result []int) {
    for _, elemi := range input {
        for _, elemj := range elemi {
            result = append(result, elemj)
        }
    }
    return result
}

func UniqueInts(input []int) (result []int) {
    for _, elem := range input {
        if !stringInSlice(elem, result) {
            result = append(result, elem)
        }
    }
    return result
}

func stringInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func PrintIni(ini map[string]map[string]string) {
    groups := make([]string, 0, len(ini))
    for group := range ini {
        groups = append(groups, group)
    }
    sort.Strings(groups)
    for i, group := range groups {
        fmt.Printf("[%s]\n", group)
        keys := make([]string, 0, len(ini[group]))
        for key := range ini[group] {
            keys = append(keys, key)
        }
        sort.Strings(keys)
        for _, key := range keys {
            fmt.Printf("%s=%s\n", key, ini[group][key])
        }
        if i+1 < len(groups) {
            fmt.Println()
        }
    }
}
