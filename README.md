# go-get-music

This is my first Go project so lower your expectations...

I want a command line that gives me the best new music of the past few weeks. Something like...

    go-get-music --since <ISO date> --include <genre[]> --exclude <genre[]> --save=<bool> --limit=<int>

    #  |  ARTIST  |  ALBUM  |  GENRE  |  URL


That table can aggregate results from pitchfork, all music etc. And I'll probably just sort alphabetically.

Hopefully it can filter out some of the fast-food type pop/EDM/top-40 crap

Future work that could be cool is creating a playlist for spotify or apple music based on results, updating periodically by means of something like a cron job
