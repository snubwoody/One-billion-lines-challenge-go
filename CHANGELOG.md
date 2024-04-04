### V 0.1.0
I took the standard simple approach for this: loopping throught the file using a buffer scanner. 
I stored had a station struct and an array of these structs, I then compared the values for each processed line. This gave me a runtime of under 5 minutes(which is quite impressive).