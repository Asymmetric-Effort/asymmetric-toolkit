package dictionary

type Word struct {
	flags byte
	// On disk we store size in 64-bit bytes
	Word string
	//ToDo: Add metadata later.
	Parent int64 // file address of parent node (0 is invalid)
	Lhs    int64 //file address of Lhs tree node
	Rhs    int64 //file address of Rhs tree node
}
