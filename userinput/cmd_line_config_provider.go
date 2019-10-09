package userinput

import (
	"flag"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type CmdLineConfigProvider struct {
	clientConfig *models.ClientConfig
	serverConfig *models.ServerConfig
}

func NewCmdLineConfigProvider() CmdLineConfigProvider {
	clientConfig, serverConfig := parseCmdLine()
	return CmdLineConfigProvider{clientConfig, serverConfig}
}

func parseCmdLine() (*models.ClientConfig, *models.ServerConfig) {
	serverAddressPtr := flag.String("a", "http://localhost:2725", "address of server to connect to")
	hashQueueBufferPtr := flag.Uint64("hb", 10000, "buffer size for hash queue")
	passwordQueueBufferPtr := flag.Uint64("pb", 10000, "buffer size for password queue")
	flushToFilePtr := flag.Bool("f", true, "flush computed hashes to file if hash buffer becomes full")
	computedHashOverFlowPathPtr := flag.String("cp", "output/computed_hash_overflow.txt", "path to file to flush computed hashes to")
	wordListPathPtr := flag.String("wp", "wordlist.txt", "path to wordlist file")
	hashListPathPtr := flag.String("hp", "hashlist.txt", "path to file containing hashes to crack")
	hashFunctionPtr := flag.String("h", "sha256", "name of hash to use - currently supported: sha256")
	apiPortPtr := flag.Uint("p", 2725, "port to expose for api")
	flag.Bool("client", false, "placeholder to allow checking of client arg in main")

	flag.Parse()

	clientConfig := &models.ClientConfig{
		ServerAddress:            *serverAddressPtr,
		HashQueueBuffer:          *hashQueueBufferPtr,
		PasswordQueueBuffer:      *passwordQueueBufferPtr,
		FlushToFile:              *flushToFilePtr,
		ComputedHashOverflowPath: *computedHashOverFlowPathPtr,
	}
	serverConfig := &models.ServerConfig{
		WordlistPath:             *wordListPathPtr,
		HashlistPath:             *hashListPathPtr,
		HashFunction:             *hashFunctionPtr,
		ApiPort:                  uint16(*apiPortPtr),
		PasswordQueueBuffer:      *passwordQueueBufferPtr,
		HashQueueBuffer:          *hashQueueBufferPtr,
		FlushToFile:              *flushToFilePtr,
		ComputedHashOverflowPath: *computedHashOverFlowPathPtr,
	}

	return clientConfig, serverConfig
}

func IsClient(args []string) bool {
	for _, b := range args {
		if b == "-client" {
			return true
		}
	}
	return false
}

func (c *CmdLineConfigProvider) GetClientConfig() *models.ClientConfig {
	return c.clientConfig
}

func (c *CmdLineConfigProvider) GetServerConfig() *models.ServerConfig {
	return c.serverConfig
}
