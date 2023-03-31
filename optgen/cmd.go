package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strings"

	_ "github.com/RangelReale/trcache"
	"golang.org/x/tools/go/packages"
)

func main() {
	err := runMain()
	if err != nil {
		panic(err)
	}
}

func runMain() error {

	cfg := packages.Config{
		Mode: packages.NeedName | packages.NeedImports | packages.NeedDeps | packages.NeedTypes |
			packages.NeedSyntax | packages.NeedTypesInfo,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			return parser.ParseFile(fset, filename, src, parser.ParseComments)
		},
	}

	pkgs, err := packages.Load(&cfg, "./")
	if err != nil {
		return fmt.Errorf("cannot load %q: %w", "./", err)
	}
	if len(pkgs) != 1 {
		return fmt.Errorf("packages.Load returned %d packages, not 1", len(pkgs))
	}

	pkg := pkgs[0]

	type specValue struct {
		filename string
		comment  *ast.Comment
	}
	specs := make(map[*ast.TypeSpec]specValue)

	for _, syntax := range pkg.Syntax {
		s := getTaggedComments(syntax, "+troptgen")
		for spt, sp := range s {
			specs[spt] = specValue{
				filename: pkg.Fset.Position(syntax.Package).Filename,
				comment:  sp,
			}
		}
	}

	// for tin, _ := range pkg.TypesInfo.Types {
	// 	if t, ok := tin.(*ast.GenDecl); ok {
	// 		s := getTaggedComments(t, "+troptgen")
	// 		for spt, sp := range s {
	// 			specs[spt] = sp
	// 		}
	// 	}
	// }

	for stype, spec := range specs {
		obj := pkg.Types.Scope().Lookup(stype.Name.Name)
		if obj == nil {
			continue
		}
		fmt.Println(strings.Repeat("=", 20))
		fmt.Println(spec.filename)
		fmt.Println(obj.Name(), spec.comment)

		namedType, ok := obj.Type().(*types.Named)
		if !ok {
			continue
		}

		interfaceType, ok := namedType.Underlying().(*types.Interface)
		if !ok {
			continue
		}

		fmt.Println(namedType.String())
		for i := 0; i < namedType.TypeParams().Len(); i++ {
			fmt.Println(namedType.TypeParams().At(i).String(), namedType.TypeParams().At(i).Underlying().String())
		}
		for i := 0; i < interfaceType.NumMethods(); i++ {
			fmt.Printf("\t%s\n", interfaceType.Method(i).String())
		}
	}

	return nil
}

// getTaggedComments walks the AST and returns types which have directive comment
// returns a map of TypeSpec to directive
func getTaggedComments(pkg ast.Node, directive string) map[*ast.TypeSpec]*ast.Comment {
	specs := make(map[*ast.TypeSpec]*ast.Comment)

	ast.Inspect(pkg, func(n ast.Node) bool {
		g, ok := n.(*ast.GenDecl)

		// is it a type?
		// http://golang.org/pkg/go/ast/#GenDecl
		if !ok || g.Tok != token.TYPE {
			// never mind, move on
			return true
		}

		if g.Lparen == 0 {
			// not parenthesized, copy GenDecl.Doc into TypeSpec.Doc
			g.Specs[0].(*ast.TypeSpec).Doc = g.Doc
		}

		for _, s := range g.Specs {
			t := s.(*ast.TypeSpec)

			if c := findAnnotation(t.Doc, directive); c != nil {
				specs[t] = c
			}
		}

		// no need to keep walking, we don't care about TypeSpec's children
		return false
	})

	return specs
}

// findDirective return the first line of a doc which contains a directive
// the directive and '//' are removed
func findAnnotation(doc *ast.CommentGroup, directive string) *ast.Comment {
	if doc == nil {
		return nil
	}

	// check lines of doc for directive
	for _, c := range doc.List {
		l := c.Text
		// does the line start with the directive?
		t := strings.TrimLeft(l, "/ ")
		if !strings.HasPrefix(t, directive) {
			continue
		}

		// remove the directive from the line
		t = strings.TrimPrefix(t, directive)

		// must be eof or followed by a space
		if len(t) > 0 && t[0] != ' ' {
			continue
		}

		return c
	}

	return nil
}
