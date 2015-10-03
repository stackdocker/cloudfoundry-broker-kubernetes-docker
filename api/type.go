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

type MetadataKey struct {
    Id              string    `json:",inline" yaml:",inline"`
    Name            string 
    Description     string 
}

type DashboardClientV2 struct {
    Id              string    `json:"id,omitempty" yaml:"id,omitempty"`
    Secret          string    `json:"secret,omitempty" yaml:"secret,omitempty"`
    RedirectUri     string    `json:"redirect_uri,omitempty" yaml:"redirect_uri,omitempty"`
}

/*
[catalog metadata documentation: plan meatadata fields]( http://docs.cloudfoundry.org/services/catalog-metadata.html )
               "metadata":{
                  "bullets":[
                     "20 GB of messages",
                     "20 connections"
                  ],
                  "costs":[
                     {
                        "amount":{
                           "usd":99.0,
                           "eur":49.0
                        },
                        "unit":"MONTHLY"
                     },
                     {
                        "amount":{
                           "usd":0.99,
                           "eur":0.49
                        },
                        "unit":"1GB of messages over 20GB"
                     }
                  ],
                  "displayName":"Big Bunny"
               }
*/

type PayAmount struct {
    Usd             float64             `json:"usd,omitempty" yaml:"usd,omitempty"`
    Eur             float64             `json:"eur,omitempty" yaml:"eur,omitempty"`
}

type ServiceCosts struct {
    Amount          PayAmount           `json:"amount,omitempty" yaml:"amount,omitempty"`
    Unit            string              `json:"unit,omitempty" yaml:"unit,omitempty"`
}

type PlanMetadata struct {
    Bullets         []string            `json:"free,omitempty" yaml:"free,omitempty"`
    Costs           ServiceCosts        `json:"costs,omitempty" yaml:"costs,omitempty"`
    DisplayName     string              `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}

type ServicePlanV2 struct {
    Id              string              `json:"id" yaml:"id"`
    Name            string              `json:"name" yaml:"name"`
    Description     string              `json:"description" yaml:"description"`
    Metadata        PlanMetadata        `json:"metadata" yaml:"metadata"`
    Free            bool                `json:"free,omitempty" yaml:"free,omitempty"`
}

/*
[catalog metadata documentation: service meatadata fields]( http://docs.cloudfoundry.org/services/catalog-metadata.html )
         "metadata":{
            "displayName":"CloudAMQP",
            "imageUrl":"https://d33na3ni6eqf5j.cloudfront.net/app_resources/18492/thumbs_112/img9069612145282015279.png",
            "longDescription":"Managed, highly available, RabbitMQ clusters in the cloud",
            "providerDisplayName":"84codes AB",
            "documentationUrl":"http://docs.cloudfoundry.com/docs/dotcom/marketplace/services/cloudamqp.html",
            "supportUrl":"http://www.cloudamqp.com/support.html"
         },
*/

type ServiceMetadata struct {
    DisplayName         string              `json:"displayName,omitempty" yaml:"displayName,omitempty"`
    ImageUrl            string              `json:"imageUrl,omitempty" yaml:"ImageUrl,omitempty"`
    LongDescription     string              `json:"logDescription,omitempty" yaml:"longDescription,omitempty"`
    ProviderDisplayName string              `json:"providerDisplayName,omitempty" yaml:"providerDisplayName,omitempty"`
    DocumentationUrl    string              `json:"documentationUrl,omitempty" yaml:"DocumentationUrl,omitempty"`
    SupportUrl          string              `json:"supportUrl,omitempty" yaml:"supportUrl,omitempty"`
}

type ServiceV2 struct {
    Id              string                  `json:"id" yaml:"id"`
    Name            string                  `json:"name" yaml:"name"`
    Description     string                  `json:"description" yaml:"description"`
    Bindable        bool                    `json:"bindable" yaml:"bindable"`
    Tags            []string                `json:"tags,omitempty" yaml:"tags,omitempty"`
    Metadata        []ServiceMetadata       `json:"metadata,inline" yaml:"metadata,inline"`
    Requires        string                  `json:"requires,omitempty" yaml:"requires,omitempty"`
    PlanUpdatable   bool                    `json:"plan_updateable,omitempty" yaml:"plan_updateable,omitempty"`
    Plans           []ServicePlanV2         `json:"plans" yaml:"plans"`
    DashboardClient DashboardClientV2       `json:"dashboard_client,omitempty" yaml:"dashboard_client,omitempty"`  
}

type CatalogV2 struct {
    Services        []ServiceV2             `json:"services" yaml:"services"`
}