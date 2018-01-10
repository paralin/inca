# Inca: Design Doc

This is a design outline for the initial prototype of Inca.

Inca is loosely based on the [Tendermint](https://github.com/tendermint/tendermint/wiki/Byzantine-Consensus-Algorithm) algorithm (Apache 2 licensed).

... but this implementation will be kick-ass.

## Core Concepts

 - A block contains zero or more transactions.
 - Deciding the next block is a process consisting of one or more rounds.
 - A validator may sign at most one block proposal per round.
 - Objects are stored in IPFS, encrypted on default.
 - Votes are timestamped. Timestamps cannot be in the future, and actions must be sequential.
 - Evidence foul play can be presented against a validator.
 - +2/3 consensus on a proposal for a round triggers the next round.

The validator set can be updated at the end of each block. The validator set change is included in the block, but must be derived deterministically based on transactions in the block.

The first proposal that receives +2/3 votes will be accepted as a committed block. It is the responsibility of the proposer to mint the final block. The first block that is minted (by timestamp) is used.

Each message sent over pubsub is also stored in IPFS, and contains a hash-link to the previous message sent (like a nonce, but with a history). In this way every participant forms their own "self-blockchain" or "message-chain." This forces sequential timestamps, sequential message processing, and makes timing or multi-sign attacks more difficult.

So, the basic rule set:

 - Sign one proposal per round
 - Node messages contain hash links to the previous message (sequential)
 - Node messages are timestamped and must be sequential in time
 - Blocks are timestamped and must be sequential in time
 - Blocks must have a timestamp within +/- 2 second of the message they are sent in
 - Objects are stored as hash-linked references (a DAG). If an object does not change, don't change its hash link.
 
Block lifecycle:

 - Proposer proposes a BlockHeader, putting it into the hash-linked IPFS db.
 - Validators issue Vote 
 
For some background, in the following examples these verbs are used:

 - Fetch: retrieve the encrypted object from IPFS, decrypt it with known key.

Validating a block requires

 - Fetch Block object
 - Fetch inner BlockHeader object
 - Verify previous block reference matches last trusted block (if known)
 - Verify timestamp seems sane for the height and round given the last known height and round.
 - Validate vote set attached to block (fetch and verify Vote references) sums to +2/3 of previous block validators
 - Mark new block as trusted, move to next block.
 
Downloading the blockchain involves traversing the chain backward until a known base:



## Lookup Tables

These are the lookup tables that we might want:

 - KeyMultihash: key multihashes to public keys
 - 

## Encryption and Storage

Objects are encoded to IPFS using [pbobject](https://github.com/aperturerobotics/pbobject). This allows for each object to identify a unique encryption mechanism.

These are the kinds of encrypted objects:

 - **Genesis**: AES-256 bit PSK, required as a prerequisite for interacting with the chain at all.
 
Each object has its default encryption specified in a table, but this can be overridden by the user if desired without modifying Inca.

The encryption mechanism is recognized at read-time, so a blockchain could be migrated off of an insecure encryption mechanism gracefully without issue or re-encoding the chain.

## Chain Objects

Status: TODO - incomplete

### Genesis

Status: TODO - incomplete

```proto
// Genesis is the object containing the blockchain identity.
message Genesis {
  // ChainId is used to differentiate between chains, but could be set to anything.
  // It is an opaque value and ignored by the system.
  string chain_id = 1;
  // Timestamp contains the time the genesis block was formed.
  timestamp.Timestamp timestamp = 2;
  // TODO: validator set? app hash? encryption key?
}
```

The genesis block contains the initial state of the blockchain, including the chain ID, init timestamp, initial validator set, and initial encryption key.


## Policies

Status: TODO - proto incomplete

The system will automatically submit a transaction with any evidence of foul-play in peer validators.

Allow the application developer to specify a **Policy** which determines what happens when:

 - **DoubleSign**: validator signed two proposals for a round.
 - **NonSequentialSign**: validator signed a previous round after endorsing a later round.
 - **OutOfTurnPropose**: proposer proposed for the wrong round.
 - **DoublePropose**: a proposer proposed two proposals for a round.
 - **MistimedRound**: a round was started before it should have been.
 - **InvalidChainHint**: a chain hint contained data for the wrong blockchain.
 
Additionally, timing-based detections can be added later. Each detection has a configurable response by the user, so experimental detections can yield a warning in the beginning.
