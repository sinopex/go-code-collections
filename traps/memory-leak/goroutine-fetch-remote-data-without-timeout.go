// COPY FROM: http://www.gaoxuan1989.com/2019/04/13/avoiding-memory-leak-in-golang-api/
package memory_leak

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type UserInfo struct {
	Data interface{}
	Err  error
}

type BadUserInfoFetcher struct{}

func (u *BadUserInfoFetcher) Get(id int64) []UserInfo {
	userChan := make(chan UserInfo, 3)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		userChan <- u.getDataFromGoogle(id) // just example of function
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		userChan <- u.getDataFromFacebook(id)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		userChan <- u.getDataFromTwitter(id)
	}()

	wg.Wait()
	close(userChan)

	result := make([]UserInfo, 0)

	for userInfo := range userChan {
		result = append(result, userInfo)
	}

	return result
}

func (u *BadUserInfoFetcher) getDataFromGoogle(id int64) UserInfo {
	time.Sleep(time.Millisecond * 10)
	return UserInfo{Data: id, Err: nil}
}

func (u *BadUserInfoFetcher) getDataFromTwitter(id int64) UserInfo {
	time.Sleep(time.Millisecond * 20)
	return UserInfo{Data: id, Err: nil}
}

func (u *BadUserInfoFetcher) getDataFromFacebook(id int64) UserInfo {
	return UserInfo{Data: id, Err: nil}
}

type BetterUserInfoFetcher struct{}

func (u *BetterUserInfoFetcher) BetterGetUserInfo(ctx context.Context, id int64) ([]UserInfo, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	userChan := make(chan UserInfo, 3)
	defer close(userChan)
	go func() {
		userChan <- u.getDataFromGoogle(ctx, id) // just example of function
	}()

	go func() {
		userChan <- u.getDataFromFacebook(ctx, id)
	}()

	go func() {
		userChan <- u.getDataFromTwitter(ctx, id)
	}()

	result := make([]UserInfo, 0)
	for loop := 0; loop < 3; loop++ {
		select {
		case userInfo := <-userChan:
			result = append(result, userInfo)
			// ============================================================
			// CATCH IF THE CONTEXT ALREADY EXCEEDED THE TIMEOUT
			// FOR AVOID INCONSISTENT DATA, WE JUST SENT EMPTY ARRAY TO
			// USER AND ERROR MESSAGE
			// ============================================================
		case <-ctx.Done(): // To get the notify signal that the context already exceeded the timeout
			err := fmt.Errorf("Timeout to get sample id: %d. ", id)
			result = make([]UserInfo, 0)
			return result, err
		}
	}

	return result, nil
}
func (u *BetterUserInfoFetcher) getDataFromGoogle(ctx context.Context, id int64) UserInfo {
	req, err := http.NewRequest("GET", "https://facebook.com", nil)
	if err != nil {
		return UserInfo{Err: err}
	}
	// ============================================================
	// THEN WE PASS THE CONTEXT TO OUR REQUEST.
	// THIS FEATURE CAN BE USED FROM GO 1.7
	// ============================================================
	if ctx != nil {
		req = req.WithContext(ctx) // NOTICE THIS. WE ARE USING CONTEXT TO OUR HTTP CALL REQUEST
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return UserInfo{Err: err}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{Err: err}
	}
	return UserInfo{
		Data: body,
		Err:  nil,
	}
}

func (u *BetterUserInfoFetcher) getDataFromTwitter(ctx context.Context, id int64) UserInfo {
	time.Sleep(time.Millisecond * 10)
	return UserInfo{Data: id, Err: nil}
}

func (u *BetterUserInfoFetcher) getDataFromFacebook(ctx context.Context, id int64) UserInfo {
	return UserInfo{Data: id, Err: nil}
}
