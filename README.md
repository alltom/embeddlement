embeddlement
============

Gets the HTML for a static image preview of a URL.

`EmbedHtmlImage`: uses a `HEAD` request to check if the URL has an `image/*` MIME type. If so, just point to the image directly.

`EmbedHtmlEmbedly`: uses embed.ly's ombed API (you need an account to use this)

`EmbedHtmlAll`: tries both of those in order

All three return an HTML snippet like this (with width and height attributes on the image when possible):

    <a href="http://example.com/"><img src="http://example.com/example.jpg" /></a>
