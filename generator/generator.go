package generator

import (
	"database/sql"
	"fmt" 
	"go/build" 
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/sumuttekno/gostruct/generator/extractor/mysql"
	models "github.com/sumuttekno/gostruct/generator/model"
)

type Generator struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

const (
	MYSQL_TYPE = "mysql"
)

func (g *Generator) Start() {

	switch strings.ToLower(g.Type) {
	case MYSQL_TYPE:
		g.mysqlGenerator()
	}

}
func (m *Generator) Dsn() string {
	dsn := m.User + `:` + m.Password + `@tcp(` + m.Host + `:` + m.Port + `)/` +
		m.DBName + `?parseTime=1&loc=Asia%2FJakarta`
	return dsn
}

func (g *Generator) mysqlGenerator() {

	dbConn, err := sql.Open(`mysql`, g.Dsn())
	if err != nil {

		fmt.Println(err)
	}
	defer dbConn.Close()

	mysqlExtractor := mysql.MysqlExtractor{DBCon: dbConn}
	schemaList, err := mysqlExtractor.FetchSchema(g.DBName)
	if err != nil {
		panic("Unexpected Error " + err.Error())
	}

	listModels := mysqlExtractor.ExtractModel(schemaList)
	g.generateStruct(listModels)

}

func (g *Generator) generateStruct(list []models.DataGenerator) {
	if len(list) < 1 {
 
		panic("Table Not Exist")
 
	}

	for i := 0; i < len(list); i++ {
		g.generateFile(&list[i])
	}

}

func (g *Generator) generateFile(dataSend *models.DataGenerator) {
 
	projectPackages := "github.com/sumuttekno/gostruct"
	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles(build.Default.GOPATH + "/src/" + projectPackages + "/generator/struct_template.tpl")
 
 
	if err != nil {
		panic("Unknown Error " + err.Error())
	}

	pathP := "models/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}
	f, err := os.Create(pathP + strings.ToLower(dataSend.ModelName) + ".go")
	if err != nil {
		panic("Unknown Error " + err.Error())
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "struct_template.tpl", dataSend)

	if err != nil {
		panic("Unknown Error " + err.Error())
	}
}
