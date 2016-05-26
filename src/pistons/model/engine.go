package model

import (
	"log"
)

type Engine struct {
	Id          int
	Description string
}

func SearchForEngines(modelId, yearId int) ([]Engine, error) {
	result := []Engine{}

	rows, err := db.Query("SELECT e.id, e.description "+
		"FROM engine e "+
		"JOIN model_year_engine mye "+
		" ON e.id = mye.engine_id "+
		"WHERE mye.model_id = $1 AND mye.year_id = $2 "+
		"ORDER BY e.description", modelId, yearId)

	if err != nil {
		log.Println(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			engine := Engine{}
			rows.Scan(&engine.Id, &engine.Description)
			result = append(result, engine)
		}
	}
	return result, err
}

func GetEngine(engineId int) (*Engine, error) {
	result := Engine{}

	row := db.QueryRow("SELECT e.id, e.description FROM engine e WHERE e.id = $1", engineId)
	err := row.Scan(&result.Id, &result.Description)

	return &result, err
}
