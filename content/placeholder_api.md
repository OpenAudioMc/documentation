[//]: # (TITLE:PlaceholderAPI Integration)
[//]: # (ICON:fas fa-sliders-h)
[//]: # (DESCRIPTION:Hooking into PAPI with OpenAudioMc)
[//]: # (TAGS:papi,placeholders,configs,technical)

# PAPI Integration
::warningstart::<strong>WARNING!</strong> Our PlaceholderAPI provider is only registered for the spigot instance ::warningend::
OpenAudioMc supports the [PlaceholderAPI](https://www.spigotmc.org/resources/placeholderapi.6245/) plugin. This means that you can display data from OpenAudioMc in your chat, scoreboard, tablist and more.

Be careful not to expose your token to the public, as it can be used to impersonate you.

## Placeholders
All our placeholders use the `oa_` prefix, and return values can be configured in our `config.yml` file.

| Placeholder         | Description                                                          | Example                       |
|---------------------|----------------------------------------------------------------------|-------------------------------|
| `oa_connect_client` | Connection state of the client in context                            | `Connected!`                  |
| `oa_connect_vc`     | Connection state to voicechat in context                             | `Not conneced to voice chat!` |
| `oa_count_client`   | The amount of players with the client open                           | `5`                           |
| `oa_count_vc`       | The amount of players in voice chat                                  | `3`                           |
| `oa_count_vc_max`   | The maximum amount of voicechat slots. Useful to displays like 5/10  | `8`                           |
| `oa_token`          | Shows the personal token of a player, should never be public in tab. | `38dDy`                       |

## Testing
You can easily validate if your placeholders are working by using the `/papi parse` command. For example, to test the `oa_token` placeholder, you would run `/papi parse <player> %oa_token%`.