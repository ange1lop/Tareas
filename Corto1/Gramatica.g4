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
        dir float64
        tipo float64
        cadena string
    }

    var desp int = 0

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
    : listas
    | EOF
    ;

listas
    : t=listaA PUNTO id=listaA                  {
                                        sim :=  $t.nodo.dir + $id.nodo.tipo+0.00
                                        cad := fmt.Sprint(sim)
                                        gen(cad)
                                    }
    | t=listaA  {
                cad := fmt.Sprintf("%f", $t.nodo.dir)
                gen(cad)
    }
    ;

listaA returns[ *Nodo nodo]
    : l=listaA el=dA { $nodo = &Nodo { }
                        $nodo.dir = $l.nodo.dir *16+$el.nodo.dir
                        divisor := $l.nodo.tipo/16
                        $nodo.tipo = $l.nodo.tipo + $el.nodo.dir *divisor

    }
    | el=dA { 
            $nodo = &Nodo { }
            din2 := 1/16.00
            $nodo.tipo = $el.nodo.tipo * din2
            $nodo.dir = $el.nodo.dir
    }
    ;

dA returns [*Nodo nodo]
    : e=NUM { $nodo = &Nodo { }
              strVar := $e.text
	          intVar, err := strconv.ParseFloat(strVar,8)
              if(err != nil){
                  gen("error en conversion")
              }
              $nodo.dir = intVar
              $nodo.tipo = intVar
              }
    | e=LA  { 
            $nodo = &Nodo { }
            $nodo.dir = 10.00
            $nodo.tipo = 10.00
    }
    | e=LB  { 
            $nodo = &Nodo { }
            $nodo.dir = 11.00
            $nodo.tipo = 11.00
    }
    | e=LC  { 
            $nodo = &Nodo { }
            $nodo.dir = 12.00
            $nodo.tipo = 12.00
    }
    | e=LD  { 
            $nodo = &Nodo { }
            $nodo.dir = 13.00
            $nodo.tipo = 13.00
    }
    | e=LE  { 
            $nodo = &Nodo { }
            $nodo.dir = 14.00
            $nodo.tipo = 14.00
    }
    | e=LF  { 
            $nodo = &Nodo { }
            $nodo.dir = 15.00
            $nodo.tipo = 15.00
    }
    ;

// Tokens
MAS: '+';
MEN: '-';
POR: '*';
DIV: '/';
PARI: '(';
PARD: ')';
PUNTO: '.';
NUM: [0-9];
LA: 'A';
LB: 'B';
LC: 'C';
LD: 'D';
LE: 'E';
LF: 'F';
WHITESPACE: [ \r\n\t]+ -> skip;