package main

type alert struct {
  ID int
  Subject string
  Message string
}

var alerts map[int]alert

func addAlert(sub string, mes string) error {
  if alerts == nil {
    alerts = make(map[int]alert)
  }
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
