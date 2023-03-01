package utils

import (
	"strconv"
	"strings"

	"github.com/go-co-op/gocron"
)

const Header = `
<html>
	<head>
		<title>Go Scheduler - Dashboard</title>
		<script src="https://unpkg.com/htmx.org@1.8.5" integrity="sha384-7aHh9lqPYGYZ7sTHvzP1t3BAfLhYSTy9ArHdP3Xsr9/3TlGurYgcPBoFmXX2TX/w" crossorigin="anonymous"></script>
		<script src="https://cdn.tailwindcss.com"></script>
	</head>
`
const BodyStart = `
<body class="bg-slate-200">
	<div class="flex items-center border-b border-slate-300">
		<img class="h-8 w-8 mx-4 my-6" src="data:image/png;base64,` + favicon + `" />
		<span class="text-xl">Go Scheduler - Dashboard</span>
	</div>
	<div class="px-8 py-4" hx-get="/jobs" hx-trigger="load" hx-swap="outerHTML">
`
const BodyEnd = "</div></body>"
const Footer = "</html>"

var Count = 0

// BuildJobOutput will build the html required to display the job data
func BuildJobOutput(job *gocron.Job) string {
	running := "No"
	if job.IsRunning() {
		running = "Yes"
	}
	err := "N/A"
	if job.Error() != nil {
		err = job.Error().Error()
	}
	return `
	<div class="px-8 py-4" hx-get="/jobs" hx-trigger="load delay:15s" hx-swap="outerHTML">
		<div class="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
			<div>Tags: ` + strings.Join(job.Tags(), ", ") + `</div>
			<div>Running: ` + running + `</div>
			<div>Last Run: ` + job.LastRun().String() + `</div>
			<div>Next Run: ` + job.NextRun().String() + `</div>
			<div>Times Run: ` + strconv.Itoa(job.RunCount()) + `</div>
			<div>Last Error: ` + err + `</div>
			<div>Count: ` + strconv.Itoa(Count) + `</div>
		</div>
	</div>
	`
}
