# Qntfy-Technical-Assignment
At the present, this runs via:
cd src/main/
go run *.go ./keywords.txt ./files/

The two args are optional.


**Work/Time Log:**

9/1:

12:00 PM: Received email

4:30 PM: Was able to review contents of email; began looking into Golang. Research continued throughout this project

6:00 PM: Started writing code - focused on the parsing of individual lines, then opening files, then writing to files

7:00 PM: Researched better methods for storing the data; limited ingest of data to 1MB per file being read

9:00 PM: Completed initial draft of file-reading and parsing code, stubbed code for writing output file, and called it a night

9/2:

9:00 AM: Began researching Docker while working on output file writing code

10:00 AM: Found a library that handles the statistical analysis - no need to reinvent the wheel

11:00 AM: Uploaded initial prototype to GitHub. Installed Docker, began work on Dockerfile

12:00 PM: Continued working with Docker while also looking for performance improvements in main body of code

1:00 PM: Lunch

2:00 PM: First successful build from Dockerfile.

2:30 PM: First successful/not crashing build using docker-compose up

3:00 PM: Completed basic benchmarking of main file and of component that is most likely to take the longest to run.

3:15 PM: Created this timeline.


**Benchmark Results** - 

I am not sure if these results are valid, or if these benchmarks are properly implemented. I do not have a Go-knowledgeable friend to check my work and my research has been fairly inconclusive. For these benchmarks (and for my own testing), I used an old version of the code itself as the test files - after all, code is still multiline text!

Main, running on sample files: 0.005s

readFiles, running on sample files: 0.0005s
