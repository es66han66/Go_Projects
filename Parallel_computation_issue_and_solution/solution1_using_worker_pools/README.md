A worker pool allows us to manage the number of go routines we have, keeping memory print low.  
  
We limited the number of worker pools to 100 and for each one, create a go routine