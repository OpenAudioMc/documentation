[//]: # (TITLE:Vistas API)
[//]: # (ICON:fab fa-java)
[//]: # (DESCRIPTION:Using vistas to scale OpenAudioMc in extremely large deployments)
[//]: # (TAGS:java,api,technical,documentation,events,packets,maven,gradle,nerds)

# 1. Introduction - The problem with state management
Vistas is OpenAudioMc’s official and production ready deployment setup for larger networks relying on an unknown amount of bungee cord proxies. Normal OpenAudioMc bungee cord installations use the proxy plugin instance (as wrapped in `OpenAudioMcBungee.java`) to track player sessions, states between servers and to wrap outgoing communication towards the OpenAudioMc backend infrastructure through one single Craftmend account. This is a perfect solution for smaller network, but causes some serious problems on large scale deployments, like (but not limited to);
1. Every proxy instance is seen as their own OpenAudioMc account, meaning that players from server A aren’t able to join voice chat sessions with players on server B
2. It requires manual setup for every proxy instance, making automatic scaling/templating impossible
3. Poses a compatibility issue for non-bungeecord servers, notable example is Imaginefun, which uses Lilipad

Vistas is a system that works around this problem, by completely skipping the proxy all together. Instead, it uses one single stand alone OpenAudioMc instance (as its own java process) which acts as the main relay node for the server, all Minecraft servers will directly communicate with your instance over Redis, completely mitigating the problems with the original project setup.

# 2. Important notes about this system
This new system does have a few major tradeoffs in its current form
- The used OpenAudioMc instance (`vistas-server`) is a single point of failure. The process **is** hot-swappable, meaning that a reboot of the service **does not** require any interactions in the Minecraft sub servers to recover player sessions. A scalable redundant system is planned for late 2022
- Some commands are now redundant or more confusing. `/oa *` commands will only affect your current Minecraft server (example, `/oa state` will only show the statistics of players in your current lobby), but admins can evaluate commands on the standalone instance by prefixing it with `vistas-eval` (like `/oa vistas-eval state`, which will show the state of the shared stand alone instance that all servers are using)
- A reliable Redis instance with decent Pub/Sub performance is a must. **Redis storage can be ignored as its currently not used**
- Only the stand alone OpenAudioMc instance (`vistas-server`) needs a voice chat license and to be linked to a craftmend account.
- Ratelimiting is still applicable to all sub servers, which can be a problem when sub servers are scaling on one physical host. Please contact Mats with a list of your IP Addresses to discuss a possible whitelist or excemption to the default rate limit configuration.

# 3. Setup requirements
> Prerequisites
1. A Redis instance with decent performance, preferably within a private network
2. A container or runtime configuration with java 9 or higher, capable of running standalone Jar files (with at least 2Gb of memory, and max 6). Make sure that only one of these services can be running at once!
3. A persistent Craftmend fingerprint (Contact Mats to request one)
4. OpenAudioMc 6.7 or higher
5. All servers must be running the same OpenAudioMc version, and use a compatible vistas-server and vistas-client module


# 4. Setting up the Vistas-Server
First, prepare your runtime environment following the requirements of chapter 3.
1. Download the latest `vistas-server.jar` from [OpenAudioMc/modules at master · Mindgamesnl/OpenAudioMc · GitHub](https://github.com/Mindgamesnl/OpenAudioMc/tree/master/modules), or build one yourself using the provided maven modules.
2. Download the default config file from [OpenAudioMc/config.yml at master · Mindgamesnl/OpenAudioMc · GitHub](https://github.com/Mindgamesnl/OpenAudioMc/blob/master/plugin/src/main/resources/config.yml) and paste it in the same folder as the `vistas-server.jar`
3. Configure Redis
    1. **Keep redis.enabled set to false to prevent default show behaviour**
    2. Set `redis.host` to the host address of your Redis server
    3. Set `redis.port` if your Redis server uses anything but the default port
    4. Set `redis.password` to the password of your Redis server. You **MUST** have authentication enabled as its enforced for security reasons
4. Configure startup arguments
    1. Go to the fingerprints page in your Craftmend account [Log in!](https://account.craftmend.com/account/fingerprint)
    2. Copy your **PERSISTENT FINGERPRINT** that you activated from the add-on you requested in chapter 3
    3. Run the vistas server using `java -Dfingerprint=aESA1Pzk6DchDuQLYPPMx92pXyiOfhw630arvmwA -jar vistas-server.jar`, where the fingerprint is obviously your own
5. Run the vistas server

# 5. Setting up the Vistas-Client
1. Install the OpenAudioMc plugin in the `plugins/` folder of your server
2. Create a sub directory within your plugins folder, called `OpenAudioMc/modules/`
3. Download the `vistas-client.jar` from [OpenAudioMc/modules at master · Mindgamesnl/OpenAudioMc · GitHub](https://github.com/Mindgamesnl/OpenAudioMc/tree/master/modules), and place it in the `plugins/OpenAudioMc/modules/` directory
4. Copy the `config.yml` file from your vistas-server, and place it in the `plugins/OpenAudioMc/` folder. **MAKE SURE THAT YOUR REDIS AND MESSAGES ARE VALID!**
5. Create a file called `data.yml` in the `plugins/OpenAudioMc/` folder and set its contents to
```yaml
# If the user agreed to the TOS and Privacy statement.
# Having any value representing "true" filled (either manually changed or through software after manual trigger or third party input) means
# That you read, understood and agree to everything mentioned in
# https://github.com/Mindgamesnl/OpenAudioMc/blob/master/LICENCE_and_PRIVACY.md
legal:
  accepted: true

regions:
aliases:
speakers:
keyset:
  key-version: -1
  private: not-set
  link-mode: true
  server-ip: not-set
  server-cc: not-set
debug:
  log-state-changes: false
```
6. Deploy the server as a template. All files that you have configured so far (the `config.yml` and `data.yml` are safe to be used as templates for scaleable sub servers)
7. Start the sub server (at least one)
8. Make sure that the vistas-server process is running for the initial test


