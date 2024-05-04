 package main 

import (
    "fmt"
    "encoding/json"
    "net/http"
    "flag"
    "strconv"
) 
type Quote struct{
    ID string `json:"_id"`
    Content string `json:"content"`
    Author string `json:"author"`
    Tags []string `json:"tags"`

}
func main(){
    numN := flag.Int("n", 1, "Numbers of quotes");
    flag.Parse()
    numNs := strconv.Itoa(*numN);
    apiUrl :=  "https://api.quotable.io/quotes/random?limit="+numNs
    if *numN<1{
        fmt.Println("Number of quotes needs to be positive! ")
        
    }
    quote(numN,apiUrl)    
  
}

func quote(num *int, apiUrl string) error{
    response, err := http.Get(apiUrl)
    
    if err != nil{
        return err
    }
    defer response.Body.Close()
    var quotedata []Quote 
    
    if err := json.NewDecoder(response.Body).Decode(&quotedata); err!=nil{
        return err
    }
    for i := 0;i<*num;i++{
        fmt.Println(quotedata[i].Content)
        fmt.Println("-",quotedata[i].Author) 
    } 
    return nil
}
