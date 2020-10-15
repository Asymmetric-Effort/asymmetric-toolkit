package cli

func (o *Configuration) LoadDefault() {
	o.Concurrency = defaultConcurrency
	o.Debug = false
	o.Delay = 0
	o.Depth = defaultDepth
	o.Force = false
	o.Pattern.Set(defaultFilterPattern)
	o.RecordTypes.Set(defaultDnsRecordTypes)
	o.Timeout = defaultTimeout
}
