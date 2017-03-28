package main

import (
	"github.com/davecgh/go-spew/spew"
	a "github.com/ranjib/adafruitio"
	"log"
)

func main() {
	c := a.NewClient("c5dd91426920e46324cc452cc1a7861f5253402f")
	/*
				u, err := c.GetUser()
				f := a.Feed{
					Name: "test-feed",
				}
				err := c.CreateFeed("ranjib", f)
				activities, err := c.ListActivities("ranjib", a.ListActivitiesOptions{})
				res, err := c.ListFeeds("ranjib")
				f, err := c.GetFeed("ranjib", "test-feed")
				f.License = "Apache 2"
				f.UpdatedAt = ""
				f.CreatedAt = ""
				f.ID = 0
				f.Visibility = "public"
				o := a.ListDataOptions{}
				res, err := c.ListData("ranjib", "test-feed", o)
				d := a.Data{
					Value:     rand.Int() + 200,
					Lattitude: 1,
					Longitude: 2,
				}
				c.SubmitData("ranjib", "test-feed", d)
		  	groups, err := c.ListGroups("ranjib")
				g := a.Group{Name: "awesomeGroup2"}
				c.CreateGroup("ranjib", g)
	*/
	g, err := c.GetGroup("ranjib", "awesomegroup2")
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(g)
}
