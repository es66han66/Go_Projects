package main

import "sync"

type sampleChannel struct {
	Data *Sample
	Err  error
}

func (u *usecase) GetSample(id int64, someparam string, anotherParam string) ([]*Sample, error) {

	chanSample := make(chan sampleChannel, 3)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		chanSample <- u.getDataFromGoogle(id, anotherParam) // just example of function
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		chanSample <- u.getDataFromFacebook(id, anotherParam)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		chanSample <- u.getDataFromTwitter(id, anotherParam)
	}()

	wg.Wait()
	close(chanSample)

	result := make([]*Sample, 0)

	for sampleItem := range chanSample {
		if sampleItem.Error != nil {
			logrus.Error(sampleItem.Err)
		}
		if sampleItem.Data == nil {
			continue
		}
		result = append(result, sampleItem.Data)
	}

	return result
}

func main() {

}
