[//]: # (TITLE:geyser)
[//]: # (ICON:fas fa-mobile-alt)
[//]: # (DESCRIPTION:Using OpenAudioMc with geyser/bedrock)
[//]: # (TAGS:installation,geyser,bedrock,pe)
# GeyserMC support
OpenAudioMc has technical support for Geyser, and handles it like a normal BungeeCord proxy (check the [installation](installation.md) instructions).

## Tokens
Minecraft Bedrock Edition *doesn't* support clickable links in chat, for, whatever reason, this means that bedrock players need to manually enter their client token when they want to open the web client.
You can add `{token}` as a placeholder in the `click-to-connect` message in your OA config (`plugins/OpenAudioMc/config.yml`), this will show the token key in chat. People have to manually enter it on https://client.openaudiomc.net/.
Make sure you are using clear messaging, like `click-to-connect: "Navigate to client.openaudiomc.net, your login token is {token}"` is recommended.

## Note on browsers
Keep in mind that mobile browsers (chrome mobile, ios safari, etc) aren't completely supported due to their resource saving measures, preventing OpenAudioMc from loading audio files normally and causing WebRTC stability issues. It's still recommended using a desktop (class) browser, even when playing on a mobile device.
