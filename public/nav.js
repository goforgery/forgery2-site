$(document).ready(function() {

    var time = null;
    var head = $(".head");
    var menu = $(".api ul").first();

    $(menu).hover(function() {
        clearTimeout(time);
    });

    $(window).scroll(function() {
        clearTimeout(time);
        time = setTimeout(function () {
            menu.animate({
                top: head.offset().top + head.outerHeight() + 30
            });
        }, 1000);
    });
});