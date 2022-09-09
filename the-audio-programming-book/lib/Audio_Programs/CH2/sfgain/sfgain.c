/* sfgain.c: change level of soundfile */
/* USAGE: sfgain infile outfile buffer limit N [dBval | -a ampfac] */ 
#include <stdio.h>
#include <stdlib.h>
#include <portsf.h>
#include <math.h>

enum {ARG_PROGNAME, ARG_INFILE, ARG_OUTFILE, ARG_BUFF, ARG_LIMIT, ARG_N, ARG_DB, ARG_OP=6, ARG_NOPS, ARG_AMP=7, ARG_OPS};

int main(int argc, char**argv)
{
	PSF_PROPS props;
	long framesread, totalread; 
	DWORD nFrames;	
	int size; 
	int limit;
	int N; // copy infile N times
	float ampfac; 
	float dbval;

	/* init all resource vals to default states */ 
	int ifd=-1, ofd=-1;
	int error=0;
	psf_format outformat = PSF_FMT_UNKNOWN; 
	PSF_CHPEAK* peaks = NULL;
	float* buffer = NULL;

	/* init flags for command-line options */ 
	int isamp=0; // default scale factor is in dBs 

	printf ("SFGAIN: change level of soundfile.\n");

	if ((argc<ARG_NOPS)||(argc>ARG_OPS))
	{
		printf("insufficient arguments.\n"
		       "USAGE:\tsfgain infile outfile buffer limit N [dBval | -a ampfac]\n"
		       "dBval must be <= 0 or ampfac must be > 0\n"); 
		return 1;
	}

	/* check for command-line options */
	if (argc==ARG_OPS)
	{
		if (argv[ARG_OP][0]=='-')	
		{
			if (argv[ARG_OP][1]=='a')
				isamp=1;
			else
			{
				printf("ERROR: %s is not a valid command-line option.\n"
				       "USAGE:\tsfgain infile outfile buffer limit N [dBval | -a ampfac]\n"
				       "dBval must be <= 0 or ampfac must be > 0\n", argv[ARG_OP]); 
				return 1;
			}
		}
	}
	
	/* startup portsf */
	if(psf_init())
	{
		printf("ERROR: unable to start portsf\n");
		return 1;
	}

	/* initialize buffer */
	nFrames = (DWORD)atoi(argv[ARG_BUFF]);
	if (nFrames < 1)
	{
		printf("ERROR: buffer size must be at least 1\n");
		return 1;
	}

	/* initialize limit */
	limit = atoi(argv[ARG_LIMIT]);
	if (limit<1)
	{
		printf("ERROR: size limit must be positive.\n");
		return 1;
	}

	/* initialize N */ 
	N = atoi(argv[ARG_N]);
	if (N<1)
	{
		printf("ERROR: N must be at least one.\n");
		return 1;
	}

	/* initialize dBval or ampfac */
	if (isamp)
	{	
		ampfac = atof(argv[ARG_AMP]);
		if (ampfac <= 0.0)
		{
			printf("ERROR: ampfac must be positive.\n");
			return 1;
		}		
		if (ampfac == 1.0)
		{
			printf("ERROR: an ampfac of 1 creates an outfile "
			       "       identicle to the infile\n");
			return 1;
		}
	}	
	else
	{
		dbval = atof(argv[ARG_DB]);
		if (dbval > 0.0)
		{
			printf("ERROR: dBval cannot be positive.\n");
			return 1;
		}
		if (dbval==0.0)
		{
			printf("ERROR: dBval of 0 creates an outfile "
						 "identicle to the infile\n");  
			return 1;
		}
		/* convert dB to amps */
		ampfac = pow(10.0, dbval/20.0);
	}

	/* open infile */ 
	ifd = psf_sndOpen(argv[ARG_INFILE], &props, 0);
	if (ifd<0)
	{
		printf("ERROR: unable to open infile \"%s\"\n",argv[ARG_INFILE]);
		return 1;
	}

	/* we now have a resource, so we use goto hereafter
		 on hitting any error */
	/* get number of frames from infile */
	size = psf_sndSize(ifd);
	if(size<0)
	{
		printf("ERROR: unable to obtain the size of \"%s\"\n",argv[ARG_INFILE]);
		error++;
		goto exit;
	}
	/* check if copy limit is less than size */
	if(size<limit)
	{
		printf("ERROR: infile size is less than the copy limit.\n"
		       "infile:\t%s\n"
		       "infile size:\t%d frames\n"
		       "copy limit:\t%d frames\n",
		        argv[ARG_INFILE], size, limit);
		error++;
		goto exit;
	}
	
	/* check if infile uses 8-bit samples*/ 
	if (props.samptype==PSF_SAMP_8)
	{
		printf("ERROR: sfgain does not support 8-bit format.\n");
		error++;
		goto exit;
	}

	/* display infile properties */
	if(!psf_sndInfileProperties(argv[ARG_INFILE],ifd,&props))
	{
		error++;
		goto exit;
	}	

	/* check if outfile extension is one we know about */	
	outformat = psf_getFormatExt(argv[ARG_OUTFILE]);
	if (outformat == PSF_FMT_UNKNOWN)
	{
		printf("Outfile name \"%s\" has unknown format.\n"
		       "Use any of .wav .aiff .aif .afc .aifc\n",
		        argv[ARG_OUTFILE]);
		error++;
		goto exit;
	}
	props.format = outformat;

	/* create outfile */
	ofd = psf_sndCreate(argv[ARG_OUTFILE], &props, 0, 0, PSF_CREATE_RDWR);
	if (ofd<0)
	{
		printf("ERROR: unable to create outfile \"%s\"\n",argv[ARG_OUTFILE]);
		error++;
		goto exit;
	}
	/* allocate space for sample frames */
	if (limit<nFrames)
		nFrames = (DWORD)limit;
	buffer= (float*)malloc(props.chans*sizeof(float)*nFrames);
	if (buffer==NULL)
	{
		puts("No memory!\n");
		error++;
		goto exit;
	}

	/* and allocate space for PEAK info */
	peaks = (PSF_CHPEAK*)malloc(props.chans*sizeof(PSF_CHPEAK));
	if (peaks==NULL)
	{
		puts("No memory!\n");
		error++;
		goto exit;
	}

	printf("copying...\n");

	int update=0; 
	int loop;
	int i;
	int j;

	/* copy the infile N times */
	for (loop=0; loop<N; loop++)
	{

		/* make sure to set nFrames to the correct value
			 every time you pass through the for loop */ 
		if (limit < atoi(argv[ARG_BUFF]))
			nFrames = (DWORD)limit;
		else
			nFrames = (DWORD)atoi(argv[ARG_BUFF]);

		totalread = 0; /* running count of sample frames */
		if(psf_sndSeek(ifd,0,PSF_SEEK_SET))
		{
			printf("ERROR: cannot reset infile\n");
			error++;
			goto exit;
		}
		/* nFrames loop to do copy, report any errors */
		framesread = psf_sndReadFloatFrames(ifd,buffer,nFrames);
		while (framesread>0&&totalread<limit)
		{
			update++;
			
			/* update copy status after refreshing the buffer every 100 times */
			if (update%100==0)
				printf("%ld samples copied...  %ld%%\r",totalread,100*totalread/size);

			totalread+=framesread;
		
			/* change sample values by amp factor */
			for (j=0; j<nFrames; j++)
				for (i=0; i<props.chans; i++)	
					buffer[props.chans*j+i] *= ampfac; 
			if(psf_sndWriteFloatFrames(ofd,buffer,nFrames)<0)
			{
				printf("Error writing to outfile.\n");
				error++;
				break;
			}
			/* make sure not to copy frames past the limit */ 
			if (nFrames+totalread > limit)
				nFrames = (DWORD)(limit-totalread); 

			framesread = psf_sndReadFloatFrames(ifd,buffer,nFrames);

		}
	}	
	totalread *= N; /* total number of frames copied */ 

	if(framesread<0)
	{
		printf("Error reading infile. Outfile is incomplete.\n");
		error++;
	}
	else
		printf("Done. %ld sample frames copied to %s\n",
						totalread, argv[ARG_OUTFILE]);

	/* report PEAKS to user */
	if (psf_sndReadPeaks(ofd,peaks,NULL)>0)
	{
		long i;
		double peaktime;
		double peakDB;
		printf("Peak information:\n");
		for (i=0; i<props.chans; i++)
		{
			peaktime = (double)peaks[i].pos/props.srate; 
			if (peaks[i].val == 0.0)
				peaks[i].val = 1.0e-4;
			peakDB = log10(peaks[i].val);
			printf("CH %ld:\t%.4f\t(%.4f dB) at %.4f secs\n",
			        i+1, peaks[i].val, peakDB, peaktime);	
		}
	}
	
	/* do all the cleanup */
	exit:
	if (ifd>=0)
		psf_sndClose(ifd);
	if (ofd>=0)
		psf_sndClose(ofd);
	if (buffer)
		free(buffer);
	if (peaks)
		free(peaks);
	psf_finish();
	 	
	return error;
}
