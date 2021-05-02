

$(document).ready(function() {

    // header
    $(".modal-background").click(function () {
        $(".modal").removeClass("is-active")
    });

    $(".ip-scan-cancel").click(function() {
        $(".ip-scan").toggleClass("is-active");
    });

    $(".port-scan-cancel").click(function() {
        $(".port-scan").toggleClass("is-active");
    });

    $(".domain-scan-cancel").click(function() {
        $(".domain-scan").toggleClass("is-active");
    });

    $(".dis-set-cancel").click(function() {
        $(".dis-set").toggleClass("is-active");
    });


    // header submit
    $(".ip-submit").click(function () {
        var ips = $(".ip-input").val();
        window.location = "/ipResult?ips=" + ips;
    });

    $(".port-submit").click(function () {
        var ip = $(".port-input-ip").val()
        var ports = $(".port-input-ports").val()
        var strs = ports.split("-")
        window.location = "/portResult?ip=" + ip + "&from=" + strs[0] + " &to=" + strs[1];
    });

    $(".domain-submit").click(function () {
        var domain = $(".domain-input").val();
        window.location = "/domainResult?domain=" + domain;
    });


    // home page
    function checkAndScan(){
        var domain = $(".home-input").val();
        if (domain === ""){
            alert("请输入要扫描的域名！");
        }else{
            window.location = "/result?domain=" + domain;
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

