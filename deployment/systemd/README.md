# Deployer as a SystemD as a service

1. Create ``/etc/systemd/system/deployer.service`` from ``deployer.service``
2. ``systemctl daemon-reload``
3. ``systemctl start deployer.service``
3. ``systemctl enable deployer.service``