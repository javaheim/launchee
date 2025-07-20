---
sidebar_position: 1
---

# Delete Shortcuts

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

You can delete a shortcut at user-level config by:

- specifying the unique name of the shortcut to be deleted
- adding mandatory `$patch: delete` directive

```yaml title="~/.config/launchee/launchee.yml"
shortcuts:
  - name: "Weather"
    $patch: delete
```

The merged configuration will result in:

```yaml
title: "My Launchee"
shortcuts:
  - name: "Terminal"
    icon: "/opt/kitty/lib/kitty/logo/kitty-128.png"
    command: "kitty"
```
