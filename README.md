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
