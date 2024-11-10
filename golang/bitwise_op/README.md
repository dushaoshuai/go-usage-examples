# Bitwise Operators 按位运算符

## Bitwise Logical Operators 按位逻辑运算符

这些操作符把他们的操作数当作位模式（0 和 1 组成的比特串），没有算术进位和正负的概念。

### Bitwise Negation, One's Complement, Bitwise Complement

0 -> 1, 1 -> 0

C: ~
Go: ^

### Bitwise And &

* 0 & 0 = 0
* 1 & 0 = 0
* 0 & 1 = 0
* 1 & 1 = 1

### Bitwise Or |

* 0 | 0 = 0
* 0 | 1 = 1
* 1 | 0 = 1
* 1 | 1 = 1

### Bitwise Xor ^

* 0 ^ 0 = 0
* 0 ^ 1 = 1
* 1 ^ 0 = 1
* 1 ^ 1 = 0

## Bitwise Shift Operators 按位移位运算符

移位运算符把一个整数的比特统一向左或向右移动指定位数。

### `>>`

### `<<`