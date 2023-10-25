// Code generated from Sfpl.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Sfpl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SfplListener is a complete listener for a parse tree produced by SfplParser.
type SfplListener interface {
	antlr.ParseTreeListener

	// EnterPolicy is called when entering the policy production.
	EnterPolicy(c *PolicyContext)

	// EnterDefs is called when entering the defs production.
	EnterDefs(c *DefsContext)

	// EnterPrule is called when entering the prule production.
	EnterPrule(c *PruleContext)

	// EnterSrule is called when entering the srule production.
	EnterSrule(c *SruleContext)

	// EnterPfilter is called when entering the pfilter production.
	EnterPfilter(c *PfilterContext)

	// EnterSfilter is called when entering the sfilter production.
	EnterSfilter(c *SfilterContext)

	// EnterDrop_keyword is called when entering the drop_keyword production.
	EnterDrop_keyword(c *Drop_keywordContext)

	// EnterPmacro is called when entering the pmacro production.
	EnterPmacro(c *PmacroContext)

	// EnterPlist is called when entering the plist production.
	EnterPlist(c *PlistContext)

	// EnterPreq is called when entering the preq production.
	EnterPreq(c *PreqContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOr_expression is called when entering the or_expression production.
	EnterOr_expression(c *Or_expressionContext)

	// EnterAnd_expression is called when entering the and_expression production.
	EnterAnd_expression(c *And_expressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterItems is called when entering the items production.
	EnterItems(c *ItemsContext)

	// EnterActions is called when entering the actions production.
	EnterActions(c *ActionsContext)

	// EnterTags is called when entering the tags production.
	EnterTags(c *TagsContext)

	// EnterPrefilter is called when entering the prefilter production.
	EnterPrefilter(c *PrefilterContext)

	// EnterSeverity is called when entering the severity production.
	EnterSeverity(c *SeverityContext)

	// EnterEnabled is called when entering the enabled production.
	EnterEnabled(c *EnabledContext)

	// EnterWarnevttype is called when entering the warnevttype production.
	EnterWarnevttype(c *WarnevttypeContext)

	// EnterSkipunknown is called when entering the skipunknown production.
	EnterSkipunknown(c *SkipunknownContext)

	// EnterFappend is called when entering the fappend production.
	EnterFappend(c *FappendContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterBinary_operator is called when entering the binary_operator production.
	EnterBinary_operator(c *Binary_operatorContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// ExitPolicy is called when exiting the policy production.
	ExitPolicy(c *PolicyContext)

	// ExitDefs is called when exiting the defs production.
	ExitDefs(c *DefsContext)

	// ExitPrule is called when exiting the prule production.
	ExitPrule(c *PruleContext)

	// ExitSrule is called when exiting the srule production.
	ExitSrule(c *SruleContext)

	// ExitPfilter is called when exiting the pfilter production.
	ExitPfilter(c *PfilterContext)

	// ExitSfilter is called when exiting the sfilter production.
	ExitSfilter(c *SfilterContext)

	// ExitDrop_keyword is called when exiting the drop_keyword production.
	ExitDrop_keyword(c *Drop_keywordContext)

	// ExitPmacro is called when exiting the pmacro production.
	ExitPmacro(c *PmacroContext)

	// ExitPlist is called when exiting the plist production.
	ExitPlist(c *PlistContext)

	// ExitPreq is called when exiting the preq production.
	ExitPreq(c *PreqContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOr_expression is called when exiting the or_expression production.
	ExitOr_expression(c *Or_expressionContext)

	// ExitAnd_expression is called when exiting the and_expression production.
	ExitAnd_expression(c *And_expressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitItems is called when exiting the items production.
	ExitItems(c *ItemsContext)

	// ExitActions is called when exiting the actions production.
	ExitActions(c *ActionsContext)

	// ExitTags is called when exiting the tags production.
	ExitTags(c *TagsContext)

	// ExitPrefilter is called when exiting the prefilter production.
	ExitPrefilter(c *PrefilterContext)

	// ExitSeverity is called when exiting the severity production.
	ExitSeverity(c *SeverityContext)

	// ExitEnabled is called when exiting the enabled production.
	ExitEnabled(c *EnabledContext)

	// ExitWarnevttype is called when exiting the warnevttype production.
	ExitWarnevttype(c *WarnevttypeContext)

	// ExitSkipunknown is called when exiting the skipunknown production.
	ExitSkipunknown(c *SkipunknownContext)

	// ExitFappend is called when exiting the fappend production.
	ExitFappend(c *FappendContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitBinary_operator is called when exiting the binary_operator production.
	ExitBinary_operator(c *Binary_operatorContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)
}
