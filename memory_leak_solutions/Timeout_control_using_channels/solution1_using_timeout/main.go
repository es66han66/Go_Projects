package main

func main() {
	func (u *usecase) GetSample(id int64, someparam string, anotherParam string) ([]*Sample, error) {
	chanSample := make(chan sampleChannel, 3)
	defer close(chanSample)
 	go func() {
       		chanSample <- u.getDataFromGoogle(id, anotherParam) // just example of function
	}()

 
	go func() {
    		chanSample <- u.getDataFromFacebook(id, anotherParam)
	}()

   
	go func() {
   		chanSample <- u.getDataFromTwitter(id,anotherParam)
	}()
 

	result := make([]*feed.Feed, 0)
	timeout := time.After(time.Second * 2)
		for loop := 0; loop < 3; loop++ {
			select {
			case sampleItem := <-chanSample:
				if sampleItem.Err != nil {
					logrus.Error(sampleItem.Err)
					continue
				}
				if feedItem.Data == nil {
					continue
				}
				 result = append(result,sampleItem.Data)
			case <-timeout:
			      err := fmt.Errorf("Timeout to get sample id: %d. ", id)
			      result = make([]*sample, 0)
			      return result, err
			}
		}

	return result, nil;
}
}