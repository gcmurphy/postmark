package main

import (
    "github.com/gcmurphy/postmark"
    "os"
    "fmt"
)

func main(){

    apiKey := os.Getenv("POSTMARK_API")
    if len(apiKey) <= 0 {
        fmt.Println("Set the POSTMARK_API environment varible to run this demo")
        os.Exit(1)
    }
    p := postmark.NewPostmark(os.Getenv("POSTMARK_API"))
    r, e := p.Send(&postmark.Message{
                From: "gm@bleepit.net",
                To: "grantcmurphy@gmail.com",
                Subject: "Test postmark",
                TextBody: "Hello World!"})

    if e != nil {
        fmt.Println("ERROR: ", e.String())
        os.Exit(1)
    }
    fmt.Println("Response :" , r.String())
}
