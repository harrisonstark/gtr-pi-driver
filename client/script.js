$(document).ready(function() {
    $('.car-button').click(function() {
        const buttonId = $(this).attr('id');
        $.ajax({
            url: `https://gtr-pi-car.harrisonstark.net/push_event?event=${buttonId}`,
            method: 'POST',
            success: function(response) {
                console.log(`Action for ${buttonId} was successful, ${JSON.stringify(response)}`);
            },
            error: function(xhr, status, error) {
                console.error(`Error occurred: ${status} - ${error}`);
            }
        });
    });
});