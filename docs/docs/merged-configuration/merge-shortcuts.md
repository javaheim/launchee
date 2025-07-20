---
sidebar_position: 2
---

# Merge Shortcuts

Given the system-config:

```yaml title="/etc/launchee/launchee.yml"
title: "My Launchee"
shortcuts:
  - name: "Terminal"
    icon: "/opt/kitty/lib/kitty/logo/kitty-128.png"
    command: "kitty"
  - name: "Weather"
    icon: "/usr/share/icons/Yaru/48x48@2x/apps/weather-app.png"
    url: "https://www.windy.com"
```

You can override a shortcut at user-level config by:

- specifying the unique name of the shortcut to be overridden
- at least one field which you want to override
- adding mandatory `$patch: merge` directive

```yaml title="~/.config/launchee/launchee.yml"
shortcuts:
  - name: "Weather"
    url: "https://www.accuweather.com/en/es/madrid/308526/weather-forecast/308526?city=madrid"
    $patch: merge
```

The merged configuration will result in:

```yaml
title: "My Launchee"
shortcuts:
  - name: "Terminal"
    icon: "/opt/kitty/lib/kitty/logo/kitty-128.png"
    command: "kitty"
  - name: "Weather"
    icon: "/usr/share/icons/Yaru/48x48@2x/apps/weather-app.png"
    url: "https://www.accuweather.com/en/es/madrid/308526/weather-forecast/308526?city=madrid"
```
