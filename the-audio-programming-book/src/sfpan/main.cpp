#include <iostream>
#include <stdio.h>
#include <stdlib.h>

#include "breakpoints.h"
#include "portsf.h"

#define NFRAMES 100

/* panning function returns a PANPOS object */
PANPOS simplepan(double position);
/* replacement for simplepan: constant power panning function */
PANPOS constpower(double position);

enum {ARG_PROGNAME, ARG_INFILE, ARG_OUTFILE, ARG_BRKFILE, ARG_NARGS};

int main(int argc, char**argv)
{
    testBreakpoint();

    int error = 0;
    int framesread;
    int ifd = -1, ofd = -1;

    long size;

    double timeincr, sampletime;

    float* inframe = NULL;
	float* outframe = NULL;

	BREAKPOINT* points = NULL;
	FILE* fp = NULL;
	PANPOS pos;
	PSF_PROPS inprops, outprops;
	psf_format outformat = PSF_FMT_UNKNOWN;

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
        printf("Error: unable to open breakpoint file %s\n", argv[ARG_BRKFILE]);
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

    /* check if breakpoint values are in range */
    if (!inrange(points, -1, 1.0, size)) {
        printf("Error in breakpoint file: values out of range -1 to +1 \n");
        error++;
        goto exit;
    }

    /* start up portsf */
    if (psf_init()) {
        printf("Error: unable to start portsf.\n");
        return 1;
    }

    /* open infile */
    ifd = psf_sndOpen(argv[ARG_INFILE], &inprops, 0);
    if (ifd < 0) {
        printf("Error: unable to open \"%s\"\n", argv[ARG_INFILE]);
        return 1;
    }

    /* check if infile is 8-bit */
	if (inprops.samptype == PSF_SAMP_8) {
		printf("ERROR: portsf does not support 8-bit soundfiles.\n");
		error++;
		goto exit;
	}

	/* check if infile is in mono */
	if (inprops.chans != 1) {
		printf("ERROR: infile must be mono.\n");
		error++;
		goto exit;
	}

	/* properties of infile and outfile will be the same except infile is mono and outfile is stereo */
	outprops = inprops;
	outprops.chans = 2;

	ofd = psf_sndCreate(argv[ARG_OUTFILE], &outprops, 0, 0, PSF_CREATE_RDWR);
	if (ofd < 0) {
        printf("Error: unable to create \"%s\"\n", argv[ARG_OUTFILE]);
        error++;
        goto exit;
	}

    /* check if outfile extension is one we know about */
	outformat = psf_getFormatExt(argv[ARG_OUTFILE]);
	if (outformat == PSF_FMT_UNKNOWN) {
		printf("Outfile name \"%s\" has unknown format.\n"
		       "Use any of .wav .aiff .aif .afc .aifc\n",
		        argv[ARG_OUTFILE]
        );
		error++;
		goto exit;
	}

    /* allocate space for input frame buffer */
	inframe = (float*)malloc(sizeof(float)*inprops.chans*NFRAMES);
	/* allocate space for output frame buffer */
	outframe = (float*)malloc(sizeof(float)*outprops.chans*NFRAMES);

    /* init time position counter for reading envelope */
    timeincr = 1.0 / inprops.srate;
    sampletime = 0.0;

    // inframe est le buffer, nframe le nombre qu'on veut lire, i.e. : 100 ou moins
    // framesread = le nombre lu
    while ((framesread = psf_sndReadFloatFrames(ifd, inframe, NFRAMES)) > 0) {
        int i, out_i;
        double stereopos;

        for (i = 0, out_i = 0; i < framesread; i++) {
            // on récupère la valeur du point stéréo au sampletime
            stereopos = val_at_brktime(points, size, sampletime);

            // don't know what is that ???
            /* replacement for simplepan: constant power panning function */
            pos = constpower(stereopos);

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
	if (points)
		free(points);
	if (ifd)
		psf_sndClose(ifd);
	if (ofd)
		psf_sndClose(ofd);
	if (inframe)
        free(inframe);
	if (outframe)
		free(outframe);
	psf_finish();

	return error;
}
