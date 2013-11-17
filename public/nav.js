$(document).ready(function() {

    var time = null;
    var head = $(".head");
    var menu = $(".api ul").first();

    function SlideMenu() {
        clearTimeout(time);
        time = setTimeout(function () {
            menu.animate({
                top: head.offset().top + head.outerHeight() + 30
            });
        }, 1000);
    }

    $(menu).mouseover(function() {
        clearTimeout(time);
    });

    $(menu).mouseout(function() {
        SlideMenu();
    });

    $(window).scroll(function() {
        SlideMenu();
    });
});