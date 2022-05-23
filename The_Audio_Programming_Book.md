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

int a[5];
int a[5] = {1,2,3,4,5};

Strings are arrays of characters, and the C programming language uses the convention that the end of a string of characters is marked by a null character (ASCII code 0).

"A" is a string constant and 'A' is a character constant

char name[40] = "hello";

Pointers

int *pa;

int *pa, a;
declares pa, a pointer to int, and an int, and the instruction
pa=&a;
stores the address of a in pa. We say that pa is pointing at a.

The operator ‘*’ is the indirection operator.

```c++
a = 10;
b = *pa; /* b is now also 10 */
*pa = 12; /* a is now 12 */
```

In summary:

A pointer is declared by a ‘*’ in front of its name.
The address of a variable is given by a ‘&’ in front of its name.
To obtain or store the value of a variable (pointed at), use ‘*’ in front of a pointer.


int a[10];

you are also declaring a pointer a to the first element in the array. Here, a is equivalent to &a[0].

The only difference between a and a pointer variable is that the array name is a constant pointer and cannot be used as a variable.

In this sense, a[i] and *(a+i) are equivalent, which makes possible what is called pointer arithmetic, adding integers to pointers to step through a memory block.

The compiler will know how many bytes to move forward when you add a certain number of memory slots to it.

If it is an int, it will jump four bytes (system-dependent of course) each step, if it is a double, then it will jump eight bytes.

This is one of the reasons why the compiler needs to know the type of a variable when you declare it.

```c++
int randarray(int *pa, int n)
{
	int i;
	for (i=0; i < n; i++)
	{
		*pa = rand()%n + 1;
		pa++;
	}
}

// OU

for(i=0; i<n; i++) *(pa+i)=rand()%n+1;

// OU ENCORE

for(i=0; i<n; i++) pa[i]=rand()%n+1;

```

char a[10], b[10];
b = a;
does not copy characters, just pointer b points to same set of char than pointer a
we need functions for manipulation : strcopy, strcat, strcmp, etc.
we can't do : a = "hello" but strcopy(a, "hello");

Pointers to Functions

```
void (*pf)();
void message(){ printf("my message\n"); }
pf = message;
// and call the function using a pointer:
(*pf)(); /* will call message() */
// or, even simpler,
pf();
```

return-type (*pointer-name) (arguments);

it's usefull for callbacks

void message_printer(int times, void (*callback)(char *msg), char *user_mess);

with

void my_important_message(char *mess);

0.13 Structures

P. 62

struct note
{
	char name[3];
	int duration;
	char intensity[5];
}

struct note first;

typedef struct _note
{
} note;

note first; (to avoid using struct everywhere)

note first = { "Ab", 80, "mf" };

passing big struct by value use a lot of memory, so you can use pointers instead

person *ptr

(*ptr).age

prt->age

---

typedef struct comp {
	double real, imag;
	void (*incr)(struct comp *p);
} complex;

void incr1(complex *p){ p->real++; p->imag++; }

complex a = { 0, 0, incr1 };
a.incr(&a);

0.14 Dynamic Memory Allocation

P. 66

pa = malloc(size)

sizeof(...)

pa = malloc(sizeof(int)*N);

free(pa) [to free memory]

Command-Line Arguments

int main(int argc, char **argv)

```c++
#include <stdio.h>
#include <stdlib.h>

int mod12(int note) {
    while (note >= 12) note -= 12;
    while (note < 0) note += 12;
    return note;
}

int main(int argc, char** argv) {
    int series[12][12], offset;
    int n, m, i;
    char* table[12] = {"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"};
    
    if (argc != 13) {
        printf("usage: %s note1 note2 ... note12\n", argv[0]);
        return -1;
    }
    
    /* loop until all available notes are entered */
    for (n = 0; n < 12; n++)
        series[0][n] = mod12(atoi(argv[n+1]));
        
    /* create inversion in column 1 */
    for (m = 1; m < 12; m++)
        series[m][0] = mod12(series[m-1][0] + series[0][m-1] - series[0][m]);
        
    /* create all transpositions */
    for (m = 0; m < 12; m++) 
        for (n = 1; n < 12; n++)
            series[m][n} = mod12(series[0][n] + series[m][0] - series[0][0]);
            
    for (m = 0; m < 12; m++) {
        /* print the pitch classes, row by row, using the translation table */
        for (n = 0; n < 12; n++) printf(" %s ", table[series[m][n]]);
        printf("\n");
    }
    
    return 0;
}
```

P. 69

Moving to C++

- variable declaration anywhere in the code
- default value for arguments in functions
- memory management: new and delete

```cpp
float *a = new float;
int *b = new int;
MyStruct *s = new MyStruct;

int size = 512;
float *array = new float[size];

delete a;
delete[] array;
```

- structures and Data Types (no need for typedef...)

```cpp
struct Osc {
    // dataspace
    float *table;
    float phase;
    // ...
    
    // methodspace
    Osc(float *tab, float ph=0.f, int len=def_len);
    -Osc() {
        delete[] output;
    }
};
```

-Osc is a destructor auto called

- Data Abstraction

struct_name::member_func()

Osc::Osc(float *tab, float ph, int len, int vsize, int sr) {
    table = tab;
    phase = ph;
    length = len;
    vecsize = vsize;
    ndx = 0.f;
    rate = sr;
    output = new float[vecsize];
}

float *Osc::Proc(float amp, float freq) {
    float incr = freq*length/rate;
    for(int i=0; i < vecsize; i++) {
        output[i] = amp*table[(int)ndx];
        ndx += incr;
        while(ndx >= length) ndx -= length;
        while(ndx < 0) ndx += length;
    }
    return output;
}

- Function Overloading

73