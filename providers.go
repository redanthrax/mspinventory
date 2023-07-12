package main

const (
  azure = iota
  itglue
  sophos
)

type provider struct {
  id int
  logo string
  parameters interface{}
}

