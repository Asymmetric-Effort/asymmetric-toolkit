package persistentTreeFlags

type PersistentTreeFlag uint16

type Flags struct {
	raw uint16
}

const (
	Compression PersistentTreeFlag = 0
	Encryption  PersistentTreeFlag = 1
	//Reserved 2...15
)

func (flags *Flags) Get(flagName PersistentTreeFlag) (result bool) {
	return (flags.raw & uint16(1<<flagName)) != 0

}

func (flags *Flags) Set(flagName PersistentTreeFlag) {
	var flagValue uint16 = 1 << uint16(flagName)
	flags.raw = flags.raw | flagValue
}

func (flags *Flags) Clear(flagName PersistentTreeFlag) {
	var flagValue uint16 = 0
	if (flags.raw & uint16(1<<flagName)) != 0 {
		flagValue = 1 << uint16(flagName)
	}
	flags.raw = flags.raw ^ flagValue
}
