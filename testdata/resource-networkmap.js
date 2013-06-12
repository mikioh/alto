{
    "meta": {
	"redistribution": {
	    "service-id": "12ab34cd",
	    "request-uri": "http://foo.bar.com/",
	    "request-body": {
		"cost-mode" : "numerical",
		"cost-type" : "routingcost",
		"pids" : {
		    "srcs" : [ "pid1" ],
		    "dsts" : [ "pid1", "pid2", "pid3" ]
		}
	    },
	    "media-type": "application/alto-costmap+json",
	    "expires": "20141231"
	}
    },
    "data": {
	"map-vtag": "1266506139",
	"map": {
	    "pid1": {
		"ipv4": [
		    "192.0.2.0/24",
		    "198.51.100.0/25"
		]
	    },
	    "pid2": {
		"ipv4": [
		    "198.51.100.128/25"
		]
	    },
	    "pid3": {
		"ipv4": [
		    "0.0.0.0/0"
		],
		"ipv6": [
		    "::/0"
		]
	    }
	}
    }
}
