/*
Copyright 2017 The Kubernetes Authors.

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

package testprovider

import (
	"errors"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
)

// Provider is a test provider which allows testing of components relying on [dns] provider interface method invokation
type Provider struct {
	ReturnError    bool
	onRecords      func(err error)
	onApplyChanges func(err error)
}

var ( //TODO: Might need to be exposed
	errRecords      = errors.New("records errored")
	errApplyChanges = errors.New("apply changes errored")
)

var _ provider.Provider = &Provider{}

// NewProvider returns new test provider
func NewProvider() *Provider {
	return &Provider{}
}

//Records calls registered callback
func (p *Provider) Records(zone string) ([]endpoint.Endpoint, error) {
	if p.ReturnError {
		p.onRecords(errRecords)
		return nil, errRecords
	}
	p.onRecords(nil)
	return nil, nil
}

//ApplyChanges calls registered callback
func (p *Provider) ApplyChanges(zone string, changes *plan.Changes) error {
	if p.ReturnError {
		p.onApplyChanges(errApplyChanges)
		return errApplyChanges
	}
	p.onApplyChanges(nil)
	return nil
}
