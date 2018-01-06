# Inca: Design Doc

This is a design outline for the initial prototype of Inca.

Inca is loosely based on the [[Tendermint][https://github.com/tendermint/tendermint/wiki/Byzantine-Consensus-Algorithm]] algorithm (Apache 2 licensed).

## Core Concepts

 - A block contains zero or more transactions.
 - Deciding the next block is a process consisting of one or more rounds.
 - A validator may sign at most one block proposal per round.
 - Objects are stored in IPFS, encrypted on default.


## Encryption

Objects are encoded to IPFS using [pbobject](https://github.com/aperturerobotics/pbobject). This allows for each object to identify a unique encryption mechanism, and for each object to refer to what key it was encrypted with in a concise way.

These are the kinds of encrypted objects:

 - **Genesis**: AES-256 bit PSK, required as a prerequisite for interacting with the chain at all.
 
Each object has its default encryption specified in a table, but this can be overridden by the user if desired without modifying Inca.

The encryption mechanism is recognized at read-time, so a blockchain could be migrated off of an insecure encryption mechanism gracefully without issue or re-encoding the chain.
