#!/bin/sh
#
# My sample start scripts for multiple terminal windows
gnome-terminal –geometry=157×22+0+670 &
gnome-terminal –geometry=80×20+700 &
gnome-terminal –geometry=100×20+0 &

# My sample start applications for multiple terminal windows at set positions
#gnome-terminal --title="Mock Singpass Mobile" --geometry=80x15+20+25 --command='sh -c "cd ~/dev/singpass/sg-verify-demo-app/mock-server; npm run mock-spm"'

#gnome-terminal --title="Mock Sg-Verify" --geometry=80x15+850+25 --command='sh -c "cd ~/dev/singpass/sg-verify-demo-app/mock-server; npm run mock-sg-verify"'

#gnome-terminal --title="Company Callback Webhook" --geometry=80x15+20+400 -e "npm run start:webhook"

#gnome-terminal --title="Company Web Server" --geometry=80x15+850+400 -e "npm run start:client"

