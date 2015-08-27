// Names scores
// Problem 22
// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.
//
// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.
//
// What is the total of all the name scores in the file?

package main

import "fmt"
import "strings"
import "io/ioutil"

func main() {
    file, err := ioutil.ReadFile("./names.txt")

    // Reading file
    if err != nil {
        fmt.Println(err)
        return
    }

    // moving to an array
    names := strings.Split(string(file), ",")

    // Cleaning names
    s := make([]string, len(names))
    for i, name := range names {
        s[i] = strings.Replace(name, "\"", "", 2)
    }

    // Sorting
    names, s = sort(s), nil

    // todo: Names already cached and sorted, we need to do the calculation now
    fmt.Println(names)
}


// MergeSort ***************************************************

func sort(list []string) []string {
    if len(list) < 2 {
        return list[:]
    } else {
        return merge(sort(list[:len(list) / 2]), sort(list[len(list) / 2:]))
    }
}

func merge(list1 []string, list2 []string) []string {
    out := make([]string, len(list1) + len(list2))

    i, j, k := 0, 0, 0

    for ; i < len(list1) && j < len(list2); k += 1 {
        if strings.Compare(list1[i], list2[j]) == -1 {
            out[k] = list1[i]
            i += 1
        } else {
            out[k] = list2[j]
            j += 1
        }
    }

    for ; i < len(list1); i, k = i + 1, k + 1 {
        out[k] = list1[i]
    }

    for ; j < len(list2); j, k = j + 1, k + 1 {
        out[k] = list2[j]
    }

    return out
}
