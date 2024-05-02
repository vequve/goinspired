package main 

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "os"
    "strconv"
) 

const defN = 1
const apiUrl ="https://api.quotable.io/random"


type Quote struct {
   Content string `json:"content"`
   Author string `json:"author"`
}
func main(){
    
    numN := defN
    if(len(os.Args)>1){    
        arg, err := strconv.Atoi(os.Args[1])
       if err != nil{
            fmt.Println(err)
        }else{
            numN = arg
        }
    }
    for i:=0;i<numN;i++{
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
    body, err := ioutil.ReadAll(response.Body)  
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
