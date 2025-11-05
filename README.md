| Benchmark | glisp (ns/op) | goja (ns/op) | lua (ns/op) | zygo (ns/op) | Winner |
|:---|:---:|:---:|:---:|:---:|:---:|
| Factorial Calculation | 6546 | 3357 | **1905** | 42486 | lua |
| Regular Expression Matching | **583.8** | 877.5 | 828.8 | 2149 | glisp |
| Complex Conditions | 771.0 | **490.0** | 499.7 | 4692 | goja |
| Time Formatting | **1165** | 1661 | 1342 | N/A | glisp |
| Hash Write | 755.0 | 542.7 | **523.6** | 2583 | lua |
| Hash Delete | 535.1 | 434.6 | **432.4** | 1933 | lua |
| Hash Access | 493.3 | **448.8** | 506.5 | 2055 | goja |
| JSON Parsing and Modification | 4444 | 7088 | **3029** | 12230 | lua |
| String Concat | 616.5 | 769.2 | **403.1** | 11528 | lua |

**Conclusion:**

*   **Lua:** Performs best in computationally intensive tasks (factorial and JSON operations), hash writes, hash deletes and string concat.
*   **glisp:** Excels in string operations (regular expressions) and time formatting.
*   **Goja:** Fastest in handling complex logical conditions and hash accesses.
*   **Zygo:** Performed the worst in all tests.

**glisp vs. zygo Performance Discussion:**

glisp and zygo are developed based on the same kernel (https://github.com/zhemao/glisp). According to the performance test results, glisp significantly outperforms zygo in all test scenarios. This is mainly due to targeted optimizations made in glisp in the following areas:

*   **Built-in Function Optimization:** glisp has rewritten and optimized commonly used built-in functions, reducing unnecessary type conversions and memory allocations.
*   **Virtual Machine Instruction Set Optimization:** glisp has optimized the virtual machine's instruction set, allowing certain operations to execute faster.
*   **Compiler Optimization:** The glisp compiler performs optimizations when generating bytecode, reducing redundant instructions.

In summary, although glisp and zygo share the same core, glisp's extensive optimizations give it a significant performance advantage. This indicates that in the implementation of scripting languages, even with the same kernel, higher-level optimizations are crucial.
