/*
Copyright 2016 The ContainerOps Authors All rights reserved.

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

package utils

import (
	"crypto/sha1"
	"fmt"
	"reflect"
	"time"
)

type Meta struct {
	Name string
	Hash string

	Created time.Time
	Expired time.Time
}

const (
	//The default life circle for a software is half a year
	defaultLifecircle = time.Hour * 24 * 180
)

func GenerateMeta(file string, contentByte []byte) (meta Meta) {
	meta.Name = file
	meta.Hash = fmt.Sprintf("%x", sha1.Sum(contentByte))
	meta.Created = time.Now()
	meta.Expired = meta.Created.Add(defaultLifecircle)
	return
}

func (a Meta) GetHash() string {
	return a.Hash
}

func (a Meta) IsExpired() bool {
	return a.Expired.Before(time.Now())
}

func (a Meta) GetCreated() time.Time {
	return a.Created
}

func (a *Meta) SetCreated(t time.Time) {
	a.Created = t
}

func (a Meta) GetExpired() time.Time {
	return a.Expired
}

func (a *Meta) SetExpired(t time.Time) {
	a.Expired = t
}

func (a Meta) Compare(b Meta) int {
	if reflect.DeepEqual(a, b) {
		return 0
	}

	if a.Created.Before(b.Created) {
		return -1
	}

	return 1
}