[//]: # (COMMANDS:/oa debug,Returns current state of your OA plugin. State, Account Tags, Platform, TimeData, Throughput, Env, VoiceRTP, Build, RestDirect, Connected clients, loaded regions, loaded speakers, aliases, srvVersion)
[//]: # (COMMANDS:/oa state,Returns current state of your OA plugin. State, Account Tags, Platform, TimeData, Throughput, Env, VoiceRTP, Build, RestDirect, Connected clients, loaded regions, loaded speakers, aliases, srvVersion)
[//]: # (COMMANDS:/oa restart,Restarts the OpenAudioMc plugin. Be warned, this will kick ALL connected web client users.)
[//]: # (COMMANDS:/oa reload,Reloads the OpenAudioMc plugin. Be warned, this will kick ALL connected web client.)
[//]: # (COMMANDS:/oa help,Returns lis of all commands, clicktrough enabled for more specific info.)

# Fixing Voicechat
Not able to login to voicechat as expected? Are your friends unable to hear you even though you're shouting at the top of your lungs?? DID THE COMPUTER MACHINE GO BROKEN?
If you answered "YES" to any of those questions above, or are having any other shenanigans, hopefully this guide will help you get things running again.

### Check your Microphone
Your browser could be blocking your microphone access, but you can manually enable this [with this chrome menu](https://support.google.com/chrome/answer/2693767?co=GENIE.Platform%3DDesktop&hl=en) or by playing [in this menu in firefox](https://support.mozilla.org/en-US/kb/how-manage-your-camera-and-microphone-permissions). You might also want to check your system preferences if you run MacOS, since it could be blocking your browser from using your microphone on that level.

### Trying in a clean window
Still no luck? Why not try opening the voice chat window in a new Chrome incognito tab (or private browsing tab).
This forces your browser to grab the latest client build, which might help things go along a bit smoother.

### Check your server connection
Is everyone on your server having issues with voicechat? Then please check your server console (or proxy if you have one) for any errors or warnings related to "RTC", you can just try again later if this is the case. (shouldn't take more than 1 minute)

### Check the discord for maintenance announcements, or report a problem
Okay, you're still reading it so its properly scuffed huh? there's a good chance that I'm just doing some routine maintenance and updates on our end, so you can just check [our discord](https://discord.openaudiomc.net/) for notes about that. You can always file a bug report in `#openaudiomc-dev` if you really think that its broken, please provide as many details as possible (your browser console, plugin version, etc)
