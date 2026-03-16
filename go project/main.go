package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Company struct {
	Name       string   `json:"name"`
	CareerLink string   `json:"careerLink"`
	Jobs       []string `json:"jobs"`
}

var companies = []Company{
	{
		Name:       "Google",
		CareerLink: "https://careers.google.com",
		Jobs:       []string{"Software Engineer", "C++ Engineer", "Data Scientist", "Product Manager"},
	},
	{
		Name:       "Meta",
		CareerLink: "https://www.metacareers.com",
		Jobs:       []string{"Frontend Engineer", "Backend Engineer", "C++ Developer", "Machine Learning"},
	},
	{
		Name:       "Antmicro",
		CareerLink: "https://antmicro.com/careers",
		Jobs:       []string{"Embedded Software Engineer", "C++ Systems Engineer", "FPGA Engineer"},
	},
	{
		Name:       "AdaCore",
		CareerLink: "https://www.adacore.com/careers",
		Jobs:       []string{"Compiler Engineer", "Safety-Critical Software Engineer", "C++ Developer"},
	},
	{
		Name:       "Trail of Bits",
		CareerLink: "https://www.trailofbits.com/careers",
		Jobs:       []string{"Security Researcher", "C++ Security Engineer", "Reverse Engineer"},
	},
	{
		Name:       "Vector35",
		CareerLink: "https://vector35.com/careers",
		Jobs:       []string{"Reverse Engineering Engineer", "C++ Developer", "Binary Analysis Engineer"},
	},
	{
		Name:       "System76",
		CareerLink: "https://system76.com/careers",
		Jobs:       []string{"Linux Kernel Engineer", "C++ Developer", "Systems Developer"},
	},
	{
		Name:       "Microsoft",
		CareerLink: "https://careers.microsoft.com",
		Jobs:       []string{"Software Engineer", "React Software Engineer", "Python Developer", "Fullstack"},
	},
	{
		Name:       "Amazon",
		CareerLink: "https://amazon.jobs",
		Jobs:       []string{"Software Development Engineer", "C++ Engineer", "DevOps", "AWS Solutions Architect"},
	},
	{
		Name:       "Replit",
		CareerLink: "https://replit.com/careers",
		Jobs:       []string{"Backend Engineer", "Systems Developer", "Software Engineer"},
	},
	{
		Name:       "Sourcegraph",
		CareerLink: "https://about.sourcegraph.com/careers",
		Jobs:       []string{"Backend Engineer", "DevOps", "Software Engineer"},
	},
	{
		Name:       "Tailscale",
		CareerLink: "https://tailscale.com/careers",
		Jobs:       []string{"Backend Engineer", "Security Researcher", "Systems Developer"},
	},
	{
		Name:       "Fly.io",
		CareerLink: "https://fly.io/careers",
		Jobs:       []string{"Backend Engineer", "DevOps", "Systems Developer"},
	},
	{
		Name:       "Fathom Analytics",
		CareerLink: "https://usefathom.com/careers",
		Jobs:       []string{"Backend Engineer", "Software Engineer"},
	},
	{
		Name:       "ScyllaDB",
		CareerLink: "https://www.scylladb.com/careers",
		Jobs:       []string{"C++ Developer", "Database Engineer", "Systems Developer"},
	},
	{
		Name:       "ClickHouse",
		CareerLink: "https://clickhouse.com/company/careers",
		Jobs:       []string{"C++ Engineer", "Database Engineer", "Systems Developer"},
	},
	{
		Name:       "Snyk",
		CareerLink: "https://snyk.io/careers",
		Jobs:       []string{"Security Researcher", "Backend Engineer", "DevOps"},
	},
	{
		Name:       "PostHog",
		CareerLink: "https://posthog.com/careers",
		Jobs:       []string{"Backend Engineer", "Software Engineer", "Data Engineer"},
	},
	{
		Name:       "Temporal",
		CareerLink: "https://temporal.io/careers",
		Jobs:       []string{"Backend Engineer", "Distributed Systems Engineer", "Software Engineer"},
	},
	{
		Name:       "Netlify",
		CareerLink: "https://www.netlify.com/careers",
		Jobs:       []string{"React Developer", "Frontend Engineer", "JavaScript Developer"},
	},
	{
		Name:       "Vercel",
		CareerLink: "https://vercel.com/careers",
		Jobs:       []string{"React Developer", "React Software Engineer", "Frontend Engineer"},
	},
	{
		Name:       "Gatsby",
		CareerLink: "https://www.gatsbyjs.com/careers",
		Jobs:       []string{"React Developer", "Frontend Engineer", "JavaScript Developer"},
	},
	{
		Name:       "Contentful",
		CareerLink: "https://www.contentful.com/careers",
		Jobs:       []string{"React Developer", "Frontend Engineer", "Fullstack"},
	},
	{
		Name:       "Storyblok",
		CareerLink: "https://www.storyblok.com/careers",
		Jobs:       []string{"React Developer", "JavaScript Developer", "Frontend Engineer"},
	},
	{
		Name:       "Elastic",
		CareerLink: "https://www.elastic.co/careers",
		Jobs:       []string{"Java Developer", "Backend Engineer", "Software Engineer"},
	},
	{
		Name:       "Neo4j",
		CareerLink: "https://neo4j.com/careers",
		Jobs:       []string{"Java Developer", "Backend Engineer", "Systems Developer"},
	},
	{
		Name:       "Camunda",
		CareerLink: "https://camunda.com/careers",
		Jobs:       []string{"Java Developer", "Backend Engineer", "Software Engineer"},
	},
	{
		Name:       "Sonatype",
		CareerLink: "https://www.sonatype.com/careers",
		Jobs:       []string{"Java Developer", "Backend Engineer", "DevOps"},
	},
	{
		Name:       "Hazelcast",
		CareerLink: "https://hazelcast.com/careers",
		Jobs:       []string{"Java Developer", "Systems Developer", "Backend Engineer"},
	},
	{
		Name:       "Ghost",
		CareerLink: "https://ghost.org/careers",
		Jobs:       []string{"JavaScript Developer", "Backend Engineer", "Fullstack"},
	},
	{
		Name:       "Strapi",
		CareerLink: "https://strapi.io/careers",
		Jobs:       []string{"JavaScript Developer", "Node.js Developer", "Backend Engineer"},
	},
	{
		Name:       "Supabase",
		CareerLink: "https://supabase.com/careers",
		Jobs:       []string{"JavaScript Developer", "Fullstack", "Frontend Engineer"},
	},
	{
		Name:       "Appwrite",
		CareerLink: "https://appwrite.io/careers",
		Jobs:       []string{"JavaScript Developer", "Backend Engineer", "DevOps"},
	},
	{
		Name:       "Directus",
		CareerLink: "https://directus.io/careers",
		Jobs:       []string{"JavaScript Developer", "Fullstack", "Frontend Engineer"},
	},
	{
		Name:       "Apple",
		CareerLink: "https://jobs.apple.com",
		Jobs:       []string{"Software Engineer", "Java Developer", "iOS Engineer", "Hardware Engineer"},
	},
	{
		Name:       "Tesla",
		CareerLink: "https://www.tesla.com/careers",
		Jobs:       []string{"Software Engineer", "Golang Developer", "Autopilot Engineer", "Manufacturing"},
	},
}

func main() {

	app := fiber.New()

	app.Static("/public", "./public")

	app.Get("/api/companies", func(c *fiber.Ctx) error {

		searchQuery := strings.TrimSpace(strings.ToLower(c.Query("search", "")))

		var filtered []Company

		if searchQuery == "" {
			return c.JSON(companies)
		}

		for _, company := range companies {
			for _, job := range company.Jobs {
				if strings.Contains(strings.ToLower(job), searchQuery) {
					filtered = append(filtered, company)
					break
				}
			}
		}

		return c.JSON(filtered)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	fmt.Println("Server running at http://localhost:3000")

	log.Fatal(app.Listen(":3000"))
}
