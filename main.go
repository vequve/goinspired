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
    if *numN>50{
        for  i:=0;i<*numN/50;i++{
            quote(numN,apiUrl)
        }
    }else{
       quote(numN,apiUrl)   
    }
   
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

            for  j := 0;j<*num;j++{
                fmt.Println(quotedata[j].Content)
                fmt.Println("-",quotedata[j].Author)
                if j>=49{
                    break
                }
      
            }
     
    return nil
}
