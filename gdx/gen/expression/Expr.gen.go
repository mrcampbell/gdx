package expression

import "github.com/mrcampbell/gdx/gdx"

type Expr interface {
	Accept(v Visitor) interface{}
}

type Binary struct {
  Left Expr
  Operator *gdx.Token
  Right Expr
}
func (t *Binary) Accept(v Visitor) interface{} {
  return v.VisitBinaryExpr(t)
}
type Grouping struct {
  Expression Expr
}
func (t *Grouping) Accept(v Visitor) interface{} {
  return v.VisitGroupingExpr(t)
}
type Literal struct {
  Value interface{}
}
func (t *Literal) Accept(v Visitor) interface{} {
  return v.VisitLiteralExpr(t)
}
type Unary struct {
  Operator *gdx.Token
  Right Expr
}
func (t *Unary) Accept(v Visitor) interface{} {
  return v.VisitUnaryExpr(t)
}
type Visitor interface {
  VisitBinaryExpr(t *Binary) interface{}
  VisitGroupingExpr(t *Grouping) interface{}
  VisitLiteralExpr(t *Literal) interface{}
  VisitUnaryExpr(t *Unary) interface{}
}
