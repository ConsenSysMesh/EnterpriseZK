# Future Value Circuit Service
Go (Golang) API REST Service with Gin Framework and Gnark Zero-Knowledge Fututure Value Circuit

This folder serves as a fully funcitoning service that allows the user to generate and verify zero-knowledge proofs through the API.
The ZK circuit takes in the 4 necessary values for calculating future value (Future Value, Present Value, Intrest Rate, Number of years). The circuit will then prove that the user correctly calculated the future value, given the other inputs, and generate a proof and proving key. Anyone can then take that proof and proving key, and provide it to the prover function along with the other public values (intrest rate and number of years) and the proving key, and cryptographically prove correcness of the end value, without needing the private future and present values. 

This overall service will be expanded further with more circuits to create a full suite of financial calculations that enterprise users, such as banks, can utilize to eliminate the need for a middle-man to be their third party verifier. The benefits of which are privicy and speed, and accuracy. 

## 1. Clone and Change To FutureValue Directory

```shell script
git clone ...
cd DCF-CIRCUIT/futureValue
```

2. **Test the circuit**

```shell script
make test-circuit
```

3. **Run the service**

```shell script
make run
```

3. **Run Postman Collection**

```
Import futurevalue.postman_collection.json into postman and send requests

Can be found under internal/api/collection
```

**Note**: Currenty the circuit does not support number of years input greater than 18. There is potential for future upgrades if the need is found. 