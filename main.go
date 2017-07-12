package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sumuttekno/gostruct/generator"
	"github.com/sumuttekno/gostruct/generator/config"

)

func main() {

	var RootCmd = &cobra.Command{
		Use:   "gclean",
		Short: "This is console for go clean",
		Long: `Before you use this, make sure you already understand the
          architecture used here. With this, your CRUD will automatically generated
          based on your schema.json
              `,
	}

	addCommand(RootCmd)
	RootCmd.Execute()
}

func addCommand(root *cobra.Command) {

	var cmdGenerate = &cobra.Command{
		Use:   "generate ",
		Short: "Generate your Golang projects",
		Run:   generate,
	}

	root.AddCommand(cmdGenerate)

}

func generate(cmd *cobra.Command, args []string) {


	v := config.NewViperConfig()

	source := v.GetString("type")
	user := v.GetString("user")
	password := v.GetString("password")
	host := v.GetString("host")
	port := v.GetString("port")
	dbname := v.GetString("dbname")

	fmt.Println("Generating Your Struct.......")
	generateStruct(source, user, password, host, port, dbname)
}
func generateStruct(source string, user string, password string, host string, port string, dbname string) {
	g := &generator.Generator{
		Type:     source,
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		DBName:   dbname,
	}
	g.Start()
}
