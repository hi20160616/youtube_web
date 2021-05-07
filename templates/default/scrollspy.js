{{define "scrollspy_js"}}
<script charset="utf-8">
(function ($) {
    // Refer: https://blog.csdn.net/whu_zhangmin/article/details/42776207
    var pos = 0;
    var LIST_ITEM_SIZE = {{len .Data.Items}}
    //滚动条距底部的距离
    var BOTTOM_OFFSET = 0;
    $(document).ready(function () {
	$(window).scroll(function () {
	    var $currentWindow = $(window);
	    //当前窗口的高度
	    var windowHeight = $currentWindow.height();
	    // console.log("current widow height is " + windowHeight);
	    //当前滚动条从上往下滚动的距离
	    var scrollTop = $currentWindow.scrollTop();
	    // console.log("current scrollOffset is " + scrollTop);
	    //当前文档的高度
	    var docHeight = $(document).height();
	    // console.log("current docHeight is " + docHeight);

	    //当 滚动条距底部的距离 + 滚动条滚动的距离 >= 文档的高度 - 窗口的高度
	    //换句话说：（滚动条滚动的距离 + 窗口的高度 = 文档的高度）  这个是基本的公式
	    if ((BOTTOM_OFFSET + scrollTop) >= docHeight - windowHeight) {
		nextVideos();
	    }
	});
    });

    function nextVideos() {
	var next = $(".pagination__next")
	var req = $.trim(next.attr("href"))
	if (next.length == 0) { // first request
	    $.get("../cidNext/?cid={{(index .Data.Items 0).Snippet.ChannelId}}&p={{.Data.NextPageToken}}", function(data, status){
		$("#items").append(data)
	    });
	} else { // infinite request
	    var isNil = getP(req);
	    console.log(isNil)
	    if (isNil != "") { // work if only NextPageToken is non-nil
		$(".pagination__next").remove()
		$.get(req, function(data, status){
		    $("#items").append(data)
		})
	    }
	}
    }

    function getP(url) {
	var pos = url.search("&p=")
	return url.substr(pos+3)
    }
})(jQuery);
</script>
{{end}}
