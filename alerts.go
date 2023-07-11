package main

import "log"

type alert struct {
  ID int
  Subject string
  Message string
}

var alerts map[int]alert

func testAlerts() {
  for i := 1; i < 11; i++ {
    err := addAlert("Test Subject", "Test Message of some longer stuff and things")
    if err != nil {
      log.Println(err)
    }
  }
}

func addAlert(sub string, mes string) error {
  //will do database stuff
  id := len(alerts) + 1
  a := alert {
    ID: id,
    Subject: sub,
    Message: mes,
  }

  alerts[id] = a 
  return nil
}
