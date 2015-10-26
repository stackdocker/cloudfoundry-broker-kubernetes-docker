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

package api

import (
    "net/http"
    "io"
    _ "bytes"
    _ "errors"
    _ "os"
    "log"
    "strings"
    //log "github.com/Sirupsen/logrus"
    "github.com/Sirupsen/logrus"
    "github.com/gorilla/mux"
    _ "github.com/gorilla/schema"
    "github.com/fatih/structs"
)

// Create a new instance of the logger. You can have any number of instances.
var logger *logrus.Logger

func init() {
    // do something here to set environment depending on an environment variable
    // or command-line flag
    if Environment == "debug" {
        //logger := log.New(os.Stdout, "logger: ", log.Lshortfile)
        log.SetPrefix("logger: ")
        log.SetFlags(log.Lshortfile)
    } else {
        logger = logrus.New()
        if Environment == "production" {
            // Log as JSON instead of the default ASCII formatter.
            //log.SetFormatter(&log.JSONFormatter{})
            logger.Formatter = &logrus.JSONFormatter{}
            // Output to stderr instead of stdout, could also be a file.
            //log.SetOutput(os.Stderr)
            
            // Only log the warning severity or above.
            //log.SetLevel(log.WarnLevel)
            logger.Level = logrus.WarnLevel
        } else {
            // The TextFormatter is default, you don't actually have to do this.
            //log.SetFormatter(&log.TextFormatter{})
            //log.SetOutput(os.Stderr)
            //log.SetLevel(log.InfoLevel)
            //logger.Formatter = &logrus.TextFormatter{}
            //logger.Level = log.InfoLevel
            //logger.Out = os.Stderr
        }
    }
}

/*
curl http://username:password@broker-url/v2/service_instances/:instance_id/service_bindings/:binding_id -d '{
  "plan_id":      "plan-guid-here",
  "service_id":   "service-guid-here",
  "app_guid":     "app-guid-here",
  "parameters":        {
    "parameter1-name-here": 1,
    "parameter2-name-here": "parameter2-value-here"
  }
}' -X PUT
*/

// for curl testing
const BINDING_ID = "00000000-0000-0000-0000-000000000000"

type RedisCredentials struct {
    Network    string       `json:"network,omitempty" schema:"network"`
    Address    string       `json:"address,omitempty" schema:"address"`
    Password   string       `json:"password,omitempty" schema:"password"`
    Db         int          `json:"db,omitempty" schema:"db"`
}

type SentinelCredentials struct {
    MasterName string       `json:"mastername,omitempty" schema:"mastername"`
    SentinelAddrs []string  `json:"sentineladdrs,omitempty" schema:"sentineladdrs"`
    Password   string       `json:"password,omitempty" schema:"password"`
    Db         int64        `json:"db,omitempty" schema:"db"`
}

var RedisConf = RedisCredentials{
        Network: "tcp",
        Address: ":6379",
        Password: "",
        Db: 0,
    }

var SentinelConf = SentinelCredentials{
        MasterName: "mymaster",
        SentinelAddrs: []string{},
        Password: "",
        Db: 0,
    }

type RedisBinder struct {
    *Binder
}

func (b *RedisBinder) Do(body io.ReadCloser) {
    b.Binder.Do(body)
    credentials := &RedisConf
    m := structs.Map(credentials)
    b.Binder.Bound.Credentials = make(map[string]interface{}, len(m))
    for k, v := range m {
        b.Binder.Bound.Credentials[strings.ToLower(k)] = v
    }
}

type SentinelBinder struct {
    *Binder
}

func (b *SentinelBinder) Do(body io.ReadCloser) {
    b.Binder.Do(body)
    credentials := &SentinelConf
    credentials.Password = RedisConf.Password
    credentials.Db = int64(RedisConf.Db)
    m := structs.Map(credentials)
    b.Binder.Bound.Credentials = make(map[string]interface{}, len(m))
    for k, v := range m {
        b.Binder.Bound.Credentials[strings.ToLower(k)] = v
    }
}

func (binder *RedisBinder) bind(w http.ResponseWriter, r *http.Request) {
    
}

func (binder *RedisBinder) unbind(w http.ResponseWriter, r *http.Request) {
    
}

func (binder *SentinelBinder) bind(w http.ResponseWriter, r *http.Request) {
    
}

func (binder *SentinelBinder) unbind(w http.ResponseWriter, r *http.Request) {
    
}

func HandleBinding(w http.ResponseWriter, r *http.Request) {
    log.Print(r.Method, " ", r.URL.Path)

    vars := mux.Vars(r)
    instanceId := vars["colon_instance_id"]
    bindingId := vars["conon_binding_id"]
    
    sentinelbinder := &SentinelBinder {
        Binder : NewBinder(instanceId, bindingId),
    }
    
    if r.Method == "PUT" {
        w.Header().Set("Content-Type", "application/json")
        sentinelbinder.Do(r.Body)
        if buf, err := sentinelbinder.Result(); err == nil {
            //w.Write(buf.Bytes)
            io.Copy(w, buf)        
        }
    } else if r.Method == "DELETE" {
        log.Print("not implemented")
    } else {
        log.Print("not implemented")
    }
    /*
    redisbinder := &RedisBinder {
        Binder : NewBinder(instanceId, bindingId),
    }
    
    if r.Method == "PUT" {
        w.Header().Set("Content-Type", "application/json")
        redisbinder.Do(r.Body)
        if buf, err := redisbinder.Result(); err == nil {
            //w.Write(buf.Bytes)
            io.Copy(w, buf)        
        }
    } else if r.Method == "DELETE" {
        log.Print("not implemented")
    } else {
        log.Print("not implemented")
    }
    */
}
