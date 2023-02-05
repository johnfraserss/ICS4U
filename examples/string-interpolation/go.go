// String interpolation!
package main

import (
    "fmt"
    "strconv"
)

func main() {
    name := "John Doe"
    age := 20

    // Methods of string interpolation:

    // 1. Concatenation
    // Strings can be concatenated by using plus signs. Make sure to convert other
    // types to strings before concatenating them.
    // You will generally want to use this form of interpolation when you just need to format strings, as
    // it can get annoying to constantly run functions to convert non-string vars to strings.
    str := "Hello! My name is " + name + " and I am " + strconv.Itoa(age) + " years old."
    fmt.Println(str)

    // 2. Printf / Sprintf
    // Similar to C, you can also format strings using these two functions. Printf 
    // will directly print the output, while Sprintf returns the output as a string.
    // 
    // These two functions need to be given a format string, and the values you want
    // in the formatted string. To indicate where to put a variable, you can include a verb.
    // A verb is basically a placeholder symbol that indicates how the value / type of a 
    // variable should be formatted in the string. There are many verbs, but the main two are %v
    // and %T.
    // 
    // You will generally want to use this form of interpolation when you want to format non-string
    // variables, as the verbs manage the conversion and formatting.

    // %v shows the value of the variable in its default format. This can be used
    // to format the values of most variable types. 
    fmt.Printf("Hello! My name is %v and I am %v years old.\n", name, age)

    // %T shows the type of the variable.
    str2 := fmt.Sprintf("Name: %T, Age: %T", name, age)
    fmt.Println(str2)
}
