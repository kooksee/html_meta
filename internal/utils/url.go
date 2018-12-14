package utils

import (
	"encoding/hex"
	url2 "net/url"
	"strings"
)

func UrlImageCheck(u1, url string) string {
	us, _ := url2.Parse(u1)

	url = strings.TrimSpace(url)
	url = strings.Replace(url, " ", "", -1)

	if strings.HasPrefix(url, "data") {
		return ""
	}

	if strings.HasPrefix(url, "//") {
		return us.Scheme + ":" + url
	}

	if strings.HasPrefix(url, "/") {
		return us.Scheme + "://" + us.Host + url
	}

	if strings.HasPrefix(url, "..") {
		u2 := strings.Split(u1, "/")
		u2 = u2[:len(u2)-1]
		return strings.Join(append(u2, url), "/")
	}

	if strings.HasPrefix(url, "./") {
		u2 := strings.Split(u1, "/")
		u2 = u2[:len(u2)-1]
		return strings.Join(append(u2, url), "/")
	}

	if (!strings.HasPrefix(url, "http")) &&
		(!strings.HasPrefix(url, "//")) &&
		(!strings.HasPrefix(url, "/")) &&
		(!strings.HasPrefix(url, ".")) &&
		(!strings.HasPrefix(url, "..")) {
		u2 := strings.Split(u1, "/")
		u2 = u2[:len(u2)-1]
		return strings.Join(append(u2, url), "/")
	}

	return url
}

func UrlCheck(u1, url string) string {
	us, _ := url2.Parse(u1)
	u := us.Scheme + "://" + us.Host

	url = strings.TrimSpace(url)
	url = strings.Replace(url, " ", "", -1)

	if url == "" {
		return ""
	}

	if strings.HasPrefix(url, "data") {
		return ""
	}

	if strings.HasPrefix(url, "//") {
		url = us.Scheme + "://" + strings.Trim(url, "//")
	}

	if strings.HasPrefix(url, "..") {
		u2 := strings.Split(u1, "/")
		u1 = strings.Join(u2[:len(u2)-2], "/")
		url = u1 + strings.Trim(url, "..")
	}

	if strings.HasPrefix(url, ".") {
		u2 := strings.Split(u1, "/")
		u1 = strings.Join(u2[:len(u2)-1], "/")
		url = u1 + strings.Trim(url, ".")
	}

	if !strings.HasPrefix(url, "http") {
		if strings.HasPrefix(url, "/") {
			url = u + url
		} else {
			u2 := strings.Split(u1, "/")
			u1 = strings.Join(u2[:len(u2)-1], "/")
			url = u1 + url
		}
	}
	return url

}
func UrlHash(u1, url string) string {
	url = UrlCheck(u1, url)
	dt := strings.Split(url, "?")
	dt1 := strings.Split(dt[0], ".")
	return hex.EncodeToString(Sha256([]byte(url))) + "." + dt1[len(dt1)-1]
}

func ImageUrlHash(url string) string {
	url = strings.TrimSpace(url)
	url = strings.Replace(url, " ", "", -1)
	dt := strings.Split(url, ".")
	return hex.EncodeToString(Sha256([]byte(url))) + "." + dt[len(dt)-1]
}

func HtmlUrlHash(url string) string {
	url = strings.TrimSpace(url)
	url = strings.Replace(url, " ", "", -1)
	dt := strings.Split(url, ".")
	return hex.EncodeToString(Sha256([]byte(url))) + "." + dt[len(dt)-1]
}
