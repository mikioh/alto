{
    "meta": {
	"cost-types": {
            "num-routing": {
		"cost-mode":  "numerical",
                "cost-metric": "routingcost",
                "description": "My default"
	    },
            "num-hop": {
		"cost-mode":  "numerical",
                "cost-metric": "hopcount"
	    },
            "ord-routing": {
		"cost-mode":  "ordinal",
                "cost-metric": "routingcost"
	    },
            "ord-hop": {
		"cost-mode":  "ordinal",
                "cost-metric": "hopcount"
	    }
	}
    },
    "resources": [
	{
            "uri": "http://alto.example.com/networkmap",
            "media-type": "application/alto-networkmap+json"
	},
	{
            "uri": "http://alto.example.com/costmap/num/routingcost",
            "media-type": "application/alto-costmap+json",
            "capabilities": {
		"cost-type-names": [
		    "num-routing"
		]
            }
	},
	{
            "uri": "http://alto.example.com/costmap/num/hopcount",
            "media-type": "application/alto-costmap+json",
            "capabilities": {
		"cost-type-names": [
		    "num-hop"
		]
            }
	},
	{
            "uri": "http://custom.alto.example.com/maps",
            "media-type": "application/alto-directory+json"
	},
	{
            "uri": "http://alto.example.com/endpointprop/lookup",
            "media-type": "application/alto-endpointprop+json",
            "accepts": "application/alto-endpointpropparams+json",
            "capabilities": {
		"prop-types": [
		    "pid"
		]
            }
	},
	{
            "uri": "http://alto.example.com/endpointcost/lookup",
            "media-type": "application/alto-endpointcost+json",
            "accepts": "application/alto-endpointcostparams+json",
            "capabilities": {
		"cost-constraints": true,
		"cost-type-names": [
		    "num-routing", "num-hop",
                    "ord-routing", "ord-hop"
		]
            }
	}
    ]
}
