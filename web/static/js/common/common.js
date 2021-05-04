

function GetURLParameter(sParam) {
    var sPageURL = window.location.search.substring(1);
    var sURLVariables = sPageURL.split('&');
    for (var i = 0; i < sURLVariables.length; i++) {
        var sParameterName = sURLVariables[i].split('=');
        if (sParameterName[0] == sParam) {
            return sParameterName[1];
        }
    }
}


$(document).ready(function () {

    // header
    $(".modal-background").click(function () {
        $(".modal").removeClass("is-active")
    });

    $(".ip-scan-cancel").click(function () {
        $(".ip-scan").toggleClass("is-active");
    });

    $(".port-scan-cancel").click(function () {
        $(".port-scan").toggleClass("is-active");
    });

    $(".domain-scan-cancel").click(function () {
        $(".domain-scan").toggleClass("is-active");
    });

    // $(".dis-set-cancel").click(function () {
    //     $(".dis-set").toggleClass("is-active");
    // });


    // header submit
    $(".ip-submit").click(function () {
        var ips = $(".ip-input").val();
        var target = "/ipResult?ips=" + ips;
        window.location = "/loading?url=" + encodeURIComponent(target)
    });

    $(".port-submit").click(function () {
        var ip = $(".port-input-ip").val()
        var ports = $(".port-input-ports").val()
        var strs = ports.split("-")
        var target = "/portResult?ip=" + ip + "&from=" + strs[0] + " &to=" + strs[1];
        window.location = "/loading?url=" + encodeURIComponent(target)
    });

    $(".domain-submit").click(function () {
        var domain = $(".domain-input").val();
        var target = "/domainResult?domain=" + domain;
        window.location = "/loading?url=" + encodeURIComponent(target)
    });


    // home page
    function checkAndScan() {
        var domain = $(".home-input").val();
        if (domain === "") {
            alert("请输入要扫描的域名！");
        } else {
            var target = "/result?domain=" + domain;
            window.location = "/loading?url=" + encodeURIComponent(target)
        }
    }

    $(".home-submit").click(function () {
        checkAndScan();
    });

    $(".home-input").keypress(function (e) {
        if (e.which === 13) {
            checkAndScan();
        }
    });

});

