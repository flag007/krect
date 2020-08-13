package main

import "fmt"
import "regexp"
import "net/url"
import "bufio"
import "os"



func main() {

	result := []string{}


	sc := bufio.NewScanner(os.Stdin)
	seen := make(map[string]bool)
	for sc.Scan() {
		uri := sc.Text()
		u, _ := url.Parse(uri)
		p1 := "=https://google.com&"
		out := regexp.MustCompile(`=http[^&]*&`).ReplaceAllString(u.RawQuery+"&" , p1);

		if _, ok := seen[u.Hostname()+u.EscapedPath()]; !ok {
			seen[u.Hostname()+u.EscapedPath()] = true
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+out)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"next"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"r"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"url"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"target"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"rurl"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"dest"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"destination"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"redir"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"redirect_uri"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"redirect_url"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"redirect"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"image_url"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"go"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"return"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"returnTo"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"return_to"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"checkout_url"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"continue"+p1)
			result = append(result, u.Scheme+"://"+u.Hostname()+u.EscapedPath()+"?"+"return_path"+p1)
		}

	}

	for _, i := range result {
		fmt.Println(i)
	}

}





