# Generate pdfs from folders of pictures

This tool has the single purpose of generating a pdf from a folder of images.
Image formats supported are: webp, jpg, png

Command line params:

| Param | Description |
| ----- | --- |
| directory | The directory of images, or a directory of directories of images|
| outputDirectory|Output directory, by default will be the directory that was used as input |
| marginTopBottom |Padding to be added to top and bottom (effective padding can be larger due to aspect ratio) |
| marginLeftRight |Padding to be added to top and bottom (effective padding can be larger due to aspect ratio) |
| jpgOnly |Convert all images to JPG before adding them to the pdf (weird bug on Sony DPT-RP1 with png images) |
| v | Verbose mode |

The tool also supports table of contents. If in a folder of images a **toc.json** is found it will generate a pdf according to it. 

```json
{
	"title": "Simple pdf",
	"entries": [
		{
			"name": "Chapter 1",
			"file": "chapter1"
		}, {
			"name": "Chapter 2",
			"file": "chapter2"
		}
	]
}
```
*file* is the file name where the link will be created

*name* is the name that will be displayed in the table of contents page

# FAQ

### Should I use this?

Probably no. This was made for a very specific use case and with a particular formatting in mind.