package dictionary

type Header struct {
	FileVersion  uint8
	ScoreVersion uint8
	RootWord     int64
	RootTag      int64
}
