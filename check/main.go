package main

import (
  "fmt"
  "os"
  "regexp"
  "github.com/gutengo/shell"
  "github.com/PuerkitoBio/goquery"
  "github.com/blang/semver"
)
var pd = fmt.Println

func main() {
  if len(os.Args) < 4 {
    shell.ErrorExit("USAGE: gutaur-check <version> <url> <regexp>")
  }

  Check(os.Args[1], os.Args[2], os.Args[3] )
}

func Check(curVersionRaw, url, pattern string) {
  curVersion, err := semver.Parse(curVersionRaw)
  if err != nil {
    shell.ErrorExit(err)
  }

  pat, err := regexp.Compile(pattern)
  if err != nil {
    shell.ErrorExit(err)
  }

  doc, err := goquery.NewDocument(url)
  if err != nil {
    shell.ErrorExit(err)
  }

  versionsRaw := []string{}
  doc.Find("a").Each(func(i int, s *goquery.Selection) {
    href := s.AttrOr("href", "")
    matches := pat.FindStringSubmatch(href)
    if len(matches) == 0 {
      return
    }
    shell.Say2("gutaur-check: %s %s\n", href, matches[1])
    versionsRaw = append(versionsRaw, matches[1])
  })

  versions := []semver.Version{}
  for _, raw := range versionsRaw {
    v, _ := semver.Parse(raw)
    versions = append(versions, v)
  }
  semver.Sort(versions)
  latestVersion := versions[len(versions)-1]

  if curVersion.LT(latestVersion) {
    shell.Say("%s", latestVersion)
  } else {
    shell.Say("%s", "")
  }
}
