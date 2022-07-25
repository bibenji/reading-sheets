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

P. 73

- Data Hiding and Encapsulation

```cpp
struct Osc {
    private:
    float *table;
    float phase;
    float ndx;
    // ...
    
    // methodspace needs to be accessible
    public:
    Osc(float *tab, float ph=0.f, int len=def_len, ...);
}
```

- Classes

a version of struct in which all members are private by default

```cpp
class Dodecaphonic {
    protected:
    int series[12];
    int mod12(int note) {
        while (note >= 12) note -= 12;
        while (note < 0) note += 12;
        return note;
    }
    
    public:
    Dodecaphonic() {
        for (int i=0; i < 12; i++) {
            series[i] = 0;
        }
    }
    Dodecaphonic(int *notes) {
        for (int i=0; i < 12; i++) series[i] = mod12(notes[i]);
    }

    int get(int index) {
        return series[mod12(index)];
    }
    void set(int note, int index) {
        series[mod12(index)] = mod12(note);
    }
    
    /* The three basic operations. */
    Dodecaphonic transpose(int interval);
    Dodecaphonic invert();
    Dodecaphonic retrograde();
};

/* Defining the operations. */
Dodecaphonic Dodecaphonic::transpose(int interval) {
    Dodecaphonic transp;
    for (int i=0; i < 12; i++)
        transp.set(mod12(series[i]+interval), i);
    return transp;
}

Dodecaphonic Dodecaphonic::invert() {
    Dodecaphonic inv;
    inv.set(series[0], 0);
    for (int i=1; i < 12; i++)
        inv.set(mod12(inv.get(i-1) + series[i-1] - series[i]), i);
    return inv;
}

Dodecaphonic Dodecaphonic::retrograde() {
    Dodecaphonic retr;
    for (int i=0; i < 12; i++)
        retr.set(series[i], 11-i);
    return retr; 
}
```

```cpp
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv)
{
    Dodecaphonic row, res;
    int interval, n;
    
    if (argc != 14 || argv[1][0] != '-') {
        printf("usage: %s [-oN | -rN | -iN | -irN ] ", "note1 note2 ... note12\n", argv[0]);
        return -1;
    }
    
    for (n = 0; n < 12; n++)
        row.set(atoi(argv[n+2]), n);
        
    switch(argv[1][1]) {
        case 'o': /* original transposed */
            interval = atoi(argv[1]+2);
            res = row.transpose(interval);
            break;
        case 'r': /* retrograde */
            interval = atoi(argv[1]+2);
            res = row.retrograde().transpose(interval);
            break;
        case 'i': /* inverted */
            if (argv[1][3] != 'r') {
                interval = atoi(argv[1]+2);
                res = row.invert().transpose(interval);
            } else { /* inverted retrograde */
                interval = atoi(argv[1]+3);
                res = row.invert().retrograve().transpose(interval);
            }
            break;
        default:
            printf("unrecognized option \n");
            return -1;
    }
    
    for (n = 0; n < 12; n++)
        printf("%d ", res.get(n));
        
    printf("\n");
    return 0;
}
```

- Inheritance

P. 79

class Osci : public Osc {
    // ...
};

virtual functions (can be overridable)

class Osc {
    (...)
    virtual float *Proc(float amp, float freq);
};

class Osci : public Osc {
    (...)
    virtual float *Proc(float amp, float freq);
};

0.18.5 Overloaded Operators
The operators =, +, -, *, /, <<, and >> can be overloaded.

class Osc {
    (...)
    float *operator*(float val) {
        // adds val to every sample in the output block
        for (int i=0; i < vecsize; i++) output[i] += val;
        // returns the audiio block
        return output;
    }
};

Now using ‘+’ with an Osc object and a float has a definite meaning: add that number to every sample in the output audio block of the Osc object (and return a block of samples).

Osc oscil(...);
float a = 1000.f;
float *buffer;
(...)
for(...) {
    (...)
    buffer = oscil + a;
    (...)
}

1.1.1 Your First C Program

P. 85

float takes 4 bytes
double takes 8 bytes

16 bits integer for numbers from 0 to 65,535
16 bits signed integer for numbers from -32,768 to +32,767 (how audio samples are stored)

32 bits signed integer => -2,147,483,648 to +2,147,483,647

char => 1 byte
short => 2 bytes
long => 4 bytes
int => * (dependent on the CPU integer size)

name "long long" for 64 bits integer

1.2.7 The sizeof() Operator

P. 94

int intsize = sizeof(int); // 2 or 4, depending on the machine

midi2freq
```c++
#include <stdio.h>
#include <math.h>

int main()
{
    double semitone_ratio;
    double c0; /* for frequency of MIDI Note 0 */
    double c4; /* for frequency of Middle C */
    double frequency; /* which we wanto to find */
    int midinote; /* given this note */
    
    semitone_ratio = pow(2, 1/12.0); /* approx. 1.0594631 */
    c5 = 220.0 * pow(semitone_ratio, 3); /* middle C, three semitones above low A = 220 */
    c0 =c5 * pow(0.5, 5); /* MIDI Note 0 is C, 5 octaves below Middle C */
    
    midinote = 73; /* C# above A = 440 */
    frequency = c° * pow(semiton_ratio, midinote)
    
    printf("MIDI Note %d has frequency %f\n", midinote, frequency);
    
    return 0;
}
```

a good rule of thumb is to try to avoid names longer than 32 characters or so

1.2.9 Initializing Variables and Reducing the Line Count

P. 99

a = b = c;
is equivalent to
b =c;
a = b;

freq2midi

```c++
#include <stdio.h>
#include <math.h>

int main()
{
    /* find the nearest MIDI note to a given frequency in Hz */
    /* uses the log rule: log_a(N) = log_b(N) / log_b(a) to find the log of a value to base 'semitone_rati' */
    
    frequency = 400.0
    fracmidi = log(frequency / c0) / log(semitone_ratio);

    /* round fracmidi to the nearest whole number */
    midinote = (int) (fracmidi + 0.5);
    
    printf("The nearest MIDI note to the frequency %f is %d\n", frequency, midinote);
}
```

1.3 Introduction to Pointers, Arrays, and the Problem of User Input

P. 104

pointer to char
`char* message;`
`char* message = "one two three four";`

```
char* two = "two";
char* four = "four";
char* six = "six";
char* message = "%s %s %s eight";
printf(message, two, four, six);
```

P. 106

char message[256];
// illegal
char message[];
char message[0];

int FullChord[4] = {60,64,67,72}; /* MIDI notes for C Major */
int root = 0, third = 1, rootnote;
FullChord[third] = 63; /* change chord to C minor */
rootnote = FullChord[root]; /* Middle C */

the special value NULL (defined in stdlib.h)

A subtle problem arises using indexing on strings used to initialize pointers:
char* name = "John";
name[2] = 'a'; /* DANGER: may crash! */

1.3.5 Converting Number Strings to Variables: The const Keyword

3 conversion functions: atof(), atoi(), atol()
convert into double, int and long
in stdlib.h

```
double atof(const char*);
int atoi(const char*);
long atol(const char*);
```

char* gets( char *buffer );

strtod() aussi ?

const char *name = "John";
name[2] = 'a'; /* trigger a warning */

P. 111

1.3.6 The if Keyword

The key aspect to remember here is that a pointer to char can itself signify an array, in the
sense of a character string

1.4.3 The ** Notation for argv

int main(int argc, char** argv)
where ** means that argv is a ‘‘pointer to a pointer.’’

1.5 Controlling Repetition: Looping and Counting

P. 121

P. 129

for(;;) {
    sampsread = readsamps("tubularbells.wav");
    if(sampsread == 0)
        break; /* end of file */
    process_samps(sampsread);
    write_samps("plasticbells.wav," sampsread);
}

or

while(1) {
    // same code
}

1.5.7 Writing a Program to Create Unusual Musical Scales

P. 130

1.5.8 Exercises

P. 135

Check: Csound [...]

https://csound.com/

1.6 Using Pointer Arithmetic and Adding Options to a Program

double* ptr

double buffer[1024];
ptr = buffer;
In C, the name of an array is in effect a pointer to the first element of it (buffer[0]).

indirection

buffer[0] = 1.0594631;

we can now assign this value indirectly, via the pointer:
*ptr = 1.0594631;

double* ptr; /* pointer; currently uninitialized */
double val; /* a simple number variable */
ptr = buffer; /* now it has something to point to: the first element of buffer, i.e. buffer[0] */
*ptr = 1.0594631; /* buffer[0] now contains 1.0594631 */

/* read the contents of ptr: */
val = *ptr;
/* val now = 1.0594631 */

This works because the name of the array is equivalent to the ‘‘address of’’ the first element
of it.

double* ptr;
double val;
ptr = &val; /* find the address of val, and assign it to ptr */

we can initialize val indirectly, via the pointer:
*ptr = 1.0594631;
printf("%f,"val);

So we can replace the assignment
ptr = buffer;
with the exactly equivalent (if rather more awkward-looking) statement
ptr = &buffer[0];

1.6.3 Moving Pointers Around

P. 139

for (i = 0; i < 1024; i++) {
    *ptr = 0.0;
    ptr++; /* move pointer to next element of the array */
}

it results in faster code than the array index notation

for(i = 0; i < 1024; i++)
    *ptr++ = 0.0;

=> such combinations are often available as single machine instructions

More generally, C supports the use of the addition, subtraction, and comparison opera-
tions on pointers.

Pointer arith-
metic is both one of the most widely exploited features of C, and also one of the most
criticized.

```
double* ptr = buffer + 1024; /* point to last element */
double maxval = 0.0;
unsigned long pos = 0;

while(--ptr != buffer) {
    if (*ptr >= maxval) {
        maxval = *ptr;
        pos = ptr - buffer;
    }
}

printf("the maximum sample is %f, at position %d\n", maxval, pos);
```

[...]

It is often the case that auto-decrement is faster than auto-increment, and similarly
that a comparison against zero is faster than comparison with some non-zero value.

Reprendre P. 141

increments tend to be post, while decrements tend to be pre

It is often the case that auto-decrement is faster than auto-increment, and similarly
that a comparison against zero is faster than comparison with some non-zero value.

1.6.4 Extending the Command Line with Optional Arguments

1.6.5 Dealing with Alternatives: if . . . else

P. 146

```
float buffer[1024];
float* bufptr = buffer;
int ascending = 0;

if (ascending) {
    int i;
    for (i = 0; i < 1024; i++)
        *bufptr++ = (float) i;
}
else {
    int i;
    for (i = 1024; i; i--)
        *bufptr-- = (float) i;
}
```

1.6.7 Creating a Text File—the Special FILE Type

These all use a pointer to a strange object called FILE.

// Open a file for reading or writing.
FILE *fopen(const char* name, const char* mode);

// Write formatted text to FILE.
int fprintf(FILE *fp, const char* format, ...);

// Close an open FILE.
int fclose(FILE *fp);

an instance of
a file, that it must always be typed in upper case, and that it is always referred to indirectly
via a pointer

fprintf() like printf() but with an initial FILE pointer argument
same format

fopen() for read and for write

perror(), can be used to display the cause of the error

Listing 1.6.10

P. 153

```
#include <stdio.h>
#include <stdlib.h>
#include <math.h>

/* usage iscale [-m][-i] N startval [outfile.txt]
    -m : sets format of startval as MIDI note
    -i : prints the calculated interval as well as the abs freq
    outfile: optional text filename for output data
*/

int main(int argc, char* argv[])
{
    int notes, i;
    int ismidi = 0;
    int write_interval = 0;
    int err = 0;
    double startval, basefreq, ratio;
    FILE* fp;
    doublie intervals[25];
    
    while (argc > 1) {
        if (argv[1][0] == '-') {
            if (argv[1][1] == 'm')
                ismidi = 1;
            else if (argv[1][1] == 'i')
                write_interval = 1;
            else {
                printf("error: unrecognized option %s\n", argv[1]);
                return 1;
            }
            
            argc--;
            argv++;
        }
        else
            break;
    }
    
    if (argc < 3) {
        printf("insufficient arguments\n");
        printf("Usage: itable [-m][-i] N startval [outfile.txt]\n");
        return 1;
    }
    
    notes = atoi(argv[1]);
    if (notes < 1 || notes > 24) {
        printf("error: N out of range. Must be between 1 and 24.\n");
        return 1;
    }
    
    startval = atof(argv[2]);
    if (ismidi) {
        if (startval > 127.0) {
            printf("error: MIDI startval must be <= 127.\n");
            return 1;
        }

        /* for MIDI, startval = 0 is legal */
        if (startval < 0.0) {
            printf("error: MIDI startval must be >= 0.\n");
            return 1;
        }
    }
    else {
        if (startval <= 0.0) {
            printf("error: frequency startval must be positive.\n");
            return 1;
        }
    }
    
    fp = NULL;
    if (argc == 4) {
        fp = fopen(argv[3], "w");
        if (fp == NULL) {
            printf("WARNING: unable to create file %s\n",argv[3]);
            perror("");
        }
    }
    
    if (ismidi) {
        double c0, c5;
        ratio = pow(2.0, 1.0 / 12.0);
        c5 = 220.0 * pow(ratio, 3);
        c0 = c5 * pow(0.5, 5);
        basefreq = c0 * pow(ratio, startval);
    }
    else
        basefreq = startval;
        
    ratio = pow(2.0, 1.0 / notes);
    for (i = 0; i <= notes; i++) {
        intervals[i] = basefreq;
        basefrequ *= ratio;
    }
    
    for (i = 0; i <= notes; i++) {
        if (write_interval)
            printf("%d:\t%f\t%f\n", i, pow(ratio,i), intervals[i]);
        else
            printf("%d:\t%f\n", i, intervals[i]);
            
        if (fp) {
            if (write_interval)
                err = fprintf(fp,"%d:\t%f\t%f\n", i, pow(ratio,i), intervals[i]);
            else
                err = fprintf(fp,"%d:\t%f\n", i, intervals[i]);
                
            if (err < 0)
                break;
        }
    }
    
    if (err < 0)
        perror("There was an error writing the file.\n");
        
    if (fp)
        fclose(fp);
        
    return 0;
}
```

1.6.9 Testing the Program

Exercises [...]

1.7 Creating Types, Functions, and Memory

P. 157

1.7.1 The Time/Value Breakpoint File

breakpoint file

1.7.2 Defining Custom Objects in C using the struct Keyword

```c
struct breakpoint {
	double time;
	double value;
};

struct tickpoint {
	unsigned long ticks;
	double value;
};
```

struct breakpoint point;

struct breakpoint {
	double time;
	double value;
} point;

point.time = 0.0;
point.value = 0.5;

struct breakpoint point = {0.0,0.5};

An important operation allowed for structs is assignment or copy:
struct breakpoint point1,point2 = {0.0,0.5};
point1 = point2; /* point1 now contains {0.0,0.5} */

struct breakpoint point3;
point3 = point1 + point2; /* error: + not allowed with structs */

1.7.3 Defining Custom Types using the typedef Keyword

typedef <existing type> <typename> ;

typedef unsigned short WORD;
typedef unsigned long DWORD;

long size = sizeof(DWORD);

/* define an object describing a soundfile */
typedef struct soundfile_info {
	DWORD nSamples;
	DWORD samplerate;
	WORD nChannels;
	char* name;
} SFINFO;

/* create an instance of the SFINFO object, and initialize it */
SFINFO info;
info.nSamples = 1234567;
info.samplerate = 96000;
info.name = "main title";
info.nChannels = 6;

typedef struct {
	DWORD nSamples;
	DWORD samplerate;
	WORD nChannels;
	char* name;
} SFINFO;

typedef struct breakpoint {
	double time;
	double value;
} BREAKPOINT;

array of BREAKPOINTS:
BREAKPOINT points[64];

point[0].time = 0.0;
point[0].value = 1.0;

1.7.4 Text Substitution: Introducing the #define Preprocessor Directive

`#define BUFFERSIZE 1024`

Note that there is no semicolon terminating the line—this is not C code as such, but instruc-
tions to the preprocessor.

short sampbuf[BUFFERSIZE];

`#define SAMPLE float`

`#define SAMPLE double`

prototype declaration
at the top
or in a header file included with #

```
/* The prototype*/
BREAKPOINT maxpoint(const BREAKPOINT* points, long npoints);

/* input: points = array of BREAKPOINTS,
   npoints gives the length of the array */
/* output: copy of the BREAKPOINT containing largest value */
/* the function definition */

BREAKPOINT maxpoint(const BREAKPOINT* points, long npoints)
{
	int i;
	BREAKPOINT point;
	point.time = points[0].time; /* initialize from first point */
	point.value = points[0].value;
	for (i=0; i < npoints; i++) {
		if (point.value < points[i].value) {
			point.value = points[i].value;
			point.time = points[i].time;
		}
	}
	return point;
}
```

1.7.6 The void and size_t Types: Creating Memory and Dealing with Variable Quantities

P. 139

dynamic memory allocation

void* malloc(size_t size);
void* calloc(size_t nobj, size_t size);
void* realloc(void *p, size_t size);
void free(void *p);

size_t: This is a symbol #defined by the compiler (in <stdio.h>) according to the target machine architecture

void*: Of the two types, this is by far the more important. As a C programmer you will
be working with the void keyword a great deal of the time. A memory request to the system
is simply for a block of bytes, and all the system can do is return a generic or type-agnostic
pointer to this block

[...] P. 166

```
/* request enough memory to hold a lot of samples */
#define SAMPLELEN 1048576

float* sampbuf;

sampbuf = (float*) malloc(SAMPLELEN * sizeof(float));
if (sampbuf == NULL) {
	puts("Sorry - not that much memory available!\n");
	exit(1);
}

/* do something useful with all this memory... */
process_sample(sampbuf);

/* and remember to return it at the end */
free(sampbuf);
```

how NULL is defined: #define NULL ((void*) 0)
integer 0 cast to a 'pointer to void'
so, it can be compared to any pointer type
that's why we can do: if (sampbuf == NULL) (and not if (sampbuf == (float*) NULL))

printf() needs memory allocation, so puts() is better to used to report memory allocation errors

1.7.7 Expanding Allocated Memory on Demand: Reading Breakpoints from a File

realloc

char *fgets(char *line, int maxline, FILE *fp);
int *sscanf(char *string, const char* format, ...);

The continue Keyword

1.7.10 Exercises

P. 173

After csound, discover Gnuplot!

Commands for gnuplot:
- pwd
- cd "/home/[folder path]"
- gnuplot "envelope.txt" with lines

One general formula for an exponential decay is:
x = 1⁄4 ae ^ (- k/T )

```
/* expdecay.c */
/* implement formula x[t] = a * exp(-k/T) */

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main(int argc, char** argv)
{
    int i, nsteps;
    double step, x, a, T, k;
    double dur;
    
    if (argc != ARG_NARGS) {
        printf("usage: expdecay dur T steps\n");
        return 1;
    }
    
    dur = atof(argv[1]);
    T = atof(argv[2]);
    nsteps = atoi(argv[3]);
    
    k = dur / nsteps; /* the constant time increment */
    a = exp(-k / T); /* calc the constant ratio value */
    x = 1.0; /* starting value for the decay */
    
    step = 0.0;
    for (i = 0; i < nsteps; i++) {
        printf("%.4lf\t%.8lf\n", step, x);
        x = a * x;
        step += k;
    }
    
    return 0;
}
```

P. 178

expdecay 1 0.5 200 > expdecay.txt

1.8.4 The Exponential Attack and the stdout and stderr Output Streams

```
/* expbrk.c generate exponential attack or decay breakpoint data */
#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main(int argc, char** argv)
{
    int i, npoints;
    double startval, endval;
    double dur, step, start, end, thisstep;
    double fac, valrange, offset;

    const double verysmall = 1.0e-4; /* ~-80dB */
    
    if (argc != 5) {
        fprintf(stderr, "Usage: expbrk duration npoints startval endval\n");
        return 1;
    }
    
    dur = atof(argv[1]);
    
    if (dur <= 0.0) {
        fprintf(stderr, "Error: duration must be positive.\n");
        return 1;
    }
    
    npoints = atoi(argv[2]);
    
    if (npoints <= 0) {
        fprintf(stderr,"Error: npoints must be positive!\n");
        return 1;
    }
    
    step = dur / npoints;

    startval = atof(argv[3]);
    endval = atof(argv[4]);

    valrange = endval - startval;
    
    if (valrange == 0.0) {
        fprintf(stderr, "warning: start and end values are the same!\n");
    }
    
    /* initialize normalized exponential as attack or decay */
    if (startval > endval) {
        start = 1.0;
        end = verysmall;
        valrange = -valrange;
        offset = endval;
    }
    else {
        start = verysmall;
        end = 1.0;
        offset = startval;
    }
    
    thisstep = 0.0;
    
    /* make normalized curve, scale output to input values, range */
    fac = pow(end / start, 1.0 / npoints);

    for (i = 0; i < npoints; i++) {
        fprintf(stdout, "%.4lf\t%.8lf\n" , thisstep,  offset + (start * valrange));
        start *= fac;
        thisstep += step;
    }
    
    /* print final value */
    fprintf(stdout,"%.4lf\t%.8lf\n", thisstep, offset + (start * valrange));
    fprintf(stderr,"done\n");
    return 0;
}
```

1.8.5 Not All Curves Are Exponentials: The log10 Gnuplot Test

P. 181




































