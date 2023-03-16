package main

import (
	"C"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)
import "fmt"

// var knowledgeBase *ast.KnowledgeBase

//export LoadRulesIntoKnowledgebase
// func LoadRulesIntoKnowledgebase(knowledgeBaseName string, version string, ruleResourceUrl string) {
// 	urlRes := pkg.NewBytesResource([]byte(ruleResourceUrl))
// 	lib := ast.NewKnowledgeLibrary()
// 	rb := builder.NewRuleBuilder(lib)
// 	_ = rb.BuildRuleFromResource(knowledgeBaseName, version, urlRes)
// 	knowledgeBase = lib.NewKnowledgeBaseInstance(knowledgeBaseName, version)
// }

func Match(knowledgeBaseName string, version string, ruleResourceUrl string, factName, factJSON string) []*ast.RuleEntry {
	urlRes := pkg.NewBytesResource([]byte(ruleResourceUrl))
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	_ = rb.BuildRuleFromResource(knowledgeBaseName, version, urlRes)
	knowledgeBase := lib.NewKnowledgeBaseInstance(knowledgeBaseName, version)

	engine := engine.NewGruleEngine()
	//Fact
	dataCtx := ast.NewDataContext()
	_ = dataCtx.AddJSON(factName, []byte(factJSON))

	ruleEntries, err := engine.FetchMatchingRules(dataCtx, knowledgeBase)
	if err != nil {
		panic(err)
	}

	return ruleEntries
}

func main() {
	fmt.Println("Hello")
}
