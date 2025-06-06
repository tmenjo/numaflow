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

package fixtures

import (
	"os"
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/yaml"

	dfv1 "github.com/numaproj/numaflow/pkg/apis/numaflow/v1alpha1"
	flowpkg "github.com/numaproj/numaflow/pkg/client/clientset/versioned/typed/numaflow/v1alpha1"
)

type Given struct {
	t                     *testing.T
	isbSvcClient          flowpkg.InterStepBufferServiceInterface
	pipelineClient        flowpkg.PipelineInterface
	servingPipelineClient flowpkg.ServingPipelineInterface
	vertexClient          flowpkg.VertexInterface
	monoVertexClient      flowpkg.MonoVertexInterface
	isbSvc                *dfv1.InterStepBufferService
	pipeline              *dfv1.Pipeline
	servingPipeline       *dfv1.ServingPipeline
	monoVertex            *dfv1.MonoVertex
	restConfig            *rest.Config
	kubeClient            kubernetes.Interface
}

// ISBSvc creates an ISBSvc based on the parameter, this may be:
//
// 1. A file name if it starts with "@"
// 2. Raw YAML.
func (g *Given) ISBSvc(text string) *Given {
	g.t.Helper()
	g.isbSvc = &dfv1.InterStepBufferService{}
	g.readResource(text, g.isbSvc)
	l := g.isbSvc.GetLabels()
	if l == nil {
		l = map[string]string{}
	}
	l[Label] = LabelValue
	g.isbSvc.SetLabels(l)
	g.isbSvc.SetName(ISBSvcName)
	return g
}

// Pipeline creates a Pipeline based on the parameter, this may be:
//
// 1. A file name if it starts with "@"
// 2. Raw YAML.
func (g *Given) Pipeline(text string) *Given {
	g.t.Helper()
	g.pipeline = &dfv1.Pipeline{}
	g.readResource(text, g.pipeline)
	l := g.pipeline.GetLabels()
	if l == nil {
		l = map[string]string{}
	}
	l[Label] = LabelValue
	g.pipeline.SetLabels(l)
	g.pipeline.Spec.InterStepBufferServiceName = ISBSvcName
	return g
}

// ServingPipeline creates a Pipeline based on the parameter, this may be:
//
// 1. A file name if it starts with "@"
// 2. Raw YAML.
func (g *Given) ServingPipeline(text string) *Given {
	g.t.Helper()
	g.servingPipeline = new(dfv1.ServingPipeline)
	g.readResource(text, g.servingPipeline)
	l := g.servingPipeline.GetLabels()
	if l == nil {
		l = map[string]string{}
	}
	l[Label] = LabelValue
	g.servingPipeline.SetLabels(l)
	g.servingPipeline.Spec.Pipeline.InterStepBufferServiceName = ISBSvcName
	return g
}

// MonoVertex creates a MonoVertex based on the parameter, this may be:
//
// 1. A file name if it starts with "@"
// 2. Raw YAML.
func (g *Given) MonoVertex(text string) *Given {
	g.t.Helper()
	g.monoVertex = &dfv1.MonoVertex{}
	g.readResource(text, g.monoVertex)
	l := g.monoVertex.GetLabels()
	if l == nil {
		l = map[string]string{}
	}
	l[Label] = LabelValue
	g.monoVertex.SetLabels(l)
	return g
}

func (g *Given) WithPipeline(p *dfv1.Pipeline) *Given {
	g.t.Helper()
	g.pipeline = p
	l := g.pipeline.GetLabels()
	if l == nil {
		l = map[string]string{}
	}
	l[Label] = LabelValue
	g.pipeline.SetLabels(l)
	g.pipeline.Spec.InterStepBufferServiceName = ISBSvcName
	return g
}

func (g *Given) WithMonoVertex(mv *dfv1.MonoVertex) *Given {
	g.t.Helper()
	g.monoVertex = mv
	l := g.monoVertex.GetLabels()
	if l == nil {
		l = map[string]string{}
	}
	l[Label] = LabelValue
	g.monoVertex.SetLabels(l)
	return g
}

func (g *Given) readResource(text string, v metav1.Object) {
	g.t.Helper()
	var file string
	if strings.HasPrefix(text, "@") {
		file = strings.TrimPrefix(text, "@")
	} else {
		f, err := os.CreateTemp("", "numaflow-e2e")
		if err != nil {
			g.t.Fatal(err)
		}
		_, err = f.Write([]byte(text))
		if err != nil {
			g.t.Fatal(err)
		}
		err = f.Close()
		if err != nil {
			g.t.Fatal(err)
		}
		file = f.Name()
	}

	f, err := os.ReadFile(file)
	if err != nil {
		g.t.Fatal(err)
	}
	err = yaml.Unmarshal(f, v)
	if err != nil {
		g.t.Fatal(err)
	}
}

func (g *Given) When() *When {
	return &When{
		t:                     g.t,
		isbSvcClient:          g.isbSvcClient,
		pipelineClient:        g.pipelineClient,
		servingPipelineClient: g.servingPipelineClient,
		vertexClient:          g.vertexClient,
		monoVertexClient:      g.monoVertexClient,
		isbSvc:                g.isbSvc,
		pipeline:              g.pipeline,
		servingPipeline:       g.servingPipeline,
		monoVertex:            g.monoVertex,
		restConfig:            g.restConfig,
		kubeClient:            g.kubeClient,
	}
}
