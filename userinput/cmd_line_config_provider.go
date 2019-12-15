package userinput

import (
	"flag"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type CmdLineConfigProvider struct {
	config *models.Config
}

func NewCmdLineConfigProvider() interfaces.ConfigProvider {
	config := parseCmdLine()
	return &CmdLineConfigProvider{
		config: config,
	}
}

func parseCmdLine() *models.Config {
	supportedHashes := "md4, md5, sha1, sha256, sha512, ripemd160, sha3_224, sha3_256, sha3_384, sha3_512, sha512_224, sha512_256"

	apiPortPtr := flag.Uint("port", 2725, "port to expose for api")
	computedHashOverFlowPathPtr := flag.String("overflow-path", "output/computed_hash_overflow.txt", "path to file to flush computed hashes to")
	flushToFilePtr := flag.Bool("flush", true, "flush computed hashes to file if hash buffer becomes full")
	hashFunctionPtr := flag.String("hash", "sha256", fmt.Sprintf("name of hash to use - currently supported: %s", supportedHashes))
	hashListPathPtr := flag.String("hashlist-path", "hashlist.txt", "path to file containing hashes to crack")
	hashQueueBufferPtr := flag.Uint64("hash-buffer", 1000000, "buffer size for hash queue")
	logFrequencyInSecondsPtr := flag.Uint64("log-frequency", 60, "time interval for logging stats")
	logPathPtr := flag.String("log-path", "crowdcrack_log.txt", "path to log file")
	passwordQueueBufferPtr := flag.Uint64("password-buffer", 1000000, "buffer size for password queue")
	passwordRequestSizePtr := flag.Uint64("request-size", 100000, "number of passwords to request per request")
	requestBackupPath := flag.String("request-backup", "request-backup.json", "file to flush request queue to upon unexpected client shutdown")
	serverAddressPtr := flag.String("saddress", "http://localhost:2725", "address of server to connect to")
	submissionBackupPath := flag.String("submission-backup", "submission-backup.json", "file to flush submission queue to upon unexpected client shutdown")
	threadsPtr := flag.Uint("threads", 3, "number of threads to be made available")
	verbosePtr := flag.Bool("verbose", true, "print log messages to console")
	wordListPathPtr := flag.String("wordlist-path", "wordlist.txt", "path to wordlist file")
	flag.Bool("client", false, "placeholder to allow checking of client arg in main")

	flag.Parse()

	validateUint16(*apiPortPtr, "ApiPort")
	validateUint16(*threadsPtr, "Threads")

	return &models.Config{
		ApiPort:                  uint16(*apiPortPtr),
		ComputedHashOverflowPath: *computedHashOverFlowPathPtr,
		FlushToFile:              *flushToFilePtr,
		HashFunction:             *hashFunctionPtr,
		HashlistPath:             *hashListPathPtr,
		HashQueueBuffer:          *hashQueueBufferPtr,
		LogFrequencyInSeconds:    *logFrequencyInSecondsPtr,
		LogPath:                  *logPathPtr,
		PasswordQueueBuffer:      *passwordQueueBufferPtr,
		PasswordRequestSize:      *passwordRequestSizePtr,
		RequestBackupPath:        *requestBackupPath,
		ServerAddress:            *serverAddressPtr,
		SubmissionBackupPath:     *submissionBackupPath,
		Threads:                  uint16(*threadsPtr),
		Verbose:                  *verbosePtr,
		WordlistPath:             *wordListPathPtr,
	}
}

func validateUint16(configValue uint, configValueName string) {
	maxUint16 := uint(1<<16 - 1)
	if configValue > maxUint16 {
		err := fmt.Errorf("value: %d given for %s is outside the range of acceptable values (0, %d)", configValue, configValueName, maxUint16)
		panic(err)
	}
}

func IsClient(args []string) bool {
	for _, b := range args {
		if b == "-client" {
			return true
		}
	}
	return false
}

func (c *CmdLineConfigProvider) GetConfig() *models.Config {
	return c.config
}
