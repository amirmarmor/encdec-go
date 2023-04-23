package log

type PrefixLogger struct {
	prefix       string
	forceVerbose *int
}

func ProducePrefixLogger(
	prefix string,
) *PrefixLogger {
	return &PrefixLogger{
		prefix: prefix,
	}
}

func (p *PrefixLogger) ForceVerbose(verbose int) {
	p.forceVerbose = &verbose
}

func (p *PrefixLogger) GetVerbose() int {
	return *p.forceVerbose
}

func (p *PrefixLogger) V1(format string, v ...interface{}) {
	if p.GetVerbose() < 1 {
		return
	}
	writeRecord(levelVerb, p.prefix+format, v...)
}

func (p *PrefixLogger) V2(format string, v ...interface{}) {
	if p.GetVerbose() < 2 {
		return
	}
	writeRecord(levelVerb, p.prefix+format, v...)
}

func (p *PrefixLogger) V5(format string, v ...interface{}) {
	if p.GetVerbose() < 5 {
		return
	}
	writeRecord(levelVerb, p.prefix+format, v...)
}

func (p *PrefixLogger) Info(format string, v ...interface{}) {
	writeRecord(levelInfo, p.prefix+format, v...)
}

func (p *PrefixLogger) Warn(format string, v ...interface{}) {
	writeRecord(levelWarn, p.prefix+format, v...)
}

func (p *PrefixLogger) Fatal(format string, v ...interface{}) {
	writeRecord(levelFatal, p.prefix+format, v...)
}
