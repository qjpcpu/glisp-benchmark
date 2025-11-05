| Benchmark | glisp (op/ms) | goja (op/ms) | lua (op/ms) | zygo (op/ms) | Winner |
|:--- |:---:|:---:|:---:|:---:|:---:|
| Factorial Calculation | 6624 | 3453 | **1959** | 43131 | lua |
| Regular Expression Matching | **631** | 892 | 873 | 2265 | glisp |
| Complex Conditions | 621 | **508** | 515 | 4977 | goja |
| Time Formatting | **1104** | 1736 | 1356 | N/A | glisp |
| Hash Write | 636 | 527 | **515** | 2576 | lua |
| Hash Delete | 475 | **426** | 433 | 1901 | goja |
| Hash Access | 530 | **416** | 474 | 2023 | goja |
| JSON Parsing and Modification | 4450 | 6897 | **3060** | 12830 | lua |
| String Concat | **239** | 773 | 427 | 12342 | glisp |

**Conclusion:**

*   **Lua:** Performs best in computationally intensive tasks (factorial and JSON operations) and hash writes.
*   **glisp:** Excels in string operations (regular expressions), time formatting and string concat.
*   **Goja:** Fastest in handling complex logical conditions, hash accesses and hash deletes.
*   **Zygo:** Performed the worst in all tests.

**glisp vs. zygo Performance Discussion:**

glisp and zygo are developed based on the same kernel (https://github.com/zhemao/glisp). According to the performance test results, glisp significantly outperforms zygo in all test scenarios. This is mainly due to targeted optimizations made in glisp in the following areas:

*   **Built-in Function Optimization:** glisp has rewritten and optimized commonly used built-in functions, reducing unnecessary type conversions and memory allocations.
*   **Virtual Machine Instruction Set Optimization:** glisp has optimized the virtual machine's instruction set, allowing certain operations to execute faster.
*   **Compiler Optimization:** The glisp compiler performs optimizations when generating bytecode, reducing redundant instructions.

In summary, although glisp and zygo share the same core, glisp's extensive optimizations give it a significant performance advantage. This indicates that in the implementation of scripting languages, even with the same kernel, higher-level optimizations are crucial.
