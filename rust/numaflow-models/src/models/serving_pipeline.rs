/*
Copyright 2022 The Numaproj Authors.

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

// Code generated by Openapi Generator. DO NOT EDIT.

#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ServingPipeline {
    /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    #[serde(rename = "apiVersion", skip_serializing_if = "Option::is_none")]
    pub api_version: Option<String>,
    /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    #[serde(rename = "kind", skip_serializing_if = "Option::is_none")]
    pub kind: Option<String>,
    #[serde(rename = "metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<k8s_openapi::apimachinery::pkg::apis::meta::v1::ObjectMeta>,
    #[serde(rename = "spec")]
    pub spec: Box<crate::models::ServingPipelineSpec>,
    #[serde(rename = "status", skip_serializing_if = "Option::is_none")]
    pub status: Option<Box<crate::models::ServingPipelineStatus>>,
}

impl ServingPipeline {
    pub fn new(spec: crate::models::ServingPipelineSpec) -> ServingPipeline {
        ServingPipeline {
            api_version: None,
            kind: None,
            metadata: None,
            spec: Box::new(spec),
            status: None,
        }
    }
}
