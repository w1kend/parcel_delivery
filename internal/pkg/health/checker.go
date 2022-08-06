package health

// Status - the status of the service
type Status string

const (
	StatusOK    Status = "OK"
	StatusWarn  Status = "WARN"
	StatusCrit  Status = "CRIT"
	StatusError Status = "ERROR"
)

// Result - health check result for a single dependency
type Result struct {
	// AppName - application name
	AppName string `json:"app_name"`
	// DepName - dependency name (e.g. postgres, redis, kafka etc.)
	DepName string `json:"dep_name"`
	// Status - current status of the system
	Status Status `json:"status"`
	// Description - description of the status
	Description string `json:"description"`
}

// State - health check for a single dependency
type State struct {
	// Status - current status of the dependency
	Status Status `json:"status"`
	// Description - description of the status
	Description string `json:"description"`
}

// CheckFn - function to check the health of a dependency
type CheckFn func() State

// Checker - is a health checker
type Checker struct {
	appName string
	deps    map[string]CheckFn
}

func NewChecker(appName string) *Checker {
	return &Checker{
		appName: appName,
		deps:    make(map[string]CheckFn, 5),
	}
}

// Track - adds a dependency to the health checker
func (c *Checker) Track(name string, dep CheckFn) {
	c.deps[name] = dep
}

func (c *Checker) Untrack(name string) {
	delete(c.deps, name)
}

// Check - checks the health of the dependencies
func (c Checker) Check() []Result {
	var results []Result
	for name, fn := range c.deps {
		state := fn()

		results = append(results, Result{
			AppName:     c.appName,
			DepName:     name,
			Status:      state.Status,
			Description: state.Description,
		})
	}
	return results
}
