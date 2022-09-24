
# Why?

Why even bother using this instead of default LeoConsole plugin API?

## Advantages

 - more languages for LC plugins, maybe you like Go more that C#
 - not required to install Microsoft software for development, especially
   important if you use Linux 🙃
 - less boilerplate code (compare [gilc example](./index.html) and
   [LC plugin template](https://github.com/lcpluginmaker/PluginTemplate))
 - it is possible that additional library functions will be implemented that are
   not available in default ILeoConsole, which will simplify your development

## Disadvantages

 - bigger file sizes (`.dll`s are in KB-area, Go binaries are several MBs big)
 - worse performance: a new process has to be started every time you execute a
   plugin (the plugin is also executed multiple times for `PluginMain`,
   `PluginShutdown` and the command call itself), which has to parse the data
   etc. If you have a lot of Go plugins installed, it can take a little longer
   at start, because all of them have to run `PluginMain`
 - some ILeoConsole functions are not available (e. g. message formatting,
   however alternatives like [arrowprint](https://github.com/alexcoder04/arrowprint)
   exist)

