.PHONY: aqara

aqara:
	 buildah bud -f ./build/aqara/Dockerfile -t aqara:0.0.1 --log-level=info
	 podman image save aqara:0.0.1 -o build/aqara/img-aqara.0.0.1
	 scp build/aqara/img-aqara.0.0.1 tix@192.168.1.30:/home/tix/images
	 // #/img-aqara.0.0.1


