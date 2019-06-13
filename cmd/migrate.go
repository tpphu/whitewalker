package cmd

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/urfave/cli"
)

func migrateAction(appContext *cli.Context) {
	// db := newDB(appContext)
	// defer db.Close()
	// subPackage := "model"
	// names := getAllNamesOfSubPackage(subPackage)
	// db.AutoMigrate(&model.Note{}, &model.User{}, &model.Department{})
}

func getAllNamesOfSubPackage(subPackage string) []string {
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, subPackage, nil, 0)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}
	names := []string{}
	for _, pack := range packs {
		for _, f := range pack.Files {
			for _, d := range f.Decls {
				typeDecl := d.(*ast.GenDecl)
				spec := typeDecl.Specs[0]
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)
					switch typeSpec.Type.(type) {
					case *ast.StructType:
						names = append(names, typeSpec.Name.Name)
					}
				}
			}
		}
	}
	return names
}

// Migrate is a definition of cli.Command used to migrate schema to database
var Migrate = cli.Command{
	Name:  "migrate",
	Usage: "Migrate db",
	Action: func(appContext *cli.Context) error {
		migrateAction(appContext)
		return nil
	},
}
