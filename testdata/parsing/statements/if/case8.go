package main
func main() {
    if e1, e2 *= 2, 3; e1 {
    }
}
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  FunctionDeclaration(main)
    PsiElement(KEYWORD_FUNC)('func')
    PsiWhiteSpace(' ')
    LiteralIdentifierImpl
      PsiElement(IDENTIFIER)('main')
    PsiElement(()('(')
    PsiElement())(')')
    PsiWhiteSpace(' ')
    BlockStmtImpl
      PsiElement({)('{')
      PsiWhiteSpace('\n')
      PsiWhiteSpace('    ')
      IfStmtImpl
        PsiElement(KEYWORD_IF)('if')
        PsiWhiteSpace(' ')
        AssignStmtImpl
          ExpressionListImpl
            LiteralExpressionImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('e1')
            PsiElement(,)(',')
            PsiWhiteSpace(' ')
            LiteralExpressionImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('e2')
          PsiWhiteSpace(' ')
          PsiElement(*=)('*=')
          PsiWhiteSpace(' ')
          ExpressionListImpl
            LiteralExpressionImpl
              LiteralIntegerImpl
                PsiElement(LITERAL_INT)('2')
            PsiElement(,)(',')
            PsiWhiteSpace(' ')
            LiteralExpressionImpl
              LiteralIntegerImpl
                PsiElement(LITERAL_INT)('3')
        PsiElement(;)(';')
        PsiWhiteSpace(' ')
        LiteralExpressionImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('e1')
        PsiWhiteSpace(' ')
        BlockStmtImpl
          PsiElement({)('{')
          PsiWhiteSpace('\n')
          PsiWhiteSpace('    ')
          PsiElement(})('}')
      PsiWhiteSpace('\n')
      PsiElement(})('}')
