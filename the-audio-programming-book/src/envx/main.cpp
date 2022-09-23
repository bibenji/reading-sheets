/* envx.c : extract amplitude envelope from mono soundfile*/
/* USAGE: envx [-wN] [-n] insndfile outfile.brk */

#include <stdio.h>
#include <stdlib.h>
#include <portsf.h>
#include <breakpoints.h>
#include <math.h>

#define DEFAULT_WINDOW_MSECS (15)

enum {ARG_PROGRAMME, ARG_INFILE, ARG_OUTFILE, ARG_NARGS};

int main(int argc, char**argv)
{
	PSF_PROPS inprops;

	long framesread,    /* number of frames copied to buffer */
         totalread;     /* running count of sample frames */

    unsigned long npoints; /* number of breakpoints generated */
    unsigned long winsize; /* number of infile buffer frames */

    char flag;

    double brktime; /* holds the time for the current breakpoint time */
	double windur = DEFAULT_WINDOW_MSECS; /* duration of the window in msecs */

	int isnorm = 0; /* -n option for normalizing breakpoints */

	double peak; /* peak value of the soundfile */
	double scalefac = 1.0; /* normalizing scale factor */

	/* init all resource vals to default states */
	int ifd = -1;
	int error = 0;
	float* inframe = NULL;
	FILE *fp = NULL;

	/*STAGE 2 */
	printf("ENVX: extract amplitude envelope from mono soundfile.\n");

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

	/* display infile properties */
	if (!psf_sndInfileProperties(argv[ARG_INFILE], ifd, &inprops)) {
		error++;
		goto exit;
	}

	// STAGE 4 ---

	/* create output breakpoint file */
	fp = fopen(argv[ARG_OUTFILE], "w");
	if (fp == NULL) {
        printf("ERROR: unable to created breakpoint file: %s\n", argv[ARG_OUTFILE]);
        error++;
        goto exit;
	}

	// STAGE 3 ---
	/* set buffersize to the required envelope window size */
    windur /= 1000.0;

    /* check if window duration is smaller than duration between two samples */
    if (windur < 1./inprops.srate) {
        printf("ERROR: window size cannot be smaller than the\n"
		       "       time length between two sample frames.\n");
		error++;
		goto exit;
    }

    /* convert to secs */
    winsize = (unsigned long)(windur * inprops.srate);

    if (isnorm) {
        peak = psf_sndPeakValue(ifd, &inprops);

		if (peak < 0) {
			printf("Error normalizing breakpoints: peak value cannot be read from soundfile.\n");
			error++;
			goto exit;
		}

		if (peak == 0.0) {
			printf("ERROR: soundfile is silent!\n");
			error++;
			goto exit;
		}

		scalefac = 1.0 / peak;
    }

    /* allocate memory for infile buffer */
    inframe = (float*) malloc(winsize * sizeof(float));

    if (inframe == NULL) {
    	puts("No memory!\n");
    	error++;
    	goto exit;
    }

	printf("creating breakpoint file...\n");

	/* initialize running count of sample frames */
	totalread = 0;
	/* initialize breakpoint time */
	brktime = 0.0;
	/* initialize count of breakpoints */
	npoints = 0;

	/* loop every time winsize numbers of frames are copied, report any errors */
	while ((framesread = psf_sndReadFloatFrames(ifd, inframe, winsize)) > 0) {
        totalread += framesread;

        /* extract breakpoint value */
    	double amp;
    	/* find peak sample of this block */
    	amp = scalefac * maxsamp(inframe, framesread);

    	/* store brktime and amp as a breakpoint */
    	/* write breakpoint time and value to outfile */
    	if (fprintf(fp, "%f\t%f\n", brktime, amp) < 2) {
        	printf("Failed to write to breakpoint file %s\n",argv[ARG_OUTFILE]);
        	error++;
        	goto exit;
        }

        npoints++;

    	brktime += windur;
    }

	if (framesread < 0) {
    	printf("Error reading infile. Outfile is incomplete.\n");
    	error++;
    }
    else
    	printf("Done: %d errors%s\n"
                "breakpoint file created: %s\n"
                "breakpoints written: %lu\n",
                error, (error == 1) ? "" : "s", argv[ARG_OUTFILE], npoints
        );

    return 0;

	// STAGE 7 ---
	/* do all cleanup */
	exit:
	if (ifd >= 0)
        if (psf_sndClose(ifd))
            printf("envx: failed to close infile: %s\n", argv[ARG_INFILE]);

	if (fp) {
        if (fclose(fp)) {
            printf("envx: failed to close breakpoint file %s\n", argv[ARG_OUTFILE]);
        }
        else if (error) {
            printf("\nThere was an error while processing the breakpoint file.\n"
                "Deleting outfile: %s ...\n", argv[ARG_OUTFILE]);
            if (remove(argv[ARG_OUTFILE])) {
                printf("Error: failed to deleted %s\n", argv[ARG_OUTFILE]);
            }
            else {
                printf("%s successfully deleted.\n", argv[ARG_OUTFILE]);
            }
        }
	}

    if (inframe)
        free(inframe);

    psf_finish();

    return error;
}
