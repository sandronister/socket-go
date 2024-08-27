package sugar

func (s *Logger) Debug(message string) {
	s.sugar.Debug(message)
}

func (s *Logger) Error(message string) {
	s.sugar.Error(message)
}

func (s *Logger) Fatal(message string) {
	s.sugar.Fatal(message)
}

func (s *Logger) Info(message string) {
	s.sugar.Info(message)
}

func (s *Logger) Warn(message string) {
	s.sugar.Warn(message)
}
