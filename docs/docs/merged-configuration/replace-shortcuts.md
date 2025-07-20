---
sidebar_position: 3
---

# Replace Shortcuts

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
- all other fields
- adding optional `$patch: replace` directive

```yaml title="~/.config/launchee/launchee.yml"
shortcuts:
  - name: "Terminal"
    icon: "/usr/share/icons/Yaru/48x48@2x/apps/terminal-app.png"
    command: "gnome-terminal"
    $patch: replace
```

or just:

```yaml title="~/.config/launchee/launchee.yml"
shortcuts:
  - name: "Terminal"
    icon: "/usr/share/icons/Yaru/48x48@2x/apps/terminal-app.png"
    command: "gnome-terminal"
```

The merged configuration will result in:

```yaml
title: "My Launchee"
shortcuts:
  - name: "Terminal"
    icon: "/usr/share/icons/Yaru/48x48@2x/apps/terminal-app.png"
    command: "gnome-terminal"
  - name: "Weather"
    icon: "/usr/share/icons/Yaru/48x48@2x/apps/weather-app.png"
    url: "https://www.windy.com"
```
