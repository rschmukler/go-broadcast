# go-broadcast

## Example

```go
package main

import "github.com/rschmukler/go-broadcast"

func main() {
  bc := broadcast.NewBroadcaster()
  
  wg := sync.WaitGroup{}
  wg.Add(5)
  
  for i := 0; i < 5; i++ {
    go func() {
      <-bc.Listen() // nil will come out of the channel
      wg.Done()
    }()
  }
  
  bc.Broadcast(nil)
  wg.Wait()
  
}
```
