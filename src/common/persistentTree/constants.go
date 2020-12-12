package persistentTree

const (
	// See 16-bit file version number
	fileHeaderVersion uint16 = 0

	//See 16-bit flags
	fileHeaderFlagCompressionEnabled = 1
	fileHeaderFlagEncryptionEnabled = 2

	filePermissions = 0600
)

