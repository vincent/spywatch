package scrapper

import (
	neturl "net/url"
	"os/exec"
	"slices"
	"strings"

	metaparser "github.com/ammit/go-metaparser"
)

type WebsiteLinksResult struct {
	Title          string   `json:"title"`
	Logo           string   `json:"logo"`
	OwnLinks       []string `json:"ownLinks"`
	SocialNetworks []string `json:"socialNetworks"`
}

func GetLinksFromWebsite(url string) (WebsiteLinksResult, error) {

	parsedUrl, _ := neturl.Parse(url)

	var base = parsedUrl.Scheme + "://" + parsedUrl.Host
	var domain = strings.Split(strings.Split(url, "://")[1], "/")[0]
	var logo = "https://icons.duckduckgo.com/ip2/" + domain + ".ico"
	var title = ""

	// get metadata
	p := metaparser.New()
	b, err := p.FetchHTML(url)
	if err != nil {
		return WebsiteLinksResult{}, err
	}
	defer b.Close()

	err = p.ParseHTML(b)
	if err == nil {
		title = p.Title
		if len(p.Favicons) > 0 {
			if strings.HasPrefix(p.Favicons[0].URL, "//") {
				logo = parsedUrl.Scheme + ":" + p.Favicons[0].URL
			} else if strings.HasPrefix(p.Favicons[0].URL, "http") {
				logo = p.Favicons[0].URL
			} else {
				logo = base + "/" + p.Favicons[0].URL
			}
		}
	}

	// execute the lynx command to get all links from the website
	cmd := exec.Command("lynx", "-dump", "-listonly", "-nonumbers", url)
	output, err := cmd.Output()
	if err != nil {
		return WebsiteLinksResult{}, err
	}

	lines := append(strings.Split(string(output), "\n"), url)

	var ownLinks []string
	var ownDic = make(map[string]bool)
	var socialLinks []string
	var socDic = make(map[string]bool)
	var prefix = strings.Trim(url, "/")
	var strict = false

	for _, line := range lines {
		link := strings.TrimSpace(line)
		if !strings.HasPrefix(link, "http") || strings.Contains(link, "#") || strings.Contains(link, "rss") || strings.Contains(link, ".atom") {
			continue
		}
		if logo == "" && strings.Contains(link, "favicon") {
			logo = link
		}
		if isSocialNetworkLink(link) {
			socialLinks = addLink(socDic, socialLinks, link)
		} else if !strict || strings.HasPrefix(link, prefix) {
			ownLinks = addLink(ownDic, ownLinks, link)
		}
	}

	slices.Sort(ownLinks)
	slices.Sort(socialLinks)

	return WebsiteLinksResult{
		Title:          title,
		Logo:           logo,
		OwnLinks:       ownLinks,
		SocialNetworks: socialLinks,
	}, nil
}

func isSocialNetworkLink(link string) bool {
	return strings.Contains(link, "facebook.com") ||
		strings.Contains(link, "tripadvisor.com") ||
		strings.Contains(link, "instagram.com") ||
		strings.Contains(link, "capterra.com/") ||
		strings.Contains(link, "linkedin.com") ||
		strings.Contains(link, "twitter.com") ||
		strings.Contains(link, "youtube.com") ||
		strings.Contains(link, "//g.page/") ||
		strings.Contains(link, "//x.com/")
}

func addLink(m map[string]bool, a []string, s string) []string {
	if m[s] {
		return a
	}
	m[s] = true
	return append(a, s)
}
