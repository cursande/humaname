# Humaname

Various packages can provide sample test data for different programming language ecosystems.

Sometimes for manual testing or a simple script, you just need **one** regular sounding hoo-man name. 

## Installation

### Linux (via curl)
``` shell
sudo curl -Lo /usr/local/bin/humaname https://github.com/cursande/humaname/releases/download/v0.1.1/humaname_linux_amd64
sudo chmod +x /usr/local/bin/humaname
```

### Mac/OSX
``` shell
# amd64
sudo curl -Lo /usr/local/bin/humaname https://github.com/cursande/humaname/releases/download/v0.1.1/humaname_darwin_amd64
sudo chmod +x /usr/local/bin/humaname
# arm64
sudo curl -Lo /usr/local/bin/humaname https://github.com/cursande/humaname/releases/download/v0.1.1/humaname_darwin_arm64
sudo chmod +x /usr/local/bin/humaname

```

## Usage

Help:
``` shell
humaname help
```

Any random first and last name:
``` shell
humaname b
humaname both
```

Any random female first name:
``` shell
humaname -g=f f
humaname -g=f first
```

Any random last name (either gender):
``` shell
humaname l
humaname last
```
