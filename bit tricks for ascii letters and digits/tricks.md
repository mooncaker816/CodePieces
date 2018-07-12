```bash
     0123456789ABCDEF0123456789ABCDEF
    +--------------------------------
0x20| !"#$%&'()*+,-./0123456789:;<=>?
0x40|@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_
0x60|`abcdefghijklmnopqrstuvwxyz{|}~
```

1. ASCII 码大写字母转小写字母  
c | 0x20 (0010 0000)
2. ASCII 码小写字母转大写字母  
c & 0xDF (1101 1111)
3. ASCII 码大小字母翻转（大转小，小转大）  
c XOR 0x20 (0010 0000)
4. 由数字（0-9）字符转为该字符代表的数值（不是该字符本身的数值）  
c & 0x0F (0000 1111)

低四位的数值正好是该数字字符代表的数值

|LETTER	| DECIMAL VALUE	| BINARY VALUE|
|-------|:-------------:|:-----------:|
|0	    |48	            |0011 0000    |
|1	    |49	            |0011 0001    |
|2	    |50	            |0011 0010    |
|3	    |51	            |0011 0011    |
|4	    |52	            |0011 0100    |
|5	    |53	            |0011 0101    |
|6	    |54	            |0011 0110    |
|7	    |55	            |0011 0111    |
|8	    |56	            |0011 1000    |
|9	    |57	            |0011 1001    |

参考：  
https://blog.cloudflare.com/the-oldest-trick-in-the-ascii-book/  
https://www.guidodiepen.nl/2017/03/clever-bit-tricks-with-ascii/