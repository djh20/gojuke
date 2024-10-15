# importing vlc module
import vlc

# importing time module
import time

# creating vlc media player object
media_player = vlc.MediaPlayer()

source = "/var/lib/jukebox/example.mp3"

# media object
media = vlc.Media(source)

# setting media to the media player
media_player.set_media(media)
media.parse()
#media_player.play()

#media_player.set_position(0)

#time.sleep(1)
# setting volume
#media_player.audio_set_volume(10)


device = media_player.audio_output_device_enum()

while device:
    print(device.contents.device)

    #p.audio_output_device_set(None, device.contents.device)

    device = device.contents.next


# start playing video
media_player.play()

vol = 80

while media_player.audio_get_volume() != vol:
  media_player.audio_set_volume(vol)
  time.sleep(0.1)

input("Press any key to stop")