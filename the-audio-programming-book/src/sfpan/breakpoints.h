#ifndef BREAKPOINTS_H_INCLUDED
#define BREAKPOINTS_H_INCLUDED

typedef struct breakpoint
{
    double time;
    double value;
} BREAKPOINT;

void testBreakpoint();

BREAKPOINT* get_breakpoints(FILE* fp, long* psize);

BREAKPOINT maxpoint(const BREAKPOINT* points, long npoints);

int inrange(
    const BREAKPOINT* points,
    double minval,
    double maxval,
    unsigned long npoints);

double val_at_brktime(
    const BREAKPOINT* points,
    unsigned long npoints,
    double time);

#endif // BREAKPOINTS_H_INCLUDED
