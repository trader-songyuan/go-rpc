package constant

// rpc message constants
const (
	MAGIC_NUMBER = 99
	VERSION = 2
	HEAD_LENGTH = 66
)
// rpc message type
const (
	RPC_REQUEST byte = 1
	RPC_RESPONSE byte = 2
)
// rpc response status code
const (
	SUCCESS int32 = 1
	INVOKE_ERROR int32 = 2
)
const (
	RPC_MESSAGE = "rpcMessage"
)
const (
	NETWORK = "tcp"
	SERVER_PORT = "29880"
)