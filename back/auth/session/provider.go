package session

type Provider interface {
	SessionInit(sid string) (Session, error) // Initializes a session and returns it in case of success
	SessionRead(sid string) (Session, error) // Returns an existing session or a new one there was no session with this SID
	SessionDestroy(sid string) error         // Destroys the targeted session if it exists
	SessionGC(maxLifeTime int64)             // Deletes expired sessions
}

var provides = make(map[string]Provider)

// Register makes a session provider available by the provided name.
// If a Register is called twice with the same name or if the driver is nil, it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}
