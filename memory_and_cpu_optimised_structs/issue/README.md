But how come 88 bytes, 16 +16 + 1 + 16 + 1+ 16 + 4 = 70 bytes, where is this additional 18 bytes coming from ?  
  
When it comes to memory allocation for structs, they are always allocated contiguous, byte-aligned blocks of memory, and fields are allocated and stored in the order that they are defined. The concept of byte-alignment in this context means that the contiguous blocks of memory are aligned at offsets equal to the platforms word size.  
  
