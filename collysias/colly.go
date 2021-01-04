package collysias

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func NewCollector(){
    
	c := colly.NewCollector(
		 colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement){
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
		
		val :=  fmt.Sprintf(`
		{%q:"%s"},
		`,e.Text,link)
		f, err := os.OpenFile("testl1.json",os.O_APPEND|os.O_WRONLY|os.O_CREATE , 0644)
		if err != nil{
			panic(err)
		}
		defer f.Close()

		if _, err = f.WriteString(val); err != nil {
			panic(err)
		}
	})

	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting", r.URL.String())
	})

	a := c.Visit("https://hackerspaces.org/")
	fmt.Println(a)
}	