# Combat Calculator

The idea is to calculate your expected profit for each zone.

Inputs:

- KPH (from combat calculator spreadsheet)
- Drop tables (either good source or scrape wiki)
- Marketplace prices (idlescape.xyz)

## Marketplace prices

### https://api.idlescape.xyz/prices

curl https://api.idlescape.xyz/prices > prices.json

{
  items: Item[],
  discord: string,
}

Item {
  id: int,
  name: string,
  image: string,
  heat: int,
  price: int
}

E.g.

~~~
{
  id: 50,
  name: "Book",
  image: "/images/misc/book.png",
  heat: 50,
  price: 43950,
}
~~~

price is the current lowest price on the market

### https://api.idlescape.xyz/hourly

{ prices: map[itemId]map[timeType]Price[]}

timeType is "hourly"
Price {
  timestamp,
  price: int,
}

## Drop Tables

The data is on this page: https://idlescape.wiki/index.php?title=Template:Combat_Logs&action=edit

Not sure if there's a JSON representation of this anywhere.

You can get reasonable HTML by going to this page: https://idlescape.wiki/p/Special:ExpandTemplates

and pasting something like this in it:

```
<p>Small Rat</p>

{{Combat Logs|Small Rat}}

<p>Lesser Demon</p>

{{Combat Logs|Lesser Demon}}
```

To generate the wiki test for the full list of monsters, see `cmd/generateWikiText/main.go`
