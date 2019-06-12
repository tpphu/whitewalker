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
	subPackage := "model"
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, subPackage, nil, 0)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}

	funcs := []*ast.FuncDecl{}
	for _, pack := range packs {
		for _, f := range pack.Scopes() {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.StructType); isFn {
					funcs = append(funcs, fn)
				}
			}
		}
	}
	//db.AutoMigrate(&model.Note{}, &model.User{}, &model.Department{})
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
