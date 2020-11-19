package dictionary

type Configuration struct {
	FileName       string // Dictionary path / filename
	Overwrite      bool	  // Flag to indicate if any existing file is to be overwritten.
	FormatVersion  Version
	ScoreVersion   Version
	passphrase     []byte
	compressionAlg CompressionAlgorithms
}
