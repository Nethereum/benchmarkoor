package client

// nethereumSpec is the client specification for Nethereum's harness node
// (Nethereum.Node.HarnessServer), which boots from a standard genesis.json and
// serves JSON-RPC plus a JWT-authenticated Engine API.
type nethereumSpec struct{}

// NewNethereumSpec creates a new Nethereum client specification.
func NewNethereumSpec() Spec {
	return &nethereumSpec{}
}

// Ensure interface compliance.
var _ Spec = (*nethereumSpec)(nil)

func (s *nethereumSpec) Type() ClientType {
	return ClientNethereum
}

func (s *nethereumSpec) DefaultImage() string {
	return "nethereum/harness-node:local"
}

func (s *nethereumSpec) DefaultCommand() []string {
	return []string{
		"--datadir=/data",
		"--http.addr=0.0.0.0",
		"--http.port=8545",
		"--authrpc.addr=0.0.0.0",
		"--authrpc.port=8551",
		"--authrpc.jwtsecret=/tmp/jwtsecret",
	}
}

func (s *nethereumSpec) GenesisFlag() string {
	return "--genesis="
}

func (s *nethereumSpec) RequiresInit() bool {
	return false
}

func (s *nethereumSpec) InitCommand() []string {
	return nil
}

func (s *nethereumSpec) DataDir() string {
	return "/data"
}

func (s *nethereumSpec) GenesisPath() string {
	return "/network-config/genesis.json"
}

func (s *nethereumSpec) JWTPath() string {
	return "/tmp/jwtsecret"
}

func (s *nethereumSpec) RPCPort() int {
	return 8545
}

func (s *nethereumSpec) EnginePort() int {
	return 8551
}

func (s *nethereumSpec) MetricsPort() int {
	return 0
}

func (s *nethereumSpec) DefaultEnvironment() map[string]string {
	return nil
}

// RPCRollbackSpec returns nil — the Nethereum harness node does not expose
// debug_setHead; use a container-level rollback strategy (container-recreate).
func (s *nethereumSpec) RPCRollbackSpec() *RPCRollbackSpec {
	return &RPCRollbackSpec{
		Method:    RollbackMethodSetHeadHex,
		RPCMethod: "debug_setHead",
	}
}

func (s *nethereumSpec) DefaultConfigFiles() map[string]string {
	return nil
}

func (s *nethereumSpec) SnapshotPrepareArgs() []string {
	return nil
}

func (s *nethereumSpec) DBMaintenanceCommands(_ string) *DBMaintenanceCommands {
	return nil
}
