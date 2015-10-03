/*
Copyright 2014 The Kubernetes Authors All rights reserved.
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
    _ "regexp"
    "encoding/json"
    _ "path"
    "github.com/gorilla/mux"
    "fmt"
    "io"
    "bytes"
)

var emptyJO []byte = []byte("{}")

type ServiceCatalog struct {
    instance    ServiceInstance
}

const (
    ORGID = "a924baea-363a-457e-8708-0eda152e76c5"
    SPACEID = "1fde0510-cc7e-42c0-8321-1505d51f0c29"
    INSTANCEID = "00000000-0000-0000-0000-000000000000"
    SERVICEID = "9c372bbc-1e7b-472b-bcb6-eeda5b21eb35"
    PLANID = "8cfbbaf5-efdb-41c1-89ab-f797185f7818"
)

/*
curl -i http://127.0.0.1:8080/v2/service_instances/00000000-0000-0000-0000-000000000000 -d '{"organization_guid": "a924baea-363a-457e-8708-0eda152e76c5", "plan_id": "8cfbbaf5-efdb-41c1-89ab-f797185f7818", "service_id": "9c372bbc-1e7b-472b-bcb6-eeda5b21eb35", "space_guid": "1fde0510-cc7e-42c0-8321-1505d51f0c29", "parameters": {"parameter1": 1, "parameter2": "value"}}' -X PUT -H "X-Broker-API-Version: 2.6" -H "Content-Type: application/json"; echo
curl http://username:password@broker-url/v2/service_instances/00000000-0000-0000-0000-000000000000 -d '{ 
  "organization_guid": "a924baea-363a-457e-8708-0eda152e76c5", 
  "plan_id":           "8cfbbaf5-efdb-41c1-89ab-f797185f7818", 
  "service_id":        "9c372bbc-1e7b-472b-bcb6-eeda5b21eb35", 
  "space_guid":        "1fde0510-cc7e-42c0-8321-1505d51f0c29", 
  "parameters":        { 
    "parameter1": 1, 
    "parameter2": "value" 
  } 
}' -X PUT -H "X-Broker-API-Version: 2.6" -H "Content-Type: application/json"; echo
*/

func (h *ServiceCatalog) Provision(w http.ResponseWriter, r *http.Request) {
    h.instance.Id = mux.Vars(r)["colon_instance_id"]

    dec := json.NewDecoder(r.Body)
    for {
        if err := dec.Decode(&h.instance); err == io.EOF {
            break
        } else if err != nil {
            http.Error(w, "Failed to decode request body", http.StatusInternalServerError)
            return
        }
    }
    
    if h.instance.Id == INSTANCEID {
        //fmt.Println("testing: ", h.instance.Id, " ", h.instance.InstancePlan.ServiceId, " ",
        //    h.instance.InstancePlan.PlanId, " ", h.instance.OrganizationGuid, " ", h.instance.SpaceGuid)
        fmt.Println("testing: ", h.instance.Id, " ", h.instance.ServiceId, " ",
            h.instance.PlanId, " ", h.instance.OrganizationGuid, " ", h.instance.SpaceGuid)
        for k, v := range h.instance.Parameters {
            fmt.Println (k, " ", v)
        }
    }
    
    w.Header().Set("Content-Type", "application/json") 
    //b, err := json.Marshal()
	//if err != nil {
	//    http.Error(w, "Failed to marshal empty json object", http.StatusInternalServerError)
	//}
	w.Write(emptyJO)
}

func (h *ServiceCatalog) UpdateInstance(w http.ResponseWriter, r *http.Request) {
    //enc := json.NewEncoder(w)
    //enc.Encode(catalog)
    js, err := json.Marshal(catalog)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    //w.Write(js)
    io.Copy(w, bytes.NewBuffer(js))        
}

func (h *ServiceCatalog) DeleteInstance(w http.ResponseWriter, r *http.Request) {
    
}

func HandleServiceInstance(w http.ResponseWriter, r *http.Request) {
    /*
    validMethod := regexp.MustCompile(`PUT|PATCH|DELETE`)
    if !validMethod.MatchString(r.Method) {
        http.Error(w, "Api only support PUT|PATCH|DELETE method.", http.StatusMethodNotAllowed)
        return
    }
    */
    if v := r.Header["X-Broker-API-Version"]; len(v) > 0 && v[0] != "2.6" {
        http.Error(w, "Unmatched API version.", http.StatusPreconditionFailed)
        return
    }
    
    if !catalog.authenticate(w, r) {
        http.Error(w, "Not authorized", http.StatusUnauthorized)
        return
    }
    
    var serviceCatalog ServiceCatalog
    
    switch r.Method {
    case "PUT":
        serviceCatalog.Provision(w, r)
    case "PATCH":
        serviceCatalog.UpdateInstance(w, r)
    case "DELETE":
        serviceCatalog.DeleteInstance(w, r)
    default:
        http.Error(w, "Api only support PUT|PATCH|DELETE method.", http.StatusMethodNotAllowed)
    }
}