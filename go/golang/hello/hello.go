package main

// Multiple import statements can be clubbed.
import (
	"fmt"
	"log"
	"strings"

	"example.com/greetings"
	"rsc.io/quote"

	xj "github.com/basgys/goxml2json"
)

func main() {

	log.SetPrefix("greetings: ")
	// Removes line, time and source file name.
	log.SetFlags(0)

	fmt.Println("Hello World!")
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())

	message, err := greetings.Hello("Kartik")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{
		"Kartik",
		"Iyyer",
		"S",
	}
	messages, err := greetings.Hellos(names)

	fmt.Println(messages)

	for name, message := range messages {
		fmt.Println(name, ": ", message)
	}

	// message, err = greetings.Hello("")
	// if err != nil {
	// 	// Since err will not be nil, program will end here not continue to next line.
	// 	log.Fatal(err)
	// }

	// xml is an io.Reader
	xml := strings.NewReader(`<?xml-stylesheet type="text/xsl" href="/universal_parse.xsl"?>
	<bgp_peer type="rlist">
	<BgpNeighborReq type="sandesh">
	<search_string type="string">0</search_string></BgpNeighborReq>
	<ShowBgpNeighborSummaryReq type="sandesh">
	<search_string type="string">0</search_string></ShowBgpNeighborSummaryReq>
	<ShowRouteReq type="sandesh">
	<routing_table type="string">0</routing_table><routing_instance type="string">0</routing_instance><prefix type="string">0</prefix><longer_match type="bool">0</longer_match><shorter_match type="bool">0</shorter_match><count type="u32">0</count><start_routing_table type="string">0</start_routing_table><start_routing_instance type="string">0</start_routing_instance><start_prefix type="string">0</start_prefix><source type="string">0</source><protocol type="string">0</protocol><family type="string">0</family></ShowRouteReq>
	<ShowRouteSummaryReq type="sandesh">
	<search_string type="string">0</search_string></ShowRouteSummaryReq>
	<ShowRoutingInstanceReq type="sandesh">
	<search_string type="string">0</search_string></ShowRoutingInstanceReq>
	<ShowRoutingInstanceSummaryReq type="sandesh">
	<search_string type="string">0</search_string></ShowRoutingInstanceSummaryReq>
	<ShowRtGroupReq type="sandesh">
	<search_string type="string">0</search_string></ShowRtGroupReq>
	<ShowRtGroupSummaryReq type="sandesh">
	<search_string type="string">0</search_string></ShowRtGroupSummaryReq>
	<ShowRtGroupPeerReq type="sandesh">
	<search_string type="string">0</search_string></ShowRtGroupPeerReq>
	<ShowPathResolverSummaryReq type="sandesh">
	<search_string type="string">0</search_string></ShowPathResolverSummaryReq>
	<ShowPathResolverReq type="sandesh">
	<search_string type="string">0</search_string></ShowPathResolverReq>
	<ShowRibOutStatisticsReq type="sandesh">
	<search_string type="string">0</search_string></ShowRibOutStatisticsReq>
	<ShowEvpnTableReq type="sandesh">
	<search_string type="string">0</search_string></ShowEvpnTableReq>
	<ShowEvpnTableSummaryReq type="sandesh">
	<search_string type="string">0</search_string></ShowEvpnTableSummaryReq>
	<ShowMulticastManagerReq type="sandesh">
	<search_string type="string">0</search_string></ShowMulticastManagerReq>
	<ShowMvpnManagerReq type="sandesh">
	<search_string type="string">0</search_string></ShowMvpnManagerReq>
	<ShowMvpnProjectManagerReq type="sandesh">
	<search_string type="string">0</search_string></ShowMvpnProjectManagerReq>
	<ShowBgpGlobalSystemConfigReq type="sandesh">
	</ShowBgpGlobalSystemConfigReq>
	<ShowBgpInstanceConfigReq type="sandesh">
	<search_string type="string">0</search_string></ShowBgpInstanceConfigReq>
	<ShowBgpPeeringConfigReq type="sandesh">
	<search_string type="string">0</search_string></ShowBgpPeeringConfigReq>
	<ShowBgpNeighborConfigReq type="sandesh">
	<search_string type="string">0</search_string></ShowBgpNeighborConfigReq>
	<ShowBgpRoutingPolicyConfigReq type="sandesh">
	<search_string type="string">0</search_string></ShowBgpRoutingPolicyConfigReq>
	<ShowBgpServerReq type="sandesh">
	</ShowBgpServerReq>
	</bgp_peer>`)
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	fmt.Println(json.String())
	// {"hello": "world"}

}
