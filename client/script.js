$(document).ready(function() {
    $('.car-button').click(function() {
        const buttonId = $(this).attr('id');
        $.ajax({
            url: `/push-car-event?${buttonId}`,
            method: 'POST',
            success: function(response) {
                console.log(`Action for ${buttonId} was successful.`);
            },
            error: function(xhr, status, error) {
                console.error(`Error occurred: ${status} - ${error}`);
            }
        });
    });
});