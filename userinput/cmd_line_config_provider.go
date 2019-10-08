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
	clientConfig, serverConfig := parseClient(), parseServer()
	return CmdLineConfigProvider{clientConfig, serverConfig}
}

func parseClient() *models.ClientConfig {
	serverAddressPtr := flag.String("a", "http://localhost:2725", "address of server to connect to")
	hashQueueBufferPtr := flag.Uint64("hbc", 10000, "buffer size for hash queue")
	passwordQueueBufferPtr := flag.Uint64("pbc", 10000, "buffer size for password queue")
	flushToFilePtr := flag.Bool("fc", true, "flush computed hashes to file if hash buffer becomes full")
	computedHashOverFlowPathPtr := flag.String("cpc", "output/computed_hash_overflow.txt", "path to file to flush computed hashes to")

	flag.Parse()

	return &models.ClientConfig{
		ServerAddress:            *serverAddressPtr,
		HashQueueBuffer:          *hashQueueBufferPtr,
		PasswordQueueBuffer:      *passwordQueueBufferPtr,
		FlushToFile:              *flushToFilePtr,
		ComputedHashOverflowPath: *computedHashOverFlowPathPtr,
	}
}

func parseServer() *models.ServerConfig {
	wordListPathPtr := flag.String("wp", "wordlist.txt", "path to wordlist file")
	hashListPathPtr := flag.String("hp", "hashlist.txt", "path to file containing hashes to crack")
	hashFunctionPtr := flag.String("h", "sha256", "name of hash to use - currently supported: sha256")
	apiPortPtr := flag.Uint("p", 2725, "port to expose for api")
	hashQueueBufferPtr := flag.Uint64("hbs", 10000, "buffer size for hash queue")
	passwordQueueBufferPtr := flag.Uint64("pbs", 10000, "buffer size for password queue")
	flushToFilePtr := flag.Bool("fs", true, "flush computed hashes to file if hash buffer becomes full")
	computedHashOverFlowPathPtr := flag.String("cps", "output/computed_hash_overflow.txt", "path to file to flush computed hashes to")

	flag.Parse()

	return &models.ServerConfig{
		WordlistPath:             *wordListPathPtr,
		HashlistPath:             *hashListPathPtr,
		HashFunction:             *hashFunctionPtr,
		ApiPort:                  uint16(*apiPortPtr),
		PasswordQueueBuffer:      *passwordQueueBufferPtr,
		HashQueueBuffer:          *hashQueueBufferPtr,
		FlushToFile:              *flushToFilePtr,
		ComputedHashOverflowPath: *computedHashOverFlowPathPtr,
	}
}

func IsClient(args []string) bool {
	for _, b := range args {
		if b == "--client" {
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
