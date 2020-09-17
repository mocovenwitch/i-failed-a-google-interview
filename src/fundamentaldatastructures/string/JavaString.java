package fundamentaldatastructures.string;

public class JavaString {
    public static void main(String[] args) {

        shared();
        notShared();
    }

    private static void shared() {
        String a = "shared";
        String b = "shared";

        System.out.println( "a and b the same? " + ((a == b) ? "true" : "false"));
        System.out.println(a);
    }

    private static void notShared() {
        String na = new String("hello");
        String nb = new String("hello");

        System.out.println( "na and nb the same? " + ((na == nb) ? "true" : "false"));
        System.out.println(na);
    }
}
