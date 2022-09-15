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

A standard scale for amplitude values is
the Decibel, defined by the formula:
P(dB) = 20.0 log 10 (x).

gnuplot:

`plot "longdecay.txt" using (20.0 * log10($2)) with lines`

`plot "longdecay.txt" using (20.0*log10($2+0.00001)) with lines`

1.8.6 Exercises

P. 182

musicdsp

https://www.musicdsp.org/

[...]

1.9 Toward the Soundfile: From Text to Binary

P. 184

study the appendixes in this book on computer architecture, number representations, and mathematics

1.9.2 The sin Function: Creating a Sine Wave

tuning fork

./sinetext >sine.txt

```
/* sinetext.c */
/* write sinewave as text */

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

/* conditional compilation - is M_PI defined? */
#ifndef M_PI
#define M_PI (3.141592654)
#endif

int main(int argc, char** argv)
{
    int i, nsamps;
    double samp;
    double twopi = 2.0 * M_PI;
    double angleincr;
    
    /* set number of points to create */
    nsamps = 50;

    /* make one complete cycle */
    angleincr = twopi / nsamps;
    
    for (i = 0; i < nsamps; i++) {
        samp = sin(angleincr * i);
        fprintf(stdout, "%.8lf\n", samp);
    }
    
    fprintf(stderr, "done\n");
    return 0;
}
```

When Gnuplot is
given a single column of numbers, it automatically treats then as equally spaced in this
way.

`plot "sine.txt" with impulses`

`plot "sine.txt" with steps`

1.9.3 Toward the Tuning Fork: Frequency Generation and the Keyword enum

elaborate the arithmetic to incorporate a sample rate

```
/* sinetext2.c */

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#ifndef M_PI
#define M_PI (3.141592654)
#endif

/* define our program argument list */
enum {ARG_NAME,ARG_NSAMPS,ARG_FREQ,ARG_SR,ARG_NARGS};

int main(int argc, char** argv)
{
    int i, nsamps;
    double samp, freq, srate;
    double twopi = 2.0 * M_PI;
    double angleincr;
    
    if (argc != ARG_NARGS) {
        fprintf(stderr, "Usage: sinetext2 nsamps freq srate\n");
        return 1;
    }
    
    nsamps = atoi(argv[ARG_NSAMPS]);
    freq = atof(argv[ARG_FREQ]);
    srate = atof(argv[ARG_SR]);

    angleincr = twopi * freq / srate;
    
    for (i = 0; i < nsamps; i++) {
        samp = sin(angleincr * i);
        fprintf(stdout, "%.8lf\n", samp);
    }
    
    fprintf(stderr, "done.\n");
    return 0;
}
```

1.9.4 The Utility Program text2sf: Converting Text to a Soundfile

P. 189

`text2sf infile outfile srate chans gain`

where infile = input text data file
outfile = output soundfile (.wav or .aiff formats),
srate = sample rate of the data
chans = number of interleaved audio channels in infile
and gain = amplitude factor applied to input data (1.0 = no change)

`plot "stereo.txt" using($1) with lines, "stereo.txt" using($2) with lines`

plot "stereo.txt" using($1) with lines, "stereo.txt" using($2) with
lines

1.9.5 Our First Tuning Fork Emulation

P. 191

[...] Unable to use text2sf...

1.9.6 Tuning Fork Program Version 2

P. 193 

Listing 1.9.4: Tuning fork v2 with maxsamp report

```
/* tfork2.c alternate tuning fork generator based on expbrk.c
* - decay is always to ~silence regardless of duration. */

#include <stdio.h>
include <stdlib.h>
#include <math.h>

#ifndef M_PI
#define M_PI (3.141592654)
#endif

enum {ARG_NAME,ARG_OUTFILE,ARG_DUR, ARG_HZ,ARG_SR,ARG_AMP,ARG_NARGS};

int main(int argc, char** argv)
{
	int i,sr,nsamps;
	double samp,dur,freq,srate,amp,maxsamp;
	double start,end,fac,angleincr;
	double twopi = 2.0 * M_PI;
	FILE* fp = NULL;
	
	if (argc != ARG_NARGS) {
		printf("Usage: tfork2 outfile.txt dur freq srate amp\n");
		return 1;
	}
	
	fp = fopen(argv[ARG_OUTFILE],"w");
    
    if (fp==NULL) {
    	printf("Error creating output file %s\n", argv[ARG_OUTFILE]);
    	return 1;
    }
    
    dur = atof(argv[ARG_DUR]);
    freq = atof(argv[ARG_HZ]);
    srate = atof(argv[ARG_SR]);
    amp = atof(argv[ARG_AMP]);
    nsamps = (int)(dur * srate);
    angleincr = twopi * freq / nsamps;
    start = 1.0;
    end = 1.0e-4; /* = -80dB */
    maxsamp = 0.0;
    
    fac = pow(end / start, 1.0 / nsamps);
    
    for (i=0;i < nsamps; i++) {
    	samp = amp * sin(angleincr*i);
    	samp *= start;
    	start *= fac;
    	fprintf(fp,"%.8lf\n",samp);
    	
    	if (fabs(samp) > maxsamp) {
    		maxsamp = fabs(samp);
    	}
    }
    
    fclose(fp);
    
    printf("done. Maximum sample value = %.8lf\n", maxsamp);
    
    return 0;
}
```

1.9.7 The Raw Binary Soundfile

the format specified for fopen (line 25) to indicate ‘binary’ mode: fp = fopen(argv[ARG_OUTFILE],"wb");

and replace fprintf (writes formatted text) by fwrite (writes arbitrary blocks of memory to disk):

size_t fwrite(const void * ptr,size_t size,size_t count, FILE* fp);

- size_t is defined by the compiler as an integer type appropriate to the platform. It can be
assumed to be at least a 32-bit unsigned type, but it may be larger (e.g. 64 bits on a 64-bit
platform).
- ptr is a pointer to the memory block to be written; being a pointer-to-void the address
can be that of any object, whether a single local variable or a large block of memory.
- Together, size and count define the size of the memory block to be written. Typically,
size will refer to the size of any standard or user-defined type such as char or int (i.e.
where its size can be found using the sizeof() operator), while count defines the number
of such elements to be written. This is also the value returned by the function.
- fp is the pointer-to-FILE to be written, as created by fopen.

```
float fsamp; /* declared at the top of the code block */

fsamp = (float) samp;

if (fwrite(&fsamp,sizeof(float),1,fp) != 1) {
	printf("error writing data to disk\n");
	return 1;
}
```

1.9.8 Platform Particulars: The Endianness Issue

Big-endian and little-endian storage.

1.9.9 A Raw Binary Version of tfork2.c

Listing 1.9.5: tforkraw.c

tforkraw > see src

```
/* tforkraw.c gen raw sfile with native endianness */
/* based on tfork2.c */

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#ifndef M_PI
#define M_PI (3.141592654)
#endif

enum {ARG_NAME,ARG_OUTFILE,ARG_DUR,ARG_HZ,ARG_SR,ARG_AMP,ARG_TYPE,ARG_NARGS};

enum samptype {RAWSAMP_SHORT,RAWSAMP_FLOAT};

/* thanks to the SNDAN programmers for this */
/* return 0 for big-endian machine, 1 for little-endian machine*/
/* so we can tell user what order the data is */
int byte_order()
{
	int one = 1;
	char* endptr = (char *)
	return (*endptr);
}

const char* endianness[2] = {"big_endian","little_endian"};

int main(int argc, char** argv)
{
	unsigned int i,nsamps;
	unsigned int maxframe = 0;
	unsigned int samptype, endian, bitreverse;
	double samp,dur,freq,srate,amp,step;
	double start,end,fac,maxsamp;
	double twopi = 2.0 * M_PI;
	double angleincr;
	FILE* fp = NULL;
	float fsamp;
	short ssamp;

	if (argc != ARG_NARGS) {
		printf("Usage: tforkraw outsfile.raw dur freq srate amp isfloat\n");
		return 1;
	}

	dur = atof(argv[ARG_DUR]);
	freq = atof(argv[ARG_HZ]);
	srate = atof(argv[ARG_SR]);
	amp = atof(argv[ARG_AMP]);
	samptype = (unsigned int) atoi(argv[ARG_TYPE]);
	
	if (samptype > 1) {
		printf("error: sampletype can be only 0 or 1\n");
		return 1;
	}

	/* create binary file: not all systems require the 'b' */
	fp = fopen(argv[ARG_OUTFILE], "wb");

	if (fp==NULL) {
		fprintf(stderr,"Error creating output file %s\n", argv[ARG_OUTFILE]);
		return 1;
	}
	
	// nb samples = duration en secondes * nb de samples par seconde	
	nsamps = (int)(dur * srate);
	
	// ??? TODO
	angleincr = twopi * freq / nsamps;
	
	// ??? à quoi ça sert ??? période ???
	step = dur / nsamps;

	/* normalized range always - just scale by amp */
	start =1.0;
	end = 1.0e-4;
	maxsamp = 0.0;
	
	// ratio end sur start puissance (1 divisé par nbre de samples)
	// pour faire rentrer la diminution de 0 à 1 
	fac = pow(end / start, 1.0 / nsamps);
	
	endian = byte_order();
	printf("Writing %d %s samples\n", nsamps, endianness[endian]);

	/* run the loop for this samptype */
	if (samptype == RAWSAMP_SHORT) {
		for(i=0;i < nsamps; i++) {
			// on calcule la valeur du son avec l'amplitude
			samp = amp * sin(angleincr*i);
			
			// on applique le decay
			samp *= start;
			
			// on passe au start suivant en multipliant par le facteur
			start *= fac;

			/* use 32767 to avoid overflow problem */
			ssamp = (short) (samp * 32767.0);
			
			if(fwrite(&ssamp,sizeof(short),1,fp) != 1) {
				printf("Error writing data to file\n");
				return 1;
			}
			
			if(fabs(samp) > maxsamp) {
				maxsamp = fabs(samp);
				maxframe = i;
			}
		}
	}
	else {
		for(i=0;i < nsamps; i++) {
			samp = amp * sin(angleincr*i);
			samp *= start;
			start *= fac;
			fsamp = (float) samp;
			
			if(fwrite(&fsamp,sizeof(float),1,fp) != 1) {
				printf("Error writing data to file\n");
				return 1;
			}
			
			if(fabs(samp) > maxsamp) {
				maxsamp = fabs(samp);
				maxframe = i;
			}
		}
	}
	
	fclose(fp);
	printf("done. Maximum sample value = %.8lf at frame %d\n",maxsamp,maxframe);
	return 0;
}
```

srate = samples per second or Hertz

P. 197

1.9.10 Auditioning Raw Soundfiles: the Audacity Soundfile Editor

Opening raw soundfiles in Audacity.

‘‘Start offset’’ entry box (the byte position
at which conversion should start) is pre-filled with the value 1. Expect to set this value to
zero for the plain raw soundfiles considered here.

float x = 1.0;
x += 2.1; (results in a double operation)
x += 2.1f; (force a float operation)

stdout /* primary output stream */
stderr /* stream to receive error messages */
stdin /* stream from which input is received e.g. teletype keyboard */

The difference between printf and
fprintf is that printf outputs only to stdout, whereas the more general fprintf can write to any
named stream. It is often used, for example, to write to stderr:
fprintf(stderr,"there was an error\n");

#define SQUARE(n) n * n
This is a small macro definition, replacing any instance of the text represented by the symbol n (which is
not a variable but a placeholder for whatever text the programmer supplies)

exit()
immediately quits the program, releasing all resources (memory, open files) in the process

```
void* stress_malloc(size_t size)
{
    size_t limit = 1024 * 1024 * 2; /* limit to 2MB */
    if(size < limit)
        return malloc(size);
    else
        return NULL;
}
```

You can then use #define to make all calls to malloc become calls to your wrapper function:
#define stress_malloc malloc

2 Audio Programming in C

P. 212

makefile

2.1 A Simple Soundfile Library: portsf

libsndfile > the pro version (to check)

ieee80.c
ieee80.h
portsf.c
portsf.h

2.1.3 Soundfile Formats—enum and typedef

`typedef enum {
STDWAVE, MC_STD, MC_MONO, MC_STEREO, MC_QUAD,MC_LCRS, MC_BFMT, MC_DOLBY_5_1, MC_WAVE_EX
} psf_channelformat;`

STD_WAVE serves as a ‘‘default’’ value, representing any ‘‘stan-
dard’’ soundfile with no special speaker positions

`typedef enum {
PSF_FMT_UNKNOWN = 0,
PSF_STDWAVE,
PSF_WAVE_EX,
PSF_AIFF,
PSF_AIFC
} psf_format;`

[...]

typedef struct psf_props
{
    long                srate;
    long                chans;
    psf_stype           samptype;
    psf_format          format;
    psf_channelformat   chformat;
} PSF_PROPS;

A soundfile contains audio data at a particular sample rate (number of samples per second).
Most audio systems support a range of standard sample rates—22,050, 44,100, 48,000,
96,000 and even 192,000 hertz for DVD audio. We can store this parameter in
PSF_PROPS structure element: srate

A soundfile can contain several channels of audio—one for mono, two for stereo, six for
Dolby 5.1, and so on. In a multi-channel file, the sample rate defines the frame rate, where
one frame contains one sample for each channel.
PSF_PROPS element: chans

soundfile can contain audio samples in a range of data formats. By far the most common
format is 16-bit, represented by the C short data type. [...]
PSF_PROPS element: samptype

A soundfile can be written in a number of file formats.

soundfile can contain audio channels associated with specific speaker positions. This
property is relatively new, but has greatly increased in importance thanks to the popularity
of surround sound. With a stereo signal, it is established largely by convention that channel
1 is front left, and channel 2 is front right.

PSF_PROPS element: chformat

2.1.4 Initializing the portsf Library

P. 220

(2) The argument int clip_floats is used to set the way in which floating-point data is
written to the file. As was noted above, the unique aspect of floating-point soundfiles is that
the samples can contain over-range values. Depending on the application, you may or many
not want these to be clipped to the normal maxima, 1.0 and þ1.0. Use of this facility
depends on whether you have requested that the soundfile include the PEAK chunk, which
records the maximum values in the file. As not all applications will even know about the
PEAK chunk (and will simply ignore it when reading), the safe approach is to set this argu-
ment to 1; but for experimental purposes you may want, for example, to offer the user the
choice. Needless to say, this parameter is ignored for all other sample formats.
(3) int minheader: it is an unfortunate fact of life that many applications fail to deal
with WAVE formats that contain optional chunks before the audio data—many older
UNIX-originated programs suffer from this. By setting minheader to 1, the soundfile is cre-
ated with a ‘‘minimum header’’ containing just the required format and data chunks—this
therefore means that no PEAK data will be written to the file. Ideally, of course, minheader
should be set to 0 always, and will be in all the examples presented here.

2.1.5 Basic Soundfile Handling in portsf—the switch. . .case Keywords

2.1.6 Reading and Writing—The Sample Frame

Note that the
nFrames argument is defined as a custom DWORD type. This name is borrowed from Win-
dows usage, and signifies (on a 32-bit platform) an unsigned long.

2.1.7 Streamlining Error Handling—The goto Keyword

```
if(sndWriteFloatFrames(ofd,frame,1) != 1)
	goto cleanup;

	/* lots of code. . .. */

cleanup:
	if(frame)
		free(frame);
```

2.1.8 Using portsf for Soundfile Conversion with PEAK Support

P. 229

Listing 2.1.3: sf2float.c

P. 230

2.1.9 Building Programs with portsf

P. 233

2.1.10 Exercises

[...]

Use the
format specifier ‘‘\r’’ (carriage return) within printf to overwrite the same line with the
updating message string.

...

The cal-
culation to convert a normalized amplitude (within e 1.0) to dB is
loudness (dB) = 20.0 * log 10 (amp);

[...]

2.2 Processing Audio

P. 235

2.2.1 The First Audio Process: Changing Level

[...], to make this a truly neutral template program: delete lines 39–45, [...]

sfgain

Stage 1: Define variables.
Stage 2: Obtain and validate arguments from user.
Stage 3: Allocate memory, open infile and outfile.
Stage 4: Perform main processing loop.
Stage 5: Report to the user.
Stage 6: Close files, free memory.

```
int i;
float ampfac;

printf("SFGAIN: change level of soundfile\n");

if (argc < 4) {
    printf("insufficient arguments.\n" "usage:\n\t" "sfgain infile outfile ampfac\n" "\twhere ampfac must be > 0\n");
    return 1;
}

ampfac = (float) atof(argv[3]);
if (ampfac <= 0.0) {
    printf("Error: ampfac must be positive.\n");
    return 1;
}

while (framesread == 1) {
    totalread++;
        for (i=0; i < props.chans; i++ )
            frame[i] *= ampfac;
        
        if (psf_sndWriteFloatFrames(ofd, frame, 1) != 1)  {
            printf("Error writing to outfile\n");
            error++;
            break;
        }
    framesread = psf_sndReadFloatFrames(ifd,frame,1);
}
```

2.2.2 Amplitude vs. Loudness—the Decibel

P. 237

Exercise 2.1.5 introduced the calculation for converting an amplitude value into decibels
(dB):
loudness dB = 20 x log 10 amp
(1)
where amp is in the range 0 to 1.0. We also need the complementary formula for converting
dB to amplitude:
amp = 10 ^ ( loudness db /20 )

2.2.3 Extending sfgain—Normalization

sfnorm

Stage 1: Set up variables.
Stage 2: Obtain and validate arguments from user.
Stage 3: Allocate memory, open the infile.
Stage 4a: Read PEAK amplitude of the infile; if not found,
Stage 4b: Scan the whole infile to obtain peak value; rewind file ready for processing stage.
Stage 4c: If peak 4 0, open outfile; otherwise, quit.
Stage 5: Perform main processing loop.
Stage 6: Report to the user.
Stage 7: Close files, free memory.


```
double dbval, inpeak = 0.0;
float ampfac, scalefac;

if (argc < 4) {
	printf("insufficient arguments.\n" "usage:\n\t" "sfnorm infile outfile dBval\n" "\twhere dBval <= 0.0\n");
	return 1;
}

dbval = (atof(argv[3]));
if (dbval > 0.0) {
	printf("Error: dBval cannot be positive.\n");
	return 1;
}
ampfac = (float) pow(10.0, dbval / 20.0);

```

int abs(int val);
double fabs(double val);

function that returns the maximum absolute value of a sample buffer:
```
double maxsamp(float* buf, unsigned long blocksize)
{
	double absval, peak = 0.0;
	unsigned long i;
	
	for (i = 0; i < blocksize; i++) {
		absval = fabs(buf[i]);
		if (absval > peak)
			peak = absval;
	}
	
	return peak;
}
```

You can place this function anywhere in sfnorm.c, so long as it is outside main.

Either way, you will have to add the declaration of the function:
double maxsamp(float* buf, unsigned long blocksize);

```
framesread = psf_sndReadFloatFrames(ifd,frame,1);

while (framesread == 1) {
	double thispeak;
	blocksize = props.chans;
	thispeak = maxsamp(frame,blocksize);
	if (thispeak > inpeak)
		inpeak = thispeak;
	framesread = psf_sndReadFloatFrames(ifd,frame,1);
}
```

#define NFRAMES (1024)

unsigned long nframes = NFRAMES;

```
frame = (float*) malloc(NFRAMES * props.chans * sizeof(float));

...

framesread = psf_sndReadFloatFrames(ifd,frame,nframes);

while (framesread > 0) {
	double thispeak;
	blocksize = framesread * props.chans;
	thispeak = maxsamp(frame,blocksize);
	if (thispeak > inpeak)
		inpeak = thispeak;
	framesread = psf_sndReadFloatFrames(ifd,frame,nframes);
}
```

```
/* get peak info: scan file if required */
/* inpeak has been initialized to 0 */
if (psf_sndReadPeaks(ifd,peaks,NULL) > 0) {
	long i;
	for (i=0; i < props.chans; i++) {
	if (peaks[i].val > inpeak)
		inpeak = peaks[i].val;
	}
}
else {
	/* scan the file, and rewind */
	
	/* rewind */
	if ((psf_sndSeek(ifd, 0, PSF_SEEK_SET)) < 0) {
    	printf("Error: unable to rewind infile.\n");
    	error++;
    	goto exit;
    }
}

/* check file is not silent */
if (inpeak==0.0) {
	printf("infile is silent! Outfile not created.\n");
	goto exit;
}

/* code that create the outfile */

/* immediately above the final processing loop */
/* calculate the scaling factor from db given by user */
scalefac = (float) (ampfac / inpeak);

```

2.2.4 Exercises

P. 245

#define max(x,y) ((x) > (y) ? (x) : (y))

With this macro, code such as
	absval = fabs(buf[i]);
	if(absval > peak)
		peak = absval;

can be rewritten more concisely, but also more expressively:
	absval = fabs(buf[i]);
	peak = max(absval,peak);

2.3 Stereo Panning

P. 246

Stage 3: Allocate memory, open infile.
Stage 4a: Perform any special data pre-processing, opening of extra data files, etc.
Stage 4b: Open outfile, once all required resources are obtained.
Stage 5: Perform main processing loop.
Stage 6: Report to the user.
Stage 7: Close files, free memory.

sfpan

usage : sfpan infile outfile panpos

```
enum {ARG_PROGRAMME, ARG_INFILE, ARG_OUTFILE, ARG_PANPOS, ARG_NARGS};

typedef struct panpos
	double left;
	double right;
} PANPOS;

PANPOS simplepan(double position)
{
	PANPOS pos;
	position *= 0.5;
	pos.left = position - 0.5;
	pos.right = position + 0.5;
	return pos;
}

int main(int argc, char** argv)
{
    float * outframe = NULL; /* STAGE 1 */
    PANPOS thispos; /*STAGE 1 */
    
	pos = atof(argv[ARG_PANPOS]);
	
	if ( (pos < 1.0) || (pos > 1.0) ) {
    	printf("Error: panpos value out of range -1 to +1\n");
    	error++;
    	goto exit;
    }

	...
	
	if (inprops.chans != 1) {
    	printf("Error: infile must be mono.\n");
    	error++;
    	goto exit;
    }
    
    outprops = inprops;
    
    // to switch from mono to stereo
    outprops.chans = 2;
    
    /* create stereo output buffer */
    outframe = (float *) malloc(nframes * outprops.chans * sizeof(float));
    if (outframe == NULL) {
    	puts("No memory!\n");
    	error++;
    	goto exit;
    }
    
	...
	
	thispos = simplepan(position);
    while ((framesread = psf_sndReadFloatFrames(ifd,inframe,nframes)) > 0)
    {
    	int i, out_i;
    	for (i=0, out_i = 0; i < framesread; i++) {
    		outframe[out_i++] = (float)(inframe[i]*thispos.left);
    		outframe[out_i++] = (float)(inframe[i]*thispos.right);
		}
		if (psf_sndWriteFloatFrames(ofd,outframe,framesread) != framesread) {
        	printf("Error writing to outfile\n");
        	error++;
        	break;
        }
    }
	
	...
	
    exit:
    	if (outframe) free (outframe);
}
```

2.3.6 Extending sfpan with Breakpoint-File Support

P. 253

2.3.8 Completing and Testing the New sfpan

P. 259

Exemple of breakpoint file:
0.0 -1.0
2.0 1.0
4.0 -1.0

2.3.9 A Better Panner—The Constant-Power Function

P. 261

and intensity (which relates to the
power of the signal) is measured by reading not the amplitude but the square of the ampli-
tude

The formulae for a constant-power pan function are:

A = square(2) / 2 * (cos (Delta) + sin (Delta))

Delta = angle of the panned signal

constpower to replace simplepan
```
PANPOS constpower(double position)
{
    PANPOS pos
    /* pi/2: 1/4 cycle of a sinusoid */
    const double piovr2 = 4.0 * atan(1.0) * 0.5; // = PI / 2
    const double root2ovr2 = sqrt(2.0) * 0.5; // = qrt(2) / 2
    
    /* scale position to fit the pi/2 range */
    double thispos = position * piovr2;
    
    /* each channel uses a 1/4 of a cycle */
    double angle = thispos * 0.5;
    
    pos.left = root2ovr2 * (cos(angle) - sin(angle));
    pos.right = root2ovr2 * (cos(angle) + sin(angle));

    return pos;
}
```

optimization => merge lines together

2.3.10 Objections to sfpan

sfpan: it is extremely inefficient!

[...]

- val_at_brktime

2.3.11 Exercises 

P. 265

2.4 Envelopes as Signals—Amplitude Processing

envx and sfenv

2.4.2 The envx Program—Describing the Task

We use linear interpolation to generate a stream of amplitude values, which
are multiplied with the audio samples, thus ‘‘enveloping’’ the source. In the context of real-
time audio processing, the combination of these two procedures is usually termed envelope
following.

2.4.3 Extracting Envelope Data from a Soundfile

envx.c
```
/* envx.c : extract amplitude envelope from mono soundfile*/
#define DEFAULT_WINDOW_MSECS (15)

void main()
{
	/* duration of the window in msecs */
	// double windur;
	double windur = DEFAULT_WINDOW_MSECS;
	
	unsigned long winsize;
	
	double brktime; /* holds the time for the current breakpoint time */
	
	// to count the number of breakpoints I think
	unsigned long npoints;
	
	/*STAGE 2 */
	printf("ENVX: extract amplitude envelope from mono soundfile\n");

	// Implementation of the -wN Flag
	if (argc > 1) {
    	char flag;
    	while (argv[1][0] == '-') {
    		flag = argv[1][1];
    		
    		switch(flag) {
				/*TODO: handle any flag arguments here */
				
				case('\0'):
					printf("Error: missing flag name\n");
					return 1;
					
				case('w'):
					windur = atof(&argv[1][2]);
					if (windur <= 0.0) {
						printf("bad value for Window Duration. Must be positive.\n");
						return 1;
					}
					break;
					
				default:
					break;
    		}
    		
    		argc--;
    		argv++;
    	}
    }
	
	/* check rest of commandline */
	if (argc < ARG_NARGS) {
		printf(
			"insufficient arguments.\n"
			"usage: envx [-wN] insndfile outfile.brk\n"
			"	-wN: set extraction window size to N msecs\n"
			"		(default: 15)\n"
			"usage: envx insndfile outfile.brk\n");
	
		return 1;
	}
	

	
	/* TODO: verify infile format for this application */
	
	/* verify infile format is acceptable */
	if (inprops.chans > 1) {
		printf("Soundfile contains %d channels: must be mono.\n",inprops.chans);
		error++;
		goto exit;
	}
	
	// STAGE 3 ---
	/* set buffersize to the required envelope window size */
    windur /= 1000.0;
    /* convert to secs */
    winsize = (unsigned long)(windur * inprops.srate);
    inframe = (float*) malloc(winsize * sizeof(float));
    if (inframe == NULL) {
    	puts("No memory!\n");
    	error++;
    	goto exit;
    }
	
	// STAGE 4 ---
	
	/* create output breakpoint file */
	fp = fopen(argv[ARG_OUTFILE],"w");
	if (fp == NULL) {
		printf(
			"envx: unable to create breakpoint file %s\n",
			argv[ARG_OUTFILE]);
			
		error++;
		goto exit;
	}
	
	// don't know if it is here
	brktime = 0.0;
	npoints = 0;
	
	while ((framesread = psf_sndReadFloatFrames(ifd, inframe, winsize)) > 0) {
    	double amp;
    	/* find peak sample of this block */
    	amp = maxsamp(inframe, framesread);
    	
    	/* store brktime and amp as a breakpoint */
    	if (fprintf(fp, "%f\t%f\n", brktime, amp) < 2) {
        	printf("Failed to write to breakpoint file %s\n",argv[ARG_OUTFILE]);
        	error++;
        	break;
        }
    	
    	brktime += windur;
    	npoints++;
    }
	
	if (framesread < 0) {
    	printf("Error reading infile. Outfile is incomplete.\n");
    	error++;
    }
    else
    	printf("Done: %d errors\n",error);
    	
    printf("%d breakpoints written to %s\n", npoints, argv[ARGV_OUTFILE]);
    
    
	
	// STAGE 7 ---
	
	/*TODO: cleanup any other resources */
	if (fp)
		if (fclose(fp))
			printf("envx: failed to close output file %s\n", argv[ARG_OUTFILE]);
}
```

2.4.4 Implementation of Envelope Extraction

As a rule of thumb, a 15-millisecond window is sufficient to capture the envelope of most
sounds—this gives around 66 envelope points per second.

2.4.5 Efficient Envelope Processing—The Program sfenv

P. 275





TODO: make envx
TODO: make sfpan work

DONE: try rawsoundfile in audacity

TODO: faire tfork2:
TODO: faire sfgain et sfnorm:
TODO: faire le son en boucle:

TODO : relire 1.2.8 A Musical Computation (et peut-être un peu après) (P. 95)

TODO: relire les trucs avec les fréquences des notes dans les gammes

TODO: build with portsf the sf2float

dans sinetext :

comprendre pourquoi on fait ça : angleincr = twopi * freq / srate
(diviser frequence en Hz par number of samples recorded every second ???)

dans expdecay :

comprendre pourquoi on fait ça :
k = dur / nsteps; /* the constant time increment */
a = exp(-k / T); /* calc the constant ratio value */
avec T = ???

The exp() function in C++ returns the exponential (Euler’s number) e (or 2.71828) raised to the given argument.

One general formula for an exponential decay is
x = ae ^ -k/T
where a and k are constants and T represents the time constant—the rate of decay.




















---

TODO : chercher bouquins maths et ddl pour dvd or achat

Handbook of Mathematical Functions, With Formulas, Graphs, and Mathematical Tables

Mathématiques pour l'informatique

Toutes les mathématiques et les bases de l'informatique

Outils mathématiques pour le génie des procédés

Eléments De Mathématiques Du Signal

Mathématiques pour le traitement du signal

Mathématiques de la Terminale S à la prépa scientifique

---

https://www.musicdsp.org/

https://www.mastersynth.com/dossiers/27-les-lfo-caracteristiques-et-applications-avec-les-synthetiseurs

Oscillateur
LFO = Low Frequency Oscillator

0,1 Hz correspond à une période toutes les dix secondes, 20 Hz à vingt périodes par seconde, etc

l'égaliseur
reverb
delay
phaser
fader

https://www.easyzic.com/dossiers/les-fonctions-de-synthese,d68.html

http://subaru.univ-lemans.fr/AccesLibre/UM/Pedago/physique/02/meca/ondeprog.html#:~:text=La%20fr%C3%A9quence%20est%20%C3%A9gale%20%C3%A0,anim%C3%A9s%20d'un%20mouvement%20sinuso%C3%AFdal.

https://zestedesavoir.com/tutoriels/2451/les-signaux-sinusoidaux-en-physique/les-signaux-sinusoidaux/

































