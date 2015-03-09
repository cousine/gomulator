# Gomulator

Gomulator is simply a gameboy emulator written in Go.

Its probably not that useful nor is it as good as other well known emulators, but I wanted to put my passion for Go in something challenging and interesting.

All the code here is based on the amazing series of articles on the subject by Imran Nazar which you can find [here](http://imrannazar.com/GameBoy-Emulation-in-JavaScript%3a-The-CPU); all I did was port his code to Go improving applicable areas where Go can do a better job than JS could, either in readability or performance.

## WiP

This is still a work in progress, at the moment I've only finished the MMU and part of the Gameboy's Z80 CPU (no ops nor the dispatch loop), however over the next few days I'll be working on finishing the CPU and then move on to the GPU emulation.

## Points to keep in mind

* I dont want to use SDL for gfx and audio, so I am researching the available options to access the HW within a terminal window (contrast to having SDL summon a new window to render the GFX)
* I'm working on Mac OS however, I doubt there is any barrier for this to work on any *NIX system. Not sure about windows though so you can keep me updated if you try it on Windows.
* Once I'm done with the core of the emulator, I'll update the README to include installation and usage.
* I'm trying to keep the code documented as much as possible for my self and others since this is purely for learning and practicing purposes.

# Wacky ideas

* Maybe in the future I'd like to replicate [Tom7's "Computer program that learns to play classic NES games"](https://www.youtube.com/watch?v=xOCurBYI_gY) this time however on a gameboy :P 
