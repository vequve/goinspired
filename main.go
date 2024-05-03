package main 

import (
    "fmt"
    "io"
    "encoding/json"
    "net/http"
    "flag"
    //"strconv"
) 
const apiUrl = "https://api.quotable.io/random"
type Quote struct {
   Content string `json:"content"`
   Author string `json:"author"`
}
func main(){
    numN := flag.Int("n", 1, "Numbers of quotes");
    flag.Parse()
    if *numN<1{
        fmt.Println("Number of quotes needs to be positive! ")
        
    }
    for i:=0;i<*numN;i++{
        quote()
    }
  
}

func quote(){
    response, err := http.Get(apiUrl)
    
    if err != nil{
        fmt.Println(err)
        return
    }
    defer response.Body.Close()
   
    body, err := io.ReadAll(response.Body)  
    if err != nil{
        fmt.Println(err)
        return
    }
    
    var quotedata Quote

    err = json.Unmarshal(body, &quotedata)
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Println(quotedata.Content)
    fmt.Println("-",quotedata.Author) 
}
