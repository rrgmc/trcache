package main

import (
	"fmt"
	"go/types"

	"github.com/dave/jennifer/jen"
)

type qualFromTypeOptions struct {
	variadic bool
}

type QualFromTypeOption func(*qualFromTypeOptions)

func WithQualFromTypeVariadic(variadic bool) QualFromTypeOption {
	return func(o *qualFromTypeOptions) {
		o.variadic = true
	}
}

func QualFromType(typ types.Type, options ...QualFromTypeOption) *jen.Statement {
	var optns qualFromTypeOptions
	for _, opt := range options {
		opt(&optns)
	}

	switch t := typ.(type) {
	case nil:
	case *types.Basic:
		return jen.Id(t.Name())
	case *types.Array:
		return jen.Index(jen.Lit(t.Len())).Add(QualFromType(t.Elem()))
	case *types.Slice:
		if optns.variadic {
			return jen.Op("...").Add(QualFromType(t.Elem()))
		}
		return jen.Index().Add(QualFromType(t.Elem()))
	case *types.Struct:
		return jen.StructFunc(func(g *jen.Group) {
			for fieldIdx := 0; fieldIdx < t.NumFields(); fieldIdx++ {
				f := t.Field(fieldIdx)
				if f.Anonymous() {
					g.Add(QualFromType(f.Type()))
				} else {
					g.Id(f.Name()).Add(QualFromType(f.Type()))
				}
			}
		})
	case *types.Pointer:
		return jen.Op("*").Add(QualFromType(t.Elem()))
	case *types.Tuple:
		// TODO
	case *types.Signature:
		ret := jen.Func().TypesFunc(func(g *jen.Group) {
			for inIdx := 0; inIdx < t.TypeParams().Len(); inIdx++ {
				tp := t.TypeParams().At(inIdx)
				fid := g.Id(tp.String())

				switch tp.Underlying().String() {
				case "interface{comparable}":
					fid.Comparable()
				case "any":
					fid.Any()
				default:
					fid.Id(tp.String())
				}
			}
		}).ParamsFunc(func(g *jen.Group) {
			for inIdx := 0; inIdx < t.Params().Len(); inIdx++ {
				fname := t.Params().At(inIdx).Name()
				if fname == "" {
					fname = fmt.Sprintf("p%d", inIdx)
				}
				g.Id(fname).Add(QualFromType(t.Params().At(inIdx).Type()))
			}
		})
		if t.Results().Len() == 1 {
			return ret.Add(QualFromType(t.Results().At(0).Type()))
		} else if t.Results().Len() > 1 {
			ret.Parens(jen.ListFunc(func(g *jen.Group) {
				for outIdx := 0; outIdx < t.Results().Len(); outIdx++ {
					g.Add(QualFromType(t.Results().At(outIdx).Type()))
				}
			}))
		}
		return ret

	case *types.Union:
		// TODO
	case *types.Interface:
		return jen.InterfaceFunc(func(g *jen.Group) {
			for methodIdx := 0; methodIdx < t.NumMethods(); methodIdx++ {
				m := t.Method(methodIdx)
				g.Id(m.Name()).ParamsFunc(func(g *jen.Group) {
					for inIdx := 0; inIdx < m.Type().(*types.Signature).Params().Len(); inIdx++ {
						g.Add(QualFromType(m.Type().(*types.Signature).Params().At(inIdx).Type()))
					}
				}).ParamsFunc(func(group *jen.Group) {
					for outIdx := 0; outIdx < m.Type().(*types.Signature).Results().Len(); outIdx++ {
						group.Add(QualFromType(m.Type().(*types.Signature).Results().At(outIdx).Type()))
					}
				})
			}
		})
	case *types.Map:
		return jen.Map(QualFromType(t.Key())).Add(QualFromType(t.Elem()))
	case *types.Chan:
		switch t.Dir() {
		case types.RecvOnly:
			return jen.Op("<-").Chan().Add(QualFromType(t.Elem()))
		case types.SendOnly:
			return jen.Chan().Op("<-").Add(QualFromType(t.Elem()))
		case types.SendRecv:
			return jen.Chan().Add(QualFromType(t.Elem()))
		default:
			panic(fmt.Errorf("unexpected ChanDir: %v", t.Dir()))
		}
	case *types.Named:
		var ftypes []jen.Code
		for p := 0; p < t.TypeParams().Len(); p++ {
			fid := jen.Id(t.TypeParams().At(p).String())
			ftypes = append(ftypes, fid)
		}

		if t.Obj().Pkg() == nil {
			return jen.Id(t.Obj().Name()).Types(ftypes...)
		}
		return jen.Qual(t.Obj().Pkg().Path(), t.Obj().Name()).Types(ftypes...)
	case *types.TypeParam:
		if t.Obj() == nil {
			return jen.Id("UNNAMED")
		}
		return jen.Qual(t.Obj().Pkg().Path(), t.Obj().Name())
	}

	panic(fmt.Errorf("unknown go type kind: %v", typ))
}

func FromTypeParam(tp *types.TypeParam) *jen.Statement {
	switch tp.Underlying().String() {
	case "interface{comparable}":
		return jen.Comparable()
	case "any":
		return jen.Any()
	default:
		return jen.Id(tp.String())
	}
}

func FromTypeParams(t *types.TypeParamList) *jen.Statement {
	return jen.TypesFunc(func(g *jen.Group) {
		for inIdx := 0; inIdx < t.Len(); inIdx++ {
			tp := t.At(inIdx)
			g.Id(tp.String()).Add(FromTypeParam(tp))
		}
	})
}

func CallFromTypeParams(t *types.TypeParamList) *jen.Statement {
	return jen.TypesFunc(func(g *jen.Group) {
		for inIdx := 0; inIdx < t.Len(); inIdx++ {
			tp := t.At(inIdx)
			g.Id(tp.String())
		}
	})
}

func CallFromTypeParamsFixed(t *types.TypeParamList, fixedType string) *jen.Statement {
	return jen.TypesFunc(func(g *jen.Group) {
		for inIdx := 0; inIdx < t.Len(); inIdx++ {
			g.Id(fixedType)
		}
	})
}

func FromParams(params *types.Tuple, variadic bool) *jen.Statement {
	return jen.ParamsFunc(func(g *jen.Group) {
		for p := 0; p < params.Len(); p++ {
			prm := params.At(p)
			fname := prm.Name()
			if fname == "" {
				fname = fmt.Sprintf("p%d", p)
			}
			var qparams []QualFromTypeOption
			if variadic && p == params.Len()-1 {
				qparams = append(qparams, WithQualFromTypeVariadic(true))
			}
			g.Id(fname).Add(QualFromType(prm.Type(), qparams...))
		}
	})
}

func CallFromParams(params *types.Tuple, variadic bool) *jen.Statement {
	return jen.CallFunc(func(g *jen.Group) {
		for p := 0; p < params.Len(); p++ {
			prm := params.At(p)
			fname := prm.Name()
			if fname == "" {
				fname = fmt.Sprintf("p%d", p)
			}
			if variadic && p == params.Len()-1 {
				g.Id(fname).Op("...")
			} else {
				g.Id(fname)
			}
		}
	})
}

func ParamName(prm *types.Var, paramIdx int) string {
	fname := prm.Name()
	if fname == "" {
		fname = fmt.Sprintf("p%d", paramIdx)
	}
	return fname
}
