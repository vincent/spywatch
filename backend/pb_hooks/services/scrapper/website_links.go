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

func GetLinksFromWebsite(target string) (WebsiteLinksResult, error) {
	result := WebsiteLinksResult{}
	url, _ := neturl.Parse(target)

	result.Title = url.Hostname()
	result.Logo = defaultLogo(url)

	title, logo, err := fetchAndParseMetadata(url)
	if err != nil {
		return result, err
	}
	if title != "" {
		result.Title = title
	}
	if logo != "" {
		result.Logo = logo
	}

	links, err := fetchWebsiteLinks(url.String())
	if err != nil {
		return result, nil
	}

	ownLinks, socialLinks, logo := categorizeLinks(links, url.String())
	if logo != "" {
		result.Logo = logo
	}
	result.OwnLinks = ownLinks
	result.SocialNetworks = socialLinks

	return result, nil
}

// defaultLogo return a default DDG logo for the given URL domain
func defaultLogo(url *neturl.URL) string {
	domain := url.Hostname()
	return "https://icons.duckduckgo.com/ip2/" + domain + ".ico"
}

// fetchAndParseMetadata fetches and parses metadata, returning title and logo
// TODO: extract html fetch to ensure doing it once
func fetchAndParseMetadata(url *neturl.URL) (string, string, error) {
	p := metaparser.New()
	b, err := p.FetchHTML(url.String())
	if err != nil {
		return "", "", err
	}
	defer b.Close()

	logo := ""
	title := ""
	base := url.Scheme + "://" + url.Host

	err = p.ParseHTML(b)
	if err == nil {
		title = p.Title
		if len(p.Favicons) > 0 {
			if strings.HasPrefix(p.Favicons[0].URL, "//") {
				logo = url.Scheme + ":" + p.Favicons[0].URL
			} else if strings.HasPrefix(p.Favicons[0].URL, "http") {
				logo = p.Favicons[0].URL
			} else {
				logo = base + "/" + p.Favicons[0].URL
			}
		}
	}
	return title, logo, nil
}

// fetchWebsiteLinks executes the lynx command and returns all links from the website
func fetchWebsiteLinks(url string) ([]string, error) {
	cmd := exec.Command("lynx", "-dump", "-listonly", "-nonumbers", url)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := append(strings.Split(string(output), "\n"), url)
	return lines, nil
}

// categorizeLinks filters and categorizes links into ownLinks and socialLinks, and may update logo
func categorizeLinks(lines []string, url string) (ownLinks, socialLinks []string, logo string) {
	ownDic := make(map[string]bool)
	socDic := make(map[string]bool)
	prefix := strings.Trim(url, "/")
	strict := false
	logo = ""

	for _, line := range lines {
		link := strings.TrimRight(line, "/#")
		if !strings.HasPrefix(link, "http") || strings.Contains(link, "#") || strings.Contains(link, "rss") || strings.Contains(link, ".atom") {
			continue
		}
		if logo == "" && strings.Contains(link, "favicon") {
			logo = link
			continue
		}
		if isSocialNetworkLink(link) {
			socialLinks = addLink(socDic, socialLinks, link)
		} else if !strict || strings.HasPrefix(link, prefix) {
			ownLinks = addLink(ownDic, ownLinks, link)
		}
	}

	slices.Sort(ownLinks)
	slices.Sort(socialLinks)

	return
}

// isSocialNetworkLink returns true if the given link seems a social network one
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
