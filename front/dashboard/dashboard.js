var datasets = {
	tests_along_day: [
		{ id:0, tests:12, hour:"00"},
		{ id:1, tests:11, hour:"01"},
		{ id:2, tests:18, hour:"02"},
		{ id:3, tests:17, hour:"03"},
		{ id:4, tests:16, hour:"04"},
		{ id:5, tests:23, hour:"05"},
		{ id:6, tests:24, hour:"06"},
		{ id:7, tests:25, hour:"07"},
		{ id:8, tests:30, hour:"08"},
		{ id:9, tests:31, hour:"09"},
		{ id:10, tests:23, hour:"10"},
		{ id:11, tests:20, hour:"11"},
		{ id:12, tests:18, hour:"12"},
		{ id:13, tests:17, hour:"13"},
		{ id:14, tests:11, hour:"14"},
		{ id:15, tests:12, hour:"15"},
		{ id:16, tests:0, hour:"16"},
		{ id:17, tests:0, hour:"17"},
		{ id:18, tests:0, hour:"18"},
		{ id:19, tests:0, hour:"19"},
		{ id:20, tests:0, hour:"20"},
		{ id:21, tests:0, hour:"21"},
		{ id:22, tests:0, hour:"22"},
		{ id:23, tests:0, hour:"23"}
	],
	agents_along_day: [
		{ id:0, agents:1, hour:"00"},
		{ id:1, agents:3, hour:"01"},
		{ id:2, agents:5, hour:"02"},
		{ id:3, agents:5, hour:"03"},
		{ id:4, agents:0, hour:"04"},
		{ id:5, agents:10, hour:"05"},
		{ id:6, agents:11, hour:"06"},
		{ id:7, agents:4, hour:"07"},
		{ id:8, agents:3, hour:"08"},
		{ id:9, agents:6, hour:"09"},
		{ id:10, agents:3, hour:"10"},
		{ id:11, agents:9, hour:"11"},
		{ id:12, agents:8, hour:"12"},
		{ id:13, agents:7, hour:"13"},
		{ id:14, agents:11, hour:"14"},
		{ id:15, agents:8, hour:"15"},
		{ id:16, agents:0, hour:"16"},
		{ id:17, agents:0, hour:"17"},
		{ id:18, agents:0, hour:"18"},
		{ id:19, agents:0, hour:"19"},
		{ id:20, agents:0, hour:"20"},
		{ id:21, agents:0, hour:"21"},
		{ id:22, agents:0, hour:"22"},
		{ id:23, agents:0, hour:"23"}
	],
	sites_monitored: [
		{ id:0, site:"yahoo.com"},
		{ id:1, site:"google.com.mx"},
		{ id:2, site:"google.com"},
		{ id:3, site:"facebook.com"},
		{ id:4, site:"segundamano.mx"},
		{ id:5, site:"github.com"},
		{ id:6, site:"webix.com"},
		{ id:7, site:"spotify.com"},
		{ id:8, site:"schibsted.com"},
		{ id:9, site:"amazon.com"},
		{ id:10, site:"amazon.com.mx"},
		{ id:11, site:"giphy.com"},
		{ id:12, site:"twitter.com"},
		{ id:13, site:"slack.com"},
		{ id:14, site:"play.google.com"},
		{ id:15, site:"itunes.com"}
	],
	connected_agents: [
		{ id:1, location_text:"Ciudad de Mexico, MEX", lat:19.4328801, lon:-99.1920686},
		{ id:2, location_text:"Jalisco, MEX", lat:20.6326773, lon:-103.326049},
		{ id:3, location_text:"Paris, FRA", lat:48.86141, lon:2.3284428},
		{ id:4, location_text:"Cordoba, ESP", lat:37.8812707, lon:-4.7680814},
		{ id:5, location_text:"Illinois, USA", lat:39.8040623, lon:-89.635416},
		{ id:6, location_text:"Santiago, CHL", lat:-33.4333079, lon:-70.6483184},
		{ id:7, location_text:"Katowice, POL", lat:50.2683284, lon:19.0274129}
	]
};

var dashboard_pages = {

	dash_general: {
		title: "General Dashboard",
		body: {
			view: "layout",
			type: "space",
			rows: [
				{
					view: "layout",
					type: "space",
					cols: [
						{
							height: 120,
							css: "dash_indicator dash_green",
							template: "<div title='Tests prefomed today so far'>"
								+ "<div class='webix_icon icon fa-check-square-o'></div>"
								+ "<div class='number'>1,230</div>"
								+ "<div class='title'>Tests</div>"
								+ "</div>"
						},
						{
							height: 120,
							css: "dash_indicator dash_red",
							template: "<div title='Connected Agents'>"
								+ "<div class='webix_icon icon fa-user-secret'></div>"
								+ "<div class='number'>8</div>"
								+ "<div class='title'>Agents</div>"
								+ "</div>"
						},
						{
							height: 120,
							css: "dash_indicator dash_blue",
							template: "<div title='Locations available to test from'>"
								+ "<div class='webix_icon icon fa-globe'></div>"
								+ "<div class='number'>4</div>"
								+ "<div class='title'>Locations</div>"
								+ "</div>"
						},
						{
							id: "dash_indicator_sites",
							height: 120,
							css: "dash_indicator dash_orange",
							template: "<div title='Sites being currently monitored'>"
								+ "<div class='webix_icon icon fa-desktop'></div>"
								+ "<div class='number'>#count#</div>"
								+ "<div class='title'>Sites</div>"
								+ "</div>"
						},
					]
				},
				{
					view: "layout",
					type: "space",
					cols: [
						{
							view: "layout",
							type: "line",
							rows: [
								{
									view: "toolbar",
									height: 40,
									css: "tool_green",
									cols: [
										{ view: "label", label: "<div class='webix_icon icon fa-check-square-o'></div>&nbsp;Tests per Hour"}
									]
								},
								{
									view: "chart",
									type: "area",
									value: "#tests#",
									height: 280,
									color: "#3A3",
									alpha: 0.8,
									xAxis: {
										template: "#hour#",
										title: "Hours"
									},
									yAxis: {
										start: 0,
										title: "Tests",
										template: function(obj){
											return (obj%10?"":obj)
										}
									},
									tooltip: {
										template: "#tests# tests performed within the #hour# hours"
									},
									data: datasets.tests_along_day
								}
							]
						},
						{
							view: "layout",
							type: "line",
							rows: [
								{
									view: "toolbar",
									height: 40,
									css: "tool_red",
									cols: [
										{ view: "label", label: "<div class='webix_icon icon fa-user-secret'></div>&nbsp;Connected Agents per Hour"}
									]
								},
								{
									view: "chart",
									type: "area",
									value: "#agents#",
									height: 280,
									color: "#F33",
									alpha: 0.8,
									xAxis: {
										template: "#hour#",
										title: "Hours"
									},
									yAxis: {
										start: 0,
										title: "Agents",
										template: function(obj){
											return (obj%2?"":obj)
										}
									},
									tooltip: {
										template: "#agents# agents connected within the #hour# hours"
									},
									data: datasets.agents_along_day
								}
							]
						}
					]
				},
				{
					view: "layout",
					type: "space",
					cols: [
						{
							view: "layout",
							type: "line",
							rows: [
								{
									view: "toolbar",
									height: 40,
									css: "tool_blue",
									cols: [
										{ view: "label", label: "<div class='webix_icon icon fa-globe'></div>&nbsp;Locations online"}
									]
								},
								{
									view: "google-map",
									id: "dash_general_map",
									height: 280
								}
							]
						},
						{
							view: "layout",
							type: "line",
							rows: [
								{
									view: "toolbar",
									height: 40,
									css: "tool_orange",
									cols: [
										{ view: "label", label: "<div class='webix_icon icon fa-desktop'></div>&nbsp;Sites monitored"}
									]
								},
								{
									view: "list",
									id: "dash_general_sites_list",
									height: 280,
									select: true,
									template: "<div class='webix_icon icon fa-desktop'></div>&nbsp;[#test_id# #frequency#s] #target_url#"
								}
							]
						}
					]
				}
			]
		},
		afterLoadFn: function () {
			google.maps.event.addListenerOnce($$("dash_general_map").map, 'idle', function () {
				var bounds, map, latLng, marker, infoWindow, x, agent;
				map = this;
				bounds = new google.maps.LatLngBounds();
				for(x in datasets.connected_agents) {
					agent = datasets.connected_agents[x];
					latLng = new google.maps.LatLng(agent.lat, agent.lon);
					bounds.extend(latLng);
					infoWindow = new google.maps.InfoWindow({
						content: "ID: " + agent.id + "<br>"
							+ "Location: <b>" + agent.location_text + "</b>"
					});
					marker = new google.maps.Marker({
						position: latLng,
						map: map,
						title: agent.location_text
					});
					marker.infoWindow = infoWindow;
					marker.infoWindowShow = false;
					google.maps.event.addListener(marker, 'click', function () {
						if(this.infoWindowShow) {
							this.infoWindow.close();
							this.infoWindowShow = false;
						} else {
							this.infoWindow.open(this.map, this);
							this.infoWindowShow = true;
						}
					});
				}
				map.fitBounds(bounds);
			});			
		},
		intervals: {
			sites: setInterval(function () {
				if(!$$("dash_general_sites_list")) return false;
				webix.ajax().get("/api/sites", function(text, data, xmlHttp) {
					data = data.json();
					if(data.status == "OK") {
						$$("dash_general_sites_list").clearAll();
						$$("dash_general_sites_list").parse(data.test_sites);
						if($$("dash_indicator_sites")) {
							$$("dash_indicator_sites").parse({count: data.test_sites.length});
						}
					}
				});

			},5000)
		}
	},

	dash_sites: {
		title: "Sites Dashboard",
		body: {
			view: "layout",
			type: "space",
			rows: [
				{
					view: "form",
					elements: [
						{
							view:"combo", 
							label: 'Site',
							placeholder: "Select a site",
							options:["http://segundamano.mx", "http://dashboard.koolmonkey.xyz"]
						}
					]
				},
				{
					view: "layout",
					type: "space",
					cols: [
						{
							view: "layout",
							type: "line",
							rows: [
								{
									view: "toolbar",
									height: 40,
									css: "tool_green",
									cols: [
										{ view: "label", label: "<div class='webix_icon icon fa-check-square-o'></div>&nbsp;Tests by Site"}
									]
								},
								{
									view: "chart",
									type: "area",
									value: "#tests#",
									height: 380,
									color: "#3A3",
									alpha: 0.8,
									xAxis: {
										template: "#hour#",
										title: "Time"
									},
									yAxis: {
										start: 0,
										title: "Tests",
										template: function(obj){
											return (obj%10?"":obj)
										}
									},
									tooltip: {
										template: "#tests# tests performed within the #hour# hours"
									},
									data: datasets.tests_along_day
								}
							]
						}
					]
				}
			]
		},
		intervals: {

		}
	},

	pDefault: {
		title: "Default Page",
		body: {
			view: "layout",
			type: "space",
			rows: [
				{},
				{
					view: "layout",
					type: "space",
					cols: [
						{},
						{
							height: 200,
							width: 400,
							template: "<p>This is just a <b>demo</b> of a project that is still in development.</p>"
								+ "<p>This project started during the Gopher Gala 2016 (22-24 January).</p>"
								+ "<p>Checkout the project at Github: <a href='https://github.com/gophergala2016/kool_monkey'>https://github.com/gophergala2016/kool_monkey</a></p>"
						},
						{}
					]
				},
				{}
			]
		}
	}

};

var menu_data = [
	{
		id: "dashboards",
		icon: "dashboard",
		value: "Dashboards",
		data: [
			{ id: "dash_general", value: "General" },
			{ id: "dash_sites", value: "Sites" },
			{ id: "dash_agents", value: "Agents" }
		]
	},
	{
		id: "reports",
		icon: "line-chart",
		value: "Reports",
		data: [
			{ id: "rep_sites", value: "Sites" },
			{ id: "rep_agents", value: "Agents" }
		]
	},
	{
		id: "settings",
		icon: "wrench",
		value: "Settings",
		data: [
			{ id: "set_sites", value: "Sites" },
			{ id: "set_agents", value: "Agents" }
		]
	},
	{
		id: "help",
		icon: "life-ring",
		value: "Help"
	}
];

var dashboard_ui = {
	view: "layout",
	type: "clean",
	rows: [
		{
			view: "toolbar",
			padding: 5,
			elements: [
				{
					view: "button",
					type: "icon",
					icon: "bars",
					width: 37,
					align: "left",
					css: "menu_button",
					click: function () {
						$$("menu_bar").toggle();
					}
				},
				{
					view: "label",
					label: "<b>Kool Monkey&nbsp;...&nbsp;<div class='webix_icon icon fa-bolt' style='width:10px;'></div><div class='webix_icon icon fa-cloud'></div></b>"
				}
			]
		},
		{
			cols: [
				{
					view: "sidebar",
					id: "menu_bar",
					data: menu_data,
					on: {
						onAfterSelect: function(id){
							setDashboardPage(id);
						}
					}
				},
				{
					view: "layout",
					id: "main_layout",
					type: "line",
					rows: [
						{
							padding: 10,
							cols: [
								{
									view: "label",
									id: "sectionTitle",
									label: "<span style='font-size:1.2em'></span>"
								}
							]
						},
						{
							view: "scrollview",
							id: "sectionBody",
							scroll: "y",
							body: {
								template: "."
							}
						}
					]
				}
			]
		}
	]
};

function setDashboardPage(pageId) {
	var title, body, afterFnFlag;
	if(dashboard_pages[pageId]) {
		title = dashboard_pages[pageId].title;
		body = dashboard_pages[pageId].body;
		afterFnFlag = (typeof dashboard_pages[pageId].afterLoadFn == "function");
	} else {
		title = $$("menu_bar").getItem(pageId).value;
		body = dashboard_pages.pDefault.body;
		afterFnFlag = false;
	}
	$$("main_layout").define({
		rows: [
			{
				padding: 10,
				cols: [
					{
						view: "label",
						id: "sectionTitle",
						label: "<span style='font-size:1.2em'>" + title + "</span>"
					}
				]
			},
			{
				view: "scrollview",
				id: "sectionBody",
				scroll: "y",
				body: body
			}
		]
	});
	$$("main_layout").reconstruct();

	if(afterFnFlag) {
		setTimeout("dashboard_pages['" + pageId + "'].afterLoadFn();",10);
	}
	return true;
}

setTimeout(function () {
	setDashboardPage("dash_general");
},1000);
