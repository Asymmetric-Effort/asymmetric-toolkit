package source
/*
	Source::Setup() will consume the common/source Configuration struct and setup the internal
	state of the system.
 */
import "asymmetric-effort/asymmetric-toolkit/src/common/errors"

func (o *Source) Setup(config *Configuration) {
	/*
		Verify the config reference is not nil.
	 */
	errors.Assert(config != nil, "Encountered nil configuration in Source::Setup()")
	/*
		Map the config into the internal state (by reference) for use by common/source.
	 */
	o.config = config
	/*
		Initialize the FIFO Queue with a given size.
	 */
	o.queue.Setup(o.config.BufferSize)
	/*
		From the outset, we leave the source Generator in a paused state.  This will
		ensure that the Generator does not produce any results until it is unpaused.
	 */
	o.isPaused = true
	/*
		Depending on the current mode of operation, we will connect the appropriate
		Generator.
	 */
	switch config.Mode.Get() {
	case Dictionary:
		o.Generator =o.generateDictionary
	case Random:
		o.Generator =o.generateRandom
	case Sequence:
		o.Generator =o.generateSequence
	default:
		panic("Source::Setup() encountered Mode NotSet")
	}
}
