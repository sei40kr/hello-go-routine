package main

import (
  "time"
  "log"
  "sync"
)

func sleepForSeconds(seconds int, wg *sync.WaitGroup) {
  log.Println("Running routine to sleep for", seconds, "sec.")
  time.Sleep(time.Duration(seconds) * time.Second);
  log.Println("==> Ended a sleep for", seconds, "sec.");
  wg.Done()
}

func main() {
  const N = 3

  wg := sync.WaitGroup{}
  start := time.Now()

  for i := 0; i < N; i++ {
    // Don't call wg.Add in routine.
    // It may cause an error to call wg.Wait without other working threads.
    wg.Add(1)
    seconds := i + 1
    go sleepForSeconds(seconds, &wg)
  }

  wg.Wait()

  end := time.Now()

  log.Printf("All routines ended in %.3f sec.\n", end.Sub(start).Seconds())
}
