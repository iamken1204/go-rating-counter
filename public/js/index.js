var host_prefix = 'http://localhost:1234';
$(document).ready(function() {
    // JSONP version - add 'callback=?' to the URL - fetch the JSONP response to the request
    $("#jsonp-button").click(function(e) {
        e.preventDefault();
        // The only difference on the client end is the addition of 'callback=?' to the URL
        var url = host_prefix + '/viewtest?callback=?';
        $.getJSON(url, function(jsonp) {
            console.log(jsonp);
            $("#jsonp-response").html(JSON.stringify(jsonp, null, 2));
        });
    });
});
