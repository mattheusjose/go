package main

import (
	"fmt"
)

type Terminal struct {
	id          int
	fctNumber   string
	assetNumber string
}

type TerminalRepository interface {
	Save(terminal Terminal)
	Delete(terminal Terminal)
	FindByAssetNumber(assetNumber string) Terminal
}

type MySQLRepository struct {
	driver   string
	url      string
	username string
	password string
}

type PostgresRepository struct {
	driver   string
	url      string
	username string
	password string
}

func (mysqlRepository MySQLRepository) Save(terminal Terminal) {

	fmt.Printf("Saving with %s driver with id %d\n",
		mysqlRepository.driver, terminal.id)
}

func (mysqlRepository MySQLRepository) Delete(terminal Terminal) {
	fmt.Printf("Deleting with %s driver with id %d\n",
		mysqlRepository.driver, terminal.id)
}

func (mysqlRepository MySQLRepository) FindByAssetNumber(assetNumber string) Terminal {

	fmt.Printf("Getting with %s driver and asset number of %s\n",
		mysqlRepository.driver, assetNumber)
	terminal := Terminal{
		id:          1,
		assetNumber: assetNumber,
		fctNumber:   "7377366",
	}

	return terminal
}

func (postgresRepository PostgresRepository) Save(terminal Terminal) {

	fmt.Printf("Saving with %s driver and asset number of %s\n",
		postgresRepository.driver, terminal.assetNumber)
}

func (postgresRepository PostgresRepository) Delete(terminal Terminal) {

	fmt.Printf("Deleting with %s driver and asset number of %s\n",
		postgresRepository.driver, terminal.assetNumber)
}

func (postgresRepository PostgresRepository) FindByAssetNumber(assetNumber string) Terminal {

	fmt.Printf("Getting with %s driver and asset number of %s\n",
		postgresRepository.driver, assetNumber)
	terminal := Terminal{
		id:          1,
		assetNumber: assetNumber,
		fctNumber:   "328825",
	}

	return terminal
}

func Save(repository TerminalRepository, terminal Terminal) {
	repository.Save(terminal)
}

func Delete(repository TerminalRepository, terminal Terminal) {
	repository.Delete(terminal)
}

func FindByAssetNumber(repository TerminalRepository, assetNumber string) Terminal {
	return repository.FindByAssetNumber(assetNumber)
}

func main() {

	mysql := MySQLRepository{
		driver:   "mysql",
		url:      "jdbc:mysql://localhost:3306/db",
		username: "admin",
		password: "123456789",
	}

	postgres := PostgresRepository{
		driver:   "postgres",
		url:      "jdbc:postgres://localhost:5475/db",
		username: "admin",
		password: "123456789",
	}

	terminalA := FindByAssetNumber(mysql, "583720")
	Save(mysql, terminalA)
	Delete(mysql, terminalA)

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-")

	terminalB := FindByAssetNumber(postgres, "583720")
	Save(postgres, terminalB)
	Delete(postgres, terminalB)
}
