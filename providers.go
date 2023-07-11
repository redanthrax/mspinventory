package main

const (
  azure = iota
  itglue
  sophos
)

type provider struct {
  ID int
  Type int
  Key string
}

var providers map[int]provider

func addProvider(t int, key string) {

}
