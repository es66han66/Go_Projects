package main

func main() {
	func (u *usecase) GetSample(c context.Context, id int64, someparam string, anotherParam string) ([]*Sample, error) {
	
  if c== nil {
    c= context.Background()
  }
  
  ctx, cancel := context.WithTimeout(c, time.Second * 2)
  defer cancel()
  
  chanSample := make(chan sampleChannel, 3)
  defer close(chanSample)
 	go func() {
       		chanSample <- u.getDataFromGoogle(ctx, id, anotherParam) // just example of function
	}()

 
	go func() {
    		chanSample <- u.getDataFromFacebook(ctx, id, anotherParam)
	}()

   
	go func() {
   		chanSample <- u.getDataFromTwitter(ctx, id,anotherParam)
	}()
 

	result := make([]*feed.Feed, 0)
	for loop := 0; loop < 3; loop++ {
	  select {
		case sampleItem := <-chanSample:
			if sampleItem.Err != nil {
				continue
			}
			if feedItem.Data == nil {
				continue
			}
			 result = append(result,sampleItem.Data)
		  // ============================================================
		  // CATCH IF THE CONTEXT ALREADY EXCEEDED THE TIMEOUT
		  // FOR AVOID INCONSISTENT DATA, WE JUST SENT EMPTY ARRAY TO 
		  // USER AND ERROR MESSAGE
		  // ============================================================
      		case <-ctx.Done(): // To get the notify signal that the context already exceeded the timeout
		      err := fmt.Errorf("Timeout to get sample id: %d. ", id)
		      result = make([]*sample, 0)
		      return result, err
		}
	}

	return result, nil;
}
}