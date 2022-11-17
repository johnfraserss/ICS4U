import 'dart:io';

int fibonacciSeq(int n, var memoTable) {
  if (n == 0 || n == 1) return 1;

  if (memoTable[n] == null)
    memoTable[n] =
        fibonacciSeq(n - 2, memoTable) + fibonacciSeq(n - 1, memoTable);
  return memoTable[n];
}

void main() {
  // Example use case
  var result = fibonacciSeq(5, {});
  print("6th term is " + result.toString());

  // Print first 20 items in sequence
  for (int i = 0; i < 20; i++)
    stdout.write(fibonacciSeq(i, {}).toString() + (i == 19 ? '' : ', '));
}
