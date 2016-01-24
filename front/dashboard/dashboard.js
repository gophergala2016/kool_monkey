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
					data: menu_data
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
									label: "<span style='font-size:1.2em'>General Dashboard</span>"
								}
							]
						},
						{
							view: "scrollview",
							scroll: "y",
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
												url: "asdasd",
												css: "dash_indicator dash_green",
												template: "<div title='Tests prefomed today so far'>"
													+ "<div class='webix_icon icon fa-check-square-o'></div>"
													+ "<div class='number'>1,230</div>"
													+ "<div class='title'>Tests</div>"
													+ "</div>"
											},
											{
												height: 120,
												url: "asdasd",
												css: "dash_indicator dash_red",
												template: "<div title='Connected Agents'>"
													+ "<div class='webix_icon icon fa-user-secret'></div>"
													+ "<div class='number'>8</div>"
													+ "<div class='title'>Agents</div>"
													+ "</div>"
											},
											{
												height: 120,
												url: "asdasd",
												css: "dash_indicator dash_blue",
												template: "<div title='Locations available to test from'>"
													+ "<div class='webix_icon icon fa-globe'></div>"
													+ "<div class='number'>4</div>"
													+ "<div class='title'>Locations</div>"
													+ "</div>"
											},
											{
												height: 120,
												url: "asdasd",
												css: "dash_indicator dash_orange",
												template: "<div title='Sites being currently monitored'>"
													+ "<div class='webix_icon icon fa-desktop'></div>"
													+ "<div class='number'>19</div>"
													+ "<div class='title'>Sites</div>"
													+ "</div>"
											},
										]
									},
									{
										view: "layout",
										type: "space",
										cols: [
											{ template: "5"},
											{ template: "6"},
										]
									},
									{
										view: "layout",
										type: "space",
										cols: [
											{ template: "5"},
											{ template: "6"},
										]
									}
								]
							}
						}
					]
				}
			]
		}
	]
};
