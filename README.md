# wc-tool
a toy wc - word count tool written in go
```
msm002@test wc % go run wc.go             
     
  Help for wc command flags. 
  wc <file-name> or wc -<flag> <file-name>

  -c      The number of bytes in each input file is written to the standard output. 
  -l      The number of lines in each input file is written to the standard output. 
  -m      The number of characters in each input file is written to the standard output. 
  -w      The number of words in each input file is written to the standard output.
  
msm002wc % go run wc.go --version   
0.0.1
msm002@test wc % 
msm002@test wc % go run wc.go testfile.txt
7 91  574 testfile.txt
msm002@test wc % go run wc.go -c testfile.txt 
byte count: 574
msm002@test wc % go run wc.go -l testfile.txt
line count: 7
msm002@test wc % go run wc.go -m testfile.txt
character count: 574
msm002@test wc % go run wc.go -w testfile.txt
word count: 91
msm002@test wc % go run wc.go -ww testfile.txt
     
  Help for wc command flags. 
  wc <file-name> or wc -<flag> <file-name>

  -c      The number of bytes in each input file is written to the standard output. 
  -l      The number of lines in each input file is written to the standard output. 
  -m      The number of characters in each input file is written to the standard output. 
  -w      The number of words in each input file is written to the standard output.

msm002@test wc
```
