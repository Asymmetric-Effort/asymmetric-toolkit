package definition

import "asymmetric-effort/asymmetric-toolkit/src/common/encryption"

func (o *Record) String (key *encryption.Key) string {
	return o.Get(key)
}