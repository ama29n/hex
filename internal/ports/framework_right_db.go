package ports

// right side framework port for DB
// on driven side, application service depends upon port and adapter implements port interface 

type DBPort interface {
	CloseDBConnection() 
	AddToHistory(answer int32, operation string) error
}
