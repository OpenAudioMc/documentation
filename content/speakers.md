[//]: # (TITLE:Speakers)
[//]: # (ICON:fas fa-volume-up)
[//]: # (DESCRIPTION:Placing speaker blocks to have spatial in game audio)
[//]: # (TAGS:spatial,speaker,speakers,music,media,region,audio,media)
[//]: # (COMMANDS:/oa speaker {source} [radius],gives you a speaker player head with your source, you can change the speaker settings by right-clicking )
[//]: # (COMMANDS:/oa speaker gc,Activates the speakers garbage collector. Warning this will load all chunks with speakers in your world. You may experience some massive lagg on a big server. type /oa speaker gc confirm)
[//]: # (COMMANDS:/oa speaker menu,Opens a GUI menu with all nearby speakers, you can destroy, modify and teleport to the speaker from this menu.)
[//]: # (COMMANDS:/oa speaker remove {world} {x} {y} {z},Removes speaker on specified place on server.)
[//]: # (COMMANDS:/oa speaker set {world} {x} {y} {z} {url},Sets a speaker in the server with the media you want. Speakers always face north, if you want to have speakers in other directions please use /oa speaker <source>(use world for overworld)

# Speakers
Speakers are little blocks you can place in your Minecraft world. Just like real speakers in a theme park, all of them play synchronized and can be hidden. The thing that makes them special is that they share some properties just like real audio sources, You can't hear them from too far away, and the volume depends on your location relative to a speaker.

There's only one command that gives you the speaker, after executing you will receive a Speaker in your inventory. To remove them, you just go up to one and destroy the block like normal.

The command follows the following format:
```
/openaudio speaker <source> [radius]
```

Example, to receive a speaker playing the park entrance sound, you should do
```
/openaudio speaker https://soundcloud.com/mrfijiwiji/imalright 10
```

# Settings and modes
You can click on a speaker (with your right mouse button) to open the Speaker GUI.
You can use this GUI to check the source of the speaker, change the range/volume and modes.
Speakers have two modes, these are,
- **2D** Two dimensional audio is the default, this means that only the volume will change based on your distance with the speaker.
- **3D** Three dimensional audio is like surround sound. OpenAudioMc will work out where you are relative to the closest speaker and properly mix the sound based on your relative location.

You can change these settings at any time, but you may have to walk for a bit before they take effect (or, reload your client).

Notes:
- You can't use the same source URL in **2D** and **3D** when they are nearby, they have to be out of range from each other.
- Speakers will default to **2D** when places, unless there's a **3D** speaker with the similar source nearby.