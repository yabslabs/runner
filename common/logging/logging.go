/*
Copyright 2018 tribock.

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
package logging

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	fieldLogID = "abraxasLogID"
)

func WithError(id string, err error) *logrus.Entry {
	return WithID(id).WithError(err)
}

func WithID(id string) *logrus.Entry {
	return logrus.WithField(fieldLogID, id)
}

func WithIDFields(id string, fields ...interface{}) *logrus.Entry {
	m := Pairs(fields...)
	m[fieldLogID] = id
	return logrus.WithFields(m)
}

func Pairs(kv ...interface{}) map[string]interface{} {
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("Pairs got the odd number of input pairs for metadata: %d", len(kv)))
	}

	v := map[string]interface{}{}
	var key string
	for i, s := range kv {
		if i%2 == 0 {
			key = fmt.Sprint(s)
			continue
		}

		v[key] = s
	}
	return v
}
