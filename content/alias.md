[//]: # (TITLE:Alias Command)
[//]: # (ICON:fas fa-external-link-alt)
[//]: # (DESCRIPTION:Using the Alias command to make shortcuts for media)
[//]: # (TAGS:commands,command,alias,aliases,aliasses,a)
[//]: # (COMMANDS:/oa Alias {alias name} {source},gives you the option of entering aliases for media ex. /openaudio alias test https://soundcloud.com/mrfijiwiji/imalright this media can now be accessed using the command /openaudio play @a a:test)

# Media Aliases
Source URL'S can get pretty long and may need replacement. Luckily, you can shorten them into aliases. This is like a shortcut for any given source.

To start, you need to define an alias, for this example, we'll make a sound called test

```
/openaudio alias test https://craftmend.com/mysound.mp3`
```

This means that every time we write a:test, we actually mean https://craftmend.com/mysound.mp3. You can use this alias as a source everywhere, regions, speakers, traincarts signs, play commands, you name it. Such an example play command is
```
/openaudio play @a a:test
```