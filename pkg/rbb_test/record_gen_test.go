// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rbb_test

import (
	"fmt"
	"otel-arrow-adapter/pkg/air"
	"otel-arrow-adapter/pkg/air/rfield"
)

func GenSimpleRecord(ts int64) *air.Record {
	record := air.NewRecord()
	record.I64Field("ts", ts)
	record.StringField("a", "a")
	record.StringField("b", "b")
	record.StringField("c", "c")
	return record
}

func GenComplexRecord(ts int64) *air.Record {
	record := air.NewRecord()
	record.I64Field("ts", ts)
	record.StructField("a", rfield.Struct{
		Fields: []*rfield.Field{
			{Name: "b", Value: &rfield.String{Value: "b"}},
			{Name: "c", Value: &rfield.String{Value: "c"}},
		},
	})
	record.StringField("b", "b")
	record.StringField("c", "c")
	return record
}

func GenRecord(ts int64, value_a, value_b, value_c int) *air.Record {
	record := air.NewRecord()
	record.I64Field("ts", ts)
	record.StringField("c", fmt.Sprintf("c___%d", value_c))
	record.StringField("a", fmt.Sprintf("a_%d", value_a))
	record.StringField("b", fmt.Sprintf("b__%d", value_b))
	record.StructField("d", rfield.Struct{
		Fields: []*rfield.Field{
			{Name: "a", Value: &rfield.String{Value: fmt.Sprintf("a_%d", value_a)}},
			{Name: "b", Value: &rfield.String{Value: fmt.Sprintf("b_%d", value_b)}},
		},
	})
	return record
}
