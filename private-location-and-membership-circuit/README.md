## **Proof of Location and Membership Circuit**
This folder showcases the zero knowledge circuit work, utilizing the GROTH16 prover system, testing on the BN254 curve, performed to generate and verify flat distance and membership proofs within a specified maximum distance from a specified center point.

This circuit was originally to prove that a vehicle is on a specific car dealer lot and belongs to the fleet of the dealer on that dealer's lot to eliminate the need for auditors to manually visit the dealership for audits.

This zero knowledge circuit can be used in many scenarios. Can you think of a use case where you need to prove that something belongs to a group (membership) and is within specific coordinates (location)? Now imagine verifiably proving it without sharing any information.

Read our blogs about the work here
- [Zero Knowledge Cryptography in the Mobility Industry](https://www.mesh.xyz/insights/applying-zero-knowledge-cryptography-in-the-mobility-industry)
- [Zero Knowledge IoT- Securing Data at the Edge](https://www.mesh.xyz/insights/zero-knowledge-iot)

## **Run Tests**

Change directory into the location and membership circuit and run the following make commands to run the circuit tests

test-circuit will verify the circuits correctness under several predefined sample inputs.

test-setup will export the circuits proving key, rank one constraint system, verifying key, and solidity contract files. 

```
cd .\private-location-and-membership-circuit\location\

make test-circuit

make test-setup
```
