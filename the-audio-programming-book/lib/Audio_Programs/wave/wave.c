#include <stdio.h>
#include <stdlib.h>
#include <wave.h>
#include <math.h>


/** static oscil function declarations **/


/* oscil creation function */
static OSCIL* oscil_create();

/* oscil initialization function */
static void oscil_init(OSCIL* p_osc, unsigned long srate);


/** wave oscil function definitions **/


/* oscil creation, returns NULL if no oscil was created */
static OSCIL* oscil_create()
{
	OSCIL* p_osc;
	p_osc = (OSCIL*)malloc(sizeof(OSCIL));
	return p_osc;
}

/* oscil initialization */
static void oscil_init(OSCIL* p_osc, unsigned long srate)
{
	p_osc->twopiovrsr = TWOPI / (double) srate;
	p_osc->curfreq = 0.0;
	p_osc->curphase = 0.0;
	p_osc->incr = 0.0;
}

/* a combined OSCIL creation and initialization function */ 
OSCIL* new_oscil(unsigned long srate)
{
	OSCIL* p_osc = oscil_create();	

	if (p_osc==NULL)
		return NULL;

	oscil_init(p_osc, srate);

	return p_osc;
}

/* a combined OSCIL creation and initialization function 
   the phase argument is a fraction between 0 and 1 which 
   sets initial phase of the oscillator */
OSCIL* new_oscilp(unsigned long srate, double phase)
{
	OSCIL* p_osc = oscil_create();	
	
	if (p_osc==NULL)
		return NULL;

	oscil_init(p_osc, srate);

	/* make sure the inputted phase doesn't exceed the range 
	   by taking the fractional part of the phase value */ 
	if (phase > 1.0)
		phase = phase - (int)phase;
	if (phase < 0.0)
		phase = (phase + (int)phase) * -1.0;
	/* phase offset is from 0 to 2*PI */
	p_osc->curphase = TWOPI*phase;

	return p_osc;
} 

/* tick function for a sine waveform */
double sinetick(OSCIL* p_osc, double freq)
{
	double val;
	
	val = sin(p_osc->curphase);
	UPDATE_FREQ;
	p_osc->curphase += p_osc->incr;
	if (p_osc->curphase >= TWOPI)
		p_osc->curphase -= TWOPI;
	if (p_osc->curphase < 0.0)
		p_osc->curphase += TWOPI;
	
	return val;
}

/* tick function for a square waveform */
double sqtick(OSCIL* p_osc, double freq)
{
	double val;
	 	
	UPDATE_FREQ; 
	val = (p_osc->curphase <= M_PI)?1.0:-1.0;	
	p_osc->curphase += p_osc->incr;
	if (p_osc->curphase >= TWOPI)
		p_osc->curphase -= TWOPI;
	if (p_osc->curphase < 0.0)
		p_osc->curphase += TWOPI;

	return val;
}

/* tick function for a square waveform
   with dynamic pulse wave modulation */
double pwmtick(OSCIL* p_osc, double freq, double pwmod)
{
	double val;

	if (pwmod < 1.0) /* if pwmod < 1%, set minimum frequency by a factor of 0.02 */
		freq *= 0.02;	
	else if (pwmod > 99.0) /* if pwmod > 99%, set maximum frequency by a factor of 1.98 */
		freq *= 1.98;
	else
		freq *= pwmod/50.0; /* normal square wave is 50% */

	UPDATE_FREQ;
	val = (p_osc->curphase <= M_PI)?1.0:-1.0;	
	p_osc->curphase += p_osc->incr;

	if (p_osc->curphase >= TWOPI)
		p_osc->curphase -= TWOPI;
	if (p_osc->curphase < 0.0)
		p_osc->curphase += TWOPI;

	return val;
}

/* tick function for a downward sawtooth waveform */
double sawdtick(OSCIL* p_osc, double freq)
{
	double val;

	UPDATE_FREQ;	
	val = 1.0 - 2.0 * (p_osc->curphase * (1.0 / TWOPI) ); 
	p_osc->curphase += p_osc->incr;
	if (p_osc->curphase >= TWOPI)
		p_osc->curphase -= TWOPI;
	if (p_osc->curphase < 0.0)
		p_osc->curphase += TWOPI;

	return val;
}

/* tick function for an upward sawtooth waveform */
double sawutick(OSCIL* p_osc, double freq)
{
	double val;

	UPDATE_FREQ;	
	val = (2.0 * (p_osc->curphase * (1.0 / TWOPI) )) - 1.0;	
	p_osc->curphase += p_osc->incr;
	if (p_osc->curphase >= TWOPI)
		p_osc->curphase -= TWOPI;
	if (p_osc->curphase < 0.0)
		p_osc->curphase += TWOPI;

	return val; 
}

/* tick function for a triangle waveform */
double tritick(OSCIL* p_osc, double freq)
{
	double val;

	UPDATE_FREQ;	
	/* rectified sawtooth */
	val = (2.0 * (p_osc->curphase * (1.0 / TWOPI) )) - 1.0;
	if (val < 0.0)
		val = -val;
	val = 2.0 * (val - 0.5);
	p_osc->curphase += p_osc->incr;
	if (p_osc->curphase >= TWOPI)
		p_osc->curphase -= TWOPI;
	if (p_osc->curphase < 0.0)
		p_osc->curphase += TWOPI;

	return val;
}
