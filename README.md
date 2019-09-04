# Qntfy-Technical-Assignment

**Work/Time Log:**

9/1:

12:00 PM: Received email

4:30 PM: Was able to review contents of email; began looking into Golang. Research continued throughout this project

6:00 PM: Started writing code - focused on the parsing of individual lines, then opening files, then writing to files

7:00 PM: Researched better methods for storing the data; limited ingest of data to 1MB per file being read

9:00 PM: Completed initial draft of file-reading and parsing code, stubbed code for writing output file, and called it a night

9/2:

10:00 AM: Began researching Docker while working on output file writing code. While researching, found a library that handles the statistical analysis (median, standard deviation) - no need to reinvent the wheel

11:00 AM: Uploaded initial prototype to GitHub. Installed Docker, began work on Dockerfile

12:00 PM: Continued working with Docker while also looking for performance improvements in main body of code

1:00 PM: Lunch

2:00 PM: First successful build from Dockerfile.

2:30 PM: First successful/not crashing build using docker-compose up

3:00 PM: Completed basic benchmarking of main file and of component that is most likely to take the longest to run.

3:15 PM: Created this timeline.

9/3:

10:00 AM: Dug back into docker-compose up issues. While I can run the app using go run \*.go, I cannot get it to run via docker-compose up. "go run: cannot run non-main package".

9/4:

1:00 PM: Gave the Dockerfile another go - was able to successfully prove that the code is running by viewing the output in the terminal log, but not able to locate the file in the Container. Additionally, looked into using sync.Map instead of maps to reduce wait times for locked maps.

**Benchmark Results** - 

I am not sure if these results are valid, or if these benchmarks are properly implemented. I do not have a Go-knowledgeable friend to check my work and my research has been fairly inconclusive. For these benchmarks (and for my own testing), I used an old version of the code itself as the test files - after all, code is still multiline text!

Main, running on sample files: 0.005s

readFiles, running on sample files: 0.0005s


**Note**

I was unsure of the best way to bring in files for testing, and as I'm new to Golang, I went with a simple approach that I could validate on my own. I do not know Docker well enough to determine if there is a better, simpler, or perhaps an obvious way I should have gone about it, but I will be looking into it in the coming days.

The keywords.txt file contains the sample keywords I used for my testing. You can add the path to a different keyword file as the first argument when using the terminal to execute the program. The /files/ directory contains the sample files I used for my testing, and can be overridden using the second argument.
