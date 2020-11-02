package google

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lu-moreira/shouldgo/context/gocontextsample/userip"
)

type Results []Result
type Result struct {
	Title string
	URL   string
}

func Search(ctx context.Context, query string) (Results, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.google.com/search", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("q", query)

	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}

	req.URL.RawQuery = q.Encode()

	var results Results

	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		var data struct {
			ResponseData struct {
				Results []struct {
					TitleNoFormatting string
					URL               string
				}
			}
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}

		for _, res := range data.ResponseData.Results {
			results = append(results, Result{
				Title: res.TitleNoFormatting,
				URL:   res.URL,
			})
		}

		return nil
	})

	return results, err
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() {
		c <- f(http.DefaultClient.Do(req))
	}()

	select {
	case <-ctx.Done():
		<-c // wait for f to return
		return ctx.Err()
	case err := <-c:
		return err
	}
}
