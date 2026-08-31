#!/bin/bash

INPUTFILE=Sternwarte.png

./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_000.png -cutoff 0.0
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_001.png -cutoff 0.01
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_002.png -cutoff 0.02
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_005.png -cutoff 0.05
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_010.png -cutoff 0.1
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_020.png -cutoff 0.2
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_030.png -cutoff 0.3
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_040.png -cutoff 0.4
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_050.png -cutoff 0.5
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_060.png -cutoff 0.6
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_070.png -cutoff 0.7
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_080.png -cutoff 0.8
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_090.png -cutoff 0.9
./pngFrequencyFilter -in=$INPUTFILE -out=output_cutoff_100.png -cutoff 1.0

./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_000.png -obstruction 0.0
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_001.png -obstruction 0.01
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_002.png -obstruction 0.02
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_005.png -obstruction 0.05
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_010.png -obstruction 0.1
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_020.png -obstruction 0.2
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_030.png -obstruction 0.3
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_040.png -obstruction 0.4
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_050.png -obstruction 0.5
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_060.png -obstruction 0.6
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_070.png -obstruction 0.7
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_080.png -obstruction 0.8
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_090.png -obstruction 0.9
./pngFrequencyFilter -in=$INPUTFILE -out=output_obstruction_100.png -obstruction 1.0
