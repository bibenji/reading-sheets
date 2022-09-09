#include <iostream>
#include "breakpoints.h"

void testBreakpoint()
{
    printf("%s\n", "Hello world!");
}

BREAKPOINT* get_breakpoints(FILE* fp, long* psize)
{
    int got; // use to put result of sscanf

    long npoints = 0; // number of points?
    long size = 64; // the size of...

    double lasttime = 0.0;

    BREAKPOINT * points = NULL; // pointer to points

    char line[80]; // use to put a line read by fgets

    if (fp == NULL)
        return NULL;

    // we allocate memory for BREAKPOINT * size
    // pointer to a BREAKPOINT (like array?)
    points = (BREAKPOINT*)(malloc(sizeof(BREAKPOINT) * size));

    if (points == NULL)
        return NULL;

    while (fgets(line, 80, fp)) {
        got = sscanf(line, "%lf %lf", &points[npoints].time, &points[npoints].value);

        if (got < 0)
            continue; // empty line

        if (got == 0) {
            printf("Line %d has non-numeric data\n", npoints + 1);
            break;
        }

        if (got == 1) {
            printf("Incomplete breakpoint found at point %d\n", npoints + 1);
            break;
        }

        if (points[npoints].time < lasttime)  {
			printf("Data error at point %ld: time not increasing.\n", npoints + 1);
			break;
		}

        lasttime = points[npoints].time;

        if (++npoints == size) {
            BREAKPOINT* tmp;
            size += npoints;
            tmp = (BREAKPOINT*)(realloc(points, sizeof(BREAKPOINT) * size));

            if (tmp == NULL) {
                /* have to release the memory, and return NULL to caller */
                npoints = 0;
                free(points);
                points = NULL;
                break;
            }

            points = tmp;
        }
    }

    if (npoints)
        *psize = npoints;

    return points;
}

BREAKPOINT maxpoint(const BREAKPOINT* points, long npoints)
{
    int i;

	BREAKPOINT point;

	point.time = points[0].time; /* initialize from first point */
	point.value = points[0].value;

	for (i = 0; i < npoints; i++) {
		if (point.value < points[i].value) {
			point.value = points[i].value;
			point.time = points[i].time;
		}
	}

	return point;
}

int inrange(
    const BREAKPOINT* points,
    double minval,
    double maxval,
    unsigned long npoints) {
    unsigned long i;
    int range_OK = 1;

    for (i = 0; i < npoints; i++) {
        if (points[i].value < minval || points[i].value > maxval) {
            range_OK = 0;
            break;
        }
    }

    return range_OK;
}

double val_at_brktime(
    const BREAKPOINT* points,
    unsigned long npoints,
    double time) {
    unsigned long i;
    BREAKPOINT left, right;
    double frac, val, width;

    /* scan until we find a span containing our time */
    for (i = 1; i < npoints; i++) {
        if (time <= points[i].time)
        break;
    }

    /* maintain final value if time beyond end of data */
    if (i == npoints) {
        return points[i-1].value;
    }

    left = points[i-1];
    right = points[i];

    /* check for instant jump - two points with same time */
    width = right.time - left.time;
    if (width == 0.0)
        return right.value;

    /* get value from this span using linear interpolation */
    frac = (time - left.time) / width;
    val = left.value + (( right.value - left.value) * frac);
    return val;
}
