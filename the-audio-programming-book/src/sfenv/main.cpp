/* sfenv.c: apply an amplitude envelope to a mono soundfile */
/* USAGE: sfenv [-n] [normalizing value] insndfile infile.brk outsndfile */

#include <stdio.h>
#include <stdlib.h>
#include <../portsf/include/portsf.h>
#include <../breakpoints/include/breakpoints.h>
#include <math.h>

#define NFRAMES 100

enum {ARG_PROGRAMME, ARG_INSNDFILE, ARG_INBRKFILE, ARG_OUTSNDFILE, ARG_NARGS};

int main(int argc, char**argv)
{
    brk_test();

	PSF_PROPS inprops, outprops;

	long framesread,    /* number of frames copied to buffer */
         totalread;     /* running count of sample frames */

    BREAKPOINT* points;

    long size; /* number of breakpoints */

    psf_format outformat = PSF_FMT_UNKNOWN;

    double timeincr, sampletime;

    char flag; /* check for any command line options */
    int isnorm = 0; /* -n option will normalize breakpoints */
    double scalefac = 1.0; /* scale factor used for normalizing breakpoints */
    double normvalue = 1.0; /* normalizing value is set to 1 by default */
    BREAKPOINT max; /* breakpoint with the maximum value, used for scaling */

    /* init all resource vals to default states */
    int ifd = 1, ofd = -1;
    int error = 0;
    float* buffer = NULL;
    FILE* fp = NULL;

    printf("SFENV: apply an amplitude envelope to a mono soundfile.\n");

    /* check for any command-line options */
	if (argc > 1) {
		while (argv[1][0] == '-') {
			flag = argv[1][1];
			switch(flag) {
				case('\0'):
					printf("ERROR:\tyou did not specify an option after using '-'\n"
					       "USAGE:\tsfenv [-n] [normalizing value] insndfile infile.brk outsndfile\n"
					       "      \t-n option normalizes breakpoint values\n"
					       "      \tbreakpoint values are normalized to 1 by default.\n"
					      );

					return 1;

				case('n'):
					if (argv[1][2] == '\0') {
						isnorm = 1;
						/* check for a normalizing value after -n option */
						if (
                            (argc >= 6)
                            && (
                                ( (argv[2][0] >= '0') && (argv[2][0] <= '9') ) || /* check if the first character is [0-9] */
                                ( (argv[2][0] == '-') && ( (argv[2][1] >= '0') && (argv[2][1] <= '9') ) ) || /* or check if the first two characters are -[0-9] OR .[0-9] */
                                ( (argv[2][0] == '-') && (argv[2][1] == '.' ) ) || /* or check if the first two characters are -. */
                                ( (argv[2][0] == '.') && ( (argv[2][1] >= '0') && (argv[2][1] <= '9') ) )
                            )
                        ) {
                            /* a normalizing value was specified */
							if ((normvalue = atof(argv[2])) <= 0.0) {
								printf("ERROR:\tnormalizing value must be positive.\n"
								       "USAGE:\tsfenv [-n] [normalizing value] insndfile infile.brk outsndfile\n"
								       "      \t-n option normalizes breakpoint values\n"
								       "      \tbreakpoint values are normalized to 1 by default.\n"
								      );

								return 1;
							}

							argc--;
							argv++;
						}

						break;
					}

				default:
					printf("ERROR:\t%s is not a valid option.\n"
					       "USAGE:\tsfenv [-n] [normalizing value] insndfile infile.brk outsndfile\n"
					       "      \t-n option normalizes breakpoint values\n"
					       "      \tbreakpoint values are normalized to 1 by default.\n",
					       argv[1]
					      );

					return 1;
			}

			argc--;
			argv++;
		}
	}

    if (argc != ARG_NARGS) {
		printf("ERROR:\tinsufficient arguments.\n"
		       "USAGE:\tsfenv [-n] [normalizing value] insndfile infile.brk outsndfile\n"
		       "      \t-n option normalizes breakpoint values\n"
		       "      \tbreakpoint values are normalized to 1 by default.\n"
		      );

		return 1;
	}

    /* startup portsf */
	if (psf_init()) {
		printf("ERROR: unable to start portsf\n");
		return 1;
	}

    /* open infile */
	ifd = psf_sndOpen(argv[ARG_INSNDFILE], &inprops, 0);
	if (ifd < 0) {
		printf("ERROR: unable to open infile \"%s\"\n", argv[ARG_INSNDFILE]);
		return 1;
	}

    /* we now have a resource, so we use goto hereafter on hitting any error */

    /* check if infile uses 8-bit samples*/
	if (inprops.samptype == PSF_SAMP_8) {
		printf("ERROR: sfenv does not support 8-bit format.\n");
		error++;
		goto exit;
	}

	/* check if infile is mono */
	if (inprops.chans != 1) {
		printf("ERROR: infile has %d channels, must be mono.\n", inprops.chans);
		error++;
		goto exit;
	}

	/* display infile properties */
	if (!psf_sndInfileProperties(argv[ARG_INSNDFILE], ifd, &inprops)) {
		error++;
		goto exit;
	}

	/* open input breakpoint file */
	fp = fopen(argv[ARG_INBRKFILE], "r");
	if (fp == NULL) {
		printf("ERROR: unable to open breakpoint file: %s\n", argv[ARG_INBRKFILE]);
		error++;
		goto exit;
	}

	/* get breakpoint data */
	points = get_breakpoints(fp, &size);
	if (points == NULL) {
		printf("ERROR: no breakpoints read.\n");
		error++;
		goto exit;
	}

    if (size < 2) {
		printf("ERROR: at least two breakpoints required.\n");
		error++;
		goto exit;
	}

	/* make sure the first breakpoint starts at 0.0 */
	if (points[0].time != 0.0) {
		printf("Error in breakpoint data: first time must be 0.0\n");
		error++;
		goto exit;
	}

	/* if specified by the user, normalize all the breakpoint values */
	if (isnorm) {
		max = maxpoint(points, size);
		scalefac = normvalue / max.value;
	}

	/* outfile properties are identicle to infile properties */
	outprops = inprops;

	/* create outfile */
	ofd = psf_sndCreate(argv[ARG_OUTSNDFILE], &outprops, 0, 0, PSF_CREATE_RDWR);
	if (ofd < 0) {
		printf("ERROR: unable to create \"%s\"\n", argv[ARG_OUTSNDFILE]);
		error++;
		goto exit;
	}

    /* check if oufile extension is one we know about */
	outformat = psf_getFormatExt(argv[ARG_OUTSNDFILE]);
	if (outformat == PSF_FMT_UNKNOWN) {
		printf("Outfile name \"%s\" has unknown format.\n"
		       "Use any of .wav .aiff .aif .afc .aifc\n",
		       argv[ARG_OUTSNDFILE]);
		error++;
		goto exit;
	}

	/* allocate memory for read write buffer */
	buffer= (float*)malloc(NFRAMES * sizeof(float));
	if (buffer == NULL) {
		puts("No memory!\n");
		error++;
		goto exit;
	}

    printf("copying outfile...\n");

	/* initialize running count of sample frames */
	totalread = 0;

	/* initialize time position counter for reading envelope */
	timeincr = 1.0 / inprops.srate;
	sampletime = 0.0;
	unsigned long pointnum = 1;
	/* keep track of any missed breakpoint values */
	unsigned long lastpoint, misscount = 0;

    /* loop every time NFRAMES are copied, report any errors */
	while ((framesread = psf_sndReadFloatFrames(ifd, buffer, NFRAMES)) > 0) {
		int i;

		totalread += framesread;

		/* extract breakpoint values */
		double thisamp;

		/* when processing each sample value, each breakpoint
		   span is calculated by taking the "rightmost" breakpoint time
		   to the left of the current sampletime and the "leftmost"
		   breakpoint time to the right of the current sampletime */
		for (i = 0; i < framesread; i++, sampletime += timeincr) {
			lastpoint = pointnum;
			thisamp = val_at_brktime(points, size, &pointnum, sampletime);
			while (pointnum - lastpoint > 2) {
				lastpoint++;
				misscount++;
			}
			buffer[i] = (float)(buffer[i] * thisamp * scalefac);
		}

		if (psf_sndWriteFloatFrames(ofd, buffer, framesread) != framesread) {
			printf("Error writing to outfile.\n");
			error++;
			break;
		}
	}

	printf(
        "\nDone: %d error%s\n"
        "soundfile created: %s\n"
        "samples copied: %lu\n\n",
        error, (error == 1) ? "" : "s", argv[ARG_OUTSNDFILE], totalread
    );

	/* tell user if breakpoint values weren't read */
	int plural = misscount > 1;
	if (misscount)
		printf(
            "Warning: %lu breakpoint value%s %s not read.\n"
            "         %s breakpoint%s could either have \n"
            "         the same time value%s as other adjacent breakpoints\n"
            "         or %s time difference is smaller than the \n"
            "         time difference of each sample.\n\n",
            misscount, (plural) ? "s" : "", (plural) ? "were" : "was",
            (plural) ? "These" : "This", (plural) ? "s": "", (plural) ? "s" : "",
            (plural) ? "their" : "its"
        );

    return 0;

	/* do all the cleanup */
	exit:
    if (ifd >= 0)
        if (psf_sndClose(ifd))
            printf("sfenv: failed to close infile: %s\n\n",argv[ARG_INSNDFILE]);

    if (ofd >= 0) {
        if (psf_sndClose(ofd))
            printf("sfenv: failed to close outfile: %s\n\n",argv[ARG_OUTSNDFILE]);
        else if (error) {
            printf("There was an error while processing the sound file.\n"
                   "Deleting outfile: %s ...\n", argv[ARG_OUTSNDFILE]);

            if (remove(argv[ARG_OUTSNDFILE]))
                printf("Error: failed to delete %s\n\n", argv[ARG_OUTSNDFILE]);
            else
                printf("%s successfully deleted.\n\n", argv[ARG_OUTSNDFILE]);
        }
    }

	if (fp)
		if(fclose(fp))
			printf("sfenv: failed to close breakpoint file: %s\n\n",argv[ARG_INBRKFILE]);

	if (buffer)
		free(buffer);

	psf_finish();

	return error;
}

