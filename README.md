| Benchmark | glisp (ns/op) | goja (ns/op) | lua (ns/op) | zygo (ns/op) | Winner |
|:---|:---:|:---:|:---:|:---:|:---:|
| Factorial Calculation | 6959 | 3427 | **1943** | 43054 | lua |
| Regular Expression Matching | **603.7** | 887.3 | 821.1 | 2180 | glisp |
| Complex Conditions | 726.8 | **492.0** | 517.8 | 4608 | goja |
| Time Formatting | **1179** | 1646 | 1355 | N/A | glisp |
| Hash Write | 689.5 | 593.7 | **548.4** | 2535 | lua |
| Hash Delete | 497.8 | **414.7** | 449.0 | 1917 | goja |
| Hash Access | 523.7 | **443.5** | 565.6 | 1984 | goja |
| JSON Parsing and Modification | 4667 | 6968 | **3090** | 12898 | lua |

**Conclusion:**

*   **Lua:** Performs best in computationally intensive tasks (factorial and JSON operations) and hash writes.
*   **glisp:** Excels in string operations (regular expressions) and time formatting.
*   **Goja:** Fastest in handling complex logical conditions, hash deletes, and hash accesses.
*   **Zygo:** Performed the worst in all tests.

**glisp vs. zygo Performance Discussion:**

glisp and zygo are developed based on the same kernel (https://github.com/zhemao/glisp). According to the performance test results, glisp significantly outperforms zygo in all test scenarios. This is mainly due to targeted optimizations made in glisp in the following areas:

*   **Built-in Function Optimization:** glisp has rewritten and optimized commonly used built-in functions, reducing unnecessary type conversions and memory allocations.
*   **Virtual Machine Instruction Set Optimization:** glisp has optimized the virtual machine's instruction set, allowing certain operations to execute faster.
*   **Compiler Optimization:** The glisp compiler performs optimizations when generating bytecode, reducing redundant instructions.

In summary, although glisp and zygo share the same core, glisp's extensive optimizations give it a significant performance advantage. This indicates that in the implementation of scripting languages, even with the same kernel, higher-level optimizations are crucial.
