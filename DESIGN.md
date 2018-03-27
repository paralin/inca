# Inca: Design Doc

This is a design outline for the initial prototype of Inca.

Inca is loosely based on the [Tendermint](https://github.com/tendermint/tendermint/wiki/Byzantine-Consensus-Algorithm) algorithm (Apache 2 licensed).

## Core Concepts

 - Deciding the next block is a process consisting of one or more rounds.
 - The state is a pointer to any structure in IPFS. 
 - Transactions are not managed by Inca, but by the application itself.
 - A validator may sign at most one block proposal per round.
 - Objects are stored locally unencrypted, remotely encrypted.
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

 - Proposer proposes a BlockHeader.
 - Validators issue Vote if they agree on the proposal.
 - Proposer issues a Block when requisite votes received.
 
For some background, in the following examples these verbs are used:

 - Fetch: retrieve the encrypted object from IPFS, decrypt it with known key.
 - Verify: return an error if a condition is not met that invalidates the object.
 - Mark: store an attribute on the object in the localdb key/value store.

Validating a block requires

 - Fetch Block object
 - Fetch inner BlockHeader object
 - Verify previous block reference matches last trusted block (if known)
 - Verify timestamp seems sane for the height and round given the last known height and round.
 - Validate vote set attached to block (fetch and verify Vote references) sums to +2/3 of previous block validators
 - Mark new block as trusted, move to next block.
 
Downloading the blockchain involves traversing the chain backward until a known base:

## Node Chains

Each node has its own block-chain of messages, like `Vote` objects. Each message has a reference to the previous message. A node must make messages sequentially, and must timestamp each message.

Nodes will be slashed for the following infractions:

 - Signing two `NodeMessage` with the same parent reference
 - Signing a `NodeMessage` with a timestamp before the previous message.
 - Signing two `NodeMessage` with `is_genesis=true` and the same genesis object reference.

Sending a pubsub message to the network involves:

 - Wrap the message object in a `pbobject.Object` and sign/encrypt it appropriately, storing in storage.
 - Build a `NodeMessage` object, timestamped, with a reference to the previous NodeMessage and the inner object.
 - Wrap the `NodeMessage` in a `pbobject.Object` and sign/encrypt it, storing in storage.
 - Send the hash of the wrapped NodeMessage to the pubsub channel.

## Merkle Tries and Storage

Inca stores its data in content-addressed distributed storage. The first implementation uses IPFS, but is pluggable, and can use any storage. 

Storage is powered by the [pbobject](https://github.com/aperturerobotics/pbobject) object-based encryption and storage mechanism. Each object has its own encryption algorithms and behaviors, tweakable by the user.

Data moves through various caching layers, described as:

 - In-memory cache: use a LRU to store recently looked up objects
 - Database cache: use the local key-value DB to store all looked up objects in the current state that are not outdated
 - IPFS: store all objects encrypted, "pin" some or all of the blockchain data depending on the operating mode

The conceptual stack of data is then:

 - **pbobject.Object**: instance of the object in memory. (in-memory cache)
 - **pbobject.ObjectWrapper**: encoded, possibly encrypted object wrapped in metadata. (database cache)
 - **ipfs.Object**: encoded object wrapper stored in IPFS, addressed by hash (IPFS storage)
 - **ipfs.Block**: block of data of fixed size, used as a unit when transporting data.
 
Writing involves going down the stack, reading involves reading up the stack.

The object types we store in storage are:

 - **Genesis**: required as a prerequisite for interacting with the chain at all, contains chain ID, chain mint timestamp, and pointer to first block.
 - **NodeMessage**: a message in the chain of messages coming from a node.

## Chain Structure

The system frequently will encounter a block that is further in the future than the local HEAD. Processing this requires traversing the chain backwards to the locally known HEAD (also known as the "root").

Each hash is approximately 50-60 bytes unencoded and around 90 bytes encoded and encrypted.

The state contains a linked merkle-list with pointers to the encrypted block and the digest of the decrypted block data.

To traverse forward quickly, a node can broadcast a "hint request." 

These are the lookup tables that we might want:

 - KeyMultihash: key multihashes to public keys
 
## Encryption Modes

Here are the supported encryption modes:

 - **ConvergentImmutable**: convergent encryption, immutable history (hash-links stored in-band).
 - **ConvergentMutable**: TODO-not implemented, 
 
## Alpha mode: ConvergentImmutable

This mode is named immutable because the history is not mutable. In this mode, storage references are encoded including the multihash of the object with an IPFS reference to the encrypted object.

Objects are stored as signed EncryptedBlob in SecretBox mode. The pre-shared key is static, not stored in the storage at all, and immutable (cannot be changed, even by transactions - this is a temporary limitation).

The SecretBox parameters are:

 - Preshared key: the chain pre-shared key
 - Nonce: the last 24 bytes of the object multihash, stored in the storage reference.

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
 - **NodeChainFork**: a node signed two objects at the same height.
 
Additionally, timing-based detections can be added later. Each detection has a configurable response by the user, so experimental detections can yield a warning in the beginning.


