// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"go.infratographer.com/x/gidx"
)

// CreateResourceProviderInput represents a mutation input for creating resourceproviders.
type CreateResourceProviderInput struct {
	Name        string
	Description *string
	OwnerID     gidx.PrefixedID
}

// Mutate applies the CreateResourceProviderInput on the ResourceProviderMutation builder.
func (i *CreateResourceProviderInput) Mutate(m *ResourceProviderMutation) {
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetOwnerID(i.OwnerID)
}

// SetInput applies the change-set in the CreateResourceProviderInput on the ResourceProviderCreate builder.
func (c *ResourceProviderCreate) SetInput(i CreateResourceProviderInput) *ResourceProviderCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateResourceProviderInput represents a mutation input for updating resourceproviders.
type UpdateResourceProviderInput struct {
	Name             *string
	ClearDescription bool
	Description      *string
}

// Mutate applies the UpdateResourceProviderInput on the ResourceProviderMutation builder.
func (i *UpdateResourceProviderInput) Mutate(m *ResourceProviderMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
}

// SetInput applies the change-set in the UpdateResourceProviderInput on the ResourceProviderUpdate builder.
func (c *ResourceProviderUpdate) SetInput(i UpdateResourceProviderInput) *ResourceProviderUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateResourceProviderInput on the ResourceProviderUpdateOne builder.
func (c *ResourceProviderUpdateOne) SetInput(i UpdateResourceProviderInput) *ResourceProviderUpdateOne {
	i.Mutate(c.Mutation())
	return c
}