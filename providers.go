package main

const (
  azure = iota
  itglue
  sophos
)

type provider struct {
  id int
  ptype int
  key string
}

var providers map[int]provider

func addProvider(t int, key string) {
  
}
