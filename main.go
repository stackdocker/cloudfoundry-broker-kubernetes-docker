/*
Copyright 2015 All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
    _ "fmt"
    "os"
    _ "errors"
    _ "flag"
    "github.com/spf13/pflag"
    "api"    
    "server"
)

func main() {
    pflag.StringVar(&api.RedisConf.Network, "redis-net", "tcp", "tcp or unix")
    pflag.StringVar(&api.RedisConf.Address, "redis-addr", ":6379", "ip:port")
    pflag.StringVar(&api.RedisConf.Password, "redis-pass", "", "Redis password")
    pflag.IntVar(&api.RedisConf.Db, "redis-db", 0, "Redis db")
    pflag.StringVar(&api.SentinelConf.MasterName, "master-name", "mymaster", 
        "Redis Sentinel master name")
    pflag.StringSliceVar(&api.SentinelConf.SentinelAddrs, "sentinel-addrs", 
        []string{":26379"}, "Sentinel failover addresses")
    pflag.Parse()

    port := os.Getenv("PORT")
    s, err := server.NewServer()
    if err != nil {
        return
    }
    if port != "" {
        s.Port = port
    }
    s.Start()
}