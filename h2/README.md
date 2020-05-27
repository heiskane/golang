# CryptV2

This is my homework for the intensive Go programming course taught by Tero Karvinen http://terokarvinen.com/2020/go-programming-course-2020-w22/

I learned how to deal with negative numbers and modulus over here: https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go

I was granted permission by the teacher to improve my previous script in following ways for this homework: Separate code into functions, Implement reading Files and Handle Errors. What it does now is it takes a file then reads its contents and encrypts them with a word.
Example Usage: 
```
./cryptv2 -f path/to/file -k secretKey > encrypted.txt
./cryptv2 -f encypted.txt -k secretKey > decrypted.txt
```

To build this run ```go build cryptv2.go```

![CryptV2 Script](/img/img2.png)
