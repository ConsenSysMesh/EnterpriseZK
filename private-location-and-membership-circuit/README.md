## **Proof of Location and Membership Circuit**
This folder showcases the zero knowledge circuit work, utilizing the GROTH16 prover system, testing on the BN254 curve, performed to generate and verify flat distance and membership proofs within a specificed maximum distance from a specfied center point.

This zero-knowledge circuit can be used in a many scenarios. The original scenario this circuit was developed for was to prove that a vehicle is on a specific car dealer lot, and belongs to the fleet of the dealer on that dealer's lot.

## **Run Tests**

Change directory into the location and membership circuit and run the following make commands to run the circuit tests.

test-circuit will verify the circuits correctness under several predefined sample inputs.

test-setup will export the circuits proving key, rank one constraint system, verying key, and solidity contract files. 

```
cd .\private-location-and-membership-circuit\location\

make test-circuit

make test-setup
```