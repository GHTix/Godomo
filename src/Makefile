.PHONY: aqara

aqara_version = 0.0.4
overkiz_version = 0.0.4
control_version = 0.0.4


all: aqara overkiz control


control:
	buildah bud -f ./build/control/Dockerfile -t control:${control_version} --log-level=info
	rm -f release/control_${control_version}.img
	podman image save aqara:${control_version} -o release/control_${control_version}.img
	scp release/control_${control_version}.img tix@192.168.1.30:/home/tix/images


aqara:
	buildah bud -f ./build/aqara/Dockerfile -t aqara:${aqara_version} --log-level=info
	rm -f release/aqara_${aqara_version}.img
	podman image save aqara:${aqara_version} -o release/aqara_${aqara_version}.img
	scp release/aqara_${aqara_version}.img tix@192.168.1.30:/home/tix/images


overkiz:
	buildah bud -f ./build/overkiz/Dockerfile -t overkiz:${overkiz_version} --log-level=info
	rm -f release/overkiz_${overkiz_version}.img
	podman image save overkiz:${overkiz_version} -o release/overkiz_${overkiz_version}.img
	scp release/overkiz_${overkiz_version}.img tix@192.168.1.30:/home/tix/images

