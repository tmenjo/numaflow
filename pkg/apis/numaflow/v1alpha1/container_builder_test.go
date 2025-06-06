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

package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/ptr"
)

var (
	testResources = corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			"cpu":    resource.MustParse("800m"),
			"memory": resource.MustParse("256Mi"),
		},
		Requests: corev1.ResourceList{
			"cpu":    resource.MustParse("100m"),
			"memory": resource.MustParse("64Mi"),
		},
	}
)

func Test_containerBuilder(t *testing.T) {
	c := containerBuilder{}.
		init(getContainerReq{
			resources: testResources,
		}).args("numa", "args").
		image("image").
		imagePullPolicy(corev1.PullIfNotPresent).
		command("cmd").
		appendVolumeMounts(corev1.VolumeMount{Name: "vol", MountPath: "/vol"}).
		appendEnv(corev1.EnvVar{
			Name: "env", Value: "value"}).
		appendPorts(corev1.ContainerPort{Name: "port", ContainerPort: 8080}).
		asSidecar().
		build()
	assert.Equal(t, "numa", c.Name)
	assert.Len(t, c.VolumeMounts, 1)
	assert.Equal(t, testResources, c.Resources)
	assert.Equal(t, []string{"numa", "args"}, c.Args)
	assert.Equal(t, "image", c.Image)
	assert.Equal(t, corev1.PullIfNotPresent, c.ImagePullPolicy)
	assert.Equal(t, []corev1.EnvVar{{Name: "env", Value: "value"}}, c.Env)
	assert.Equal(t, []corev1.ContainerPort{{Name: "port", ContainerPort: 8080}}, c.Ports)
	assert.Equal(t, ptr.To(corev1.ContainerRestartPolicyAlways), c.RestartPolicy)
}
