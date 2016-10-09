package main

import (
  "fmt"
  "bufio"
  "os"
  "regexp"
  "strconv"
)


func main() {
  guide := `
-----------------Simple calculator---------------
  - 가능한 연산: +, -, x, /          
  - 프로그램 종료를 원하면 exit 또는 ctrl+C 입력 
-------------------------------------------------
`
  fmt.Println(guide)

  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Println("계산을 입력하시오.")
    text, _ := reader.ReadString('\n')
    text = text[:len(text)-1]
    if len(text) < 3 {
      continue
    }

    if len(text) > 4 && text[:4] == "exit" {
      fmt.Println("계산기 프로그램을 종료합니다.")
      break 
    }

    var digit = regexp.MustCompile(`[0-9]`)
    var non_valid_input = regexp.MustCompile(`[A-Za-z!@#$%^&()_~=]`)
    if non_valid_input.MatchString(text) == true || 
       digit.MatchString(text[len(text)-1:]) == false {
      fmt.Println("잘못된 입력입니다.")
      continue
    }
    for i := 0; i < len(text)-1; i++ {
      if (text[i:i+1] == "+" || text[i:i+1] == "-" || text[i:i+1] == "*" || text[i:i+1] == "/") && (text[i+1:i+2] == "+" || text[i+1:i+2] == "-" || text[i+1:i+2] == "*" || text[i+1:i+2] == "/") {
        fmt.Println("잘못된 입력입니다.")
        return
      }
    }

    // separate symbols and numbers
    var symbol_list []string
    var num_list []string
    start := 0
    for i := 0; i < len(text); i++ {
      if digit.MatchString(text[i:i+1]) == true {
        if i+1 == len(text) {
          num_list = append(num_list, text[start:i+1])
        } else {
          continue
        }
      }
      if text[i:i+1] == "-" && (i == 0 || text[i-1:i] == "+" || text[i-1:i] == "-" || text[i-1:i] == "*" || text[i-1:i] == "/") {
        continue
      } 

      if i > 0 && text[i-1:i] != "-" && digit.MatchString(text[i-1:i]) == true && (text[i:i+1] == "+" || text[i:i+1] == "-" || text[i:i+1] == "*" || text[i:i+1] == "/") {
        symbol_list = append(symbol_list, text[i:i+1])
        num_list = append(num_list, text[start:i])
        start = i+1
      }
    }
    //fmt.Println(symbol_list)
    //fmt.Println(num_list)

    // convert to postfix
    var stack []string
    var postfix []string
    for i := 0; i < len(symbol_list); i++ {
      postfix = append(postfix, num_list[i])

      for ; len(stack) != 0; {
        if (stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/") && (symbol_list[i] == "+" || symbol_list[i] == "-") {
          postfix = append(postfix, stack[len(stack)-1])
          stack = stack[:len(stack)-1]
        } else {
          break
        }
      }
      stack = append(stack, symbol_list[i])
    }

    postfix = append(postfix, num_list[len(num_list)-1])
    for ; len(stack) != 0; {
      postfix = append(postfix, stack[len(stack)-1])
      stack = stack[:len(stack)-1]
    }


    // postfix calculation
    for i := 0; i < len(postfix); i++ {
      if postfix[i] == "+" || postfix[i] == "-" || postfix[i] == "*" || postfix[i] == "/" {
        num_2, _ := strconv.Atoi(stack[len(stack)-1])
        num_1, _ := strconv.Atoi(stack[len(stack)-2])
        stack = stack[:len(stack)-2]
        if postfix[i] == "+" {
          stack = append(stack, strconv.Itoa(num_1 + num_2))
        }
        if postfix[i] == "-" {
          stack = append(stack, strconv.Itoa(num_1 - num_2))
        }
        if postfix[i] == "*" {
          stack = append(stack, strconv.Itoa(num_1 * num_2))
        }
        if postfix[i] == "/" {
          stack = append(stack, strconv.Itoa(num_1 / num_2))
        }
      } else {
        stack = append(stack, postfix[i])
      }
    }
    results, _ := strconv.Atoi(stack[0])
    fmt.Printf("답: %d \n\n", results)
  }
}

