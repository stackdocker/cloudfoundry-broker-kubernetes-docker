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
    "fmt"
    "gopkg.in/redis.v3"
)

// The Sentinel address is orchestrated by Kubuernetes, e.g. $ kubectl get endpoints
// Should using your actual env value
func main() {
    client := redis.NewFailoverClient(&redis.FailoverOptions{
        MasterName: "mymaster",
        SentinelAddrs: []string{"172.31.33.2:26379","172.31.33.3:26379",
            "172.31.75.4:26379"},
    })
    cmd := client.Ping()
    if cmd != nil {
        fmt.Println(cmd)
    }
}
