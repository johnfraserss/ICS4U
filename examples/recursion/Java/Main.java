/**
 * Recursion
 * 
 * Recursion is essentially just calling a function in itself until a certain
 * value or expected response is returned.
 * 
 * In this example, we'll use factorials as an example. Factorials (in case you
 * weren't aware already) are mathematically represented by `N!` where N is a
 * positive integer, and ! is the factorial operator.
 * 
 * Take `4!`, this is equal to 4 x 3 x 2 x 1, or 24
 * 
 * So how do we implement this in code?
 */


class Main {

  /**
   * Recursively finds the factorial of N
   * 
   * @param N a positive integer
   * @param product the product of previous steps, 1 for the first
   * @return factorial of N
   */
  public static int recursiveFactorial(int N, int product) {
    /**
     * Start off by checking to see if N is equal to 1. If it is, we can just return
     * the product since any number times 1 is equal to that number.
     */
    if (N == 1) {
      // notice that the default value of product is set to 1, so that 1! = 1, and
      // that N*1 = N
      return product;
    } else {
      /**
       * If N isn't equal to one, there's more numbers to multiply We pass N-1 into
       * `recursiveFactorial` as N, as well as the product, which is just product*N.
       */
      return recursiveFactorial(N - 1, product * N);
    }
  }

  public static void main(String[] args) {
    /**
     * Naive approach
     * 
     * In this approach, we're given N, and we store the factorial of N in a
     * variable The factorial of N is manually calculated by typing everything out.
     * 
     * As you may have predicted, this won't hold true for all integers N, if N was
     * 5, then `N!` wouldn't be equivalent to 4*3*2*1
     */
    int N = 4;
    int factorialOfN = 4*3*2*1;

    // driver code for testing

    // note that we pass 1 as the product initially
    System.out.println(recursiveFactorial(4,1));

  }
}

// Still confused? There's a lengthy explanation below

/**
 * Let's take the driver code/example there, finding the factorial of 4.
 * 
 * We start off by passing 6 into `recursiveFactorial`, 6 !== 1,  so we pass
 * 4-1, which is 3, into recursiveFactorial, as well as the product, which is equal to 1*4, or just 4.
 * At N = 4:
 *    product = 4, N = 3
 * 
 * `recursiveFactorial` is called again with the above parameters.
 * N !== 1, which means we pass N-1, which is 2, into `recursiveFactorial` with the product of 
 * 4*3, which is 12.
 * At N = 3:
 *    product = 12, N = 2
 * 
 * `recursiveFactorial` is called again with the above parameters.
 * N !== 1, which means we pass N-1, which is 1, into `recursiveFactorial` with the product of 
 * 12*2, which is 24.
 * At N = 2:
 *    product = 24, N = 1
 * 
 * `recursiveFactorial` is called again with the above parameters.
 * N === 1, which means we return the product, 24
 * 
 * 24 === 4!
 */