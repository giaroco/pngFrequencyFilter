#!/bin/bash

INPUTFILE=Sternwarte.png

./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.0.png -cutoff 0.0
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.01.png -cutoff 0.01
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.02.png -cutoff 0.02
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.05.png -cutoff 0.05
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.1.png -cutoff 0.1
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.2.png -cutoff 0.2
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.3.png -cutoff 0.3
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.4.png -cutoff 0.4
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.5.png -cutoff 0.5
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.6.png -cutoff 0.6
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.7.png -cutoff 0.7
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.8.png -cutoff 0.8
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_0.9.png -cutoff 0.9
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_1.0.png -cutoff 1.0

./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.0.png -obstruction 0.0
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.01.png -obstruction 0.01
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.02.png -obstruction 0.02
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.05.png -obstruction 0.05
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.1.png -obstruction 0.1
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.2.png -obstruction 0.2
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.3.png -obstruction 0.3
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.4.png -obstruction 0.4
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.5.png -obstruction 0.5
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.6.png -obstruction 0.6
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.7.png -obstruction 0.7
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.8.png -obstruction 0.8
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_0.9.png -obstruction 0.9
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_1.0.png -obstruction 1.0
