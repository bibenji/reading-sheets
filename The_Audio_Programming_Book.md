Preface

Introduction

opcodes

Chapter 0

P. 30

```c++
unsigned int ua;
unsigned long ulb;
short sample;

float result;
double nextResult;
```

2L, 10L (long integer constants)
2.f is a float; 2.0 is a double
octal constant is preceded by '0' and a hexadecimal by '0x'
31 = 037 = 0x1F
decimal = octal = hexadecimal

`#define END 10000`

0.6 Standard Input and Output

P. 37

scanf("%d %d",&i,&j);
scanf("%lOd",&i)

0.7 Control of Flow

P. 41

0.9.1 Bitwise Logic

&	bitwise AND
|	bitwise inclusive OR
^	bitwise exclusive OR
~	ones complement (unary operator).

0.9.2 Bitshift Operators

```c++
<< left shift
>> right shift
```

Left shifts are equivalent to multiplication (a fast way of doing it):
x << n	multiplication by 2 n .
Right shifts are equivalent to division (with rounding):
x >> n	multiplication by 2 n .

0.11 Arrays

P. 54


