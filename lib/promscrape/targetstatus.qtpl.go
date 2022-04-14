// Code generated by qtc from "targetstatus.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line targetstatus.qtpl:1
package promscrape

//line targetstatus.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/prompbmarshal"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promrelabel"
	"time"
)

//line targetstatus.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line targetstatus.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line targetstatus.qtpl:9
func StreamTargetsResponsePlain(qw422016 *qt422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels bool) {
//line targetstatus.qtpl:11
	for _, js := range jts {
//line targetstatus.qtpl:11
		qw422016.N().S(`job=`)
//line targetstatus.qtpl:12
		qw422016.N().Q(js.job)
//line targetstatus.qtpl:12
		qw422016.N().S(`(`)
//line targetstatus.qtpl:12
		qw422016.N().D(js.upCount)
//line targetstatus.qtpl:12
		qw422016.N().S(`/`)
//line targetstatus.qtpl:12
		qw422016.N().D(js.targetsTotal)
//line targetstatus.qtpl:12
		qw422016.N().S(` `)
//line targetstatus.qtpl:12
		qw422016.N().S(`up)`)
//line targetstatus.qtpl:13
		qw422016.N().S(`
`)
//line targetstatus.qtpl:14
		for _, ts := range js.targetsStatus {
//line targetstatus.qtpl:15
			qw422016.N().S("\t")
//line targetstatus.qtpl:15
			qw422016.N().S(`state=`)
//line targetstatus.qtpl:16
			if ts.up {
//line targetstatus.qtpl:16
				qw422016.N().S(`up`)
//line targetstatus.qtpl:16
			} else {
//line targetstatus.qtpl:16
				qw422016.N().S(`down`)
//line targetstatus.qtpl:16
			}
//line targetstatus.qtpl:16
			qw422016.N().S(`,`)
//line targetstatus.qtpl:16
			qw422016.N().S(` `)
//line targetstatus.qtpl:16
			qw422016.N().S(`endpoint=`)
//line targetstatus.qtpl:17
			qw422016.N().S(ts.sw.Config.ScrapeURL)
//line targetstatus.qtpl:17
			qw422016.N().S(`,`)
//line targetstatus.qtpl:17
			qw422016.N().S(` `)
//line targetstatus.qtpl:17
			qw422016.N().S(`labels=`)
//line targetstatus.qtpl:18
			qw422016.N().S(promLabelsString(promrelabel.FinalizeLabels(nil, ts.sw.Config.Labels)))
//line targetstatus.qtpl:18
			qw422016.N().S(`,`)
//line targetstatus.qtpl:18
			qw422016.N().S(` `)
//line targetstatus.qtpl:19
			if showOriginLabels {
//line targetstatus.qtpl:19
				qw422016.N().S(`originalLabels=`)
//line targetstatus.qtpl:19
				qw422016.N().S(promLabelsString(ts.sw.Config.OriginalLabels))
//line targetstatus.qtpl:19
				qw422016.N().S(`,`)
//line targetstatus.qtpl:19
				qw422016.N().S(` `)
//line targetstatus.qtpl:19
			}
//line targetstatus.qtpl:19
			qw422016.N().S(`scrapes_total=`)
//line targetstatus.qtpl:20
			qw422016.N().D(ts.scrapesTotal)
//line targetstatus.qtpl:20
			qw422016.N().S(`,`)
//line targetstatus.qtpl:20
			qw422016.N().S(` `)
//line targetstatus.qtpl:20
			qw422016.N().S(`scrapes_failed=`)
//line targetstatus.qtpl:21
			qw422016.N().D(ts.scrapesFailed)
//line targetstatus.qtpl:21
			qw422016.N().S(`,`)
//line targetstatus.qtpl:21
			qw422016.N().S(` `)
//line targetstatus.qtpl:21
			qw422016.N().S(`last_scrape=`)
//line targetstatus.qtpl:22
			qw422016.N().FPrec(ts.getDurationFromLastScrape().Seconds(), 3)
//line targetstatus.qtpl:22
			qw422016.N().S(`s ago,`)
//line targetstatus.qtpl:22
			qw422016.N().S(` `)
//line targetstatus.qtpl:22
			qw422016.N().S(`scrape_duration=`)
//line targetstatus.qtpl:23
			qw422016.N().D(int(ts.scrapeDuration))
//line targetstatus.qtpl:23
			qw422016.N().S(`ms,`)
//line targetstatus.qtpl:23
			qw422016.N().S(` `)
//line targetstatus.qtpl:23
			qw422016.N().S(`samples_scraped=`)
//line targetstatus.qtpl:24
			qw422016.N().D(ts.samplesScraped)
//line targetstatus.qtpl:24
			qw422016.N().S(`,`)
//line targetstatus.qtpl:24
			qw422016.N().S(` `)
//line targetstatus.qtpl:24
			qw422016.N().S(`error=`)
//line targetstatus.qtpl:25
			if ts.err != nil {
//line targetstatus.qtpl:25
				qw422016.N().S(ts.err.Error())
//line targetstatus.qtpl:25
			}
//line targetstatus.qtpl:26
			qw422016.N().S(`
`)
//line targetstatus.qtpl:27
		}
//line targetstatus.qtpl:28
	}
//line targetstatus.qtpl:30
	for _, jobName := range emptyJobs {
//line targetstatus.qtpl:30
		qw422016.N().S(`job=`)
//line targetstatus.qtpl:31
		qw422016.N().Q(jobName)
//line targetstatus.qtpl:31
		qw422016.N().S(`(0/0 up)`)
//line targetstatus.qtpl:32
		qw422016.N().S(`
`)
//line targetstatus.qtpl:33
	}
//line targetstatus.qtpl:35
}

//line targetstatus.qtpl:35
func WriteTargetsResponsePlain(qq422016 qtio422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels bool) {
//line targetstatus.qtpl:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line targetstatus.qtpl:35
	StreamTargetsResponsePlain(qw422016, jts, emptyJobs, showOriginLabels)
//line targetstatus.qtpl:35
	qt422016.ReleaseWriter(qw422016)
//line targetstatus.qtpl:35
}

//line targetstatus.qtpl:35
func TargetsResponsePlain(jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels bool) string {
//line targetstatus.qtpl:35
	qb422016 := qt422016.AcquireByteBuffer()
//line targetstatus.qtpl:35
	WriteTargetsResponsePlain(qb422016, jts, emptyJobs, showOriginLabels)
//line targetstatus.qtpl:35
	qs422016 := string(qb422016.B)
//line targetstatus.qtpl:35
	qt422016.ReleaseByteBuffer(qb422016)
//line targetstatus.qtpl:35
	return qs422016
//line targetstatus.qtpl:35
}

//line targetstatus.qtpl:37
func StreamTargetsResponseHTML(qw422016 *qt422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, onlyUnhealthy bool) {
//line targetstatus.qtpl:37
	qw422016.N().S(`<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous"><title>Scrape targets</title></head><body class="py-3"><div class="container-fluid"><div class="row"><main class="col-12"><h1>Scrape targets</h1><div class="row mb-4"><div class="col-12"><button type="button" class="me-1 btn`)
//line targetstatus.qtpl:53
	qw422016.N().S(` `)
//line targetstatus.qtpl:53
	if !onlyUnhealthy {
//line targetstatus.qtpl:53
		qw422016.N().S(`btn-primary`)
//line targetstatus.qtpl:53
	} else {
//line targetstatus.qtpl:53
		qw422016.N().S(`btn-secondary`)
//line targetstatus.qtpl:53
	}
//line targetstatus.qtpl:53
	qw422016.N().S(`" onclick="location.href='targets'">All</button><button type="button" class="me-1 btn`)
//line targetstatus.qtpl:56
	qw422016.N().S(` `)
//line targetstatus.qtpl:56
	if onlyUnhealthy {
//line targetstatus.qtpl:56
		qw422016.N().S(`btn-primary`)
//line targetstatus.qtpl:56
	} else {
//line targetstatus.qtpl:56
		qw422016.N().S(`btn-secondary`)
//line targetstatus.qtpl:56
	}
//line targetstatus.qtpl:56
	qw422016.N().S(`" onclick="location.href='targets?show_only_unhealthy=true'">Unhealthy</button><button type="button" class="btn btn-primary me-1" onclick="collapse_all()">Collapse all</button><button type="button" class="btn btn-secondary me-1" onclick="expand_all()">Expand all</button></div></div><div class="row"><div class="col-12">`)
//line targetstatus.qtpl:69
	for i, js := range jts {
//line targetstatus.qtpl:70
		if onlyUnhealthy && js.upCount == js.targetsTotal {
//line targetstatus.qtpl:70
			continue
//line targetstatus.qtpl:70
		}
//line targetstatus.qtpl:70
		qw422016.N().S(`<div class="row mb-4"><div class="col-12"><h4>`)
//line targetstatus.qtpl:74
		qw422016.E().S(js.job)
//line targetstatus.qtpl:74
		qw422016.N().S(` `)
//line targetstatus.qtpl:74
		qw422016.N().S(`(`)
//line targetstatus.qtpl:74
		qw422016.N().D(js.upCount)
//line targetstatus.qtpl:74
		qw422016.N().S(`/`)
//line targetstatus.qtpl:74
		qw422016.N().D(js.targetsTotal)
//line targetstatus.qtpl:74
		qw422016.N().S(` `)
//line targetstatus.qtpl:74
		qw422016.N().S(`up)</h4><div class="row mb-2"><div class="col-12"><button type="button" class="btn btn-primary me-1"onclick="document.getElementById('table-`)
//line targetstatus.qtpl:79
		qw422016.N().D(i)
//line targetstatus.qtpl:79
		qw422016.N().S(`').style.display='none'">collapse</button><button type="button" class="btn btn-secondary me-1"onclick="document.getElementById('table-`)
//line targetstatus.qtpl:82
		qw422016.N().D(i)
//line targetstatus.qtpl:82
		qw422016.N().S(`').style.display='block'">expand</button></div></div><div id="table-`)
//line targetstatus.qtpl:86
		qw422016.N().D(i)
//line targetstatus.qtpl:86
		qw422016.N().S(`" class="table-responsive"><div class="col-4 mb-2"><input type="text" id="search-`)
//line targetstatus.qtpl:88
		qw422016.N().D(i)
//line targetstatus.qtpl:88
		qw422016.N().S(`" placeholder="Search" class="form-control"/><small class="form-text text-muted">Enter text or regexp and it will filter by two fields <spanclass="badge bg-light text-dark">endpoint</span>or<span class="badge bg-light text-dark">labels</span></small></div><table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col">Endpoint</th><th scope="col">State</th><th scope="col" title="scrape target labels">Labels</th><th scope="col" title="total scrapes">Scrapes</th><th scope="col" title="total scrape errors">Errors</th><th scope="col" title="the time of the last scrape">Last Scrape</th><th scope="col" title="the duration of the last scrape">Duration</th><th scope="col" title="the number of metrics scraped during the last scrape">Samples</th><th scope="col" title="error from the last scrape (if any)">Last error</th></tr></thead><tbody class="list-`)
//line targetstatus.qtpl:110
		qw422016.N().D(i)
//line targetstatus.qtpl:110
		qw422016.N().S(`">`)
//line targetstatus.qtpl:111
		for _, ts := range js.targetsStatus {
//line targetstatus.qtpl:113
			endpoint := ts.sw.Config.ScrapeURL
			targetID := getTargetID(ts.sw)
			lastScrapeTime := ts.getDurationFromLastScrape()

//line targetstatus.qtpl:117
			if onlyUnhealthy && ts.up {
//line targetstatus.qtpl:117
				continue
//line targetstatus.qtpl:117
			}
//line targetstatus.qtpl:117
			qw422016.N().S(`<tr`)
//line targetstatus.qtpl:118
			if !ts.up {
//line targetstatus.qtpl:118
				qw422016.N().S(` `)
//line targetstatus.qtpl:118
				qw422016.N().S(`class="alert alert-danger" role="alert"`)
//line targetstatus.qtpl:118
			}
//line targetstatus.qtpl:118
			qw422016.N().S(`><td class="endpoint"><a href="`)
//line targetstatus.qtpl:119
			qw422016.E().S(endpoint)
//line targetstatus.qtpl:119
			qw422016.N().S(`" target="_blank">`)
//line targetstatus.qtpl:119
			qw422016.E().S(endpoint)
//line targetstatus.qtpl:119
			qw422016.N().S(`</a> (<a href="target_response?id=`)
//line targetstatus.qtpl:120
			qw422016.E().S(targetID)
//line targetstatus.qtpl:120
			qw422016.N().S(`" target="_blank"title="click to fetch target response on behalf of the scraper">response</a>)</td><td>`)
//line targetstatus.qtpl:124
			if ts.up {
//line targetstatus.qtpl:124
				qw422016.N().S(`UP`)
//line targetstatus.qtpl:124
			} else {
//line targetstatus.qtpl:124
				qw422016.N().S(`DOWN`)
//line targetstatus.qtpl:124
			}
//line targetstatus.qtpl:124
			qw422016.N().S(`</td><td class="labels"><div title="click to show original labels"onclick="document.getElementById('original_labels_`)
//line targetstatus.qtpl:127
			qw422016.E().S(targetID)
//line targetstatus.qtpl:127
			qw422016.N().S(`').style.display='block'">`)
//line targetstatus.qtpl:128
			streamformatLabel(qw422016, promrelabel.FinalizeLabels(nil, ts.sw.Config.Labels))
//line targetstatus.qtpl:128
			qw422016.N().S(`</div><div style="display:none" id="original_labels_`)
//line targetstatus.qtpl:130
			qw422016.E().S(targetID)
//line targetstatus.qtpl:130
			qw422016.N().S(`">`)
//line targetstatus.qtpl:131
			streamformatLabel(qw422016, ts.sw.Config.OriginalLabels)
//line targetstatus.qtpl:131
			qw422016.N().S(`</div></td><td>`)
//line targetstatus.qtpl:134
			qw422016.N().D(ts.scrapesTotal)
//line targetstatus.qtpl:134
			qw422016.N().S(`</td><td>`)
//line targetstatus.qtpl:135
			qw422016.N().D(ts.scrapesFailed)
//line targetstatus.qtpl:135
			qw422016.N().S(`</td><td>`)
//line targetstatus.qtpl:137
			if lastScrapeTime < 365*24*time.Hour {
//line targetstatus.qtpl:138
				qw422016.N().FPrec(lastScrapeTime.Seconds(), 3)
//line targetstatus.qtpl:138
				qw422016.N().S(`s ago`)
//line targetstatus.qtpl:139
			} else {
//line targetstatus.qtpl:139
				qw422016.N().S(`none`)
//line targetstatus.qtpl:141
			}
//line targetstatus.qtpl:141
			qw422016.N().S(`<td>`)
//line targetstatus.qtpl:142
			qw422016.N().D(int(ts.scrapeDuration))
//line targetstatus.qtpl:142
			qw422016.N().S(`ms</td><td>`)
//line targetstatus.qtpl:143
			qw422016.N().D(ts.samplesScraped)
//line targetstatus.qtpl:143
			qw422016.N().S(`</td><td>`)
//line targetstatus.qtpl:144
			if ts.err != nil {
//line targetstatus.qtpl:144
				qw422016.E().S(ts.err.Error())
//line targetstatus.qtpl:144
			}
//line targetstatus.qtpl:144
			qw422016.N().S(`</td></tr>`)
//line targetstatus.qtpl:146
		}
//line targetstatus.qtpl:146
		qw422016.N().S(`</tbody></table></div></div></div>`)
//line targetstatus.qtpl:152
	}
//line targetstatus.qtpl:152
	qw422016.N().S(`</div></div>`)
//line targetstatus.qtpl:156
	for _, jobName := range emptyJobs {
//line targetstatus.qtpl:156
		qw422016.N().S(`<div><h4><a>`)
//line targetstatus.qtpl:159
		qw422016.E().S(jobName)
//line targetstatus.qtpl:159
		qw422016.N().S(`(0/0 up)</a></h4><table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col">Endpoint</th><th scope="col">State</th><th scope="col">Labels</th><th scope="col">Last Scrape</th><th scope="col">Scrape Duration</th><th scope="col">Samples Scraped</th><th scope="col">Error</th></tr></thead></table></div>`)
//line targetstatus.qtpl:175
	}
//line targetstatus.qtpl:175
	qw422016.N().S(`</main></div></div></body><script src="//cdnjs.cloudflare.com/ajax/libs/list.js/2.3.1/list.min.js"></script><script>function collapse_all() {for (var i = 0; i <`)
//line targetstatus.qtpl:184
	qw422016.N().D(len(jts))
//line targetstatus.qtpl:184
	qw422016.N().S(`; i++) {let el = document.getElementById("table-" + i);if (!el) {continue;}el.style.display = 'none';}}function expand_all() {for (var i = 0; i <`)
//line targetstatus.qtpl:193
	qw422016.N().D(len(jts))
//line targetstatus.qtpl:193
	qw422016.N().S(`; i++) {let el = document.getElementById("table-" + i);if (!el) {continue;}el.style.display = 'block';}}</script><script>var usersLists = [];for (var i = 0; i <=`)
//line targetstatus.qtpl:204
	qw422016.N().D(len(jts))
//line targetstatus.qtpl:204
	qw422016.N().S(`-1; i++) {var options = {valueNames: [ 'endpoint', 'labels' ],listClass: `)
//line targetstatus.qtpl:204
	qw422016.N().S("`")
//line targetstatus.qtpl:204
	qw422016.N().S(`list-${i}`)
//line targetstatus.qtpl:204
	qw422016.N().S("`")
//line targetstatus.qtpl:204
	qw422016.N().S(`};usersLists.push({ lists: new List(`)
//line targetstatus.qtpl:204
	qw422016.N().S("`")
//line targetstatus.qtpl:204
	qw422016.N().S(`table-${i}`)
//line targetstatus.qtpl:204
	qw422016.N().S("`")
//line targetstatus.qtpl:204
	qw422016.N().S(`, options), input: document.querySelector(`)
//line targetstatus.qtpl:204
	qw422016.N().S("`")
//line targetstatus.qtpl:204
	qw422016.N().S(`#search-${i}`)
//line targetstatus.qtpl:204
	qw422016.N().S("`")
//line targetstatus.qtpl:204
	qw422016.N().S(`)});}usersLists.forEach((userList) => {userList.lists.search('', ['endpoint'], searchFunction);userList.lists.search('', ['labels'], searchFunction);userList.input.addEventListener('keyup', debounce(function(event) {userList.lists.search(event.target.value, ['endpoint'], searchFunction);userList.lists.search(event.target.value, ['labels'], searchFunction);}));function searchFunction(searchString, columns) {for (var k = 0, kl = userList.lists.items.length; k < kl; k++) {const text = userList.lists.items[k]._values[columns[0]];userList.lists.items[k].found = regexpSearch(searchString, text);}}});function regexpSearch(searchString, text) {var searchRegex = "(?=.*"+searchString.replace(/(?:\W+)/g,")(?=.*")+").*";var searchRegexObj = new RegExp(searchRegex,"gm");return (searchString !== "") && (text.search(searchRegexObj) > -1);}function debounce(func, timeout = 300){let timer;return (...args) => {clearTimeout(timer);timer = setTimeout(() => { func.apply(this, args); }, timeout);};}</script></html>`)
//line targetstatus.qtpl:245
}

//line targetstatus.qtpl:245
func WriteTargetsResponseHTML(qq422016 qtio422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, onlyUnhealthy bool) {
//line targetstatus.qtpl:245
	qw422016 := qt422016.AcquireWriter(qq422016)
//line targetstatus.qtpl:245
	StreamTargetsResponseHTML(qw422016, jts, emptyJobs, onlyUnhealthy)
//line targetstatus.qtpl:245
	qt422016.ReleaseWriter(qw422016)
//line targetstatus.qtpl:245
}

//line targetstatus.qtpl:245
func TargetsResponseHTML(jts []jobTargetsStatuses, emptyJobs []string, onlyUnhealthy bool) string {
//line targetstatus.qtpl:245
	qb422016 := qt422016.AcquireByteBuffer()
//line targetstatus.qtpl:245
	WriteTargetsResponseHTML(qb422016, jts, emptyJobs, onlyUnhealthy)
//line targetstatus.qtpl:245
	qs422016 := string(qb422016.B)
//line targetstatus.qtpl:245
	qt422016.ReleaseByteBuffer(qb422016)
//line targetstatus.qtpl:245
	return qs422016
//line targetstatus.qtpl:245
}

//line targetstatus.qtpl:247
func streamformatLabel(qw422016 *qt422016.Writer, labels []prompbmarshal.Label) {
//line targetstatus.qtpl:247
	qw422016.N().S(`{`)
//line targetstatus.qtpl:249
	for i, label := range labels {
//line targetstatus.qtpl:250
		qw422016.E().S(label.Name)
//line targetstatus.qtpl:250
		qw422016.N().S(`=`)
//line targetstatus.qtpl:250
		qw422016.E().Q(label.Value)
//line targetstatus.qtpl:251
		if i+1 < len(labels) {
//line targetstatus.qtpl:251
			qw422016.N().S(`,`)
//line targetstatus.qtpl:251
			qw422016.N().S(` `)
//line targetstatus.qtpl:251
		}
//line targetstatus.qtpl:252
	}
//line targetstatus.qtpl:252
	qw422016.N().S(`}`)
//line targetstatus.qtpl:254
}

//line targetstatus.qtpl:254
func writeformatLabel(qq422016 qtio422016.Writer, labels []prompbmarshal.Label) {
//line targetstatus.qtpl:254
	qw422016 := qt422016.AcquireWriter(qq422016)
//line targetstatus.qtpl:254
	streamformatLabel(qw422016, labels)
//line targetstatus.qtpl:254
	qt422016.ReleaseWriter(qw422016)
//line targetstatus.qtpl:254
}

//line targetstatus.qtpl:254
func formatLabel(labels []prompbmarshal.Label) string {
//line targetstatus.qtpl:254
	qb422016 := qt422016.AcquireByteBuffer()
//line targetstatus.qtpl:254
	writeformatLabel(qb422016, labels)
//line targetstatus.qtpl:254
	qs422016 := string(qb422016.B)
//line targetstatus.qtpl:254
	qt422016.ReleaseByteBuffer(qb422016)
//line targetstatus.qtpl:254
	return qs422016
//line targetstatus.qtpl:254
}
