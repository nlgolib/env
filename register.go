package env

var RegisteredEnvironments = make(map[string]*Environment)

func New(key string) *Environment {
	env := &Environment{Key: key}
	RegisteredEnvironments[key] = env
	return env
}
