package main
import (
"fmt"
"log"
"github.com/fsnotify"
)

func main() {
scriptDone   := make(chan bool)
dirSpy , err :=  fsnotify.NewWatcher()
if err != nil {
log.Fatal(err)
}



go func () {
for {
select {
case fileChange :=  <-dirSpy.Event :
   log.Println("event for file: ",fileChange)
case err := <-dirSpy.Error:
    log.Println("Error happened ", err)

}}
}()
err = dirSpy.Watch("/mnt/hsm_asset_bin")
if err != nil {
fmt.Println(err)
}
<-scriptDone

dirSpy.Close()}
