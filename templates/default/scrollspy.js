(function ($) {
    // Refer: https://blog.csdn.net/whu_zhangmin/article/details/42776207
    var pos = 0;
    var LIST_ITEM_SIZE = 100;
    //滚动条距底部的距离
    var BOTTOM_OFFSET = 0;
    createListItems();
    $(document).ready(function () {
        $(window).scroll(function () {
            var $currentWindow = $(window);
            //当前窗口的高度
            var windowHeight = $currentWindow.height();
            console.log("current widow height is " + windowHeight);
            //当前滚动条从上往下滚动的距离
            var scrollTop = $currentWindow.scrollTop();
            console.log("current scrollOffset is " + scrollTop);
            //当前文档的高度
            var docHeight = $(document).height();
            console.log("current docHeight is " + docHeight);
 
            //当 滚动条距底部的距离 + 滚动条滚动的距离 >= 文档的高度 - 窗口的高度
            //换句话说：（滚动条滚动的距离 + 窗口的高度 = 文档的高度）  这个是基本的公式
            if ((BOTTOM_OFFSET + scrollTop) >= docHeight - windowHeight) {
                createListItems();
            }
        });
    });
 
    function createListItems() {
        var mydocument = document;
        var mylist = mydocument.getElementById("items");
        var docFragments = mydocument.createDocumentFragment();
        for (var i = pos; i < pos + LIST_ITEM_SIZE; ++i) {
            var liItem = mydocument.createElement("li");
            liItem.innerHTML = "This is item " + i;
            docFragments.appendChild(liItem);
        }
        pos += LIST_ITEM_SIZE;
        mylist.appendChild(docFragments);
    }
})(jQuery);
