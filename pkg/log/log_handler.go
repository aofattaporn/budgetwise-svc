package log

import "context"

func (l *Logger) Debug(msg string) {
	l.logger.Desugar().Debug(msg)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *Logger) Info(msg string) {
	l.logger.Desugar().Info(msg)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *Logger) Warn(msg string) {
	l.logger.Desugar().Warn(msg)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *Logger) Error(msg string) {
	l.logger.Desugar().Error(msg)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *Logger) Fatal(msg string) {
	l.logger.Desugar().Fatal(msg)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *Logger) WithCtx(ctx context.Context) ILogger {
	f := []interface{}{}
	if correlationid, ok := ctx.Value(CorrelationId{}).(string); ok {
		f = append(f, "requestuid", correlationid)
	}
	return &Logger{
		logger: l.logger.With(f...),
	}
}

func (l *Logger) WithField(field Field) *Logger {
	f := []interface{}{}
	for k, v := range field {
		f = append(f, k)
		f = append(f, v)
	}
	return &Logger{
		logger: l.logger.With(f...),
	}
}
