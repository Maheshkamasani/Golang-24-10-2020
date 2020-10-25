package main

import "fmt"

func main() {
   var StateCapitalMap map[string]string
   /* create a map*/
   StateCapitalMap = make(map[string]string)
   
   /* insert key-value Amaravathi in the map*/
   StateCapitalMap["Andhra Pradesh"] = "Amaravati"
   StateCapitalMap["Telangana"] = "Hydrebad"
   StateCapitalMap["Tamilnadu"] = "Chennai"
   StateCapitalMap["Karnataka"] = "Bangloor"
   
   /* print map using keys*/
   for State := range StateCapitalMap {
      fmt.Println("Capital of",State,"is",StateCapitalMap[State])
   }
   
   /* test if entry is present in the map or not*/
   capital, ok := StateCapitalMap["Kerala"]
   
   /* if ok is true, entry is present otherwise entry is absent*/
   if(ok){
      fmt.Println("Capital of Kerala is", capital)  
   } else {
      fmt.Println("Capital of Kerala is not present") 
   }
}