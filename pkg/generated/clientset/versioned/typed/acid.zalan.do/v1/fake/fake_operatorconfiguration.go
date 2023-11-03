/*
Copyright 2023 Compose, Zalando SE

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	acidzalandov1 "github.com/cybertec-postgresql/CYBERTEC-pg-operator/tree/v0.7.0_changeAPI/pkg/apis/cpo.opensource.cybertec.at/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeOperatorConfigurations implements OperatorConfigurationInterface
type FakeOperatorConfigurations struct {
	Fake *FakeAcidV1
	ns   string
}

var operatorconfigurationsResource = schema.GroupVersionResource{Group: "cpo.opensource.cybertec.at", Version: "v1", Resource: "operatorconfigurations"}

var operatorconfigurationsKind = schema.GroupVersionKind{Group: "cpo.opensource.cybertec.at", Version: "v1", Kind: "OperatorConfiguration"}

// Get takes name of the operatorConfiguration, and returns the corresponding operatorConfiguration object, and an error if there is any.
func (c *FakeOperatorConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *acidzalandov1.OperatorConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(operatorconfigurationsResource, c.ns, name), &acidzalandov1.OperatorConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidzalandov1.OperatorConfiguration), err
}
