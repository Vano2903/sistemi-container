- comando per runc:
    - creare il file di configuazione
    > runc spec

    - installa il filesystem root
	> mkdir rootfs
    > docker export $(docker create busybox) | tar -C rootfs -xvf -

    - creare il docker
    > sudo runc create container
