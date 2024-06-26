package shardkv

//
// Sharded key/value server.
// Lots of replica groups, each running Raft.
// Shardctrler decides which group serves each shard.
// Shardctrler may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

const (
	OK             = "OK"
	ErrNoKey       = "ErrNoKey"
	ErrWrongGroup  = "ErrWrongGroup"
	ErrWrongLeader = "ErrWrongLeader"
	ErrRepTimeout  = "ErrRepTimeout"
	ErrDefault     = ""
	ErrWrongNum    = "ErrWrongNum"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	// You'll have to add definitions here.
	Key   string
	Value string
	Op    string // "Put" or "Append"
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.

	ClerkID   int64
	RequestID int64
}

type PutAppendReply struct {
	Err Err
}

type GetArgs struct {
	Key string
	// You'll have to add definitions here.

	ClerkID   int64
	RequestID int64
}

type GetReply struct {
	Err   Err
	Value string
}

type PullShardArgs struct {
	Num      int
	ShardIDs []int
}

type PullShardReply struct {
	Err Err
	Num int

	// mapping: inShardID -> shardDB
	ShardDBs map[int]ShardDB
}

// more general command args, takes care of all Get, Put and Append client command requests
type CommandArgs struct {
	Key       string
	Value     string
	CommandOp commandOperation
	ClerkID   int64
	RequestID int64
}

type CommandReply struct {
	Err   Err
	Value string
}

type EliminateShardArgs struct {
	Num      int
	ShardIDs []int
}

type EliminateShardReply struct {
	Err Err
}

// command op from client, Get, Put and Append
type commandOperation uint8

const (
	CommandGet commandOperation = iota
	CommandPut
	CommandAppend
)

// shard status
type shardStatus uint8

const (
	StatusServing shardStatus = iota
	StatusPulling
	StatusBePulling
	StatusEliminated
)
