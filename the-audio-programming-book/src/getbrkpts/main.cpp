#include <iostream>

using namespace std;

typedef struct breakpoint {
    double time;
    double value;
} BREAKPOINT;

BREAKPOINT maxpoint(const BREAKPOINT* points, long npoints);

BREAKPOINT* get_breakpoints(FILE* fp, long* psize);

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
    points = (BREAKPOINT*) malloc(sizeof(BREAKPOINT) * size);

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

        lasttime = points[npoints].time;

        if (++npoints == size) {
            BREAKPOINT* tmp;
            size += npoints;
            tmp = (BREAKPOINT*) realloc(points, sizeof(BREAKPOINT) * size);

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

int main(int argc, char* argv[])
{
    long size;
    double dur;
    BREAKPOINT point, *points;
    FILE* fp;

    printf("breakdur: find duration of breakpoint file\n");

    if (argc < 2) {
        printf("usage: breakdur breakpoints.txt\n");
        return 0;
    }

    fp = fopen(argv[1], "r");

    if (fp == NULL) {
        return 0;
    }

    size = 0;
    points = get_breakpoints(fp, &size);

    if (points == NULL) {
        printf("No breakpoints read.\n");
        fclose(fp);
        return 1;
    }

    if (size < 2) {
        printf("Error: at least two breakpoints required\n");
        free(points);
        fclose(fp);
        return 1;
    }

    if (points[0].time != 0.0) {
        printf("Error in breakpoint data: first time must be 0.0\n");
        free(points);
        fclose(fp);
        return 1;
    }

    printf("read %d breakpoints\n", size);
    dur = points[size-1].time;
    printf("duration: %f seconds\n", dur);
    point = maxpoint(points, size);
    printf("maximum value: %f at %f secs\n", point.value, point.time);
    free(points);
    fclose(fp);

    return 0;
}
