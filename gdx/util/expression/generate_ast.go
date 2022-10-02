package main

import (
	"fmt"
	"os"

	"github.com/mrcampbell/gdx/gdx"
	"github.com/mrcampbell/gdx/gdx/gen/expression"
)

type AstTypeParams struct {
	Name   string
	Fields []string
}

func main() {
	args := os.Args
	if len(args) == 1 {
		println("Usage: generate_ast.go <output directory>")
		os.Exit(64)
	}
	outputDir := args[1]

	defineAst(outputDir, "Expr", []AstTypeParams{
		{"Binary", []string{"Left Expr", "Operator *gdx.Token", "Right Expr"}},
		{"Grouping", []string{"Expression Expr"}},
		{"Literal", []string{"Value interface{}"}},
		{"Unary", []string{"Operator *gdx.Token", "Right Expr"}},
	})

	minus := gdx.NewToken(gdx.MINUS, "-", nil, 1)
	star := gdx.NewToken(gdx.STAR, "*", nil, 1)

	exp := expression.Binary{
		Left: &expression.Unary{
			Operator: &minus,
			Right: &expression.Literal{
				Value: 123,
			},
		},
		Operator: &star,
		Right: &expression.Grouping{
			Expression: &expression.Literal{
				Value: 45.67,
			},
		},
	}
	printer := NewAstPrinter()
	fmt.Println(printer.Print(&exp))
}

func defineAst(outputDir string, baseName string, types []AstTypeParams) {
	path := outputDir + "/" + baseName + ".gen.go"
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// todo: template this shit
	file.WriteString("package expression\n\n")
	file.WriteString("import \"github.com/mrcampbell/gdx/gdx\"\n\n")
	file.WriteString(fmt.Sprintf(`type %s interface {
	Accept(v Visitor) interface{}
}

`, baseName))

	for _, t := range types {
		defineType(file, baseName, t.Name, t.Fields)
	}

	defineVisitor(file, baseName, types)
}

func defineType(file *os.File, baseName string, className string, fields []string) {
	file.WriteString("type " + className + " struct {\n")
	for _, f := range fields {
		file.WriteString("  " + f + "\n")
	}
	file.WriteString("}\n")

	file.WriteString("func (t *" + className + ") Accept(v Visitor) interface{} {\n")
	file.WriteString("  return v.Visit" + className + baseName + "(t)\n")
	file.WriteString("}\n")
}

func defineVisitor(file *os.File, baseName string, types []AstTypeParams) {
	file.WriteString("type Visitor interface {\n")
	for _, t := range types {
		fmt.Printf("%s%s\n", t, baseName)
		file.WriteString(fmt.Sprintf("  Visit%s%s(t *%s) interface{}\n", t.Name, baseName, t.Name))
	}
	file.WriteString("}\n")
}

type AstPrinter struct{}

// todo: depending on the code we're generating.  Not cool
var _ expression.Visitor = &AstPrinter{}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

func (p *AstPrinter) VisitBinaryExpr(t *expression.Binary) interface{} {
	return parenthize(t.Operator.Lexeme(), t.Left, t.Right)
}

func (p *AstPrinter) VisitGroupingExpr(t *expression.Grouping) interface{} {
	return parenthize("group", t.Expression)
}

func (p *AstPrinter) VisitLiteralExpr(t *expression.Literal) interface{} {
	return t.Value
}

func (p *AstPrinter) VisitUnaryExpr(t *expression.Unary) interface{} {
	return parenthize(t.Operator.Lexeme(), t.Right)
}

func (p *AstPrinter) Print(t expression.Expr) string {
	return t.Accept(p).(string)
}

// todo: we're depending on the code we're generating.  Not cool
func parenthize(name string, exprs ...expression.Expr) string {
	result := ""
	result += "("
	result += name
	for _, e := range exprs {
		result += " "
		result += fmt.Sprintf("%v", e.Accept(&AstPrinter{}))
	}
	return result + ")"
}
