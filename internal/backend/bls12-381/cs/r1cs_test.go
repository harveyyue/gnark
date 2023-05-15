// Copyright 2020 ConsenSys Software Inc.
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

// Code generated by gnark DO NOT EDIT

package cs_test

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"testing"
)

const n = 10000

type circuit struct {
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
}

func (circuit *circuit) Define(api frontend.API) error {
	for i := 0; i < n; i++ {
		circuit.X = api.Add(api.Mul(circuit.X, circuit.X), circuit.X, 42)
	}
	api.AssertIsEqual(circuit.X, circuit.Y)
	return nil
}

func BenchmarkSolve(b *testing.B) {

	var c circuit
	ccs, err := frontend.Compile(ecc.BLS12_381, r1cs.NewBuilder, &c)
	if err != nil {
		b.Fatal(err)
	}

	var w circuit
	w.X = 1
	w.Y = 1
	witness, err := frontend.NewWitness(&w, ecc.BLS12_381)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ccs.IsSolved(witness)
	}
}
