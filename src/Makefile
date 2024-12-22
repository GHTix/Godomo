.PHONY: aqara

aqara:
	 buildah bud -f ./build/aqara/Dockerfile -t aqara:0.0.1 --log-level=info
	 podman image save aqara:0.0.1 -o release/aqara_0.0.1.img
	 scp release/aqara_0.0.1.img tix@192.168.1.30:/home/tix/images
	 //podman image load --input aqara_0.0.1.img

overkiz:
	 buildah bud -f ./build/overkiz/Dockerfile -t overkiz:0.0.1 --log-level=info
	 rm release/overkiz_0.0.1.img
	 podman image save overkiz:0.0.1 -o release/overkiz_0.0.1.img
	 scp release/overkiz_0.0.1.img tix@192.168.1.30:/home/tix/images
	 //podman image load --input overkiz_0.0.1.img
