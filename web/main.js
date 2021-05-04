$(document).ready(function() {
    $(".btnGeneral").click(
        function(){
            let clients = $("#ajax_form").serialize()
            let arr = $(this).parent("form").find("input[type='text'],input[type='hidden']"), obj = {};
            $.each(arr, function(idx, el){
                obj[el.name] ? obj[el.name].push(el.value) : (obj[el.name] = el.value);
            });
            console.log(JSON.stringify(obj))
            sendAjaxForm("JSON",'/'+$(this).attr("data-action"), JSON.stringify(obj));
            return false;
        }
    );
    $(".show-form").on("click", function(){
        $(this).siblings(".hide").show()
    })
    $(".hide-form").on("click", function(){
        $(this).parent("form").hide()
    })
});

function sendAjaxForm(format, url, data) {
    console.log(data)
    $.ajax({
        url:     url,
        type:     "POST",
        dataType: format,
        data: data,
        statusCode: {
            201: function() {
                alert( "success" );
            }
        }
    });
}