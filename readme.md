SynergyChain is a multichain-based data sharing framework with hierarchical access control.
The consensus of SynergyChain is the Delegated Proof of Stake (DPoS), which uses a decentralized voting mechanism.
Merkel tree is mainly used for effective data comparison,and verification.
Smart contracts are used to complete the main transaction processing logic of the system.
Data sharing is mainly to submit shared data for different blockchain nodes that have been connected to the SynergyChain system, including two main parts: shared data structure and data sharing level.
SynergyChain receives the access request from the contract layer and agrees on the access request within the blockchain network, and then it will trigger the access request contract. The access request contract takes the request data from the contract layer as input and takes the process result as output.
The data synchronization verification layer is a module belonging to the client, including the SynergyChain block header synchronization sub-module and the data validator submodule. The application can customize the data requester as needed. The shared data requested from SynergyChain needs to go through a verification sub-module to ensure the reliability of the data.
The system transaction processing speed test will use the stress test tool Apache Bench (AB) to trigger 15,000 transactions with different concurrency to record the peak system performance and real-time system throughput.
The request concurrency is set to 1, 5, and 10 respectively, and send 15,000 data to the SynergyChain system.
