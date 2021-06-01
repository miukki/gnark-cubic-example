package main

import (
	"fmt"
	"log"

	"example.com/cubic"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
)

// Circuit defines a simple circuit
// x**3 + x + 5 == y
type Circuit struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x"`
	Y frontend.Variable `gnark:",public"`
}

// Define declares the circuit constraints
// x**3 + x + 5 == y
func (circuit *Circuit) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {
	x3 := cs.Mul(circuit.X, circuit.X, circuit.X)
	cs.AssertIsEqual(circuit.Y, cs.Add(x3, circuit.X, 5))
	return nil
}

func main() {
	// Get a greeting message and print it.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := cubic.Hello("Test import custom module")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	//cubic frontend init
	fmt.Println("Import Cubic Circuit")
	var cubicCircuit Circuit

	// compiles our circuit into a R1CS
	r1cs, err := frontend.Compile(ecc.BN254, backend.GROTH16, &cubicCircuit)
	fmt.Println(r1cs, err)

	//setup pk, vk as part of algorithms
	pk, vk, err := groth16.Setup(r1cs)
	fmt.Println(`err`, err)

	//frontend
	var witness Circuit
	witness.X.Assign(3)
	witness.Y.Assign(35)

	proof, err := groth16.Prove(r1cs, pk, &witness)
	fmt.Println(`buidl proof`, proof, err)

	//backend
	var publicWitness Circuit
	publicWitness.Y.Assign(35)

	errornew := groth16.Verify(proof, vk, &publicWitness)
	fmt.Println(`buidl verified!`, errornew)

	//	fmt.Println(quote.Go())

}
