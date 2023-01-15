package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

   fileContent, err := os.Open("users.json")

   if err != nil {
      log.Fatal(err)
      return
   }

   fmt.Println("The File is opened successfully...")

   defer fileContent.Close()

   byteResult, _ := ioutil.ReadAll(fileContent)

   var res map[string][]interface{} //taking an array of struct so that it can be iterate over
   json.Unmarshal(byteResult, &res)
  
   val:=res["users"] //defining the key name, which is in json file
   for _, item := range val {
       myMap := item.(map[string]interface{}) //converting the item iterface in map
      for key,value:=range myMap{ //iterating over the 
         fmt.Println(key,value)
      }
   }
}












