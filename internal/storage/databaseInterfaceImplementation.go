package storage

type executeSql struct{}

func (e executeSql) createTables() error {
	return nil
}

func (e executeSql) insertValuesInDomainAndMTA() error {
	return nil
}

func (e executeSql) createRole() error {
	return nil
}

func executesSql(e executeSqlQueries) error {
	if e.createTables() != nil {
		return e.createTables()
	}
	if e.insertValuesInDomainAndMTA() != nil {
		return e.insertValuesInDomainAndMTA()
	}
	if e.createRole() != nil {
		return e.createRole()
	}
	return nil
}
