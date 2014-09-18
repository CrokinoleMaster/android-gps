$(function() {
  var $position = $('#position')
  $position.text('lat: 37.69251 lng: -122.450866')
  var map = new GMaps({
    div: '#map',
    lat: 37.69251,
    lng: -122.450866
  })
  var marker = map.addMarker({
    lat: 37.69251,
    lng: -122.450866,
    infoWindow: {
      content: '(37, 37)'
    }
  })
  google.maps.event.addListener(map.map, 'click', function(e) {
    var pos = e.latLng
    marker.setPosition(pos)
    marker.infoWindow.content = pos.toString()
    marker.infoWindow.close()
    $position.text('lat: ' + pos.lat()  + ' lng: ' + pos.lng())
    $.ajax({
      type: 'POST',
      url: 'position',
      data: JSON.stringify({
        lat: pos.lat().toString(),
        lng: pos.lng().toString()
      }),
      dataType: 'json'
    })
  })

})
