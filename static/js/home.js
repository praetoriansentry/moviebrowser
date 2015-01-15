$(function () {

    var movies = new Bloodhound({
        datumTokenizer: Bloodhound.tokenizers.obj.whitespace('RawTitle'),
        queryTokenizer: Bloodhound.tokenizers.whitespace,
        limit: 10,
        prefetch: {
            url: '/movie/list.json'
        }
    });

    movies.initialize();

    $('.typeahead').typeahead(null, {
        name: 'movies',
        displayKey: 'RawTitle',
        source: movies.ttAdapter(),
        templates: {
            empty: '<p style="padding:10px;">No movies found!</p>',
            suggestion: function (movie) {
                return '<img width="50" src="/movie/' + movie.Ssid + '/image"> ' + movie.RawTitle;
            }
        }
    });

    $('.typeahead').on('typeahead:selected', function (evt, sug, name) {
        window.location = '/movie/' + sug.Ssid;
    })

});

