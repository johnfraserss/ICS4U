import 'dart:io';

int fibonacci(int n, var memoizationTable) {
  if (n == 0 || n == 1) return 1;

  if (memoizationTable[n] == null)
    memoizationTable[n] =
        fibonacci(n - 2, memoizationTable) + fibonacci(n - 1, memoizationTable);
  return memoizationTable[n];
}

void main() {
  // Example use case
  var result = fibonacci(5, {});
  print("6th term is " + result.toString());

  // Print first 20 items in sequence
  for (int i = 0; i < 20; i++)
    stdout.write(fibonacci(i, {}).toString() + (i == 19 ? '' : ', '));
}
