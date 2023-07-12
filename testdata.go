package main

import "log"

func testAlerts() {
  for i := 1; i < 11; i++ {
    err := addAlert("Test Subject", "Test Message of some longer stuff and things")
    if err != nil {
      log.Println(err)
    }
  }
}
