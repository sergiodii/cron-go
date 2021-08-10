package interfaces

type IDataBaseConnection interface {
	Connect()
	Disconnect()
	StartMigration(...interface{})
}
