package cli

func (o *Configuration) LoadDefault() {
	o.Concurrency = DefaultConcurrency
	o.Debug = false
	o.Delay = 0
	o.Depth = DefaultDepth
	o.Force = false
	o.Pattern.Set(DefaultFilterPattern)
	o.RecordTypes.Set(DefaultDnsRecordTypes)
	o.Timeout = DefaultTimeout
}
