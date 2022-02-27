grammar Gramatica;
/*
E		E + E
	|	E – E
	|	E * E
	|	E / E
	|	( E )
	|	NUM

*/

options {
    language='Go';
}

@parser::header {
    import "Laboratorio5/entorno"
}

@parser::members{
    type Nodo struct {
        dir string
        tipo string
    }

    var desp int = 0
    var tmp int = 1
    // var env []*entorno.Entorno
    var sup *entorno.Entorno

    // func push(e *entorno.Entorno) {
    //     env = append(env, e)
    // }

    // func pop() *entorno.Entorno {
    //     if len(env) < 1 {
    //         panic("empty env")
    //     }
    //     result := env[len(env) - 1]
    //     env = env[:len(env) - 1]
    //     return result
    // }

    func gen(out string) {
        fmt.Println(out);
    }
}

start
    : exp2
    | EOF
    ;
exp2 returns [string dir,int num]
    : e1=exp2 op=('*'|'/') e2=exp2{
        tmp = tmp - $e1.num - $e2.num
        $num = 1
        $dir = "t" + strconv.Itoa(tmp)
        gen($dir + "=" + $e1.dir + $op.text + $e2.dir)
        tmp = tmp + 1
    }
    | e1=exp2 op=('+'|'-') e2=exp2{
        tmp = tmp - $e1.num - $e2.num
        $num = 1
        $dir = "t" + strconv.Itoa(tmp)
        gen($dir + "=" + $e1.dir + $op.text + $e2.dir)
        tmp = tmp + 1
    }
    | PARI e=exp2 PARD {
        $num = $e.num
        $dir = $e.dir
    }
    | n=NUM { $num = 0
            $dir = $n.text
            }
    | n=ID { $num = 0
            $dir = $n.text
            }
    ;
/*
marcador
    : { sup = entorno.NewEntorno(nil); desp = 0; }
    ;

dec
    : t=tipo id=ID                  {
                                        sim := entorno.NewSimbolo($id.text, $t.tipoDato, desp)
                                        sup.Put($id.text, sim)
                                        desp = desp + 1
                                    }
        ';' dec
    | 
    ;

tipo returns[ string tipoDato ]
    : t='int'                       {
                                        $tipoDato = $t.text
                                    }
    ;

exp returns [ *Nodo nodo ]
    : e1=exp op=('*'|'/') e2=exp    {
                                        $nodo = &Nodo { }
                                        $nodo.dir = generador.NewTemp()
                                        cad := $nodo.dir + " = " + $e1.nodo.dir + " " + $op.text + " "
                                        cad += $e2.nodo.dir
                                        gen(cad)
                                    }
    | e1=exp op=('+'|'-') e2=exp    {
                                        $nodo = &Nodo { }
                                        $nodo.dir = generador.NewTemp()
                                        cad := $nodo.dir + " = " + $e1.nodo.dir + " " + $op.text + " "
                                        cad += $e2.nodo.dir
                                        gen(cad)
                                    }
    | '(' e=exp ')'                 {
                                        $nodo = &Nodo { }
                                        $nodo.dir = $e.nodo.dir
                                    }
    | n=NUM                         {
                                        $nodo = &Nodo { }
                                        $nodo.dir = $n.text
                                    }
    ;
*/

// Tokens
MAS: '+';
MEN: '-';
POR: '*';
DIV: '/';
PARI: '(';
PARD: ')';
NUM: [0-9]+;
ID: [_A-Za-z][_A-Za-z0-9]*;
WHITESPACE: [ \r\n\t]+ -> skip;