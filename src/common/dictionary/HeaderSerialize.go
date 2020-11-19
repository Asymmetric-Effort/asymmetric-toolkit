package dictionary

import (
	"fmt"
)

func (o *HeaderDescriptor) Serialize() string {
	return fmt.Sprintf("%d%d", o.FormatVersion, o.ScoreVersion)
}
