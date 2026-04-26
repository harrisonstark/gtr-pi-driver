var config = {
    'carBaseURL': 'https://gtr-pi-car.harrisonstark.net'
}

$(document).ready(function() {
    $('.car-button').click(function() {
        const buttonId = $(this).attr('id');
        $.ajax({
            url: `${carBaseURL}/push_event?event=${buttonId}`,
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

const videoStream = document.getElementById("video-stream");
videoStream.src = `${carBaseURL}/stream_video`;
