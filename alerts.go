package main

type Alert struct {
  ID int
  Subject string
  Message string
}

var alerts map[int]Alert

func testAlerts() {
  for i := 1; i < 11; i++ {
    alerts[i] = Alert {
      ID: i,
      Subject: "Test Subject",
      Message: "Test Message of some longer stuff and things",
    }
  }
}
