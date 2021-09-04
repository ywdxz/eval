%{
package eval
%}

%union {
    value uint64
    ast *ast
}


%type  <ast> expression

%token <value> NUMBER

%left '+' '-'
%left '*' '/' '%' '^'

%start top

%%
top        :expression
                {
                    if l, ok := yylex.(*simpleLex); ok {
                        l.ast = $1
                    }
                }
                ;

expression :'(' expression ')'
           {$$ = $2}
           | expression '+' expression
           {$$ = genOpNode('+',$1,$3)}
           | expression '-' expression
           {$$ = genOpNode('-',$1,$3)}
           | expression '*' expression
           {$$ = genOpNode('*',$1,$3)}
           | expression '/' expression
           {$$ = genOpNode('/',$1,$3)}
           | expression '%' expression
           {$$ = genOpNode('%',$1,$3)}
           | expression '^' expression
           {$$ = genOpNode('^',$1,$3)}
           | NUMBER
           {$$ = genValueNode($1)}
           ;
%%
