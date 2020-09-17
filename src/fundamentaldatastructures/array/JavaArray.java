package fundamentaldatastructures.array;

import java.util.Arrays;

class JavaArray {

    public static void main(String[] args) {
        // create a fix-sized array
        createFixSizedArray();

        // resize array
        resize();
    }

    private static void createFixSizedArray() {
        int[] firstIntArray = new int[3];

        System.out.println(firstIntArray.length);
    }

    private static void resize() {
        int[] originArray = new int[3];
        int[] newArray = Arrays.copyOf(originArray, 5);

        System.out.println(originArray.length + "->" +newArray.length);
    }
}