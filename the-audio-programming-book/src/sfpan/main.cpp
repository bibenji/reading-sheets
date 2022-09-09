#include <iostream>
#include <stdio.h>
#include <stdlib.h>

#include "breakpoints.h"

enum {ARG_PROGNAME, ARG_INFILE, ARG_OUTFILE, ARG_BRKFILE, ARG_NARGS};

int main(int argc, char**argv)
{
    testBreakpoint();

    int error = 0;
    int framesread;
    int ifd = -1, ofd = -1;

    long size;

    float* inframe = NULL;
	float* outframe = NULL;

    FILE* fp = NULL;
    BREAKPOINT* points = NULL;

    if (argc != ARG_NARGS) {
        printf(
            "ERROR: insufficient number of arguments.\n"
            "USAGE: sfpan infile outfile posfile.brk\n"
            "\tposfile.brk is a breakpoint file\n"
            "\twitch values in range -1.0 <= pos <= 1.0\n"
            "\twhere -1.0 = full Left, 0 = Centre, +1.0 = full Right"
        );

        return 1;
    }

    /* read breakpoint file and verify it */
    fp = fopen(argv[ARG_BRKFILE], "r");
    if (fp == NULL) {
        printf(
            "Error: unable to open"
            "breakpoint file %s\n",
            argv[ARG_BRKFILE]
        );
        error++;
        goto exit;
    }

    points = get_breakpoints(fp, &size);

    if (points == NULL) {
        printf("No breakpoints read.\n");
        error++;
        goto exit;
    }

    if (size < 2) {
        printf("Error: at least two breakponts required\n");
        free(points);
        fclose(fp);
        return 1;
    }

    /* we require breakpoints to start from 0 */
    if (points[0].time != 0.0) {
        printf("Error in breakpoint data: first time must be 0.0\n");
        error++;
        goto exit;
    }

    if (!inrange(points, -1, 1.0, size)) {
        printf("Error in breakpoint file: "
            "values out of range -1 to +1 \n");
        error++;
        goto exit;
    }

    // TODO: all the stuffs with portsf to open a file
    // use that: https://github.com/umnum/Audio_Programs/blob/master/CH2/sfpan/sfpan.c

    /* allocate space for input frame buffer */
	inframe = (float*)malloc(sizeof(float)*inprops.chans*NFRAMES);
	/* allocate space for output frame buffer */
	outframe = (float*)malloc(sizeof(float)*outprops.chans*NFRAMES);

    /* init time position counter for reading envelope */
    timeincr = 1.0 / inprops.srate;
    sampletime = 0.0;

    // inframe est le buffer, nframe le nombre qu'on veut lire, i.e. : 100 ou moins
    // framesread = le nombre lu
    while ((framesread = psf_sndReadFloatFrames(ifd, inframe, nframe)) > 0) {
        int i, out_i;
        double stereopos;

        for (i = 0, out_i = 0; i < framesread; i++) {
            // on récupère la valeur du point stéréo au sampletime
            stereopos = val_at_brktime(points, size, sampletime);

            // don't know what is that ???
            pos = simplepan(stereopos);

            // on set sur le buffer les trucs en stéréo pour chaque valeur de i
            outframe[out_i++] = (float)(inframe[i] * pos.left);
            outframe[out_i++] = (float)(inframe[i] * pos.right);

            // incr sampletime by the timeincr calculate on top of that
            sampletime += timeincr;
        }

        if (psf_sndWriteFloatFrames(ofd, outframe, framesread) != framesread) {
            printf("Error writing to outfile\n");
            error++;
            break;
        }
    }

    printf("Done.\n");
    return 0;

    exit:
	if (fp)
		fclose(fp);
	// if (points)
	//	free(points);
	// if (ifd)
	//	psf_sndClose(ifd);
	// if (ofd)
	//	psf_sndClose(ofd);
	// if (inbuffer)
    //    free(inbuffer);
	// if (outbuffer)
	//	free(outbuffer);
	// psf_finish();

	return error;
}
