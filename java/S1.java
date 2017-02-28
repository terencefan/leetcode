class A {

  public int a;

  public int b;

  A() {
    a = 1;
    b = 3;
  }

}

class B extends A {

  public int a;

  B() {
    a = 10;
  }

}

public class S1 {

  public static void main(String[] args) {
    B b = new B();
    A a = b;
    boolean c = a == b;
    System.out.println(c);
    System.out.println(a.a);
    System.out.println(b.a);
  }

}
