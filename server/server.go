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

package server

import (
    "fmt"
    _ "os"
    "net/http"
    _ "errors"
    "log"
    "github.com/gorilla/mux"
    "api"
)

type Route struct {
    Path string
    HandleFn func(http.ResponseWriter, *http.Request)
    Methods []string
}

type Server struct {
    Port string
    Routes []Route
}

func NewServer() (*Server, error) {
    s := Server{
        Port: "80",
        Routes: make([]Route, 0, 6),
    }
    
    s.Routes = append(s.Routes, Route{"/v2/catalog", api.HandleCatalog, []string{"GET"}})
    s.Routes = append(s.Routes, Route{`/v2/service_instances/{colon_instance_id}`, 
        api.HandleServiceInstance, []string{"PUT", "PATCH", "DELETE"}})
    s.Routes = append(s.Routes, Route{`/v2/service_instances/{colon_instance_id}/service_bindings/{colon_binding_id}`, 
        api.HandleBinding, []string{"PUT", "DELETE"}}) 
    s.Routes = append(s.Routes, Route{"/sayhello", func (w http.ResponseWriter, r *http.Request){
            fmt.Fprintf(w, "hello world")
        }, []string{"GET"}})
    
    return &s, nil
}

func (s *Server) Start() {
    router := mux.NewRouter()
    
    for _, r := range s.Routes {
        router.HandleFunc(r.Path, r.HandleFn).Methods(r.Methods...)
    }

    http.Handle("/", router)
    
    fmt.Printf("Listening on port %s\n", s.Port)
    log.Fatal(http.ListenAndServe(":" + s.Port, nil))
}

