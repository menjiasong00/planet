package iparser

import parser "planet/pkg/grule/internal/iantlr/alr"

type GengineParserVisitor struct {
	parser.BasegengineVisitor
}

func NewGengineParserVisitor() *GengineParserVisitor {
	return &GengineParserVisitor{}
}
