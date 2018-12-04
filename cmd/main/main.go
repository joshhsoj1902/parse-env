package main

import (
    "flag"
    "os"
    "strings"
    "encoding/json"
)

type env struct {
    Fullname   string       `json:"fullname"`
    Name       string       `json:"name"`
    Value      string       `json:"value"`
}

func main() {
    envsPtr := flag.String("env", "", "Envs to parse. (Required)")
    flag.Parse()

    if *envsPtr == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }

    var retEnvs []env

    envVars := os.Environ()

    for _, envVar := range envVars {
        v := strings.SplitN(envVar,"=", 2)
         if strings.Contains(v[0], *envsPtr)  {
            retEnvs = append(retEnvs, env{
                Fullname: v[0],
                Name: strings.ToLower(strings.Replace(v[0], *envsPtr, "", 1)),
                Value: v[1],
            })
         }
    }

    enc := json.NewEncoder(os.Stdout)
    enc.Encode(retEnvs)
}