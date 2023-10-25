// Code generated from Sfpl.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Sfpl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSfplListener is a complete listener for a parse tree produced by SfplParser.
type BaseSfplListener struct{}

var _ SfplListener = &BaseSfplListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSfplListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSfplListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSfplListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSfplListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPolicy is called when production policy is entered.
func (s *BaseSfplListener) EnterPolicy(ctx *PolicyContext) {}

// ExitPolicy is called when production policy is exited.
func (s *BaseSfplListener) ExitPolicy(ctx *PolicyContext) {}

// EnterDefs is called when production defs is entered.
func (s *BaseSfplListener) EnterDefs(ctx *DefsContext) {}

// ExitDefs is called when production defs is exited.
func (s *BaseSfplListener) ExitDefs(ctx *DefsContext) {}

// EnterPrule is called when production prule is entered.
func (s *BaseSfplListener) EnterPrule(ctx *PruleContext) {}

// ExitPrule is called when production prule is exited.
func (s *BaseSfplListener) ExitPrule(ctx *PruleContext) {}

// EnterSrule is called when production srule is entered.
func (s *BaseSfplListener) EnterSrule(ctx *SruleContext) {}

// ExitSrule is called when production srule is exited.
func (s *BaseSfplListener) ExitSrule(ctx *SruleContext) {}

// EnterPfilter is called when production pfilter is entered.
func (s *BaseSfplListener) EnterPfilter(ctx *PfilterContext) {}

// ExitPfilter is called when production pfilter is exited.
func (s *BaseSfplListener) ExitPfilter(ctx *PfilterContext) {}

// EnterSfilter is called when production sfilter is entered.
func (s *BaseSfplListener) EnterSfilter(ctx *SfilterContext) {}

// ExitSfilter is called when production sfilter is exited.
func (s *BaseSfplListener) ExitSfilter(ctx *SfilterContext) {}

// EnterDrop_keyword is called when production drop_keyword is entered.
func (s *BaseSfplListener) EnterDrop_keyword(ctx *Drop_keywordContext) {}

// ExitDrop_keyword is called when production drop_keyword is exited.
func (s *BaseSfplListener) ExitDrop_keyword(ctx *Drop_keywordContext) {}

// EnterPmacro is called when production pmacro is entered.
func (s *BaseSfplListener) EnterPmacro(ctx *PmacroContext) {}

// ExitPmacro is called when production pmacro is exited.
func (s *BaseSfplListener) ExitPmacro(ctx *PmacroContext) {}

// EnterPlist is called when production plist is entered.
func (s *BaseSfplListener) EnterPlist(ctx *PlistContext) {}

// ExitPlist is called when production plist is exited.
func (s *BaseSfplListener) ExitPlist(ctx *PlistContext) {}

// EnterPreq is called when production preq is entered.
func (s *BaseSfplListener) EnterPreq(ctx *PreqContext) {}

// ExitPreq is called when production preq is exited.
func (s *BaseSfplListener) ExitPreq(ctx *PreqContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSfplListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSfplListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOr_expression is called when production or_expression is entered.
func (s *BaseSfplListener) EnterOr_expression(ctx *Or_expressionContext) {}

// ExitOr_expression is called when production or_expression is exited.
func (s *BaseSfplListener) ExitOr_expression(ctx *Or_expressionContext) {}

// EnterAnd_expression is called when production and_expression is entered.
func (s *BaseSfplListener) EnterAnd_expression(ctx *And_expressionContext) {}

// ExitAnd_expression is called when production and_expression is exited.
func (s *BaseSfplListener) ExitAnd_expression(ctx *And_expressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSfplListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSfplListener) ExitTerm(ctx *TermContext) {}

// EnterItems is called when production items is entered.
func (s *BaseSfplListener) EnterItems(ctx *ItemsContext) {}

// ExitItems is called when production items is exited.
func (s *BaseSfplListener) ExitItems(ctx *ItemsContext) {}

// EnterActions is called when production actions is entered.
func (s *BaseSfplListener) EnterActions(ctx *ActionsContext) {}

// ExitActions is called when production actions is exited.
func (s *BaseSfplListener) ExitActions(ctx *ActionsContext) {}

// EnterTags is called when production tags is entered.
func (s *BaseSfplListener) EnterTags(ctx *TagsContext) {}

// ExitTags is called when production tags is exited.
func (s *BaseSfplListener) ExitTags(ctx *TagsContext) {}

// EnterPrefilter is called when production prefilter is entered.
func (s *BaseSfplListener) EnterPrefilter(ctx *PrefilterContext) {}

// ExitPrefilter is called when production prefilter is exited.
func (s *BaseSfplListener) ExitPrefilter(ctx *PrefilterContext) {}

// EnterSeverity is called when production severity is entered.
func (s *BaseSfplListener) EnterSeverity(ctx *SeverityContext) {}

// ExitSeverity is called when production severity is exited.
func (s *BaseSfplListener) ExitSeverity(ctx *SeverityContext) {}

// EnterEnabled is called when production enabled is entered.
func (s *BaseSfplListener) EnterEnabled(ctx *EnabledContext) {}

// ExitEnabled is called when production enabled is exited.
func (s *BaseSfplListener) ExitEnabled(ctx *EnabledContext) {}

// EnterWarnevttype is called when production warnevttype is entered.
func (s *BaseSfplListener) EnterWarnevttype(ctx *WarnevttypeContext) {}

// ExitWarnevttype is called when production warnevttype is exited.
func (s *BaseSfplListener) ExitWarnevttype(ctx *WarnevttypeContext) {}

// EnterSkipunknown is called when production skipunknown is entered.
func (s *BaseSfplListener) EnterSkipunknown(ctx *SkipunknownContext) {}

// ExitSkipunknown is called when production skipunknown is exited.
func (s *BaseSfplListener) ExitSkipunknown(ctx *SkipunknownContext) {}

// EnterFappend is called when production fappend is entered.
func (s *BaseSfplListener) EnterFappend(ctx *FappendContext) {}

// ExitFappend is called when production fappend is exited.
func (s *BaseSfplListener) ExitFappend(ctx *FappendContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSfplListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSfplListener) ExitVariable(ctx *VariableContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseSfplListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseSfplListener) ExitAtom(ctx *AtomContext) {}

// EnterText is called when production text is entered.
func (s *BaseSfplListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseSfplListener) ExitText(ctx *TextContext) {}

// EnterBinary_operator is called when production binary_operator is entered.
func (s *BaseSfplListener) EnterBinary_operator(ctx *Binary_operatorContext) {}

// ExitBinary_operator is called when production binary_operator is exited.
func (s *BaseSfplListener) ExitBinary_operator(ctx *Binary_operatorContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseSfplListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseSfplListener) ExitUnary_operator(ctx *Unary_operatorContext) {}
