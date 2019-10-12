package userinput

import (
	"flag"
	"fmt"
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
	supportedHashes := "md4, md5, sha1, sha256, sha512, ripemd160, sha3_224, sha3_256, sha3_384, sha3_512, sha512_224, sha512_256"

	serverAddressPtr := flag.String("saddress", "http://localhost:2725", "address of server to connect to")
	hashQueueBufferPtr := flag.Uint64("hash-buffer", 10000, "buffer size for hash queue")
	passwordQueueBufferPtr := flag.Uint64("password-buffer", 10000, "buffer size for password queue")
	flushToFilePtr := flag.Bool("flush", true, "flush computed hashes to file if hash buffer becomes full")
	computedHashOverFlowPathPtr := flag.String("overflow-path", "output/computed_hash_overflow.txt", "path to file to flush computed hashes to")
	wordListPathPtr := flag.String("wordlist-path", "wordlist.txt", "path to wordlist file")
	hashListPathPtr := flag.String("hashlist-path", "hashlist.txt", "path to file containing hashes to crack")
	hashFunctionPtr := flag.String("hash", "sha256", fmt.Sprintf("name of hash to use - currently supported: %s", supportedHashes))
	apiPortPtr := flag.Uint("port", 2725, "port to expose for api")
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
