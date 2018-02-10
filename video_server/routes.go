package main

type JSONRoute struct {
	Method      string
	Pattern     string
	HandlerFunc APIHandler
}

type JSONRoutes []JSONRoute

type RawRoute struct {
	Method      string
	Pattern     string
	HandlerFunc RawAPIHandler
}

type RawRoutes []RawRoute

var jsonRoutes = JSONRoutes{
	JSONRoute{
		"GET",
		"/list",
		getList,
	},
	JSONRoute{
		"GET",
		"/video/{id}",
		getVideo,
	},
}

var rawRoutes = RawRoutes{
	RawRoute{
		"POST",
		"/video",
		uploadVideo,
	},
}
