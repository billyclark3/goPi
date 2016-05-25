package main

import (
  "fmt"
  "math"
  "os"
  "strconv"
)

func main() {
  n, _ := strconv.Atoi(os.Args[1])
  fmt.Println("Pi, according to the Leibniz Series: ")
  fmt.Println(piLeibniz(n))
  fmt.Println();
  fmt.Println("Pi, according to the Euler Series: ")
  fmt.Println(piEuler(n))
}

func leibnizNode(ch chan float64, i float64) {
  ch <- 4 * math.Pow(-1, i) / (2*i + 1)
}

func piLeibniz(n int) float64 {
  ch := make(chan float64)
  for i := 0; i <= n; i++ {
    go leibnizNode(ch, float64(i))
  }
  sum := 0.0
  for i := 0; i <= n; i++ { // Leibniz Series starts from Node n=0
    sum += <-ch
  }
  return sum
}

func eulerNode(ch2 chan float64, i float64) {
  ch2 <- 1.0 / math.Pow(i, 2)
}

func piEuler(n int) float64 {
  ch2 := make(chan float64)
  for i := 1; i <= n; i++ {   // Euler Series starts from Node n=1
    go eulerNode(ch2, float64(i))
  }
  sum := 0.0
  for i := 1; i <= n; i++ {
    sum += <-ch2
  }
  return math.Sqrt(6 * sum)
}


