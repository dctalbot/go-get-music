# go-get-music

This was my first project in Go, so lower your expectations...

I wanted a command line that would give me the best new music of the past few weeks. Something like...

    go-get-music --since <ISO date> --include <genre[]> --exclude <genre[]> --save=<bool> --limit=<int>

    #  |  ARTIST  |  ALBUM  |  GENRE  |  URL


Aggregated from Pitchfork, All Music etc. And I'll probably just sort alphabetically.

Future work could include creating a playlist for spotify or apple music based on results, updating periodically by means of something like a cron job
