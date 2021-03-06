
//----------------------------------------------------
// The following code was generated by CUP v0.11b 20160615 (GIT 4ac7450)
//----------------------------------------------------

package analizadores;

import java_cup.runtime.*;
import java_cup.runtime.XMLElement;

/** CUP v0.11b 20160615 (GIT 4ac7450) generated parser.
  */
@SuppressWarnings({"rawtypes"})
public class Parser extends java_cup.runtime.lr_parser {

 public final Class getSymbolContainer() {
    return sym.class;
}

  /** Default constructor. */
  @Deprecated
  public Parser() {super();}

  /** Constructor which sets the default scanner. */
  @Deprecated
  public Parser(java_cup.runtime.Scanner s) {super(s);}

  /** Constructor which sets the default scanner. */
  public Parser(java_cup.runtime.Scanner s, java_cup.runtime.SymbolFactory sf) {super(s,sf);}

  /** Production table. */
  protected static final short _production_table[][] = 
    unpackFromStrings(new String[] {
    "\000\010\000\002\002\004\000\002\002\003\000\002\005" +
    "\005\000\002\005\003\000\002\003\004\000\002\003\003" +
    "\000\002\004\003\000\002\004\003" });

  /** Access to production table. */
  public short[][] production_table() {return _production_table;}

  /** Parse-action table. */
  protected static final short[][] _action_table = 
    unpackFromStrings(new String[] {
    "\000\013\000\006\005\010\006\004\001\002\000\012\002" +
    "\ufffa\004\ufffa\005\ufffa\006\ufffa\001\002\000\012\002\ufffc" +
    "\004\ufffc\005\ufffc\006\ufffc\001\002\000\004\002\015\001" +
    "\002\000\004\002\000\001\002\000\012\002\ufffb\004\ufffb" +
    "\005\ufffb\006\ufffb\001\002\000\012\002\ufffe\004\013\005" +
    "\010\006\004\001\002\000\012\002\ufffd\004\ufffd\005\ufffd" +
    "\006\ufffd\001\002\000\006\005\010\006\004\001\002\000" +
    "\010\002\uffff\005\010\006\004\001\002\000\004\002\001" +
    "\001\002" });

  /** Access to parse-action table. */
  public short[][] action_table() {return _action_table;}

  /** <code>reduce_goto</code> table. */
  protected static final short[][] _reduce_table = 
    unpackFromStrings(new String[] {
    "\000\013\000\012\002\005\003\010\004\004\005\006\001" +
    "\001\000\002\001\001\000\002\001\001\000\002\001\001" +
    "\000\002\001\001\000\002\001\001\000\004\004\011\001" +
    "\001\000\002\001\001\000\006\003\013\004\004\001\001" +
    "\000\004\004\011\001\001\000\002\001\001" });

  /** Access to <code>reduce_goto</code> table. */
  public short[][] reduce_table() {return _reduce_table;}

  /** Instance of action encapsulation class. */
  protected CUP$Parser$actions action_obj;

  /** Action encapsulation object initializer. */
  protected void init_actions()
    {
      action_obj = new CUP$Parser$actions(this);
    }

  /** Invoke a user supplied parse action. */
  public java_cup.runtime.Symbol do_action(
    int                        act_num,
    java_cup.runtime.lr_parser parser,
    java.util.Stack            stack,
    int                        top)
    throws java.lang.Exception
  {
    /* call code in generated class */
    return action_obj.CUP$Parser$do_action(act_num, parser, stack, top);
  }

  /** Indicates start state. */
  public int start_state() {return 0;}
  /** Indicates start production. */
  public int start_production() {return 0;}

  /** <code>EOF</code> Symbol index. */
  public int EOF_sym() {return 0;}

  /** <code>error</code> Symbol index. */
  public int error_sym() {return 1;}



    /**
     * M??todo al que se llama autom??ticamente ante alg??n error sintactico.
     **/ 
    public void syntax_error(Symbol s){ 
        System.out.println("Error Sint??ctico en la L??nea " + (s.left) +
        " Columna "+s.right+ ". No se esperaba este componente: " +s.value+"."); 
    } 

    /**
     * M??todo al que se llama autom??ticamente ante alg??n error sint??ctico 
     * en el que ya no es posible una recuperaci??n de errores.
     **/ 
    public void unrecovered_syntax_error(Symbol s) throws java.lang.Exception{ 
        System.out.println("Error s??ntactico irrecuperable en la L??nea " + 
        (s.left)+ " Columna "+s.right+". Componente " + s.value + 
        " no reconocido."); 
    }  

    public void mostrar(String out) {
        System.out.println(out);
    }

    public int getVal(String num) {
        return Integer.parseInt(num);
    }


/** Cup generated class to encapsulate user supplied action code.*/
@SuppressWarnings({"rawtypes", "unchecked", "unused"})
class CUP$Parser$actions {
  private final Parser parser;

  /** Constructor */
  CUP$Parser$actions(Parser parser) {
    this.parser = parser;
  }

  /** Method 0 with the actual generated action code for actions 0 to 300. */
  public final java_cup.runtime.Symbol CUP$Parser$do_action_part00000000(
    int                        CUP$Parser$act_num,
    java_cup.runtime.lr_parser CUP$Parser$parser,
    java.util.Stack            CUP$Parser$stack,
    int                        CUP$Parser$top)
    throws java.lang.Exception
    {
      /* Symbol object for return from actions */
      java_cup.runtime.Symbol CUP$Parser$result;

      /* select the action based on the action number */
      switch (CUP$Parser$act_num)
        {
          /*. . . . . . . . . . . . . . . . . . . .*/
          case 0: // $START ::= ini EOF 
            {
              Object RESULT =null;
		int start_valleft = ((java_cup.runtime.Symbol)CUP$Parser$stack.elementAt(CUP$Parser$top-1)).left;
		int start_valright = ((java_cup.runtime.Symbol)CUP$Parser$stack.elementAt(CUP$Parser$top-1)).right;
		Nodo start_val = (Nodo)((java_cup.runtime.Symbol) CUP$Parser$stack.elementAt(CUP$Parser$top-1)).value;
		RESULT = start_val;
              CUP$Parser$result = parser.getSymbolFactory().newSymbol("$START",0, ((java_cup.runtime.Symbol)CUP$Parser$stack.elementAt(CUP$Parser$top-1)), ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), RESULT);
            }
          /* ACCEPT */
          CUP$Parser$parser.done_parsing();
          return CUP$Parser$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 1: // ini ::= listas 
            {
              Nodo RESULT =null;
		int Lleft = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).left;
		int Lright = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).right;
		Nodo L = (Nodo)((java_cup.runtime.Symbol) CUP$Parser$stack.peek()).value;
		 mostrar("El numero es " + L.cad); 
              CUP$Parser$result = parser.getSymbolFactory().newSymbol("ini",0, ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), RESULT);
            }
          return CUP$Parser$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 2: // listas ::= listaA PUNTO listaA 
            {
              Nodo RESULT =null;
		int L1left = ((java_cup.runtime.Symbol)CUP$Parser$stack.elementAt(CUP$Parser$top-2)).left;
		int L1right = ((java_cup.runtime.Symbol)CUP$Parser$stack.elementAt(CUP$Parser$top-2)).right;
		Nodo L1 = (Nodo)((java_cup.runtime.Symbol) CUP$Parser$stack.elementAt(CUP$Parser$top-2)).value;
		int Aleft = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).left;
		int Aright = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).right;
		Nodo A = (Nodo)((java_cup.runtime.Symbol) CUP$Parser$stack.peek()).value;
		
                                    L1.aux = L1.cont+L1.val;
                                    L1.cad = ""+L1.aux;
                                    RESULT = L1;
                                
              CUP$Parser$result = parser.getSymbolFactory().newSymbol("listas",3, ((java_cup.runtime.Symbol)CUP$Parser$stack.elementAt(CUP$Parser$top-2)), ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), RESULT);
            }
          return CUP$Parser$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 3: // listas ::= listaA 
            {
              Nodo RESULT =null;
		int Aleft = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).left;
		int Aright = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).right;
		Nodo A = (Nodo)((java_cup.runtime.Symbol) CUP$Parser$stack.peek()).value;
		 
                                    A.cad = ""+A.cont;
                                    RESULT = A; 
              CUP$Parser$result = parser.getSymbolFactory().newSymbol("listas",3, ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), RESULT);
            }
          return CUP$Parser$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 4: // listaA ::= listaA deA 
            {
              Nodo RESULT =null;
		int L1left = ((java_cup.runtime.Symbol)CUP$Parser$stack.elementAt(CUP$Parser$top-1)).left;
		int L1right = ((java_cup.runtime.Symbol)CUP$Parser$stack.elementAt(CUP$Parser$top-1)).right;
		Nodo L1 = (Nodo)((java_cup.runtime.Symbol) CUP$Parser$stack.elementAt(CUP$Parser$top-1)).value;
		int Bleft = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).left;
		int Bright = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).right;
		Nodo B = (Nodo)((java_cup.runtime.Symbol) CUP$Parser$stack.peek()).value;
		
                                    L1.cont = L1.cont * 2+B.cont;
                                    L1.aux = L1.aux/2;
                                    L1.val = L1.val + B.val*L1.aux;
                                    RESULT = L1;
                                
              CUP$Parser$result = parser.getSymbolFactory().newSymbol("listaA",1, ((java_cup.runtime.Symbol)CUP$Parser$stack.elementAt(CUP$Parser$top-1)), ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), RESULT);
            }
          return CUP$Parser$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 5: // listaA ::= deA 
            {
              Nodo RESULT =null;
		int Aleft = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).left;
		int Aright = ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()).right;
		Nodo A = (Nodo)((java_cup.runtime.Symbol) CUP$Parser$stack.peek()).value;
		  
                                    A.aux = 0.5;
                                    A.val = A.aux*A.val;
                                    RESULT = A;
                                
              CUP$Parser$result = parser.getSymbolFactory().newSymbol("listaA",1, ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), RESULT);
            }
          return CUP$Parser$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 6: // deA ::= CERO 
            {
              Nodo RESULT =null;
		
                                    Nodo nodo = new Nodo();
                                    nodo.cont = 0;
                                    nodo.aux =0;
                                    nodo.val = 0;
                                    RESULT = nodo;
                                
              CUP$Parser$result = parser.getSymbolFactory().newSymbol("deA",2, ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), RESULT);
            }
          return CUP$Parser$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 7: // deA ::= UNO 
            {
              Nodo RESULT =null;
		
                                    Nodo nodo = new Nodo();
                                    nodo.aux =0;
                                    nodo.cont = 1;
                                    nodo.val = 1;
                                    RESULT = nodo;
                                
              CUP$Parser$result = parser.getSymbolFactory().newSymbol("deA",2, ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), ((java_cup.runtime.Symbol)CUP$Parser$stack.peek()), RESULT);
            }
          return CUP$Parser$result;

          /* . . . . . .*/
          default:
            throw new Exception(
               "Invalid action number "+CUP$Parser$act_num+"found in internal parse table");

        }
    } /* end of method */

  /** Method splitting the generated action code into several parts. */
  public final java_cup.runtime.Symbol CUP$Parser$do_action(
    int                        CUP$Parser$act_num,
    java_cup.runtime.lr_parser CUP$Parser$parser,
    java.util.Stack            CUP$Parser$stack,
    int                        CUP$Parser$top)
    throws java.lang.Exception
    {
              return CUP$Parser$do_action_part00000000(
                               CUP$Parser$act_num,
                               CUP$Parser$parser,
                               CUP$Parser$stack,
                               CUP$Parser$top);
    }
}

}
