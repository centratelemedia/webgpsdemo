class Map {
    constructor() {
        this.mapInitialized = false;
    }
    init() {
        console.log('init Map App')
        const that = this;
        this.pages = [

            {
                // dst element id/page element id
                dst: "page_map",
                // sub-html-page, which would be loaded to dst element.innerHTML
                url: "pages/map.html",
                lazy: true,
                // init will be called when the page is loaded, if there's not an url, it will be called immediately
                init: function (page) {
                    console.log("pages Map init")
                    that.initPageMap();
                },
                // onshow will be called every time when this page is displayed
                onshow: function (page) {
                   
                },
                // onshow will be called every time when this page is hided
                onhide: function (page) {
                    console.log("page_map onhide");
                },
                // bind children
                children: [],
            },
            {
                dst: "page_report",
                url: "pages/report.html",
                lazy: true,
                init: function (page) {
                    console.log("page_report init");
                },
                onshow: function (page) {
                    console.log("page_report onshow");
                },
                onhide: function (page) {
                    console.log("page_report onhide");
                }
            },
        ];

        $pm.bindPages(this.pages);
        $pm.listenRouter();
        $pm.select("page_map");


    }
}