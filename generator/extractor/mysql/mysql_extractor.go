package mysql

import (
	"database/sql"
	"time"

	models "github.com/golangid/gostruct/generator/model"
)

type MysqlExtractor struct {
	DBCon *sql.DB
}

func (m *MysqlExtractor) fetch(query string, args ...interface{}) ([]*models.ColumnSchema, error) {

	rows, err := m.DBCon.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.ColumnSchema, 0)
	for rows.Next() {

		t := new(models.ColumnSchema)
		err = rows.Scan(
			&t.TableName,
			&t.ColumnName,
			&t.IsNullable,
			&t.DataType,
			&t.CharacterMaximumLength,
			&t.NumericPrecision,
			&t.NumericScale,
			&t.ColumnType,
			&t.ColumnKey,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *MysqlExtractor) FetchSchema(dbName string) ([]*models.ColumnSchema, error) {

	query := `SELECT TABLE_NAME, COLUMN_NAME, IS_NULLABLE, DATA_TYPE,
		CHARACTER_MAXIMUM_LENGTH, NUMERIC_PRECISION, NUMERIC_SCALE, COLUMN_TYPE,
		COLUMN_KEY FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = ? ORDER BY TABLE_NAME, ORDINAL_POSITION`

	return m.fetch(query, dbName)

}

func (m *MysqlExtractor) ExtractModel(schemaList []*models.ColumnSchema) []models.DataGenerator {

	if len(schemaList) <= 0 {
		panic("Schema is Empty. Check Your DB Connection")
	}
	last := schemaList[0].TableName

	var model models.DataGenerator
	model.Type = "mysql"
	model.TimeStamp = time.Now()
	model.ModelName = last
	var modelList []models.DataGenerator
	var attrList []models.Attribute
	imports := make(map[string]models.Import)
	for i, schema := range schemaList {

		if last != schema.TableName {
			model.Attributes = attrList
			model.Imports = imports
			modelList = append(modelList, model)

			attrList = nil
			imports = make(map[string]models.Import)

			model.ModelName = schema.TableName
			last = schema.TableName
		} else {
			tipeData := ""

			switch schema.DataType {
			case "char", "varchar", "enum", "set", "text", "longtext", "mediumtext", "tinytext":
				tipeData = "string"
				break
			case "blob", "mediumblob", "longblob", "varbinary", "binary":
				tipeData = "[]byte"
				break
			case "date", "time", "datetime", "timestamp":

				imports["time"] = models.Import{
					Alias: "time",
					Path:  "time",
				}
				tipeData = "time.Time"
				break
			case "bit", "tinyint", "smallint", "int", "mediumint", "bigint":
				tipeData = "int64"
				break
			case "float", "decimal", "double":
				tipeData = "float64"
				break
			}

			a := models.Attribute{
				Name: schema.ColumnName,
				Type: tipeData,
			}
			attrList = append(attrList, a)
		}

		if i == len(schemaList)-1 {
			model.Imports = imports
			model.ModelName = last
			model.Attributes = attrList
			modelList = append(modelList, model)
			attrList = nil
			imports = nil
			last = ""
		}
	}
	return modelList
}
