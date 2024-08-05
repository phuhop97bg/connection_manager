package repository

import (
	"fmt"
	"log"

	cmap "github.com/orcaman/concurrent-map/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IConnectionManager interface {
	GetConnectionByPartnerID(partnerID string) privateDB
	GetCommonDB() ICommonDB
	OnNewPartner(partnerID string)
}

type connectionManager struct {
	commonDB          commonDB
	mapPrivateConnect cmap.ConcurrentMap[string, *privateDB]
}

var connectionManagerInstance *connectionManager

func NewConnectionManager() IConnectionManager {
	if connectionManagerInstance == nil {
		initConnectionManager()
	}
	return connectionManagerInstance
}

func (c *connectionManager) initCommonDB() {
	host := "localhost"
	port := 3306
	user := "root"
	password := "my-secret-pw"

	dbName := "cygate_common"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", user, password, host, port, dbName) + "&parseTime=true&loc=Asia%2FHo_Chi_Minh"
	connection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	c.commonDB = commonDB{
		INationalityRepository: GetNationalityRepo(connection, dbName),
		IPartnerRepository:     GetPartnerRepo(connection, dbName),
		db:                     connection,
	}
}

func (c *connectionManager) initPartnerConnection(partnerID string) {
	host := "localhost"
	port := 3306
	user := "root"
	password := "my-secret-pw"

	dbName := makeDatabaseName(partnerID)
	log.Println("Database name: ", dbName)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", user, password, host, port, dbName) + "&parseTime=true&loc=Asia%2FHo_Chi_Minh"
	connectionDB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Printf("Error when connect to database: %v", err)
		return
	}
	partnerDB := privateDB{
		IDepartmentRepository: GetDepartmentRepo(connectionDB, dbName),
		IProfileRepository:    GetProfileRepo(connectionDB, dbName),
		db:                    connectionDB,
	}
	c.mapPrivateConnect.Set(partnerID, &partnerDB)
}

func initConnectionManager() {
	connectionManagerInstance = &connectionManager{
		mapPrivateConnect: cmap.New[*privateDB](),
	}
	connectionManagerInstance.initCommonDB()
	partners, err := connectionManagerInstance.commonDB.GetAll()
	if err != nil {
		panic(err)
	}
	for _, partner := range partners {
		connectionManagerInstance.initPartnerConnection(partner.PartnerId)
	}
	log.Println("Connection manager initialized, partner connections: ", connectionManagerInstance.mapPrivateConnect.Keys())
}

// -----------------------------------------------------------------------------------------------------------------------
// GetCommonDB implements IConnectionManager.
func (c *connectionManager) GetCommonDB() ICommonDB {
	return c.commonDB
}

// GetConnectionByPartnerID implements IConnectionManager.
func (c *connectionManager) GetConnectionByPartnerID(partnerID string) privateDB {
	connection, ok := c.mapPrivateConnect.Get(partnerID)
	if !ok || connection == nil {
		log.Printf("Connection not found for partner: %s", partnerID)
		return privateDB{}
	}
	return *connection
}

// NewPartner implements IConnectionManager.
func (c *connectionManager) OnNewPartner(partnerID string) {
	c.commonDB.CreateNewDB(partnerID)
	c.initPartnerConnection(partnerID)
	privateDB, ok := c.mapPrivateConnect.Get(partnerID)
	if !ok {
		log.Printf("Error while creating new partner: %s", partnerID)
		return
	}
	log.Printf("New partner created: %s", partnerID)
	err := privateDB.CreateTables()
	if err != nil {
		log.Printf("Error while creating tables for partner: %s, %v", partnerID, err)
	}
}
func makeDatabaseName(partnerId string) string {
	return "cygate_" + partnerId
}
