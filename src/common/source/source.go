package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/fifo"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
)

type Source struct {
	config    *Configuration // Source-specific configuration.
	isPaused  bool           // Indicator whether to pause the source Generator.
	queue     fifo.Fifo      // First-in/first-out queue from the Generator to the consumer.
	counter   int            // Counter indicating the number of words generated.
	logger    *logger.Logger // Logger reference.
	Generator func()         // The Generator function pointer.
	wordSize  int            //The current generator word size.
}
