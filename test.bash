#!/bin/bash

INPUTFILE=Sternwarte.png

./pngFrequencyFilter -in=$INPUTFILE -out=output_0.0.png -cutoff 0.0
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.01.png -cutoff 0.01
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.02.png -cutoff 0.02
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.05.png -cutoff 0.05
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.1.png -cutoff 0.1
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.2.png -cutoff 0.2
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.3.png -cutoff 0.3
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.4.png -cutoff 0.4
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.5.png -cutoff 0.5
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.6.png -cutoff 0.6
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.7.png -cutoff 0.7
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.8.png -cutoff 0.8
./pngFrequencyFilter -in=$INPUTFILE -out=output_0.9.png -cutoff 0.9
./pngFrequencyFilter -in=$INPUTFILE -out=output_1.0.png -cutoff 1.0
