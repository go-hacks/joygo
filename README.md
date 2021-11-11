# joygo
## Nintendo Switch Joycon keyboard mapper for Linux

First, build with
``` sh
chmod +x build && ./build
```

Then pair your Joycons to your computer via Bluetooth.

Lastly, run with
``` sh
./joygo game.conf
```

You can make a config for each game you like and choose to use the R, L, or both Joycons together.

Certain local 2 player fighters might work nicely with this as joygo just maps each controller to
the specified keys so there's no issues with the joycons being separate controllers.

I have included the config I am using for Fell Seal (for reference) which is why I made this in the first place.

There is also a required hack to make Fell Seal work. There may be other games with the same issue.
Basically if a game has its own controller handling that conflicts with this and provides no option
to disable controllers in game then you'll need to use the additional 'fell' option and run it as root.

Example to use the included conf with Fell Seal
``` sh
sudo ./joygo FellSeal.conf fell
```

Reminder: The mouse is still active in game so make sure your cursor isn't off to the side or
the game will scroll that direction and seem like a joystick is stuck but it's not ;)

Things to come:

1. I might add mouse controls for the joysticks. Some games may work okay with it.
However, the joysticks are treated as D-Pads by the Linux driver so there's nothing I can do about this.

2. I will eventually add the ability for key combinations. Currently, you can only set single keys. 

NOTE: The key mapping format is from robotgo so if you would like a full reference, you can check it out here:

https://github.com/go-vgo/robotgo/blob/master/docs/keys.md
